// Code generated by go-bindata. DO NOT EDIT.
// sources:
// client.sh (698B)
// controller.sh (983B)
// node.sh (565B)
// registry.yaml (1.01kB)
// net/calico.yaml (16.739kB)
// net/weave.yaml (6.882kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _clientSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x31\x6f\x1b\x31\x0c\x85\x77\xfd\x8a\x57\xb7\xab\x7d\xb1\xe3\x18\x29\x9c\x76\x4b\x01\x4f\xcd\xd0\xa1\xab\x4e\x47\xfb\x88\xd3\x89\x82\x48\x1b\x2e\x82\xfc\xf7\x42\x97\xa6\x70\xda\x4e\xcd\x42\xe0\x01\x4f\x1f\x1f\x45\xbe\x7f\xd7\xb4\x9c\x9a\xd6\x6b\xef\x9c\x92\x61\x4e\xc7\xb3\x20\x73\xa6\xbd\xe7\xe8\x9c\x15\x9f\x31\x33\x39\x86\x1e\x4d\x2f\x6a\x4d\x2b\x62\x1a\x0a\x67\x9b\x77\x92\x68\x86\xfb\xef\xbb\x6f\xce\xed\x1e\x3e\x2d\x3f\xae\x16\xcb\xcd\xed\xe2\xe6\x6a\x71\xed\x38\xc3\x77\x5d\xa9\x05\x1f\x1e\x77\x0f\x4f\xcd\x6a\x8d\x8e\x4e\xa0\xa4\x6b\x57\x49\xc9\x8f\x14\x2c\x42\xc9\xe6\x2f\x1a\x21\x32\x25\x73\x2e\x78\xc3\xe7\x86\x2c\x34\x2d\x97\x6e\x2a\x8b\x20\x69\x8f\xbb\xbb\xfb\xaf\x5f\x5c\x91\xa3\x51\x01\xff\x62\x6f\x9d\xcb\x45\x4c\x82\x44\x0c\x54\x12\x45\x3c\x3a\x20\x53\x51\x56\xdb\x3a\x60\xa4\x72\x20\x64\x6f\xbd\x56\xa9\xc1\x27\x18\x8f\x84\xcd\x55\xd5\x3c\x66\x29\x86\x24\x89\xaa\xa4\xf3\x24\x7d\x8c\x5b\xf7\x74\xc1\xee\xe8\xc4\x81\x26\xf6\x6b\xc2\x2b\x13\x17\x0a\x36\x99\x38\x19\x95\xbd\x0f\x84\x59\x1d\x7b\x56\xd9\xa1\xa7\x30\x20\x72\x1a\x2e\x1a\xff\xd5\xa9\x3d\x64\x0c\xb7\x3a\x7a\xad\x73\x56\x56\x94\xe0\x23\xbc\x62\xb3\xbe\x59\x5e\xd7\xb7\xd9\xab\xf2\x69\x0a\x9c\x88\x0f\x7d\x2b\x05\x17\x3b\x58\xbe\x98\x57\xcf\x8e\xb3\xa1\x97\x0c\xa5\xb8\xff\xb3\xf3\xef\x89\x9f\x3f\xe0\x1f\x41\x92\x74\xf4\x7f\x31\x56\x6f\x88\x51\x57\xad\x3f\xd4\x68\xac\x87\x52\x48\xcd\x17\x43\x3d\x06\xf7\x33\x00\x00\xff\xff\x86\x8f\xe8\xe4\xba\x02\x00\x00"

func clientShBytes() ([]byte, error) {
	return bindataRead(
		_clientSh,
		"client.sh",
	)
}

func clientSh() (*asset, error) {
	bytes, err := clientShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1d, 0xff, 0x78, 0x58, 0xc, 0x94, 0x36, 0xb8, 0x14, 0x8, 0xfb, 0xf7, 0xb8, 0xb8, 0x84, 0xd3, 0x64, 0x30, 0xf9, 0x18, 0xca, 0x68, 0x9, 0xfb, 0xe6, 0x35, 0x28, 0x54, 0x45, 0xdf, 0xe0, 0xcd}}
	return a, nil
}

var _controllerSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6e\xdb\x30\x0c\xbd\xeb\x2b\x38\xad\x87\xed\x20\x29\xce\x8a\x6d\x30\xda\x02\x5d\x90\x0e\x45\x81\x14\x68\xb6\xde\x15\x8b\x89\x85\xc8\x92\x20\xd1\x5e\x8b\x61\xff\x3e\xc8\xf1\x1a\x74\xed\xa5\xba\xd8\x20\xf9\x1e\xf9\xc8\xf7\xfe\x9d\xda\x58\xaf\x36\x3a\xb7\x8c\x65\x24\x10\xd8\x3f\x04\x88\x36\xe2\x56\x5b\xc7\xd8\xfd\xf2\x6e\x7d\x7d\xbb\x3a\xe7\x43\x25\xab\xb9\x9c\x73\xc6\x1a\x4d\x70\xa1\xa8\x8b\x6a\xdf\x6f\x50\x9b\x4e\x36\xc1\x6f\xe1\xec\x6c\x79\x7b\xc5\x74\xb4\xf7\x98\xb2\x0d\xbe\x86\x7f\xe9\xfd\xd7\x2c\x6d\x50\x43\xa5\x5d\x6c\xf5\x27\xb6\xb7\xde\xd4\x70\xed\x2d\x2d\x82\xdf\xda\x5d\x9f\x34\xd9\xe0\xd9\x26\x04\xca\x94\x74\xfc\x11\xf6\xe8\x73\xcd\x04\x50\xf9\xab\x81\xcf\xc6\x27\x67\xff\x3d\xce\x00\x88\x5c\x0d\x7c\x7e\xda\xf2\xd2\x7c\xe9\x4d\x0c\xd6\x53\xcd\x00\xb4\x19\x30\x91\xcd\x78\x69\x4c\xc2\x9c\x6b\x38\xf9\xd0\xa2\x36\x20\x2a\x50\x6d\xc8\xa4\x6c\xfc\xc8\x7c\x30\x78\x87\x3b\x5b\x3a\x97\x31\x0a\xb2\x4c\xee\x90\x96\x0f\x94\xf4\x65\xda\xe5\x12\x03\x28\x95\xc2\xc6\xd7\x69\x84\x10\x6f\x10\xbf\x70\x7d\x26\x4c\xcf\xf5\x7b\xa4\x5f\x21\xed\xad\xdf\x95\x7e\x31\x98\x75\xbf\xf1\x48\x35\xf0\x6a\x26\x4f\xe7\x72\x26\x67\xaa\xfa\xcc\x59\xa1\x4e\x1e\x09\xf3\x53\x37\x7e\xf2\x7b\xba\xd4\x1f\xce\x9a\x03\xf9\x4a\x77\x58\x03\x1f\x6c\xa2\x5e\x17\xcc\xb8\xa0\x35\xa6\x01\xd3\x02\x13\xad\x2f\x57\xe3\x8e\x79\x35\xff\x52\xb8\x65\xc5\x5f\x55\xe1\x90\xc6\x0b\xdb\xdd\x51\xcc\x06\x49\x57\x93\x96\x9b\x43\xcd\x73\x2d\x09\x73\x70\x43\x89\xd5\xa0\x52\xef\x55\x7e\xcc\x84\x9d\x51\x87\x04\x4e\xdf\x91\x98\x15\xe3\x60\xd3\x06\xe0\xf8\x10\x43\x22\xb8\xf9\xf9\x6d\xb9\xb8\x5d\x5d\x5d\x7f\x3f\x57\x48\x8d\x3a\x2a\x56\xda\x74\xd6\x8f\x30\x0e\x17\x63\x32\xa6\xb0\xb5\x0e\xa5\x51\x65\xbc\xdc\xb2\x69\xf5\x60\xbd\x25\x10\xe2\x30\xfb\xf9\x0b\xc7\xb2\xb7\x34\x1b\x49\x1b\x72\x40\xda\x7a\x1a\xad\x90\x41\x08\xed\xdc\xc1\x16\x29\x38\x94\x47\x64\xd9\x52\xa7\xcb\x15\xc4\x13\x52\xc7\xe8\x1e\x41\x6c\x27\xd3\x68\x63\x82\xcf\xf2\x51\x77\x8e\x35\x11\x4e\x8e\x63\x4c\x05\x23\xee\xb0\x77\xea\x22\xeb\x86\x23\xf6\x79\xea\x45\x90\xfd\x0d\x00\x00\xff\xff\xbe\xd8\x64\xca\xd7\x03\x00\x00"

func controllerShBytes() ([]byte, error) {
	return bindataRead(
		_controllerSh,
		"controller.sh",
	)
}

func controllerSh() (*asset, error) {
	bytes, err := controllerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0xae, 0xfd, 0x76, 0x94, 0x28, 0x26, 0xf0, 0xde, 0xc3, 0xc5, 0x17, 0xbb, 0x45, 0x7a, 0x6b, 0xf7, 0xe9, 0xd1, 0x35, 0xc3, 0xbb, 0xe9, 0x78, 0x85, 0xd4, 0xcf, 0x3a, 0x7f, 0xb5, 0x37, 0x4b}}
	return a, nil
}

var _nodeSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x51\xab\xd3\x30\x14\xc0\xf1\xf7\x7c\x8a\x63\xef\xe0\xde\x3d\xa4\xdd\xd8\x10\x09\x9b\x50\x46\x45\x65\xae\xb2\xcd\xbd\x4a\x96\x9e\xb6\xc7\x76\x49\x48\xd2\x32\x51\xbf\xbb\xb4\x53\x91\xa9\xe0\x79\x4c\xfe\x84\x9c\xdf\xc3\xb3\xe4\x4c\x3a\x39\x4b\x5f\x33\xe6\x31\x00\xc7\xee\x6a\xc0\x92\xc5\x52\x52\xcb\xd8\x26\xdf\x1d\xf7\xf9\x76\x9b\xed\x3f\xbe\xce\x0f\xc7\x5d\xfa\x2e\x5b\x47\x93\xa7\xda\xf8\xa0\xe5\x05\xe1\x2b\xa8\x2e\x00\x2f\xe7\xc0\x8b\x47\xfe\x38\xe5\xca\xe8\xe0\x4c\xdb\xa2\x8b\x18\x53\x32\xc0\xcb\x24\x5c\x6c\xd2\x74\x67\x94\xc5\x25\x56\x46\x97\xb0\x5a\x65\xf9\x2b\x26\x2d\x9d\xd0\x79\x32\x5a\xc0\xcf\xeb\xe6\x85\x8f\xc9\x24\xfd\x5c\xb6\xb6\x96\x0b\xd6\x90\x2e\x04\xbc\x35\xa4\x37\x46\x97\x54\x75\x4e\x06\x32\x9a\x05\xd3\xa0\x16\x10\xcd\xc6\x89\x67\x77\x13\xb1\x82\xbc\x32\x3d\xba\xcf\xc7\xa1\xfc\xa0\xbd\x2c\xf1\xd0\x90\xdd\xa4\x27\x74\x54\x92\x1a\xdf\x11\x10\x5c\x87\x77\x71\xfa\xfe\xcd\x01\x5d\x8f\xce\x0b\xc6\x61\xf2\xe5\x2f\x04\xdf\xe2\xd6\x28\xd9\x8a\xe7\xcb\xe5\x82\x69\x53\xe0\x1e\x2b\xf2\xe1\xf6\x39\xc1\x60\xdc\xa7\xc5\x90\x5d\x83\x93\xa9\xab\xfc\x70\x06\x30\x94\x9c\xac\x80\xc9\x53\x8d\xb2\x00\x3e\x87\x64\x90\x4c\xc8\x4e\xd9\x40\xf2\x83\x01\x3e\x19\xd2\xc0\x07\xcb\x92\xaa\xf5\x1f\x80\xec\xe1\x2e\x1c\x39\xd6\xff\xc0\x00\xce\x7f\x2d\x78\x2b\x79\x37\x7a\x70\xdf\x90\xe5\x4a\xf2\xfe\x37\x92\xff\x58\xf8\x7b\x00\x00\x00\xff\xff\xe8\xbc\x4b\x30\x35\x02\x00\x00"

func nodeShBytes() ([]byte, error) {
	return bindataRead(
		_nodeSh,
		"node.sh",
	)
}

func nodeSh() (*asset, error) {
	bytes, err := nodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5c, 0x7c, 0x73, 0x74, 0xbf, 0x12, 0xb0, 0x21, 0xb, 0xa3, 0x42, 0x32, 0xe, 0xdf, 0xff, 0x80, 0xc8, 0x2d, 0x83, 0x47, 0xa6, 0x22, 0x2d, 0x7e, 0x94, 0x14, 0x29, 0xb, 0x5c, 0xf6, 0x78, 0xb4}}
	return a, nil
}

var _registryYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\x4d\x6b\xdb\x40\x10\xbd\xef\xaf\x18\x72\x57\x64\xb7\xf4\xb2\xb7\x10\xa9\xa9\xa1\xa9\x8d\x24\x0a\x3e\x99\xb1\x3c\xa4\x4b\xf6\x8b\xdd\xb1\xa8\x28\xfd\xef\x65\x1b\x4b\x91\x22\xb7\xb7\xcc\x49\x3c\xcd\x7b\xf3\x66\xf6\x65\x59\x26\xd0\xab\xef\x14\xa2\x72\x56\x02\xfd\x64\xb2\xe9\x33\xe6\xdd\xfa\x48\x8c\x6b\xf1\xac\xec\x49\x42\x41\x5e\xbb\xde\x90\x65\x61\x88\xf1\x84\x8c\x52\x00\x58\x34\x24\x21\xd0\x93\x8a\x1c\xfa\x0b\x10\x3d\xb6\x24\xe1\xf9\x7c\xa4\x2c\xf6\x91\xc9\x08\x00\x8d\x47\xd2\x31\x71\x00\xd0\xfb\x09\x29\x7a\x6a\x13\x1e\xc8\x6b\xd5\x62\x94\xb0\x16\x00\x4c\xc6\x6b\x64\x7a\x61\x4c\x67\xa6\x9a\xaa\x2d\x15\x13\x32\xa8\xa6\xb2\xee\x44\x35\x69\x6a\xd9\x85\x57\x4e\x42\xb3\xe0\x34\xdd\x26\xa7\xc1\x12\x53\xbc\x55\x2e\x37\x18\x99\x82\x84\x9b\x9b\x4b\x6b\xeb\x2c\xa3\xb2\x14\xc6\x81\xd9\x72\xf1\x97\x52\x06\x9f\x26\xb8\xfc\x30\xfe\x21\xdb\xbd\x8e\x1e\xf8\x55\xf9\xb0\xa9\x9b\x6a\x7f\xf8\xd2\x34\xbb\xc3\x5d\x51\x54\x63\x0b\x40\x87\xfa\x4c\x12\xe4\xa7\xd5\x6a\xf5\x6f\x66\xdd\x6c\xab\xbb\x87\xf2\xf0\x79\xf3\xb5\xac\xf7\x75\x53\x3e\x1e\xaa\xed\xb6\x29\x36\x55\x79\xdf\x6c\xab\xfd\x52\x30\xef\x30\xe4\x5a\x1d\xf3\x85\xfb\xce\xe9\xb3\xa1\x47\x77\xb6\x1c\x97\x66\xff\xee\x96\x45\x76\x81\x26\xa2\x26\x75\xef\x90\x7f\xfc\x4f\xd8\xbb\x30\x57\x1c\x4f\xba\x73\x81\x25\xcc\x56\xbc\x92\xaa\x51\x27\x38\x76\xad\xd3\x12\x9a\xfb\x9d\x98\x9a\x5e\x3c\xcd\x35\xb7\x64\x3c\xf7\x85\x0a\x12\x7e\xfd\x16\x6f\xa3\xdf\x0d\x51\xaf\x29\x74\xaa\xa5\x77\xcb\x39\xf7\x9e\x24\x7c\x73\x27\x4a\xcb\x8b\xc9\x75\xae\xc6\xca\xcf\x2f\x64\x2f\x3c\x09\x1f\x57\x03\xf6\xf6\x2c\x71\x96\xf5\xb9\x8b\x3f\x01\x00\x00\xff\xff\x44\x90\xb8\xd2\xf2\x03\x00\x00"

func registryYamlBytes() ([]byte, error) {
	return bindataRead(
		_registryYaml,
		"registry.yaml",
	)
}

func registryYaml() (*asset, error) {
	bytes, err := registryYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registry.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0x66, 0x9, 0xb0, 0x75, 0xb4, 0x7c, 0x58, 0xcc, 0x66, 0x99, 0x31, 0xd8, 0xaa, 0x58, 0x8b, 0xc1, 0xb6, 0xca, 0x1f, 0x83, 0xc0, 0xc1, 0x71, 0x9c, 0xec, 0xf1, 0xb1, 0x8f, 0x2e, 0x50, 0x5c}}
	return a, nil
}

var _netCalicoYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5b\xed\x73\x1a\x39\x93\xff\xee\xbf\xa2\x0b\x7f\xc8\xdd\x53\x1e\xb0\x37\xd9\x5c\xc2\x7d\xc2\x98\x38\x54\x6c\xa0\x0c\xce\xde\xd6\xd6\x15\x2b\x66\x1a\xd0\x59\x23\xcd\x23\x69\x70\xd8\xbb\xe7\x7f\xbf\xd2\xcb\xbc\x32\x83\xe3\xec\xc6\x59\x3e\xd9\x23\xa9\xbb\xf5\x53\xab\xdf\x24\x9d\xc2\x90\x30\x1a\x0a\xf8\x8c\x52\x51\xc1\x61\xf7\xba\xfb\xba\x7b\x71\x72\x0a\x5b\xad\x13\xd5\xef\xf5\x22\x11\xaa\x6e\x22\xc5\xff\x60\xa8\x43\xdb\xb7\x2b\xe4\xa6\x67\xfa\xf5\x24\x32\x24\x0a\xd5\xa9\x1f\xf5\x40\x79\xd4\x87\x21\x4b\x95\x46\x79\x27\x18\x9e\x90\x84\x7a\xca\x7d\x90\x2b\x12\x76\x49\xaa\xb7\x42\xd2\x3f\x88\xa6\x82\x77\x1f\xde\xa9\x2e\x15\xbd\xdd\xc5\x0a\x35\xb9\x38\x89\x51\x93\x88\x68\xd2\x3f\x01\xe0\x24\xc6\x3e\x38\x8e\x01\x17\x11\x9e\xc8\x94\xa1\x32\x4d\x01\x90\x84\x5e\x4b\x91\x26\xaa\x0f\xbf\x75\x3a\xff\x7d\x02\x00\x20\x51\x89\x54\x86\xae\x8b\xf9\x05\x96\x86\x4a\x48\x88\x2a\xff\xa4\x50\xee\x68\x88\x24\x0c\x45\xca\xb5\xfb\xbe\x43\xb9\x2a\x8d\xda\xa0\xce\xff\x66\x54\x15\xff\x3c\x12\x1d\x6e\x9f\xc3\x3f\x11\x91\xea\x29\x4d\x74\xda\xc8\x28\xf9\x16\x7a\xdf\x57\x62\x0f\xcf\x11\x2e\x5f\x4d\x0a\x79\x94\x08\x7a\x14\xe3\xaf\x5f\x48\x11\x1d\x93\xa9\x61\xe6\x69\x12\x11\x8d\x4f\x01\x81\x5f\x34\x72\xa3\x9d\xea\x08\x6f\xd4\x8f\x42\x3e\x24\x82\xd1\x90\x3e\x53\x8a\x16\xb6\x9e\x24\xe5\x1b\xbf\x03\xfe\x14\xf7\x8c\x49\x85\x7f\x8d\x63\x28\xa3\xc3\x4d\xdc\xce\x75\xc3\xc4\x8a\xb0\x35\x32\xfa\x25\x14\x7c\x4d\x37\xc5\x06\x2a\x7d\x4c\xa5\xdd\xc5\x45\xdb\x6a\x93\x24\x88\x52\xd5\xe8\xac\x36\x49\x9d\x4a\xfe\xe9\x80\x06\x4d\x12\x21\x58\x9d\x44\x13\x0c\x07\xad\x0a\xb5\x3a\x0a\x9c\xf9\x1e\x3a\xeb\x44\xf9\x5a\xc8\xb8\xc6\x7b\x2b\x94\x3e\xaa\xb6\xa1\xc4\xb2\x56\x3d\x4f\xfb\x4e\x82\x20\x38\x79\xa6\x45\x3c\xb0\xa8\x97\x94\x47\x94\x6f\x9e\x32\x95\x82\xe1\x1d\xae\x4d\x63\xa6\x06\x47\xb8\x9d\x00\x1c\x5a\xee\x26\xb2\x2a\x5d\x19\x0d\x52\xfd\x93\xc0\x8f\x98\x3b\x7b\x31\x70\xe6\xb4\x71\x10\x14\x66\xb8\x0f\x0f\xe9\x0a\x03\xb5\x57\x1a\x63\x0b\xc7\x5f\xeb\x7f\x4e\x61\xb1\xa5\x0a\x62\xc2\xe9\x1a\x95\x06\xca\x43\x96\x46\xa8\x40\x6f\x11\xd6\x82\x31\xf1\x48\xf9\x06\x42\x11\x27\x82\x23\xd7\x66\x79\xed\xee\xef\x9f\x9c\x02\x78\xa9\x7b\x46\xea\x7e\x4e\x31\xff\x1c\x72\x9a\x7d\xcd\x18\x0d\xad\x0a\xdf\x92\x04\xa8\x82\x54\x61\x04\x5a\x40\xa6\xd7\x08\x04\x14\xb2\x75\x60\xb5\x2a\xca\x26\x4a\xb9\xd2\x84\x31\x87\x7f\xb6\xbe\x19\x9d\x8a\x76\xec\x8e\x38\x44\xc7\xa4\x1d\xdb\x6c\xd0\x29\x2c\x04\x20\x27\x2b\x86\xb0\xd8\x27\x5b\x72\x06\x0a\x35\x68\x23\xbd\x16\xd0\xf1\xe4\xb4\x69\xea\xc0\x3f\x08\x8f\xfe\x61\x3b\x10\xe0\x82\x07\x7f\xa0\x14\xb0\x23\x2c\x35\xe8\x49\x47\x00\x24\x26\x8c\x86\x44\x59\xea\x2b\x64\xe2\xb1\x0b\xf0\x0b\x82\xc4\x50\xc4\x31\xf2\x08\x52\x65\x50\x76\xbd\xe9\x1a\xf6\x22\x85\x2d\xd9\x21\xc4\x42\x22\xe8\x2d\xe1\xf0\xf3\xb9\x33\xe9\x5d\x18\xac\xc4\x0e\xe1\xe2\xdc\x7f\x00\xaa\x81\x3a\xd2\xa8\x14\x72\x4d\x09\xeb\x9e\x00\x58\x01\x97\xde\x3b\x2d\x1d\x14\x1d\x2e\x38\x76\x6c\xdf\x61\x8e\xb9\x59\x69\x8f\xf4\x8a\x84\x0f\x46\x1c\x2d\xcc\xda\x18\x2a\x6e\xb6\x4b\xdf\xd0\x87\xce\x8a\xca\xa8\x73\xd2\x40\xe2\x76\x71\xef\xc7\x9d\x18\x23\xa0\xb7\xcb\x58\xa7\x7d\xe8\x5c\xbc\x79\x73\xee\x07\x2c\x0c\xa7\xc9\x38\x33\x34\x50\xb1\x67\x66\xb0\x5f\x69\x10\x1c\x90\x84\x5b\x3b\xc1\x2e\xd8\x71\x2a\xc1\x90\x12\x66\xe9\x58\x7c\x15\x50\xee\x56\xc5\x91\x81\x47\xca\x18\xac\x10\x48\xaa\x85\x31\x53\x21\x61\x6c\x0f\x89\x48\x52\x46\x34\x46\x76\x36\x9c\x2e\x3d\xf3\xa5\x1b\xd5\x87\xff\x0b\xac\xd1\xf9\x5f\x6f\x7a\x3a\x06\xa9\x4e\x1f\x3a\x0f\xef\x54\x90\x88\x28\xf0\xfd\x3b\x67\x59\x87\x90\x67\x0a\x67\xba\x9d\x77\x5f\x77\xcf\x8b\xc6\x84\xa5\x1b\xca\x55\xa7\x0f\xbf\xf9\x4f\x05\x69\xdb\x41\xef\x13\x4b\xde\x01\x9b\x0f\xb4\x6d\x4c\x6c\x96\x0c\x77\xc8\x4c\x07\x63\x6e\xab\xcd\x46\x43\x95\x16\x12\x97\x19\x11\xa3\xbf\x92\xa3\x46\x55\xed\x69\x70\xcb\xe6\xb1\x5c\x7e\xba\xbf\x1c\xdd\x4d\x46\x8b\xd1\x7c\x39\x99\x5e\x8d\x96\x93\xc1\xed\x68\xb9\xac\x8e\x88\x75\xda\xe9\xc3\x72\x39\x9c\x8c\x97\xb7\x8b\xfb\xe5\xb2\xd2\x4a\x13\x12\x77\xfa\x95\x89\x94\xa6\x62\x76\x6b\xc0\x44\x48\x58\x85\x26\x40\x47\xa5\x2b\x8e\xda\xf4\x49\x15\xce\x44\x34\xa4\x91\xec\x94\xba\xfc\xab\xc2\xc5\x3a\x9e\xfd\x01\x9f\x12\xa7\x87\x77\xaa\x7d\x78\x09\x8c\x06\x12\xa6\xd5\x2d\x79\x01\xca\x70\x3a\xf9\x30\xbe\x5e\x7e\x18\xdf\x8c\x66\x83\xc5\xc7\xe5\xb2\x42\xfc\xa4\x81\x4d\xe3\x52\x26\x42\xea\x98\x24\x55\x40\x15\x27\x66\xe2\x5a\xa6\x58\xf9\x1e\x92\x84\xac\x28\xa3\x9a\x3a\x39\xed\xe8\x5b\x92\x24\x94\x6f\x94\x1f\x50\x62\xed\xff\x72\x61\xc7\xbf\xbc\x43\x3c\xb0\xda\xce\xc9\x2a\x20\x99\x77\x39\x83\xc7\x2d\x0d\xb7\xf9\x9e\xb0\xfb\x37\x82\xd5\xde\xef\xf4\x57\xca\x1b\x9b\x88\x60\x6c\xac\xea\xa9\xff\x5f\x51\x6d\x37\xd6\x0a\xf5\x23\x22\x87\x0f\x26\x76\x01\x62\x6c\xc2\x16\x61\x30\x1b\xdb\x78\x17\xe5\x19\x48\x8c\xd2\xd0\x58\xad\x9c\x22\x13\x24\x32\x5b\xb7\xda\xb3\x7b\x52\x37\xd1\x15\x3f\xd8\x6e\xb0\xad\x01\x6b\xb7\xd7\x00\x8c\xac\x90\xf9\x50\xc3\xec\x56\x92\x24\xb5\xc1\xc6\x68\x98\x76\x83\xb1\xef\x18\xd8\x7f\xfa\xf0\xf3\x9b\xff\x78\xed\xc1\x4d\xa4\xd0\x22\x14\xac\x0f\x8b\xe1\xcc\x7f\xd3\x44\x6e\x50\xcf\x6c\xd7\x9a\x3c\xe6\xd7\x28\xa8\x42\x86\xa1\x16\xf2\x98\x44\x6e\x01\xdb\xd7\xef\x0a\x13\x26\xf6\xb1\xf1\xb4\x62\xed\xd7\x44\x3b\xc3\x6c\x71\x25\xd6\xf8\x7b\xab\x5e\x83\x96\x24\x89\xaa\x85\x41\x05\xb9\x17\xc1\xf9\x14\x26\x69\xbc\x42\x59\xc8\x9e\xb9\x3e\x63\xc5\x8f\xb9\xd5\x03\xef\x59\x38\x56\xbd\x45\x4b\xfa\xd0\xa1\xc1\x8e\x48\x6a\x29\x52\xa7\x75\x15\x57\x5f\x8a\x35\x2c\x6a\xc6\x01\x9c\x5a\x4a\xdf\xec\x7a\x9b\x7d\x6f\xe1\x77\x2d\xf5\x7f\x7b\xdc\x22\xf7\x64\x8d\x50\x9f\x72\xbb\x04\xb9\x05\xff\xf7\x2e\xc0\xbd\x42\x10\x1c\x33\x88\x6c\xbc\x80\x3b\x94\x7b\x43\x3f\xf8\xe9\xbc\xe0\x3a\xe6\x96\x70\x22\x45\x94\x86\xc6\x57\x9e\xc1\x63\x79\x0a\x32\xe5\xdc\x70\x23\x1a\x4c\x74\xa7\xe1\x75\x8e\xbb\xc1\xd6\xee\x54\xe7\xa4\x69\x9c\x90\xd0\xaa\x96\x14\x8c\x99\x31\x69\xb2\x91\x24\xb2\xe0\x64\x63\xfa\x70\x6e\xff\xdb\x51\xa3\x57\x1f\xa9\x91\x78\x7f\x43\x63\xaa\xfb\xf0\x93\x89\x2c\x30\x4e\x8c\x4b\x75\x1a\x51\xd6\x2b\xf3\x2b\x6b\x4b\xbb\xc6\xb8\x36\xc2\xb9\xd0\x2e\xa3\x28\x06\xb8\xcd\x71\x06\x84\x09\x6e\xbc\xba\xde\xba\x10\x45\x52\xeb\xd3\x07\x51\x24\xb8\x9a\x72\xb6\x07\x2d\x18\xfa\xe0\xc1\x46\x55\x67\x10\x13\xf9\xe0\x62\xd7\x44\x44\x40\xcc\x8e\x0a\xfd\xb8\x12\x7d\x12\x45\x81\x01\x11\xb9\x4a\xa5\x01\x81\x6a\x93\x98\x28\x48\x24\x15\x92\xea\x3d\xa8\x70\x8b\x51\x6a\x01\x72\xb6\x8f\x68\x30\x96\x31\x4f\xfc\x80\x48\xb3\x02\xd6\xc4\x45\x25\xd2\x74\x6d\x88\x99\x65\x74\x14\x71\x47\x43\x1f\x7c\xb8\x9f\xa7\x8c\xb2\x4b\x58\xb2\x25\xdd\xc2\x6b\x99\x24\x26\x13\xd6\x44\x1e\x7d\x78\xf5\x2a\x1f\xe6\x73\xb0\xc0\x04\x37\x2a\x24\x66\x7c\x75\xa4\x22\x6b\x0c\xb4\x08\x2c\xc3\x3e\xbc\x32\x7e\xc4\x0d\xcf\x36\xa7\xb5\x5b\x22\xc2\x79\xc5\x4c\x99\x9f\xb1\x18\x35\x72\x42\xf5\x81\x51\x9e\x7e\xf1\x9d\x8c\x8f\x9f\xb8\x48\xc8\xf9\xa8\xcc\x4e\xe6\x4b\x50\x59\xc1\x5b\x22\x1f\x9a\x97\xc1\x83\x6f\xb5\x5d\x62\x81\x73\x81\x50\x00\x0f\xb8\xef\x37\xac\x77\xc9\x8d\x8a\xc4\x70\x15\xb2\x0f\xa3\x2f\x54\xe5\x39\xec\x29\xcc\x29\x0f\xf3\x70\x36\x24\xfc\x95\xce\xa3\x4d\x62\x85\x49\xb9\xa6\x2c\xdb\xee\x0a\xd2\xc4\xee\x25\x8e\x2e\x13\x91\x29\xcf\xda\xb4\xc9\x44\x72\xba\x76\x0e\x36\xd0\xf1\xe4\x30\x32\xe4\x32\xa9\x55\x25\xa9\x9b\x34\x64\x74\x76\x11\x05\xd7\x84\x72\x94\xa5\xfc\x98\xc6\x64\x83\x7d\xf8\x67\x4a\xf6\x56\x01\x5c\xe6\x64\x37\x49\x96\x3b\x65\x93\x6e\xb4\xda\xde\x85\x15\x0e\xce\x91\xcd\x59\xcd\xea\xde\xee\x09\x4a\xcd\xee\x10\x00\xf9\xae\x5f\xea\x74\x0a\x23\x67\xcb\x5d\x9c\x0a\x4c\x6c\x36\x66\xb7\xac\xf6\x10\xe1\x9a\xa4\x4c\x77\x01\x86\xc4\x6c\x4c\x67\xc4\x05\x74\x22\x5c\xa5\x9b\x8e\x8b\xf3\x8d\xc7\x53\x68\x8b\x05\x42\x51\xbd\xef\x96\x68\x07\x5e\xbe\xc5\xaf\xb3\x8f\x83\xe5\xcd\xf4\x7a\x3e\xfa\x3c\xba\x1b\x2f\x7e\x9d\x0f\xef\x46\xa3\x49\x25\xb2\xb3\xce\x22\x8b\x96\x2b\xf2\x5d\x51\x65\x05\xcc\x24\xd3\x02\xd6\x94\xa1\xdd\xcf\x6a\xaf\x98\xd8\x80\xb2\xca\xa2\xb7\x42\x21\x44\xc2\xe8\x4a\x4c\x1e\x8c\xbc\x5c\x59\x97\x52\x58\xee\xa3\xe2\x65\xf1\x63\xa3\x60\x59\xc6\xf5\xf4\xe4\x7e\x9d\x7f\x15\x81\x53\xb8\x15\x9c\x6a\x21\xeb\xce\xc5\x84\x5c\x76\x92\x3e\x58\xe3\xb9\x2b\xce\xfc\x83\x4d\xaf\xb8\xb5\x5e\xc6\x6b\xe0\x8a\x30\xf3\x6f\x85\x78\x28\x38\x47\xeb\x64\x8e\x4c\x7a\x38\x9d\x4c\x46\xc3\xc5\x78\x3a\xb9\x1b\x5d\x0e\x6e\x06\x93\xe1\x78\x72\x7d\x3b\xbd\x1a\x35\xce\xa0\x14\x98\xb7\x52\xbc\x1a\x2c\x06\xf3\xc5\xf4\x6e\xb4\xf8\x75\xf6\xed\x54\x3e\x8e\x06\x37\x8b\x8f\xa3\xc9\xe0\xf2\x66\x74\xd5\x48\xc5\x98\xae\x2a\x9a\xf7\xdc\xb9\x51\x1b\x68\x28\x34\x56\x0f\xad\xdb\xf4\xc1\x4a\x22\x45\x8c\x7a\x8b\xa9\x32\xae\x4e\xd2\xd0\x78\x64\x67\x66\x72\x2b\x52\x35\x0c\x67\x15\xfa\x36\xc2\x11\x09\x72\x65\x2d\x90\xd4\x59\x88\x6c\x06\x65\x71\x7a\x4c\xf6\xb9\x0d\xb2\x3b\x26\x4c\x65\xd9\x6b\x00\x9c\x56\x67\x3a\xbb\x9b\xde\x8e\x16\x1f\x47\xf7\xf3\xdb\xd1\xe2\x6e\x3c\x9c\x1f\x4e\xfa\xf4\xc8\xac\x9f\x20\x36\x9b\xde\x2d\x9a\x29\xbd\x3f\x7f\xff\xba\xa0\xc4\xe8\x0e\x39\x2a\x35\x93\x62\x85\x65\xf3\x80\x5f\x0a\x9f\xe3\x7e\x06\x64\xc2\xa3\xea\xc7\xa0\xcd\x06\xd9\xa6\x2d\x86\x0f\xb5\x6f\x19\xc3\xb2\xb9\x42\x49\x45\x34\x37\xc9\x5d\xa4\xfa\xf0\xfa\xbc\xd4\x46\x39\x35\x61\xd9\x15\x32\xb2\x6f\xea\x21\x91\x44\xf4\x45\x27\x90\x73\x6c\x9f\xc1\xc5\xf9\x93\x59\xc2\x4c\x44\xc6\xbe\xc9\x34\xb1\xc1\xcf\x65\x1a\x6d\x50\x97\x0a\x4e\x26\xa0\x66\x4c\x3c\xc2\xa7\x77\x2a\x2b\x51\xc2\x20\x0f\x1d\xac\x76\x9b\x38\xa1\x9a\x3f\xb8\xfc\xbb\x96\x41\xcc\x44\x54\x70\x72\x8c\x5e\x24\x95\x88\xc9\x97\x7b\x4e\x76\x84\x32\xb3\x0b\xfb\x70\x71\x90\x60\xc5\x44\x87\xdb\x9b\x4a\xb4\xf9\x9c\x9c\xcb\xd7\x9a\x54\x29\x71\xb0\x75\xcc\xc2\x75\x9e\x19\xaf\xff\x88\x8c\x9d\x58\xff\x5f\xaa\x93\x0d\x27\x63\xf0\x95\x1e\x6b\x4a\xab\x05\x2d\x10\xfc\xe4\xd4\x15\xb0\x62\x62\xa1\x37\x7d\xac\x6d\x90\x36\x02\x33\xbe\x85\x94\x4d\xb7\x8f\xee\xb2\xe2\xe6\x95\x4d\xca\xe7\xa8\x2b\xcb\x53\x9c\xbc\x7c\xe5\xe9\xdf\x37\xae\x83\xab\x5b\xfb\x65\x78\x36\xe4\x9e\xb3\x2b\xe7\xcf\xb5\x24\x1a\x37\x7b\xd7\x59\xef\x13\xec\xc3\x9d\x4b\x3c\xee\x8b\x7a\xbf\x2c\x7f\xc9\xe8\x36\xad\xff\x9f\x4b\x3c\x4a\x81\xd8\x5f\x9d\x77\x94\x68\x1c\xc9\x40\x5a\xf3\x8e\xd2\xf0\xbf\x67\x06\xf2\x63\x52\x08\x13\x8a\xa5\x12\xcb\xeb\xe7\xa6\x94\xcd\xc0\x56\x9b\x08\x63\x3e\x51\x2e\x85\xbf\xb8\x5e\xa3\x49\x83\x26\x62\xee\xfb\x3e\x9d\x3c\xfc\xc0\xbc\xa5\x2a\xf2\xe8\x0b\x86\xa9\xfe\x0a\x89\xbf\x2e\xeb\x38\x85\x5b\xca\x69\x4c\xff\x30\x41\xee\x23\xd7\x34\x46\x88\x9c\xfe\x91\x7a\x19\x00\x84\x84\x08\x19\x9a\xd5\xf8\x4f\xd0\xc8\x58\xd9\x4c\x69\x01\x91\x00\x02\x9d\xb5\x90\x61\x41\x3e\x1b\xd0\xe9\xe7\x47\x50\x55\x4d\x88\x44\xa8\x7a\xa1\xe0\x21\x26\x5a\xf5\x8c\x02\x30\x41\x22\xd5\xb3\x87\xfd\x89\x88\x7a\xa7\x1a\x65\x4c\xb9\x55\x82\x40\xac\x8d\xee\xe5\xcb\x59\x6a\xba\x96\x24\xc4\x59\xd5\x5b\x9e\xb7\xe6\x56\x46\xb4\xbb\x94\xab\x66\xe3\x9e\x9f\x31\x94\xe6\x97\x1f\x37\xd0\xb2\x4e\x14\x43\x12\x29\x36\x92\xc4\x2a\xb7\xf7\xce\x63\xba\x68\x5a\xa4\x86\x84\xa7\x5a\x1a\x6e\xd4\xbe\xac\x24\x4d\x76\x3a\xfb\x35\x27\x83\xe5\xd3\xb5\xa2\x6f\x2d\x21\xb3\xa1\xac\x3a\x48\x08\xbc\xdf\x5a\x91\xf0\xc1\x2c\x73\x5e\x80\xea\xd6\xe2\x12\x27\x55\x1e\x85\x2f\x0f\xc2\xf0\xa7\x02\x71\x6b\x3f\x5d\x09\x39\x4d\x5c\x89\xd5\x00\x67\xf4\xcb\x15\x9e\xad\x2d\xcd\x2a\x72\xcd\xec\x3f\x8c\x6e\xc6\xff\xb5\xb4\x31\xe9\xa7\x77\xf3\xf9\xe8\xee\xf3\x78\x38\x9a\x0c\x6e\x1b\x25\xf9\x20\x45\xdc\xaf\x35\x80\xf7\xc0\xb7\x24\xf9\x84\x7b\x7f\x9c\x5b\xff\x35\x9f\x0a\xd6\x7f\x76\x2f\x1f\x96\x1c\x6b\x33\xfe\x85\x50\x17\x77\x99\xd9\x3d\x01\xee\x2f\x83\xf1\x62\xf9\x61\x7a\x57\xe4\x3a\x2d\x00\xd7\xa3\x75\x5b\xd7\x40\x0d\x2b\xa2\x30\x2f\xb0\x3f\xbc\x73\xfa\x6a\x89\x37\x33\x9c\x4c\xaf\x9e\x89\xde\x9a\x22\x8b\x5a\x60\xb3\x6d\x33\xa2\xb7\x7d\xeb\x0e\xba\x86\xf9\xe4\x10\x91\xe1\x56\x98\x74\x3a\x53\xba\xea\x69\xe2\xa1\x8c\xc3\xc1\xcd\x78\x38\x5d\x4e\x46\x8b\x5f\xa6\x77\x9f\xc6\x93\xeb\xe5\xe5\x60\xf8\x69\x34\xb9\xfa\x51\x4b\x5e\x3d\xf0\xac\x4f\xce\x87\xd2\x26\x92\xb1\x75\x8c\x08\xb9\xa6\x6b\xa7\xdc\x51\x51\xba\x37\xed\xcd\xd3\xbd\xb9\x9f\x2f\x46\x77\x47\xb7\xd7\x3b\x75\xb6\xda\x24\x75\x05\x30\xd1\x7b\x10\xa1\xc6\xd0\x26\xaa\x70\x79\x3d\x83\xf1\xcc\x78\x25\x89\x4a\x35\x83\x3b\x9e\xb5\xf0\x20\xa9\x16\x8e\x56\x9d\x8d\xaf\xea\x8c\x67\xb5\xb1\xb5\xf5\x1a\xcf\x3e\xbf\x99\x4d\xa7\x37\xcb\x83\x8e\x05\x93\x01\x7b\x24\xfb\x03\x1b\x61\x14\xf9\x76\x71\xef\x36\x4d\xca\x39\x32\x88\x4c\x88\x82\xee\x36\x80\x89\x5e\x12\x6a\x6f\x07\xb8\xfc\x3b\x3a\x62\x29\xc6\xb3\xf1\x64\x3c\xbb\x5d\xdc\xff\x28\x65\xc9\x8e\xb9\xeb\x76\xd0\x6a\x83\xad\x82\xc1\x78\xb6\x7b\x03\x89\x10\xcc\x5e\x74\xb0\x19\x9c\xd9\xc2\x4a\x13\xa9\xd3\xc4\xcc\x97\x0b\x8e\x80\xd6\xaf\x77\x6d\x66\x37\x9e\xa9\xec\xe0\xae\x46\x38\xdc\x0a\x85\x1c\xd6\x52\xc4\xae\xae\x20\x09\xdf\x60\x17\x86\x5b\xc2\x5d\x91\xcb\x7c\x74\x07\x28\x64\x6d\x34\xb5\x7c\x77\xc2\x11\xdd\x92\x5d\x9d\x2c\x17\x3e\xfa\xe8\xba\x1c\x49\x6d\x45\xca\x22\x58\x9b\xd0\xca\xc4\xc1\x94\xc3\xef\x41\x90\x95\xa0\x43\x1a\xc9\xdf\x8f\xee\xe6\x5c\x3b\x86\xe3\xab\xbb\x16\xed\xb8\x78\xff\x53\xf7\xe2\xed\xbb\xee\x79\xf7\xbc\x77\xf1\xb6\xae\x25\x59\xf1\xce\x56\xec\xb2\x0a\x9e\x12\xf0\xbb\x3d\xc6\xd5\xcc\x7c\x53\xbf\xdb\x9c\xaa\x45\xf7\xbd\x28\x57\xe3\xf9\xe0\xf2\x66\x64\xcf\x78\x97\x37\xd3\xeb\xeb\xf1\xe4\xfa\x99\x66\xd7\x1d\x81\x66\x97\x9c\xcc\x32\x1a\x97\x9e\x2f\x30\x09\xb3\xdb\x0b\x83\xe1\x70\x34\x5b\x1c\x73\x6d\x57\xa3\x0f\x83\xfb\x9b\xc5\x68\x72\x35\x9b\x8e\x27\x8b\xc5\xf4\xe3\x74\xbe\x18\xd8\x42\x5a\xdb\x26\xb2\x34\xdb\xe0\x19\xcf\x76\x6f\x8d\x36\x35\x17\x2a\x0f\xb7\xcb\xe7\xb7\xf3\xfb\x59\xad\xb8\x53\xe6\xb6\x26\x4c\x1d\xc1\xa0\x54\x4a\x3d\x28\xb8\xd6\x99\x1d\x2f\xdb\xb6\x17\x6e\xeb\x74\xda\x4b\x7a\xed\x2b\x67\xcb\x67\x54\xef\x87\x82\x6b\xfc\xa2\xab\x9b\x3c\x91\x74\x47\x19\x6e\x30\xaa\x64\x21\xd0\x74\xad\x2f\xfb\xfc\xcf\x14\x95\x56\x75\x63\x11\x26\x69\x1f\x7e\xfa\xf9\x3c\x2e\x7d\x6f\xad\x88\x81\x8d\x8d\xaf\x51\xd7\xa9\x24\xd6\xa5\xf6\x1a\x2a\x5b\xe0\x0b\xfb\x7d\x78\x7f\xfe\xfe\x7d\xad\xc1\xe8\x60\x1f\xec\xad\x0b\xf3\x67\x75\x8a\x07\x05\xa5\x72\x6b\x63\x51\xac\xd6\x67\x4d\x28\x4b\x25\x2e\xb6\x12\xd5\x56\xb0\xa8\x0f\x6f\x2b\x38\xb5\x95\xcd\x9a\x0a\x67\x2d\xa5\x33\xb3\xd0\xbd\x15\xe5\xbd\xe6\x50\xd8\x75\x08\x56\x54\x46\x81\xe1\xb7\x3f\x6c\xb3\xf7\x2a\x1b\x1a\x8f\x4d\x7f\x27\x58\x1a\xe3\xad\xbd\xc9\x5c\x2f\xe5\xc5\xe6\xeb\x2c\x5b\x8e\x55\x2f\x16\x26\x71\xac\xaf\x88\x53\x4e\x46\x57\x41\x73\xbb\x91\xc7\x64\x7d\x07\xda\x55\x67\x21\x53\xde\xfb\xa2\xcd\x3e\x56\x5d\x26\x6a\x95\xc2\x8c\x8f\xef\x10\x34\x74\x28\x18\xd9\x6d\x7b\x84\xd3\x8e\x48\xcb\xcd\x21\xdd\xc8\x67\x47\x64\x20\x53\x1e\x34\x76\x79\x1e\x27\x03\xdd\x13\x9c\x0c\x7a\xcf\xe2\xe4\x8b\x77\x45\x2a\x56\xa9\xde\x95\xaa\x72\x2b\xca\x89\xa4\x58\xce\xdf\x4c\x76\x76\x78\xeb\xcc\xb9\x96\xca\x45\xb3\x83\x24\xcd\x33\x09\x42\x4e\x9f\x4c\xd2\x4a\x77\x1d\x8b\xae\x99\xe2\xc3\x6f\x9d\x5e\x89\x58\x57\x6d\xfd\x45\x62\xf7\x6b\x48\xe5\x4c\x64\x0d\x62\xed\xa6\x37\x19\x57\x84\xce\x03\x89\x16\xcf\x37\x19\x2f\x87\xd3\xc9\x87\x65\x5b\xec\x6f\xbc\xef\xb9\xc7\xbf\x6b\x08\x33\xaa\x0e\x02\xc1\xb9\xbb\x79\x61\xed\x8c\xbd\x6b\xf1\xdc\xec\xa3\xe9\xd6\xd9\x8b\x66\x22\xcd\xb7\x0d\x5b\xaf\x19\xb6\x62\xe9\x73\x93\xa5\xbb\x2d\xf6\xc3\xf2\x92\x83\xab\x8b\xf5\xdc\x64\x32\xb6\x81\xb5\xcb\xb1\xf3\xab\x31\xad\xd3\xfa\x7b\x85\xcd\x5f\x6b\x97\x8d\x3a\xf6\x44\xa2\xcd\x7e\x33\xde\xa3\xd1\xc4\x84\x9c\x06\x2b\xca\x83\x88\xca\xa7\x48\xa1\x0e\x2d\x29\x8e\xba\x1b\xb5\x12\xe3\xa8\x4b\xc4\x9c\xac\x95\x3a\xd3\xbd\x72\x45\x8e\x52\xc5\xe6\xd0\x9c\x34\xfb\x0c\x23\x86\x95\xa8\xea\xc3\xda\xdd\x50\xf0\x94\xc9\x3e\x4a\xb1\xc5\x17\x04\x4f\x59\xe7\x27\x89\x36\x98\xfd\xe0\xb8\x13\x3b\x4a\xf2\xa8\x77\x74\x27\x09\x1f\x28\xc3\xa9\x1c\x96\x1f\x18\xe4\x6b\x51\xda\xe5\xc3\xc9\xb8\xa1\xfe\xd6\xa8\x22\x47\x25\x6a\xd2\xba\xa0\x55\x49\x9e\x24\x57\xd5\xbc\x83\x87\x0e\xf5\x7b\x92\xd9\x7b\x81\x6f\x3c\xf2\xc9\xce\xc0\x1c\x5a\xb6\x68\x6e\x5d\x4b\xaa\xb4\x88\xef\x7c\xdc\x7b\x85\x6b\x1b\x1c\x0a\xae\xec\x69\x34\x46\x26\x4d\x2f\x5e\x1c\x94\xca\x9d\xc5\x7b\x1c\x88\xad\xb2\xd7\x6e\x22\xd2\xe2\xb4\xaa\xe5\x81\x46\x0b\xe7\xca\x04\xfd\x0c\x1b\x1e\xd0\x74\x1b\x9f\xe7\x14\xc7\x56\xa1\x30\x1a\xe2\xeb\x35\x27\x00\x1b\xf7\xa0\xa3\x79\x14\x64\xaf\x1a\x2c\xec\x1e\x43\x7f\x38\x66\x85\xb5\xe9\xcf\xb0\x2c\x80\x6d\x4c\x58\x2a\x09\x6b\x94\xcf\xb6\x2b\xca\x37\x29\x23\xb2\xa9\xc7\xe1\x8a\xff\x55\xa0\x39\xcc\xb2\x87\x45\x2f\x0b\xd4\xe5\xf5\x6c\x86\x28\x2b\xe8\x54\x9e\x38\x15\x90\xf8\xcf\x0d\x6f\x7c\xfe\x72\x20\x7e\xa4\xea\x5c\x5e\xcf\xda\x15\xa7\xf9\x4d\x57\x05\xa3\xaa\xd2\x7c\x6f\xb0\xfc\x53\xb2\x97\x85\x68\x3c\x9b\x09\xc1\x2a\xc0\x94\x9f\xb4\x15\x70\xb8\xaf\xdf\x1d\x84\xca\x9b\xb6\x97\x85\xe2\xa3\x50\x7a\xe4\x59\x57\x00\x39\x7c\x67\x57\xc0\x52\x6e\xfb\xee\xe0\x34\x3c\x06\x7c\x59\x88\x3c\x9d\x71\x21\x40\x05\xa8\xb6\xc7\x8a\x05\x5c\x87\x3d\xbe\x3b\x68\x8d\x2f\x32\x5f\x16\xb6\x6b\x2b\x82\x3f\x7f\x9f\x59\x47\x5e\xc1\xad\xfd\xd1\x68\x81\xdc\x61\x9f\xfd\xcb\x42\xa7\xf0\xa5\x37\x64\x05\xb6\x39\xea\x76\xcc\xf2\xa7\xb4\x2d\x78\x29\xfc\xfe\x9b\xf3\x9b\x34\x6c\x92\x45\x8c\xd1\x9f\x46\xab\x5d\xbd\x8e\x2b\x56\x4d\xa5\xfe\x3f\x00\x00\xff\xff\x49\x3a\xe7\xa3\x63\x41\x00\x00"

func netCalicoYamlBytes() ([]byte, error) {
	return bindataRead(
		_netCalicoYaml,
		"net/calico.yaml",
	)
}

func netCalicoYaml() (*asset, error) {
	bytes, err := netCalicoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "net/calico.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0xdd, 0x6f, 0x68, 0x11, 0xd4, 0x9c, 0xae, 0xee, 0x99, 0xce, 0x8b, 0xd, 0x91, 0xe6, 0xe3, 0x8c, 0xb7, 0xf, 0x1f, 0x58, 0xfb, 0xd7, 0x95, 0x1a, 0x51, 0x1d, 0x7e, 0x29, 0x11, 0xa5, 0x53}}
	return a, nil
}

var _netWeaveYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdf\x6f\xdb\xb6\x13\x7f\xef\x5f\x71\xf0\x4b\xbe\x5f\x6c\x92\x6c\x6f\x59\x02\x02\xc3\x96\xa5\x6d\x36\xa0\x71\x82\x38\xdd\x3b\x45\x9d\x6d\xce\x14\xa9\x91\x27\x25\x5e\x97\xff\x7d\xa0\x64\xbb\xb2\x24\x5b\xe9\xda\x0d\x09\x16\x3d\x49\xc7\xbb\x0f\x8f\xf7\xe3\x43\x52\x3c\x93\xbf\xa2\x75\xd2\x68\x06\xc5\xe8\xd5\x52\xea\x84\xc1\x3b\xe9\xe8\x95\x24\x4c\x1d\x7b\x05\x10\x40\x43\x09\x00\xa0\x52\x9c\xa2\x2d\xa4\xc0\x33\x21\x4c\xae\xa9\x1c\x48\x91\x78\xc2\x89\xb3\xf2\x0b\x40\xf3\x14\x19\xdc\x21\x2f\x30\xd0\x48\x6b\x29\xd7\xda\x10\x27\x69\xb4\xdb\x28\x02\x08\x65\xf2\x24\x2c\x55\xc3\x3b\x63\x97\x2e\x52\x3c\xd7\x62\x81\x36\x90\x7a\x66\x18\xfc\x19\x6c\x75\x01\x3e\xd4\xde\x01\x06\xc6\xca\xb9\xd4\x5c\x05\x16\x7f\xcf\xd1\xd1\x80\x35\x34\x00\x06\xb9\x55\x03\x06\x83\x68\x79\xea\xa2\x62\x14\x8e\x86\x91\x46\x0a\x57\x3c\x55\x3f\x2c\x4f\x5d\x50\x54\x6b\xfc\x7e\x14\x8e\xc6\xe1\x78\xf0\x75\xd3\x3c\xe1\x84\xde\x7e\x9a\x6b\x98\x98\x02\xc6\xc7\x30\x1e\x8e\x4e\x61\xf8\x0d\x3b\x3e\x61\xa3\x63\xb8\xb8\xbc\xfd\x6a\x38\x1c\x0e\xe1\x7f\xef\x6f\xcf\xff\x3f\xd8\xb1\x7f\xd8\x85\x1b\x60\xca\xa5\x0a\x78\x92\x58\x74\xce\xa3\xba\x3c\xcb\x8c\xa5\x1f\x6b\xeb\xaf\x23\x3c\xac\xdf\x15\x8f\x51\xd5\x82\xd6\x1d\x5f\x2f\x75\x19\x17\xc8\x60\x99\xc7\x18\xb8\x95\x23\x4c\x5b\xc9\xb4\x31\x17\x21\xcf\x69\x61\xac\xfc\xa3\x4c\x47\xb8\x3c\x75\xa1\x34\x51\x31\x8a\x91\x78\x3d\xd7\xe7\x2a\x77\x84\xf6\xc6\x28\x7c\x49\xf4\x13\x49\xb4\xcd\x15\x6e\x55\xca\xd4\x5e\x58\x93\x67\x35\x2b\x2f\x3e\x3a\xda\x7e\x5a\x74\x26\xb7\x02\x1b\x1a\x99\x49\xdc\x8e\x60\x5b\x3f\x0d\xb1\x49\x6a\x92\x02\x6d\xdc\x00\x9a\x6f\x13\x5f\x7d\x2b\xcf\x24\x75\xc1\x1d\x27\xb1\xe8\x73\x58\x23\xf9\xb0\x48\x3d\x5f\x97\x63\x9f\xff\x6b\x83\xcc\x28\x29\xe4\xbf\xe1\x61\x7f\x48\xcb\x50\x45\x8e\x38\xe5\x07\xfd\xc9\x6a\xd3\x55\x92\x3c\xf3\xf5\xf7\xa5\x7a\xf5\x27\xa9\x13\xa9\xe7\x2f\x2d\xfb\x54\x5a\xd6\x28\xbc\xc1\xd9\x46\xa9\x9b\x5b\xf7\xa6\x66\x5d\x8d\x07\xaa\xa1\xd4\x74\x79\xfc\x1b\x0a\xaa\x51\xc3\xde\xfd\x7a\xff\x64\xff\xc8\x2e\xf2\xb2\x7d\x3c\xa1\x5a\x3c\x90\xe1\xcf\xd8\x5b\x26\x1e\x73\x57\xab\x5d\x5a\x7b\x58\x53\x18\x3d\x93\xf3\x94\x67\x9f\xc4\xe1\x5b\xc6\xfc\x54\x57\xff\xd6\xec\xc2\xe2\xe7\xf2\xf3\x0b\x31\x3f\xb7\x66\xe8\x62\xed\x67\x4a\xd7\x78\x4f\xa8\xfd\xab\xeb\xa8\xcc\xd7\x1c\x53\xa3\xa7\xf8\x72\x99\x7b\x16\x75\xe9\x32\x14\x1b\xf3\x54\xea\x1b\xe4\xc9\x6a\x8a\xc2\xe8\xc4\x31\x38\x5e\x0f\x10\xa6\x99\xe2\x84\x1f\xe7\x69\xe6\xb5\xcb\x97\x43\x95\x56\x9f\xb6\x4c\xb6\xd1\xc4\xa5\x46\xdb\xb0\x0f\xea\x08\x8d\x0c\x08\x93\xa6\x5c\x27\xac\x21\xf6\x46\xd1\xc2\xa4\x18\x95\x46\xeb\xaa\x09\xdd\xa2\xa1\x87\xba\xe8\x32\xad\xe6\xfb\xf9\x6a\x7a\x3b\x39\xbb\x7c\xd3\x52\x00\x28\xb8\xca\xf1\xad\x35\x69\xdb\xda\x3f\x33\x89\x2a\xa9\x75\x7a\xf3\xe9\xf8\x15\xb2\x07\xe4\x9a\xd3\x82\x95\x91\x0a\xfd\x5d\xc0\xef\x8a\x0d\x75\x99\xf2\x39\x32\x38\x4a\x8c\x58\xa2\xf5\x3b\x45\xb9\xe4\xaa\x5b\xaa\xa0\xfb\x7c\xb3\x71\x78\x1c\x0e\x8f\x1a\xc6\x16\x79\x22\x35\x3a\x77\x6d\x4d\x8c\x6d\x77\x17\x44\xd9\x05\x52\xd7\x3a\x16\xc6\x11\x83\xd1\xf8\x24\x1c\x86\xc3\xb0\x6b\x11\x59\xe9\x7b\xf3\xf6\x52\x1b\x37\x96\x18\x7c\x77\x72\xfa\x6d\xcb\xab\x8e\x8d\x75\x33\x54\xb6\x72\xc7\x08\x80\xc8\x72\x06\xa3\x61\xda\x18\x72\x28\x72\x2b\x69\x75\x6e\x34\xe1\x7d\xc7\x5a\x32\x2b\x0b\xa9\x70\x8e\x09\x03\xb2\x79\x33\xc0\x85\x51\x79\x8a\x97\x9e\x3c\x3b\xa6\xdd\x29\xce\x24\xee\x70\x2b\xf5\x96\x55\x1e\xa3\x7d\x5a\x1b\x14\xa1\x65\x10\x4b\xdd\x83\xe2\x43\x1f\x99\x8c\xfa\x60\xc6\x8f\xc1\xf1\x5d\x72\x10\xc8\x9f\x67\x1e\x03\x84\x24\xf6\xe2\x24\x71\x67\x05\xb4\x30\x0a\x6e\x23\x25\xe3\xa8\x53\x7f\x03\xa6\x64\x1c\xa4\x26\xf1\x27\xcb\x1e\x4c\x0f\xb5\x4f\x73\x83\x76\x4f\x3c\x56\xe8\x02\x65\xc4\xb2\x07\xce\xe6\x3a\x5a\xab\x87\x2d\xf5\x60\x97\xe7\xb2\x66\x30\xfe\x1b\x54\xa3\x33\xb1\x8f\x69\x9e\x5f\x4f\x7f\x99\xd2\xf0\x85\x3d\xa9\x7e\x2d\xb5\x7c\xf1\x63\xd7\xbf\xbc\x6e\xc9\x2d\x3a\xe2\x96\xae\x8d\x92\x62\xc5\xe0\x4c\xdd\xf1\x55\xbd\x84\x0f\xae\xdf\xe1\x3b\xa9\xf3\xfb\xab\xac\x3a\x4c\xc1\x87\x87\x1d\xcb\xfa\x69\x70\xb2\x67\x6b\x06\x20\xa3\xd0\x36\x8f\x63\x55\x70\x70\x36\x43\x41\x0c\x26\x66\x2a\x16\xe8\xbb\xab\x11\x1c\x93\x79\x53\x63\x19\xbc\xb9\x97\x8e\xea\x8e\x57\x81\x3f\xb4\xc1\xb7\xd8\xb1\x8c\x91\x8f\x72\x3b\xc9\x55\xec\x37\xa4\xd1\x3e\x1f\x1c\xa6\xd5\x5e\xe0\x26\xc5\xf6\xd0\x6b\x2f\x5e\x8b\x6a\x7b\x68\xb6\x17\xb0\x49\xb9\x07\xe8\xf6\xd1\x51\x6c\xd9\xf6\xd3\x6e\x2f\xf6\x3e\x1e\x7e\x44\xa3\xf5\x62\x1f\x24\x65\xff\xd0\x2a\x43\x06\x6f\xa5\xc2\x2b\x7b\xbe\xb9\x75\xfb\xa7\xba\xf0\x4f\xc9\x72\xc2\xf9\xea\xe3\x04\x95\xc1\x8d\x51\x4a\xea\xf9\xfb\xea\xaf\xc0\x5f\x01\x00\x00\xff\xff\x43\xbd\x2c\x1a\xe2\x1a\x00\x00"

func netWeaveYamlBytes() ([]byte, error) {
	return bindataRead(
		_netWeaveYaml,
		"net/weave.yaml",
	)
}

func netWeaveYaml() (*asset, error) {
	bytes, err := netWeaveYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "net/weave.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0xcf, 0x84, 0x4, 0x52, 0x9e, 0x47, 0x53, 0xc3, 0xf0, 0xd3, 0x83, 0xb6, 0x94, 0xca, 0xf7, 0x3a, 0x84, 0xc9, 0x64, 0x1b, 0x7f, 0x5c, 0xae, 0x3f, 0x97, 0x56, 0x56, 0x6a, 0x6f, 0x5f, 0x98}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"client.sh": clientSh,

	"controller.sh": controllerSh,

	"node.sh": nodeSh,

	"registry.yaml": registryYaml,

	"net/calico.yaml": netCalicoYaml,

	"net/weave.yaml": netWeaveYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"client.sh":     &bintree{clientSh, map[string]*bintree{}},
	"controller.sh": &bintree{controllerSh, map[string]*bintree{}},
	"net": &bintree{nil, map[string]*bintree{
		"calico.yaml": &bintree{netCalicoYaml, map[string]*bintree{}},
		"weave.yaml":  &bintree{netWeaveYaml, map[string]*bintree{}},
	}},
	"node.sh":       &bintree{nodeSh, map[string]*bintree{}},
	"registry.yaml": &bintree{registryYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}