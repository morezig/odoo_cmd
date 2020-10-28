// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/css/app.css
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

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xc1\x4a\xc5\x30\x10\x45\xd7\x2f\x90\x7f\x18\x67\xf3\x56\x9a\x0f\x30\xc9\x46\xc5\xe5\x2b\x52\x10\x97\x69\x13\x9b\xe0\x98\x29\x75\x22\xd4\xaf\x17\xdb\x82\xee\x5d\x0d\x73\x2e\x1c\xb8\xd7\x5e\xdd\x5f\xee\xfa\x97\xee\x01\xb2\xbc\x93\xd7\xca\xfe\x5c\xa0\x50\x27\x87\x5f\xf9\xba\x7f\x46\xaf\xd5\xc9\xe6\x14\xa2\xd7\x4a\x2b\x00\x00\x6b\x8e\xf7\x64\x07\x8e\xab\xdf\xe9\x96\x94\x3a\x37\x01\x59\xe7\xe4\x70\x68\x22\x5c\x11\xb8\x8e\x54\xc6\x37\x87\xc4\x63\x90\xc2\xf5\x26\x2f\xe9\xd5\x9d\x0d\x47\x66\x33\x15\x99\x1b\xd1\xf9\x16\xe1\x33\x50\x4b\x0e\x1f\x8b\x40\xd7\x88\x10\xcc\xff\xd4\x4b\xfa\x90\xb0\xc8\x1f\xf5\xd3\x4e\xe0\x12\x99\x7f\xf5\xd6\x1c\x35\xac\xd9\x56\xf8\x0e\x00\x00\xff\xff\xf2\xf1\xd1\x9a\x15\x01\x00\x00")

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

	info := bindataFileInfo{name: "assets/index.html", size: 277, mode: os.FileMode(438), modTime: time.Unix(1603878614, 0)}
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
