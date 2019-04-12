// Code generated by go-bindata.
// sources:
// ../dist/straceback-main-bpf.o
// ../dist/straceback-tailcall-bpf.o
// DO NOT EDIT!

package straceback

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

var _stracebackMainBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\xbb\x6e\x13\x41\x14\x3d\xe3\xb5\x71\x88\x53\xa4\x01\x59\x02\xa4\xd0\xa5\x5a\x23\xbe\x20\x0a\xe2\x51\xb8\x88\xa0\xa1\x1b\x0d\xd6\x24\x58\x5a\xdb\xab\x99\x01\x12\x88\x04\x4d\x28\xe9\xa0\xa3\xe1\x0b\xe8\x28\xcd\x67\x50\xf2\x09\x14\x48\xb8\x62\xd1\xcc\xde\x61\xc3\xdd\x5d\x85\xf4\xb9\xc5\x5e\x9f\x33\xf7\x31\x73\xe6\xe1\xd7\x77\xc7\xf7\x3a\x42\x20\x9a\xc0\x2f\x54\xa8\xb2\xcf\xdd\xea\xf7\x0e\x7d\x37\x20\xb0\xbc\x5a\x72\x27\x00\xf6\x00\xbc\x5a\x5f\x15\x1e\x2f\x3f\x95\x7c\xbf\x03\xac\x8a\xa2\x18\xb2\xa2\x27\xa1\x17\x70\x05\xbd\x80\x55\x52\xf2\x4b\x55\xfa\x61\xa7\x1e\xbf\x01\xe0\x0b\xe1\xf7\x7f\xe7\x0b\xac\x01\xe8\x92\xbf\xb0\xd2\x12\xd2\xe4\x42\x97\x7f\xed\xfe\xde\x18\xbf\x8b\xa2\xd8\x24\x2c\x5e\x3e\xc4\xda\xf1\x40\x5c\x03\xe0\xb9\x21\xf1\x6f\xce\xa8\xf3\x21\x7c\x13\xe4\x0d\x63\x09\x92\x1a\xb7\x08\xf5\xfb\x35\x7e\x3f\xf0\xf5\x1d\x7a\x10\xf8\x5e\x8d\x7f\x1c\xf8\x4b\x35\xbe\x1f\xf8\x7a\xdf\x6d\xf2\xfe\x9e\xf4\x28\x3f\x62\x5f\xe5\x3a\xe1\x75\x00\xe1\xca\xa5\x4e\x1f\x3a\x38\xa3\x26\x3a\x5f\x4c\xe7\x4e\x4a\x7b\x64\xa5\x9e\x3b\x6d\x90\x1a\x9d\x55\x23\x23\xa3\x5e\xf8\xc1\x89\xca\x32\x3b\xaa\xa2\x66\x2a\xb7\xa3\xc9\x81\x59\x3c\xcb\xe5\x4c\xe5\x25\x76\x6a\x9a\x49\x1f\x19\x28\xf9\x5c\x1b\x3b\x5d\xcc\x21\xb3\xe9\x44\xcf\xad\x0e\xa5\x53\xfd\x54\xee\x1b\x35\xd3\x48\xad\x33\x4e\x3d\x41\x6a\x8f\x66\xde\x8f\x77\x77\x6f\xc9\xdb\x0d\x52\x9f\xdb\xde\xd2\xdd\xe0\xb6\x49\xef\xdb\x47\xc6\xf3\xb7\x50\x9c\xd2\xee\xb4\xed\xb4\xf4\xeb\x32\x7c\xf3\x8c\x7c\x7e\xee\xf8\xc9\xb8\x01\xe0\x72\x43\x9f\xef\xb4\xa8\x2d\xc2\x03\x5a\x67\xcc\x8f\xe7\xfd\x0e\xf5\xe7\x1a\x7c\x25\xcf\xdf\x69\x3e\xff\x47\x2d\xf9\xdf\xc4\xff\xe5\xe7\x2d\xf9\x3f\x3a\xcd\xf1\x5c\xff\x83\x96\xfc\x9f\x2d\xf9\x1c\x1f\x53\x3e\xfb\x7b\xc1\x8a\x88\x6d\xc6\x73\xfd\x0f\x5b\xf4\x8f\x8d\xa2\xce\x03\x8a\xe3\xfa\xbf\x6b\xe8\xed\x6d\x8b\x16\x14\xf7\x41\xb0\xfd\x8b\xef\xd2\x9f\x00\x00\x00\xff\xff\x9c\x4c\xdf\x84\xb0\x07\x00\x00")

func stracebackMainBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackMainBpfO,
		"straceback-main-bpf.o",
	)
}

func stracebackMainBpfO() (*asset, error) {
	bytes, err := stracebackMainBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-main-bpf.o", size: 1968, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stracebackTailcallBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x3d\x6c\x13\x4b\x10\xfe\xce\x3f\xb1\x9f\x5f\xde\x53\xa4\x27\x3d\x19\x43\x71\x42\x58\x0a\x05\x0e\xa1\xa0\x40\x8a\x94\x06\x68\x22\x88\xd2\x40\x83\xcc\x61\x2f\x60\x71\x71\x9c\xbd\x73\x42\x72\x48\x48\x20\x9a\xf4\x48\x11\x55\xf8\x29\x52\xa6\x73\xba\xa3\xa4\x4c\x49\x49\x49\x99\x8e\x54\x2c\x9a\xbd\x39\x9f\xbd\xe7\x53\xd2\xc3\x4a\xc9\xec\x7c\xbb\xb3\x33\xdf\xfc\x9c\x5f\xde\x5c\xba\x95\xb3\x2c\xc4\xcb\xc2\x0f\x24\x5a\xb2\xbe\x4c\x25\xfb\x45\xfe\xff\x0f\x2c\x84\xff\x47\xd8\x1b\x00\x65\x00\x61\x39\xd1\xff\x25\xbd\x94\xe8\x45\x00\x41\xe5\x48\x91\xfe\xa4\x0c\xd8\x00\x36\x59\x06\x3b\x5f\x35\x1e\x6c\x7f\xd3\x72\xcb\x89\x1e\x0a\x6a\xdf\xb5\x3e\xe0\xa0\x82\xda\x49\x74\xaf\x76\xac\x65\xf8\x3e\xc2\x4b\x16\x70\xac\x94\x1a\xe4\x80\x19\xf6\x47\x32\xfc\xc4\xe7\x05\xe0\x48\x29\x15\x3a\x91\x5e\xcd\x8d\xf3\x0b\x5f\x47\x72\x50\x04\xae\xb2\xfd\x39\xf6\x5b\x01\xf0\xac\xf6\x59\xfb\xab\x5a\x40\xdf\x16\x92\xf6\x72\xa1\xde\x0e\x6a\x87\x43\xdc\xed\xdb\x77\x68\xbf\x62\xd7\xdd\xa0\x76\x30\xc4\x6f\xd8\xbd\x0e\xed\xdb\x1a\xdf\xd7\xf8\x66\x29\xe2\x5d\x2d\x03\xbe\xd3\x71\x09\x6b\x39\xae\x1b\xec\xec\x45\xfc\xdb\x11\xff\x51\x7e\x7b\xcc\xef\x22\xe1\x1e\xc7\x5d\x4c\xf2\x3b\x35\x96\xb7\xdd\x28\x3f\x1f\xd8\x3e\x07\xec\x2a\xa5\xaa\x46\x71\x07\x79\x20\xc7\xf6\xf7\x38\x5e\x21\xe5\x02\x9d\xd5\xdb\x15\x8c\xf2\xeb\x4a\xbb\x4e\x7b\xd7\xed\xdb\xe3\xfc\xd6\xfb\xb4\x17\x7d\x61\xc7\xfc\x86\x3c\x3e\xee\xa6\xea\x14\xf3\x20\xfe\x61\x61\x3c\xfe\x01\xc7\xf5\x96\x65\x61\xe4\x0f\x85\x09\x8d\xf9\x9b\x2e\x1e\x3b\x3d\x73\x7f\xf2\x92\x2c\x9a\x25\x9a\x23\x9a\x39\x9a\x37\x9a\x49\x9a\x47\x9a\x59\x9a\xd7\x4a\xe6\xfd\xf5\x3e\xf5\x2f\xf5\xb8\x3b\x72\xfb\xf6\xf2\x12\x7e\x2a\xa5\x66\xf8\xbe\xb5\xbd\x82\xf2\x8b\xbf\xad\xf3\xfc\x7d\xa9\x32\xfe\x30\x77\x7a\x6c\x79\xe4\x53\x58\x4b\xbf\x53\x4e\xe1\x77\x35\xfe\x57\x0a\x9f\xd6\x78\x31\x85\x3f\xd0\xf8\x54\x0a\xff\x4f\xe3\x69\xbf\x87\x31\x1f\x6e\x9f\x03\x2b\xd1\xe9\xf5\x0b\x7c\x5e\x61\x0c\x0d\x5f\x3c\xf7\xb1\xea\xf4\xbc\x39\xb1\x21\xba\xbe\x07\x5f\x3a\x2d\xd1\x5b\xeb\x74\xfd\x66\xd3\xdb\xf2\x9a\xa2\xeb\x0b\x89\x86\x14\x6e\x72\x32\x27\x9d\x4d\x3a\xa4\x14\x7b\x73\xc9\xad\xe6\x86\x90\x5e\x67\xad\x1b\x3d\xa8\x53\x8f\xa6\xdb\x69\x89\xae\x27\xf4\x13\x0d\xf1\xb4\xf9\x58\x3a\xab\x02\x0d\xcf\x97\xbe\xf3\x08\x0d\x6f\x6b\x55\x4b\xb9\xd6\x76\x7c\x87\xe0\xf9\xc6\xfc\xf5\xd3\xd3\x7e\xa6\xb5\xad\xeb\x93\x5e\xb3\x9c\xd2\x77\x06\x6e\xfe\x46\x5a\xfc\x67\x56\x60\x31\xc3\x9f\x39\xb2\x57\x4e\xb1\x37\xfb\xcb\xec\x98\xcb\xc0\x84\x6e\x01\x4e\xb8\x55\x6c\xd6\xa7\x99\x67\x6c\x1f\xf7\x75\x89\xfd\x9b\x39\xd8\x63\xbf\xe6\xef\x86\x19\xff\xfd\x0c\xfb\xc3\xfc\xd9\xec\x77\xd8\xfe\x9a\x81\x1f\xf1\xc5\x65\x03\x1f\xce\x23\xcb\x76\x86\x7f\xbb\x38\xd9\x9f\x59\xbf\xe5\x0c\xfb\x4b\x19\xf6\xa6\xde\x63\x7b\xf3\x33\x30\xcb\xf6\xb3\x06\x6e\xd6\xcf\xcd\xa8\x5f\x95\x1b\x22\xe6\x3b\xcd\x33\x69\xd6\xef\xd5\x04\xdf\x9a\x17\xfb\xdf\x67\x3d\x8e\x31\xb6\x8f\xbf\x5f\xbf\x02\x00\x00\xff\xff\x9b\x77\x90\xc7\x08\x0a\x00\x00")

func stracebackTailcallBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackTailcallBpfO,
		"straceback-tailcall-bpf.o",
	)
}

func stracebackTailcallBpfO() (*asset, error) {
	bytes, err := stracebackTailcallBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-tailcall-bpf.o", size: 2568, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"straceback-main-bpf.o": stracebackMainBpfO,
	"straceback-tailcall-bpf.o": stracebackTailcallBpfO,
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
	"straceback-main-bpf.o": &bintree{stracebackMainBpfO, map[string]*bintree{}},
	"straceback-tailcall-bpf.o": &bintree{stracebackTailcallBpfO, map[string]*bintree{}},
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

