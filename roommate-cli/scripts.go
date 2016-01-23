package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _scripts_alert_applescript = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xce\xbd\x4e\xc4\x30\x10\x04\xe0\xde\x4f\x31\x31\x4d\x90\x22\x52\x51\x53\xd3\x23\x51\x6f\x12\xff\xac\x08\xb6\xb5\xbb\x06\x21\xc4\xbb\xe3\xbb\xf2\x74\xd5\x16\xfb\x8d\x66\x1e\xa6\xb5\xab\xac\x1b\x97\xb5\x2a\xe9\x2e\xdc\xcc\xb9\x5a\x20\xbd\x60\x26\x49\xfa\xe8\x80\x83\xb5\x9d\xf4\x33\x2e\x9d\x35\x61\x8e\x2c\x6a\x60\x0b\x9f\xa8\x11\x57\x85\x6f\xb6\x0c\x63\x3b\x03\x66\x0d\x7b\x2d\xc7\x3d\xc0\xe3\x81\xc8\x17\x64\x99\xe5\xd6\x6c\xdd\xac\x16\xc5\xaf\x7f\xcf\x64\x2f\x93\x5f\xe0\xdf\x32\x95\x0f\x5d\xf0\x0a\x6a\x4d\xc2\xce\x64\x61\xc4\x9e\xfc\x1f\x12\x7f\x71\x49\xe8\x0d\x14\x2d\x08\x9e\x5d\x18\xb5\x63\xfa\x7f\x00\x00\x00\xff\xff\x41\x34\x86\x55\xd7\x00\x00\x00")

func scripts_alert_applescript_bytes() ([]byte, error) {
	return bindata_read(
		_scripts_alert_applescript,
		"scripts/alert.applescript",
	)
}

func scripts_alert_applescript() (*asset, error) {
	bytes, err := scripts_alert_applescript_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "scripts/alert.applescript", size: 215, mode: os.FileMode(493), modTime: time.Unix(1453579134, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scripts_brightness_applescript = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x51\xc1\x52\xeb\x30\x0c\xbc\xe7\x2b\xf6\xe5\x5d\xda\x03\xed\xf4\x17\x18\xb8\x33\xc0\x0f\x28\x89\x5a\x34\xe3\xda\x46\x92\xd3\xe9\xdf\x63\x27\x85\xe9\x01\x4e\xd9\x8d\x76\xa5\x95\xfc\xff\xdf\xbe\x98\xee\x07\x89\xfb\x64\x64\xa3\x4a\xf6\xae\x4b\x11\x5a\x22\x36\xa4\x27\xdb\x76\x80\x73\x08\xa0\x9c\x83\x8c\xe4\x52\xab\xfd\xdb\xd5\x9c\xcf\x78\x51\x3e\xb2\x72\x1c\xd9\xfa\xaa\x03\x68\x74\x99\xc9\x79\x21\xca\x33\x53\x35\xc6\xf1\x23\x29\xfa\x49\x2c\x07\xba\xda\xd3\xfa\x7d\xa7\xa1\x47\x3a\x22\x53\x64\xc8\x84\x7e\x4c\xe7\x5d\x1b\xc2\xbb\xfc\xd3\x76\xf7\x6d\x5a\xdb\xff\x19\xe4\x79\xe6\xe8\x37\x11\x30\x71\xb5\xe0\x70\x63\xc6\x8e\x99\x42\xe1\x36\xcd\x82\x4c\xac\x38\x34\x7c\xd2\x54\xf2\x0a\x9d\x86\x7b\x7a\x91\x38\xa5\x0b\xfa\xc7\x22\xc1\x1f\x24\xe2\x95\x5d\x22\xe1\x16\x7d\xcd\xad\xa9\xae\x6d\xbf\xde\x02\x9e\xb0\x39\x8a\x9a\x43\x5a\xad\xca\xdb\x2d\x41\x86\x58\xce\x03\xeb\xb6\x5b\xc2\x71\x9c\x96\x9d\x16\xf2\x59\xc4\xbb\xbb\x7f\x0d\xd4\x67\xf8\x0a\x00\x00\xff\xff\x80\xad\xd4\xd8\xa3\x01\x00\x00")

func scripts_brightness_applescript_bytes() ([]byte, error) {
	return bindata_read(
		_scripts_brightness_applescript,
		"scripts/brightness.applescript",
	)
}

func scripts_brightness_applescript() (*asset, error) {
	bytes, err := scripts_brightness_applescript_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "scripts/brightness.applescript", size: 419, mode: os.FileMode(493), modTime: time.Unix(1453544765, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"scripts/alert.applescript": scripts_alert_applescript,
	"scripts/brightness.applescript": scripts_brightness_applescript,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"scripts": &_bintree_t{nil, map[string]*_bintree_t{
		"alert.applescript": &_bintree_t{scripts_alert_applescript, map[string]*_bintree_t{
		}},
		"brightness.applescript": &_bintree_t{scripts_brightness_applescript, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

