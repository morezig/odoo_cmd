// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/css/app.css
// assets/done.html
// assets/index.html
// assets/js/app.js
package bindata

import (
	"github.com/elazarl/go-bindata-assetfs"
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsCssAppCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func assetsCssAppCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssAppCss,
		"assets/css/app.css",
	)
}

func assetsCssAppCss() (*asset, error) {
	bytes, err := assetsCssAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/app.css", size: 0, mode: os.FileMode(438), modTime: time.Unix(1603875159, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDoneHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\xb1\x8e\xc2\x30\x10\x44\xeb\x58\xf2\x3f\xec\x6d\x93\xea\xce\x1f\x70\x6b\x17\x04\x6a\x28\x22\x21\x4a\xc7\x31\x38\x8a\x63\x47\xb0\x41\x0a\x5f\x8f\x02\x41\x4c\x33\x9a\xd7\xbc\xa1\x9f\xed\xbe\xaa\x4f\x87\x1d\x04\x1e\xa2\x91\x82\x96\x86\x68\xd3\x45\xe3\x23\xfc\xd6\x47\x34\x52\x14\x14\xbc\x6d\x8d\x14\x52\x00\x00\x90\x5a\x67\x41\x4d\x6e\x67\xf3\xa6\x4b\x6a\x7b\xeb\xa1\xca\xc3\x18\x3d\xfb\xf6\xcb\xa9\x4b\xe3\xc4\xc0\xf3\xe8\x35\x36\x13\x73\x4e\x08\x39\xb9\xd8\xb9\x5e\x63\xcc\xce\x72\x97\xd3\x5f\xb8\xfa\xb3\x2e\x55\xf9\x8f\x70\xb7\x71\xf2\x1a\x37\xd6\xf5\x08\xca\x7c\xc4\xab\x8f\xd4\xeb\xee\x33\x00\x00\xff\xff\xc0\x94\x25\x96\xbe\x00\x00\x00")

func assetsDoneHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDoneHtml,
		"assets/done.html",
	)
}

func assetsDoneHtml() (*asset, error) {
	bytes, err := assetsDoneHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/done.html", size: 190, mode: os.FileMode(438), modTime: time.Unix(1603878806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x4f\x6b\xe2\x40\x18\xc6\xef\x82\xdf\x61\x9c\x8b\xa7\x6c\xee\x6e\x92\xcb\xee\x9e\x15\x56\x58\xf6\xb4\x4c\xfe\x68\xc2\x26\x99\x30\x99\x28\xee\x69\x85\x65\x5d\xc9\x4a\x0f\xc5\x42\x8b\x20\xf6\x50\x2c\x2d\xd4\x5b\x8b\xd8\x6f\xe3\x24\xf1\x5b\x94\x71\xd4\x46\x3d\xd5\x9e\x32\x6f\x98\xf9\x3d\xef\xc3\xf3\xbe\x4a\xe9\x73\xf5\x53\xfd\x7b\xed\x0b\xb0\xa9\xe7\x6a\xc5\x82\xc2\xbf\xc0\x45\x7e\x53\x85\xbf\x6c\xa9\xfe\x0d\x6a\xc5\x02\xff\x6d\x21\x53\x9c\xe4\xd7\xa3\x8e\xcd\x8e\x56\x2c\x00\x00\x80\x52\x92\x24\xa0\x38\x7e\x10\x51\x40\x3b\x81\xa5\x42\x3d\xa2\x14\xfb\x10\x60\xdf\x70\x1d\xe3\xa7\x0a\x5d\x6c\x20\xea\x60\xff\x83\x4d\xac\x86\x5a\x96\xb1\x89\xb1\xdc\x74\x68\x10\xb9\x6e\xf9\x23\x04\x2d\xe4\x46\x96\x0a\x93\xb8\x9f\x5c\xcc\xd2\xfe\xbf\x64\x74\x9f\x4e\x63\xb6\x38\x83\x40\xd6\x80\x24\xbd\x53\x89\x58\x21\x45\x84\xe6\x94\x56\xbd\x01\x1b\x8e\x97\x8b\x79\x32\x1a\xb0\xcb\x29\x97\x11\x12\x6b\x99\xb7\x2b\x58\xa1\x11\x1c\xe1\xd9\x60\x96\x4d\xa6\xc9\xed\xf8\xc0\x45\x03\x13\x0f\x78\x16\xb5\xb1\xa9\xc2\x5a\xf5\x6b\x1d\x02\x64\x70\xa4\x0a\x05\x4d\x54\x30\xd7\x12\x7b\xf8\x9b\x5e\x2f\x2a\x7b\x9d\x05\x28\x0c\xdb\x98\x98\x10\xf8\xc8\xe3\x75\xdb\xe4\x42\x8a\x4e\x72\x0f\xb3\xbb\x78\xd5\x7d\x4a\xce\x7b\xd9\x4d\x97\x8d\x1f\xb3\xc9\xff\xf4\xea\x0f\x8b\x87\xcb\xe7\x51\x65\xff\xe6\x1e\x9b\x20\xd3\xc1\x5b\x70\x48\x5a\x3f\x90\x41\x77\xee\xb8\x59\xa8\x1d\x78\x3c\x05\xc4\xcd\x6e\x41\xbb\x2c\x4e\x01\x35\x1d\x0a\xb5\xe3\xe9\x11\x06\x73\x40\x11\xe6\x86\x18\x46\xba\xe7\xd0\x1d\x52\x54\xda\xea\x77\x97\xf5\xe6\x8a\x2c\xae\x6e\x23\x93\x79\x66\x7c\x47\xe4\xcd\xe0\xaf\xb7\x81\xef\xcd\x4b\x00\x00\x00\xff\xff\xa1\x60\xc8\xa4\x47\x03\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 839, mode: os.FileMode(438), modTime: time.Unix(1620798503, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func assetsJsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsAppJs,
		"assets/js/app.js",
	)
}

func assetsJsAppJs() (*asset, error) {
	bytes, err := assetsJsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/app.js", size: 0, mode: os.FileMode(438), modTime: time.Unix(1603875153, 0)}
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
	"assets/css/app.css": assetsCssAppCss,
	"assets/done.html":   assetsDoneHtml,
	"assets/index.html":  assetsIndexHtml,
	"assets/js/app.js":   assetsJsAppJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"app.css": &bintree{assetsCssAppCss, map[string]*bintree{}},
		}},
		"done.html":  &bintree{assetsDoneHtml, map[string]*bintree{}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{assetsJsAppJs, map[string]*bintree{}},
		}},
	}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
