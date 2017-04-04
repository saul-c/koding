// Code generated by go-bindata.
// sources:
// bootstrap.json.tmpl
// DO NOT EDIT!

package google

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bootstrapJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x5d\x6f\xda\x30\x14\x7d\xe7\x57\x58\xd6\x1e\x49\xc6\xc7\xe8\x28\xaf\x9b\xa6\x4e\xd3\xaa\x69\xed\x5b\x85\xd0\x25\xb9\xa5\x1e\xc6\x8e\xfc\x01\xd2\x50\xfe\x7b\xe5\x38\x21\x0e\x49\x4b\x5e\x2a\xb5\xe2\x25\xd7\x3e\xf7\xf8\xdc\xe3\x6b\xec\xe3\x80\x10\x9a\x29\xb9\x67\x29\x2a\xba\x20\x2e\x26\x84\x6e\xa4\xdc\x70\x3c\xc5\x84\xd0\x44\x61\x8a\xc2\x30\xe0\x9a\x2e\x08\xfd\x74\xdc\x83\x8a\x3d\x6c\x15\xcc\xe5\x74\x58\x65\x64\x4a\xfe\xc3\xc4\xb4\xd0\xe5\x78\x80\x54\xb8\x61\x52\xb4\x80\x7e\x38\xa7\x05\x2c\x1f\x10\x92\xbb\x0c\x2a\xad\xc9\xac\xa9\xb5\x6e\x65\xca\xc4\x66\x25\xd0\x1c\xa4\xda\xae\x58\x1a\xca\xde\x03\xb7\xe8\x99\x2b\xb1\x72\x97\x59\x83\x15\x3e\xf6\xe9\x31\x4b\x5b\x0b\x29\xd4\xd2\xaa\x04\xcf\x6d\x39\xa7\x08\xd7\xf3\x6c\xc1\x08\x21\x54\xc0\xae\x90\x70\x3c\xc6\xb7\x3e\xe3\x16\x76\x98\xd7\x06\x10\x42\xc1\x1a\xe9\x7c\x04\x83\x2b\x6d\xd7\x25\xb3\xa6\x64\x41\x8c\xb2\x58\x02\x73\xaf\x70\xd8\xa9\xe6\x91\x29\x3c\x00\xe7\x6d\x39\x11\x70\x2e\x0f\xd1\x93\x31\xd9\x8b\xca\x7e\x94\xd9\x4e\xda\xcd\xfd\xfd\x9f\xa6\xbc\xba\xd2\x8b\x46\x3a\xce\x46\xae\x37\x71\xa5\x40\x6c\xd0\xf5\xce\x03\x1d\xc5\xc5\xef\xf3\x88\x2e\x87\x83\xc0\x03\xa7\xb2\x21\xd0\x77\x91\x91\x89\x74\x55\x51\x93\x64\x01\xb1\x9b\x94\xca\x78\xca\xf9\x88\x2e\x4f\x33\x79\x65\xd7\xb0\xd3\x07\x96\xec\xfa\xf9\xf0\xf3\xdb\xef\x77\xe9\x43\x51\x40\xef\x6a\x85\x41\x25\x80\xf7\xab\xb8\x04\xbf\x65\xd5\xfe\x8c\x9f\x9f\xda\x6c\xff\xc5\x83\xf2\x6e\x33\x1e\x02\x33\x42\x63\x2e\xb4\x48\xa3\x49\x46\xd1\xd5\x6c\x36\x9d\x05\x9d\x12\xb8\x76\x81\xd8\xa6\x6f\x44\xdc\xdc\xcc\x7a\x3b\x09\x59\xbe\xbe\xb1\x2a\xed\xd7\xc5\x7f\xbf\xbf\xcb\x26\x7e\xe5\x30\x4f\xa7\xf3\xeb\xfe\xc7\x59\xeb\xa7\x5e\x3e\xdc\xdd\xdd\x7c\x30\x1f\x26\x93\xfe\x2e\x6c\x39\x43\x61\x7a\x19\xf1\xab\x80\x7e\x30\x2f\x66\x57\x5f\xbb\x9b\xe2\xec\xd2\xde\x83\x62\xb0\x0e\xde\x2e\xf4\xbf\x14\x85\xac\xd3\xa5\x98\xe2\x23\x58\x7e\xe6\x15\x68\x06\x11\x82\x36\x63\x27\xa4\x8e\x22\x08\x4b\x45\xab\x64\x86\xd1\x01\x4b\x5c\x18\x47\xeb\x10\x69\x75\x94\xa0\x30\x0a\x78\x01\x0c\xc2\x26\xa3\xd5\xf5\xaa\xd5\x77\x8b\xe9\xb4\x5e\xf5\x1d\x01\xed\x7c\x12\xbc\xf8\xaf\xda\x5d\x3f\x1d\x8f\xe2\xf1\x64\x5e\xec\xd5\x75\xf0\xfc\x19\xe4\x83\xe7\x00\x00\x00\xff\xff\x26\x9f\x37\x7a\x19\x0a\x00\x00")

func bootstrapJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapJsonTmpl,
		"bootstrap.json.tmpl",
	)
}

func bootstrapJsonTmpl() (*asset, error) {
	bytes, err := bootstrapJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.json.tmpl", size: 2585, mode: os.FileMode(420), modTime: time.Unix(1475345133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bootstrap.json.tmpl": bootstrapJsonTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bootstrap.json.tmpl": {bootstrapJsonTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
