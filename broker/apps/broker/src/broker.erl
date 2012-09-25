%%%-------------------------------------------------------------------
%%% File : broker.erl
%%% Author : Son Tran-Nguyen <son@koding.com>
%%% Description : A named gen_server to handle subscribe request from
%%% client. It keeps track of a supervisor to create subscriptions.
%%% It is under the supervision of another supervisor.
%%%
%%% Created : 27 August 2012 by Son Tran <sntran@koding.com>
%%%-------------------------------------------------------------------
-module(broker).
-behaviour(gen_server).
%% API
-export([start_link/0, subscribe/2, presence/3, unsubscribe/1,
            bind/2, unbind/2, trigger/4, trigger/5, rpc/3]).
%% gen_server callbacks
-export([init/1, terminate/2, code_change/3,
        handle_call/3, handle_cast/2, handle_info/2]).
-define (SERVER, ?MODULE).

-include_lib("amqp_client/include/amqp_client.hrl").
-compile([{parse_transform, lager_transform}]).

%%====================================================================
%% API
%%====================================================================
%%--------------------------------------------------------------------
%% Function: start_link() -> {ok,Pid} | ignore | {error,Error}
%% Description: Starts the server
%%--------------------------------------------------------------------
start_link() -> 
    MqHost = get_env(mq_host, "localhost"),
    MqUser = get_env(mq_user, <<"guest">>),
    MqPass = get_env(mq_pass, <<"guest">>),

    {ok, Connection} = amqp_connection:start(#amqp_params_network{
        host = MqHost, username = MqUser, password = MqPass
        }),

    gen_server:start_link({local, ?SERVER}, ?MODULE, [Connection], []).

%%--------------------------------------------------------------------
%% Function: subscribe(Exchange) -> Reply
%% Description: Init a subscription for the requesting client.
%%--------------------------------------------------------------------
subscribe(Conn, Exchange) ->
    gen_server:call(?SERVER, {subscribe, Conn, Exchange}).

unsubscribe(Subscription) when is_pid(Subscription) ->
    gen_server:call(?SERVER, {unsubscribe, Subscription}).

presence(Conn, Where, Presenter) ->
    gen_server:call(?SERVER, {presence, Conn, Where, Presenter}).

%%====================================================================
%% Wrappers for subscription gen_server
%%====================================================================
bind(Subscription, Event) ->
    subscription:bind(Subscription, Event).

unbind(Subscription, Event) ->
    subscription:unbind(Subscription, Event).

trigger(Subscription, Event, Payload, Meta) ->
    trigger(Subscription, Event, Payload, Meta, false).

trigger(Subscription, Event, Payload, Meta, NoRestriction) ->
    subscription:trigger(Subscription, 
                        Event, 
                        Payload, 
                        Meta, 
                        NoRestriction).

rpc(Subscription, RoutingKey, Payload) ->
    gen_server:call(Subscription, {rpc, RoutingKey, Payload}).

%%====================================================================
%% gen_server callbacks
%%====================================================================
%%--------------------------------------------------------------------
%% Function: init(Args) -> {ok, State} |
%% {ok, State, Timeout} |
%% ignore |
%% {stop, Reason}
%% Description: Initiates the server
%%--------------------------------------------------------------------
init([Connection]) ->
    % To know when the supervisor shuts down. In that case, this
    % terminate function will be called to give the gen_server a chance
    % to clean up.
    process_flag(trap_exit, true),
    {ok, Connection}.

%%--------------------------------------------------------------------
%% Function: %% handle_call({subscribe, Conn, Exchange}, From, State) 
%%                          -> {noreply, State} 
%% Description: Handling subscription request. Subscription supervisor
%% will create one under its supervision tree.
%%--------------------------------------------------------------------
handle_call({subscribe, Conn, Exchange}, From, Connection) ->
    Result = subscription_sup:start_subscription(Connection,
                                                From, 
                                                Conn, 
                                                Exchange),
    {reply, Result, Connection};

%%--------------------------------------------------------------------
%% Function: %% handle_call({unsubscribe, Subscription}, From, State) -> 
%%                          {noreply, State} 
%% Description: Handling unsubscription request. This will tell the
%% subscription supervisor to stop the child subscription, but not call
%% its `terminate/2` to do all clean up necessesary.
%%-------------------------------------------------------------------- 
handle_call({unsubscribe, Subscription}, _From, Connection) ->
    ok = subscription_sup:stop_subscription(Subscription),
    {reply, ok, Connection};

%%--------------------------------------------------------------------
%% Function: %% handle_call({presence, Conn, Exchange}, From, State) -> 
%%                          {noreply, State} 
%% Description: Handling presence. It will subscribe to an x-presence
%% exchange, create a binding with an empty key, and another binding
%% with the presenter key to announce the presenter's presence. 
%%--------------------------------------------------------------------
handle_call({presence, Conn, Where, Presenter}, From, Connection) ->
    PresencePrefix = get_env(presence_prefix, <<"KDPresence-">>),
    Exchange = <<PresencePrefix/bitstring, Where/bitstring>>,
    Result = subscription_sup:start_subscription(Connection,
                                                From, 
                                                Conn, 
                                                Exchange),
    case Result of
        {ok, SID} ->
            subscription:bind(SID, <<>>),
            subscription:bind(SID, Presenter),
            {reply, Result, Connection};
        {error, _Err} ->
            {reply, Result, Connection}
    end;

%%--------------------------------------------------------------------
%% Function: %% handle_call(Request, From, State) -> {reply, Reply, State} |
%% {reply, Reply, State, Timeout} |
%% {noreply, State} |
%% {noreply, State, Timeout} |
%% {stop, Reason, Reply, State} |
%% {stop, Reason, State}
%% Description: Handling call messages
%%--------------------------------------------------------------------
handle_call(_Request, _From, State) ->
    Reply = ok,
    {reply, Reply, State}.
    
%%--------------------------------------------------------------------
%% Function: handle_cast(Msg, State) -> {noreply, State} |
%% {noreply, State, Timeout} |
%% {stop, Reason, State}
%% Description: Handling cast messages
%%--------------------------------------------------------------------
handle_cast(_Msg, State) ->
    {noreply, State}.
    
%%--------------------------------------------------------------------
%% Function: handle_info(Info, State) -> {noreply, State} |
%% {noreply, State, Timeout} |
%% {stop, Reason, State}
%% Description: Handling all non call/cast messages
%%--------------------------------------------------------------------
handle_info(_Info, State) ->
    {noreply, State}.
    
%%--------------------------------------------------------------------
%% Function: terminate(Reason, State) -> void()
%% Description: This function is called by a gen_server when it is about to
%% terminate. It should be the opposite of Module:init/1 and do any necessary
%% cleaning up. When it returns, the gen_server terminates with Reason.
%% The return value is ignored.
%%--------------------------------------------------------------------
terminate(_Reason, Connection) ->
    lager:info("Broker server dies, closing connection ~p", [Connection]),
    amqp_connection:close(Connection),
    ok.

%%--------------------------------------------------------------------
%% Func: code_change(OldVsn, State, Extra) -> {ok, NewState}
%% Description: Convert process state when code is changed
%%--------------------------------------------------------------------
code_change(_OldVsn, State, _Extra) ->
    {ok, State}.

%%--------------------------------------------------------------------
%%% Internal functions
%%--------------------------------------------------------------------
get_env(Param, DefaultValue) ->
    case application:get_env(broker, Param) of
        {ok, Val} -> Val;
        undefined -> DefaultValue
    end.