// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    21280,
		modtime: 0,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
