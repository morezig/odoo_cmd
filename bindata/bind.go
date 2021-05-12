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

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd0\xb1\x4a\xc3\x31\x10\x06\xf0\xb9\x81\xbc\xc3\x79\x4b\x27\xcd\x03\x98\x64\xa9\xc5\xb1\x7f\xb4\x20\x8e\x69\x12\x9b\xe0\x99\x0b\xff\x5e\x84\xfa\xf4\x62\x5b\xd0\xc1\xad\xd3\x71\xdf\x07\xbf\xe1\xb3\x37\x0f\x9b\xd5\xf6\x75\x5a\x43\x91\x0f\xf2\x5a\xd9\x9f\x0b\x14\xda\xde\xe1\x57\xb9\xdd\xbe\xa0\xd7\x6a\x61\x4b\x0e\xc9\x6b\xa5\x15\x00\x80\x35\x97\x77\x61\x77\x9c\x8e\xfe\x9c\x9e\x9a\xda\xfa\x10\x90\x63\xcf\x0e\x77\x43\x84\x1b\x02\xb7\x48\x35\xbe\x3b\x24\x8e\x41\x2a\xb7\xbb\x32\xe7\x37\xb7\x34\x9c\x98\xcd\xbe\x4a\x1f\x44\xcb\x7b\x84\xcf\x40\x23\x3b\x7c\xac\x02\xd3\x20\x42\x30\xd7\xd1\x73\x3e\x48\x98\xe5\x0f\xfd\x74\x4e\x60\x93\x98\xaf\xe6\xf3\x21\xf6\x7f\xec\xf5\xf3\x6a\xfa\xb5\xad\xb9\x4c\x64\xcd\x69\xe1\xef\x00\x00\x00\xff\xff\x44\x0c\x9d\x83\x71\x01\x00\x00")

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

	info := bindataFileInfo{name: "assets/index.html", size: 369, mode: os.FileMode(438), modTime: time.Unix(1620795805, 0)}
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
