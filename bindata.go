// Code generated by go-bindata.
// sources:
// map.go.tmpl
// DO NOT EDIT!

package mapgen

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

var _mapGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x5b\x6f\xa3\x46\x14\x7e\x8e\x7f\xc5\xe9\x4b\x0b\x15\x82\xf7\x48\x79\xda\x58\x51\xd4\x90\x54\xbb\xd9\xf6\x21\x5a\x55\x13\x38\x98\x11\x73\x41\x33\x83\x6d\x64\xf1\xdf\xab\xc3\x40\x6c\xe2\xbb\x76\x5f\x2c\x33\x73\xae\xdf\x05\x92\x04\xbe\xe8\x1c\x61\x81\x0a\x0d\x73\x98\xc3\x7b\x0b\x92\xd5\x0b\x54\x10\x94\xce\xd5\xf6\x36\x49\x16\xdc\x95\xcd\x7b\x9c\x69\x99\x30\x29\x99\xe1\x3a\xf1\x21\x61\x04\xf7\x2f\xf0\xfc\xf2\x0a\xf3\xfb\xc7\xd7\x78\x56\xb3\xac\x62\x0b\x84\xcd\x06\xe2\xbf\x87\xff\x5d\x37\x9b\x25\x09\x6c\x36\x71\xca\xea\x67\x26\xb1\xeb\x80\x5b\x60\x3b\x2d\x5d\x69\x90\xe5\x60\x59\x81\xd4\x9b\xc2\x57\xdc\x95\x50\x61\x4b\x79\x7f\x61\xfb\xda\xd6\x94\xc7\x54\x0e\x4b\x26\x1a\xea\x10\xff\xc3\x84\x3f\x9e\xb9\xb6\xc6\x69\x03\xeb\x4c\x93\x39\xd8\xcc\x00\x80\xa6\xe1\x05\xc4\xdf\x2d\x7e\xfd\x37\x6d\x1c\xae\x69\x26\xba\xb0\xad\xca\xe2\xe1\x6c\x8c\x44\x61\x71\x72\x3f\xbd\x55\xf9\x78\xd9\xff\x24\x09\xa4\x90\x69\xe5\x18\x57\x16\x5c\x89\xd0\xa8\x1c\x8d\x68\xb9\x5a\xd0\x2a\xf1\x18\xf5\xa0\x8d\x6e\x1c\x57\x68\x61\x55\xf2\xac\x04\x96\x65\x68\x2d\xa4\x90\x73\x83\x99\x13\x2d\xd8\x52\x37\x22\x87\x52\x8b\x7c\xcc\xa2\x82\x92\xfa\xfb\x3a\x29\x95\x7c\xdb\x45\xe4\xc7\x04\x87\x0e\x60\xc0\x9a\xf6\x9d\xaf\x6b\x6d\x1c\xe6\x5d\xf7\x8c\xab\xed\x66\x8a\x1e\x50\xe5\x5d\xb7\xd9\x38\xee\x04\xc2\x0e\x6e\x06\x5d\x63\x94\x05\xa6\x80\x2b\xeb\x98\x72\xfc\x97\x30\x54\x34\x2a\xbb\x7c\x2c\x98\xce\x05\x5d\x17\x84\xf0\xe7\x84\x60\xcf\xac\x1f\x17\x7e\xdf\xbd\xf2\x37\x3d\x5c\xb7\x20\x59\x85\xc1\x49\xd0\xc2\xa8\x8f\xef\x66\x5e\xa6\x2f\x35\x2a\x60\x42\xe8\x15\x49\x34\x13\xda\x36\x06\xc1\xe9\x7e\x75\xd1\x82\xae\x7b\xcd\x82\x56\x9e\x1c\x56\xfb\xdd\x02\x39\x1d\x30\xec\x2b\x05\x05\xd0\x6d\x10\x86\xc3\xc0\x32\x7e\xd2\x59\x15\x84\xfd\x43\x8e\x05\x1a\x90\xf1\x77\x25\xb6\x87\x45\x10\x0e\xa3\x7c\xd1\x75\xfb\x61\x92\x7e\x1a\x3a\xd0\xc5\xd8\x38\x3e\xd6\x99\x12\x83\x10\x82\xec\xb4\x5c\x68\xa6\x9b\x0c\xee\x2e\x00\x29\x02\x81\x2a\x90\x71\x1a\x86\xa7\x1d\x25\xe3\xaf\x3b\x0b\x7e\xb2\xd3\x64\xf9\xad\x99\x6e\x0a\x6d\xa0\x8a\x60\x09\xb7\x77\x60\x98\x5a\x20\xc8\x38\xa5\xe1\x6e\xb2\xb7\xea\x07\xdc\xc1\x72\x76\xd3\x9d\xed\x3b\x41\x71\xaf\xf3\xde\xed\xd0\x7b\x50\x50\x36\x60\xfe\x0d\x1d\x58\x74\x84\x36\x49\xfb\x3c\xcb\xdf\xd0\x05\x9f\x4d\x10\x91\x03\xe0\x33\xd4\x7b\x08\xc8\x38\x7d\xab\xb0\xed\x17\x64\xe2\xf3\x9c\x1f\xf3\x3c\x16\xcf\xda\xcd\xd7\xdc\x1e\x99\x8c\x10\xe1\x0e\x72\x8d\x56\xfd\xe1\x00\x29\x32\xa6\xdc\x47\xf7\x61\x67\x0a\xf5\xae\xf4\x2f\x1f\x6e\xa9\xd6\x51\x01\x4d\xba\x5e\xb6\xde\xce\xd3\xa1\x5d\x97\x11\xe8\x8a\xf8\x1d\x97\xee\x4f\x79\x01\xbf\xe9\x0a\xb6\x8e\x3d\x00\x49\xe7\x89\xdc\xc6\x50\x73\x92\x84\x77\xed\x01\x76\x07\x4a\xa9\x80\x07\xf1\x1e\x05\x3a\x04\x83\x52\x2f\x71\x04\xb0\x30\x5a\x9e\x23\xd7\x27\xee\x01\x70\xc4\xcc\x7d\xac\x8c\xd3\x88\xea\x87\x87\xf9\x7c\xc0\x9e\x14\xc3\xf1\xba\x49\x1e\x0e\xc8\x2c\x84\xe0\x90\xc6\x2e\xf4\xe6\x72\x8f\x8d\xb3\x1e\x3a\x93\x7d\xc4\x62\x13\x4e\x06\x18\x9e\x50\x4d\xb4\x29\x50\x2d\x5c\xb9\xf3\x6e\x3b\x86\xc3\x13\xaa\x20\x04\xae\xce\x7e\xda\xa7\xcb\x2a\x1a\x77\x7c\x85\x5d\xbd\xec\xa1\xec\x0b\x96\x55\x5b\xce\xe7\xeb\xd3\xac\x27\x49\xff\xe1\x5c\x95\xe8\x4a\x34\x64\xe7\xde\xc6\xf6\x84\x1c\xe6\xeb\x73\x82\x88\xe0\x5d\x6b\x71\xad\x2c\xa2\xa1\xf5\xcf\xc9\xe3\x44\x95\x23\xc8\xed\xaa\x64\xcc\x1e\x00\x9c\xfb\x52\xa3\x60\x78\x31\x40\x78\x1a\x23\x9f\x75\x00\x24\x42\xe5\x3a\x50\xfe\xfb\x25\xa0\x9c\xaa\x72\x1e\x94\x0f\x44\xfe\x0f\x00\x00\xff\xff\x79\xc5\x02\xd8\xb2\x0b\x00\x00")

func mapGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_mapGoTmpl,
		"map.go.tmpl",
	)
}

func mapGoTmpl() (*asset, error) {
	bytes, err := mapGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "map.go.tmpl", size: 2994, mode: os.FileMode(436), modTime: time.Unix(1498663970, 0)}
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
	"map.go.tmpl": mapGoTmpl,
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
	"map.go.tmpl": &bintree{mapGoTmpl, map[string]*bintree{}},
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

