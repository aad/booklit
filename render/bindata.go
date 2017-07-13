// Code generated by go-bindata.
// sources:
// render/html/block.tmpl
// render/html/bold.tmpl
// render/html/code-block.tmpl
// render/html/code-inline.tmpl
// render/html/element.tmpl
// render/html/image.tmpl
// render/html/inset.tmpl
// render/html/italic.tmpl
// render/html/larger.tmpl
// render/html/link.tmpl
// render/html/list.tmpl
// render/html/note.tmpl
// render/html/page.tmpl
// render/html/paragraph.tmpl
// render/html/preformatted.tmpl
// render/html/reference.tmpl
// render/html/section.tmpl
// render/html/sequence.tmpl
// render/html/smaller.tmpl
// render/html/strike.tmpl
// render/html/string.tmpl
// render/html/subscript.tmpl
// render/html/superscript.tmpl
// render/html/target.tmpl
// render/html/toc.tmpl
// DO NOT EDIT!

package render

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

var _renderHtmlBlockTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xc9\x2c\x53\x48\xce\x49\x2c\x2e\xb6\x55\xaa\xae\xd6\x73\x06\xb1\x6a\x6b\x95\xec\x40\xec\xfc\xbc\x92\xd4\xbc\x12\x85\x1a\x85\xa2\xd4\xbc\x94\xd4\xa2\xda\x5a\x1b\xfd\x94\xcc\x32\x3b\x2e\x40\x00\x00\x00\xff\xff\xa6\x01\x03\x65\x34\x00\x00\x00")

func renderHtmlBlockTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlBlockTmpl,
		"render/html/block.tmpl",
	)
}

func renderHtmlBlockTmpl() (*asset, error) {
	bytes, err := renderHtmlBlockTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/block.tmpl", size: 52, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlBoldTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x29\xca\xcf\x4b\xb7\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x87\x0a\x73\x01\x02\x00\x00\xff\xff\x40\x12\x1c\xc4\x20\x00\x00\x00")

func renderHtmlBoldTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlBoldTmpl,
		"render/html/bold.tmpl",
	)
}

func renderHtmlBoldTmpl() (*asset, error) {
	bytes, err := renderHtmlBoldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/bold.tmpl", size: 32, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlCodeBlockTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x28\x4a\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x89\x71\x01\x02\x00\x00\xff\xff\x76\xaa\x4d\xfe\x1a\x00\x00\x00")

func renderHtmlCodeBlockTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlCodeBlockTmpl,
		"render/html/code-block.tmpl",
	)
}

func renderHtmlCodeBlockTmpl() (*asset, error) {
	bytes, err := renderHtmlCodeBlockTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/code-block.tmpl", size: 26, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlCodeInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xce\x4f\x49\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x0b\x72\x01\x02\x00\x00\xff\xff\x25\x74\x64\xb1\x1c\x00\x00\x00")

func renderHtmlCodeInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlCodeInlineTmpl,
		"render/html/code-inline.tmpl",
	)
}

func renderHtmlCodeInlineTmpl() (*asset, error) {
	bytes, err := renderHtmlCodeInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/code-inline.tmpl", size: 28, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlElementTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x48\xcc\x53\x48\xce\x49\x2c\x2e\xb6\x55\xaa\xae\xd6\x73\x06\xb1\x6a\x6b\x95\xec\x40\xec\xfc\xbc\x92\xd4\xbc\x12\x85\x1a\x85\xa2\xd4\xbc\x94\xd4\xa2\xda\x5a\x1b\x7d\x90\x72\x3b\x2e\x40\x00\x00\x00\xff\xff\x34\x0f\x2e\x87\x36\x00\x00\x00")

func renderHtmlElementTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlElementTmpl,
		"render/html/element.tmpl",
	)
}

func renderHtmlElementTmpl() (*asset, error) {
	bytes, err := renderHtmlElementTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/element.tmpl", size: 54, mode: os.FileMode(420), modTime: time.Unix(1499948718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlImageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\xcc\x4d\x57\x28\x2e\x4a\xb6\x55\xaa\xae\xd6\x0b\x48\x2c\xc9\xa8\xad\x55\x52\x48\xcc\x29\x01\xf3\x5d\x52\x8b\x93\x8b\x32\x0b\x4a\x32\xf3\xf3\x40\xc2\xfa\x76\x5c\x80\x00\x00\x00\xff\xff\x6c\xbe\xfe\x41\x2f\x00\x00\x00")

func renderHtmlImageTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlImageTmpl,
		"render/html/image.tmpl",
	)
}

func renderHtmlImageTmpl() (*asset, error) {
	bytes, err := renderHtmlImageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/image.tmpl", size: 47, mode: os.FileMode(420), modTime: time.Unix(1499948718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlInsetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xc9\x2c\x53\x28\x2e\xa9\xcc\x49\xb5\x55\xca\x4d\x2c\x4a\xcf\xcc\xb3\x52\x30\x50\x30\x4a\xcd\x55\x30\x4c\xcd\x55\x52\x48\xce\x49\x2c\x2e\xb6\x55\xca\xcc\x2b\x4e\x2d\x51\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xc9\x2c\xb3\xe3\x02\x04\x00\x00\xff\xff\xaa\x57\xac\x9f\x42\x00\x00\x00")

func renderHtmlInsetTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlInsetTmpl,
		"render/html/inset.tmpl",
	)
}

func renderHtmlInsetTmpl() (*asset, error) {
	bytes, err := renderHtmlInsetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/inset.tmpl", size: 66, mode: os.FileMode(420), modTime: time.Unix(1499974064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlItalicTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xcd\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xcd\xb5\xe3\x02\x04\x00\x00\xff\xff\x99\x57\x43\xb8\x18\x00\x00\x00")

func renderHtmlItalicTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlItalicTmpl,
		"render/html/italic.tmpl",
	)
}

func renderHtmlItalicTmpl() (*asset, error) {
	bytes, err := renderHtmlItalicTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/italic.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlLargerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x48\xcc\x53\x28\x2e\xa9\xcc\x49\xb5\x55\x4a\xcb\xcf\x2b\xd1\x2d\xce\xac\x4a\xb5\x52\x30\x34\x32\x50\x55\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x29\xb6\xe3\x02\x04\x00\x00\xff\xff\xc0\x67\x23\x56\x34\x00\x00\x00")

func renderHtmlLargerTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlLargerTmpl,
		"render/html/larger.tmpl",
	)
}

func renderHtmlLargerTmpl() (*asset, error) {
	bytes, err := renderHtmlLargerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/larger.tmpl", size: 52, mode: os.FileMode(420), modTime: time.Unix(1499971478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlLinkTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x54\xc8\x28\x4a\x4d\xb3\x55\xaa\xae\xd6\x0b\x49\x2c\x4a\x4f\x2d\xa9\xad\x55\xb2\xab\xae\xd6\x73\xce\xcf\x2b\x49\xcd\x2b\x51\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xb4\xe3\x02\x04\x00\x00\xff\xff\x30\x04\xce\x2f\x30\x00\x00\x00")

func renderHtmlLinkTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlLinkTmpl,
		"render/html/link.tmpl",
	)
}

func renderHtmlLinkTmpl() (*asset, error) {
	bytes, err := renderHtmlLinkTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/link.tmpl", size: 48, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlListTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x31\x0a\x82\x31\x10\x44\xe1\x3e\xa7\xd8\x13\x24\x17\x08\xe9\xad\x3c\x83\xb0\xa3\x04\xd6\x08\xfb\x6b\x35\xce\xdd\x05\x1b\x63\xf7\x9a\xf7\x91\xf3\x6a\xf5\x9c\x8e\x84\x4b\xfd\x11\x83\x44\x1c\x90\xfa\xeb\xdb\xcb\xa5\x42\xe6\x65\xdd\x60\xf5\xf4\xc4\xfd\x90\x8a\x59\x8f\x39\xc8\x6a\x6f\x4b\x2c\x47\x4a\xbd\xc5\x1c\xe5\xb7\xfc\xcb\x6d\xa7\xdb\x66\x7f\x02\x00\x00\xff\xff\x26\x88\xac\xef\x83\x00\x00\x00")

func renderHtmlListTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlListTmpl,
		"render/html/list.tmpl",
	)
}

func renderHtmlListTmpl() (*asset, error) {
	bytes, err := renderHtmlListTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/list.tmpl", size: 131, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlNoteTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xca\xc9\x4f\xce\x2e\x2c\xcd\x2f\x49\x55\x48\xce\x49\x2c\x2e\xb6\x55\xca\xcb\x2f\x49\x55\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x47\xa8\xb3\xe3\x02\x04\x00\x00\xff\xff\xcb\x2d\x7f\xbd\x35\x00\x00\x00")

func renderHtmlNoteTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlNoteTmpl,
		"render/html/note.tmpl",
	)
}

func renderHtmlNoteTmpl() (*asset, error) {
	bytes, err := renderHtmlNoteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/note.tmpl", size: 53, mode: os.FileMode(420), modTime: time.Unix(1499974280, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlPageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\x41\x4f\xc5\x20\x10\x84\xef\xfe\x8a\x95\x7b\xe1\x6a\x22\xed\x45\x3d\x6b\x62\x2f\x1e\xb1\x5d\x85\x04\x28\xb6\x53\xf3\x1a\xc2\x7f\x7f\xe1\xd1\x77\xda\xd9\xec\xce\x7c\xa3\x1f\x5f\xdf\x5f\xc6\xaf\x8f\x37\xb2\x08\x7e\x78\xd0\x6d\x10\x69\xcb\x66\xae\x82\x48\x07\x86\x21\x0b\xa4\x8e\xff\x76\xf7\xdf\x8b\x69\x89\xe0\x88\x0e\x47\x62\x41\xe7\xd6\x0b\xf0\x05\xaa\x06\x3c\xd3\x64\xcd\xba\x31\xfa\x1d\x3f\xdd\x93\x20\x75\x26\xc1\xc1\xf3\x90\xb3\x1c\xab\x90\x9f\x58\x5d\xfc\x2d\x45\xab\x76\xa8\x5c\x75\x07\xeb\xef\x65\x3e\x9a\x2f\x67\x70\x48\xde\x80\x49\x6c\x3c\xc1\x2d\x51\x22\x24\x2f\x48\x96\x72\x33\xb5\x5f\xad\x5a\xfd\x6b\x00\x00\x00\xff\xff\x50\x1e\xa4\x17\xd6\x00\x00\x00")

func renderHtmlPageTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlPageTmpl,
		"render/html/page.tmpl",
	)
}

func renderHtmlPageTmpl() (*asset, error) {
	bytes, err := renderHtmlPageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/page.tmpl", size: 214, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlParagraphTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\xb0\xab\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcc\x4b\x49\xad\xd0\x51\x50\x29\x4e\xcd\x2b\x49\xcd\x4b\x4e\x55\xb0\xb2\x55\xd0\xab\xad\xad\xae\xce\x4c\x83\x4a\xd6\xd6\x2a\x54\x57\xa7\xe6\xa5\x80\x04\x11\xca\x6a\x14\x8a\x52\xf3\x52\x52\x8b\x40\xa2\x60\x49\x1b\xfd\x02\x3b\x2e\x40\x00\x00\x00\xff\xff\xd3\x3e\x01\x8c\x5a\x00\x00\x00")

func renderHtmlParagraphTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlParagraphTmpl,
		"render/html/paragraph.tmpl",
	)
}

func renderHtmlParagraphTmpl() (*asset, error) {
	bytes, err := renderHtmlParagraphTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/paragraph.tmpl", size: 90, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlPreformattedTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcc\x4b\x49\xad\xd0\x51\x50\x29\x4e\xcd\x2b\x49\xcd\x4b\x4e\x55\xb0\xb2\x55\xd0\xab\xad\xad\xae\xce\x4c\x83\x4a\xd6\xd6\x72\x55\x57\xa7\xe6\xa5\x80\x04\x11\xca\x6a\x14\x8a\x52\xf3\x52\x52\x8b\x40\xa2\x60\x49\x2e\x40\x00\x00\x00\xff\xff\x53\xe4\x01\x6f\x53\x00\x00\x00")

func renderHtmlPreformattedTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlPreformattedTmpl,
		"render/html/preformatted.tmpl",
	)
}

func renderHtmlPreformattedTmpl() (*asset, error) {
	bytes, err := renderHtmlPreformattedTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/preformatted.tmpl", size: 83, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlReferenceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x54\xc8\x28\x4a\x4d\xb3\x55\xaa\xae\xd6\x0b\x49\x4c\x57\xa8\x51\x28\x2d\xca\xa9\xad\x55\xb2\xab\xae\xd6\x73\xc9\x2c\x2e\xc8\x49\xac\x54\xa8\x51\x28\x2e\x29\xca\x2c\x70\x2c\xad\x50\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xb4\xe3\x02\x04\x00\x00\xff\xff\x6c\xcd\x19\x93\x3e\x00\x00\x00")

func renderHtmlReferenceTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlReferenceTmpl,
		"render/html/reference.tmpl",
	)
}

func renderHtmlReferenceTmpl() (*asset, error) {
	bytes, err := renderHtmlReferenceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/reference.tmpl", size: 62, mode: os.FileMode(420), modTime: time.Unix(1499964046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSectionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x31\x6e\xc3\x30\x0c\x45\x77\x9f\x82\xf0\x1e\xfa\x02\x8a\x87\xb6\x73\x10\x20\xb9\x00\x1b\x31\xb1\x00\x89\x16\x24\x75\x08\x58\xde\xbd\xb0\x63\x77\x68\x27\x09\x9f\x9f\x8f\xcf\x4d\xaa\x13\x93\xe7\xf2\xc1\xb9\x4d\x80\x66\x70\x8b\x54\xeb\xb1\xaf\x7c\x6b\x61\x96\xc3\x6b\xdc\x8f\x8e\x40\x28\xf1\xb1\x57\xc5\x73\x09\x89\xca\xf3\x4a\x0f\x3c\x51\x62\xb3\x7e\x74\x03\x8d\x1d\x80\xea\x01\xc2\x1d\xf0\x4c\x85\xa5\xc1\xc1\xac\x03\x00\x70\x35\x93\xfc\x25\xcb\x57\xfa\x5c\xc8\xaa\x78\x5a\xbf\x66\xe0\x86\xa5\xb9\x93\x58\xfc\x86\x50\xc5\x6b\x68\x91\xe1\x1b\x0a\x8b\xe7\xb2\xe6\x6e\xf8\xef\x3f\x76\x9d\x2a\xbe\xcd\xfe\xf9\xdb\x35\x5b\xb2\x70\x07\x99\x1b\xe0\x25\xc7\xd0\x2e\x2f\x87\xba\xc1\x0b\xc9\x83\x01\xdf\xa7\x10\x7d\x61\xd9\xac\x55\x1b\xa7\x1c\xa9\x31\xec\xd2\xd8\x52\x8e\xfd\x72\x67\xdd\x63\xf1\x66\xdd\xfe\xfe\x04\x00\x00\xff\xff\xf1\x40\x4f\x43\x4e\x01\x00\x00")

func renderHtmlSectionTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSectionTmpl,
		"render/html/section.tmpl",
	)
}

func renderHtmlSectionTmpl() (*asset, error) {
	bytes, err := renderHtmlSectionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/section.tmpl", size: 334, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSequenceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\x2b\x49\xcd\x2b\x29\xae\xad\xad\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\x02\xf1\x52\xf3\x52\x6a\x6b\xb9\x00\x01\x00\x00\xff\xff\xfc\xbe\x50\x06\x29\x00\x00\x00")

func renderHtmlSequenceTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSequenceTmpl,
		"render/html/sequence.tmpl",
	)
}

func renderHtmlSequenceTmpl() (*asset, error) {
	bytes, err := renderHtmlSequenceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/sequence.tmpl", size: 41, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSmallerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x48\xcc\x53\x28\x2e\xa9\xcc\x49\xb5\x55\x4a\xcb\xcf\x2b\xd1\x2d\xce\xac\x4a\xb5\x52\xb0\x30\x50\x55\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\xa9\xb5\xe3\x02\x04\x00\x00\xff\xff\xd6\xad\x70\x53\x33\x00\x00\x00")

func renderHtmlSmallerTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSmallerTmpl,
		"render/html/smaller.tmpl",
	)
}

func renderHtmlSmallerTmpl() (*asset, error) {
	bytes, err := renderHtmlSmallerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/smaller.tmpl", size: 51, mode: os.FileMode(420), modTime: time.Unix(1499971488, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlStrikeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x48\xcc\x53\x28\x2e\xa9\xcc\x49\xb5\x55\x2a\x49\xad\x28\xd1\x4d\x49\x4d\xce\x2f\x4a\x2c\xc9\xcc\xcf\xb3\x52\xc8\xc9\xcc\x4b\xd5\x2d\xc9\x28\xca\x2f\x4d\xcf\x50\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x69\xb5\xe3\x02\x04\x00\x00\xff\xff\xd1\x8d\xb6\xdb\x42\x00\x00\x00")

func renderHtmlStrikeTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlStrikeTmpl,
		"render/html/strike.tmpl",
	)
}

func renderHtmlStrikeTmpl() (*asset, error) {
	bytes, err := renderHtmlStrikeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/strike.tmpl", size: 66, mode: os.FileMode(420), modTime: time.Unix(1499971598, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlStringTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\xd0\x0b\x2e\x29\xca\xcc\x4b\x57\xa8\xad\xe5\x02\x04\x00\x00\xff\xff\xc7\x02\x8e\x75\x0e\x00\x00\x00")

func renderHtmlStringTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlStringTmpl,
		"render/html/string.tmpl",
	)
}

func renderHtmlStringTmpl() (*asset, error) {
	bytes, err := renderHtmlStringTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/string.tmpl", size: 14, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSubscriptTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x4d\xb2\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x89\x71\x01\x02\x00\x00\xff\xff\x10\xe7\x39\x53\x1a\x00\x00\x00")

func renderHtmlSubscriptTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSubscriptTmpl,
		"render/html/subscript.tmpl",
	)
}

func renderHtmlSubscriptTmpl() (*asset, error) {
	bytes, err := renderHtmlSubscriptTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/subscript.tmpl", size: 26, mode: os.FileMode(420), modTime: time.Unix(1499971574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSuperscriptTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x2d\xb0\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x89\x71\x01\x02\x00\x00\xff\xff\xab\xc5\x7f\xa0\x1a\x00\x00\x00")

func renderHtmlSuperscriptTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSuperscriptTmpl,
		"render/html/superscript.tmpl",
	)
}

func renderHtmlSuperscriptTmpl() (*asset, error) {
	bytes, err := renderHtmlSuperscriptTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/superscript.tmpl", size: 26, mode: os.FileMode(420), modTime: time.Unix(1499971583, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlTargetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x54\xc8\x4b\xcc\x4d\xb5\x55\xaa\xae\xd6\x0b\x49\x4c\xf7\x4b\xcc\x4d\xad\xad\x55\xb2\xb3\xd1\x4f\xb4\xe3\x02\x04\x00\x00\xff\xff\x6d\x16\xe9\x86\x1c\x00\x00\x00")

func renderHtmlTargetTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlTargetTmpl,
		"render/html/target.tmpl",
	)
}

func renderHtmlTargetTmpl() (*asset, error) {
	bytes, err := renderHtmlTargetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/target.tmpl", size: 28, mode: os.FileMode(420), modTime: time.Unix(1499898549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlTocTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x41\x6a\xc3\x30\x10\x45\xf7\x3e\xc5\xc7\xab\x76\x23\x5f\x40\x35\x94\x40\x97\x4d\x17\xbe\x80\x52\x8f\x13\x81\x34\x36\x93\x51\x69\x19\xeb\xee\xc5\x6e\x4d\x56\x33\xbc\x0f\xff\x3f\xb3\x38\x21\xf0\x08\x77\xba\xc5\x34\x0a\x31\x9e\x78\x56\xb8\x73\x8e\x7a\xa0\x37\x99\xf3\x10\x2e\x89\xce\xd3\x69\x66\x25\xd6\xfb\x73\xad\x8d\xe7\xf0\xd5\x37\x80\x2f\x69\x3b\x66\x12\xf8\x4a\x8f\xa6\x5a\x1b\x00\xf0\x29\xf6\xfb\x03\xf8\x80\x9b\xd0\xf4\xd2\x9a\xb9\x0f\x89\x39\xc8\xcf\x10\xae\x58\x51\x24\xd5\xda\xf6\x66\xee\xbd\xe4\x0b\x49\xad\x30\x73\x43\xd4\x44\x58\x71\x57\x89\xcb\x6b\xf9\xc6\x0a\x21\x1e\xb7\xd8\x77\xa1\x6f\xfe\x5b\xcd\x94\xf2\x92\x82\x12\x5a\x9d\x3f\x9d\xe6\x25\xb5\x70\xc7\x7c\xf7\xb7\x6f\x46\x3c\xee\xcc\x77\x9b\xb0\xef\x76\xfd\x03\xff\x06\x00\x00\xff\xff\xc5\x44\xc0\xe1\x09\x01\x00\x00")

func renderHtmlTocTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlTocTmpl,
		"render/html/toc.tmpl",
	)
}

func renderHtmlTocTmpl() (*asset, error) {
	bytes, err := renderHtmlTocTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/toc.tmpl", size: 265, mode: os.FileMode(420), modTime: time.Unix(1499963838, 0)}
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
	"render/html/block.tmpl": renderHtmlBlockTmpl,
	"render/html/bold.tmpl": renderHtmlBoldTmpl,
	"render/html/code-block.tmpl": renderHtmlCodeBlockTmpl,
	"render/html/code-inline.tmpl": renderHtmlCodeInlineTmpl,
	"render/html/element.tmpl": renderHtmlElementTmpl,
	"render/html/image.tmpl": renderHtmlImageTmpl,
	"render/html/inset.tmpl": renderHtmlInsetTmpl,
	"render/html/italic.tmpl": renderHtmlItalicTmpl,
	"render/html/larger.tmpl": renderHtmlLargerTmpl,
	"render/html/link.tmpl": renderHtmlLinkTmpl,
	"render/html/list.tmpl": renderHtmlListTmpl,
	"render/html/note.tmpl": renderHtmlNoteTmpl,
	"render/html/page.tmpl": renderHtmlPageTmpl,
	"render/html/paragraph.tmpl": renderHtmlParagraphTmpl,
	"render/html/preformatted.tmpl": renderHtmlPreformattedTmpl,
	"render/html/reference.tmpl": renderHtmlReferenceTmpl,
	"render/html/section.tmpl": renderHtmlSectionTmpl,
	"render/html/sequence.tmpl": renderHtmlSequenceTmpl,
	"render/html/smaller.tmpl": renderHtmlSmallerTmpl,
	"render/html/strike.tmpl": renderHtmlStrikeTmpl,
	"render/html/string.tmpl": renderHtmlStringTmpl,
	"render/html/subscript.tmpl": renderHtmlSubscriptTmpl,
	"render/html/superscript.tmpl": renderHtmlSuperscriptTmpl,
	"render/html/target.tmpl": renderHtmlTargetTmpl,
	"render/html/toc.tmpl": renderHtmlTocTmpl,
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
	"render": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"block.tmpl": &bintree{renderHtmlBlockTmpl, map[string]*bintree{}},
			"bold.tmpl": &bintree{renderHtmlBoldTmpl, map[string]*bintree{}},
			"code-block.tmpl": &bintree{renderHtmlCodeBlockTmpl, map[string]*bintree{}},
			"code-inline.tmpl": &bintree{renderHtmlCodeInlineTmpl, map[string]*bintree{}},
			"element.tmpl": &bintree{renderHtmlElementTmpl, map[string]*bintree{}},
			"image.tmpl": &bintree{renderHtmlImageTmpl, map[string]*bintree{}},
			"inset.tmpl": &bintree{renderHtmlInsetTmpl, map[string]*bintree{}},
			"italic.tmpl": &bintree{renderHtmlItalicTmpl, map[string]*bintree{}},
			"larger.tmpl": &bintree{renderHtmlLargerTmpl, map[string]*bintree{}},
			"link.tmpl": &bintree{renderHtmlLinkTmpl, map[string]*bintree{}},
			"list.tmpl": &bintree{renderHtmlListTmpl, map[string]*bintree{}},
			"note.tmpl": &bintree{renderHtmlNoteTmpl, map[string]*bintree{}},
			"page.tmpl": &bintree{renderHtmlPageTmpl, map[string]*bintree{}},
			"paragraph.tmpl": &bintree{renderHtmlParagraphTmpl, map[string]*bintree{}},
			"preformatted.tmpl": &bintree{renderHtmlPreformattedTmpl, map[string]*bintree{}},
			"reference.tmpl": &bintree{renderHtmlReferenceTmpl, map[string]*bintree{}},
			"section.tmpl": &bintree{renderHtmlSectionTmpl, map[string]*bintree{}},
			"sequence.tmpl": &bintree{renderHtmlSequenceTmpl, map[string]*bintree{}},
			"smaller.tmpl": &bintree{renderHtmlSmallerTmpl, map[string]*bintree{}},
			"strike.tmpl": &bintree{renderHtmlStrikeTmpl, map[string]*bintree{}},
			"string.tmpl": &bintree{renderHtmlStringTmpl, map[string]*bintree{}},
			"subscript.tmpl": &bintree{renderHtmlSubscriptTmpl, map[string]*bintree{}},
			"superscript.tmpl": &bintree{renderHtmlSuperscriptTmpl, map[string]*bintree{}},
			"target.tmpl": &bintree{renderHtmlTargetTmpl, map[string]*bintree{}},
			"toc.tmpl": &bintree{renderHtmlTocTmpl, map[string]*bintree{}},
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

