package kloud

import (
	"errors"
	"strconv"
	"time"

	"github.com/koding/kite"
	"github.com/mitchellh/mapstructure"
)

// Builder is used to create and provisiong a single image or machine for a
// given Provider.
type Builder interface {
	// Prepare is responsible of configuring the builder and validating the
	// given configuration prior Build.
	Prepare(...interface{}) error

	// Build is creating a image and a machine.
	Build(...interface{}) (interface{}, error)
}

type BuildResponse struct {
	MachineName string `json:"machineName" mapstructure:"machineName"`
	MachineId   int    `json:"machineId" mapstructure:"machineId"`
	KiteId      string `json:"kiteId" mapstructure:"kiteId"`
	IpAddress   string `json:"ipAddress" mapstructure:"ipAddress"`
}

type BuildArgs struct {
	Provider     string
	SnapshotName string
	MachineName  string
	Credential   map[string]interface{}
	Builder      map[string]interface{}
}

var (
	defaultSnapshotName = "koding-klient-0.0.1"
	providers           = make(map[string]interface{})
)

func (k *Kloud) build(r *kite.Request) (interface{}, error) {
	args := &BuildArgs{}
	if err := r.Args.One().Unmarshal(args); err != nil {
		return nil, err
	}

	p, ok := providers[args.Provider]
	if !ok {
		return nil, errors.New("provider not supported")
	}

	provider, ok := p.(Builder)
	if !ok {
		return nil, errors.New("provider doesn't satisfy the builder interface.")
	}

	if err := provider.Prepare(args.Credential, args.Builder); err != nil {
		return nil, err
	}

	snapshotName := defaultSnapshotName
	if args.SnapshotName != "" {
		snapshotName = args.SnapshotName
	}

	signFunc := func() (string, string, error) {
		return createKey(r.Username, k.KontrolURL, k.KontrolPrivateKey, k.KontrolPublicKey)
	}

	machineName := r.Username + "-" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	if args.MachineName != "" {
		machineName = args.MachineName
	}

	artifact, err := provider.Build(snapshotName, machineName, signFunc)
	if err != nil {
		return nil, err
	}

	response := &BuildResponse{}
	if err := mapstructure.Decode(artifact, response); err != nil {
		return nil, err
	}

	return response, nil
}
