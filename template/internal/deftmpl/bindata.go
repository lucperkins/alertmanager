// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\xff\x4f\xdb\xb8\xfe\xf7\xfc\x15\x9f\xcb\xe9\xe9\x86\xd4\x36\x85\xdd\x4d\x47\x69\x79\xea\x4a\x18\xd1\x2b\x29\x4a\xc3\x76\xd3\xe9\x84\xdc\xc4\x6d\xbd\x25\x76\xce\x76\x28\x3d\xd6\xff\xfd\xc9\x4e\xfa\x25\x34\x85\x82\x76\xd0\xf7\x0e\xd0\x46\xe2\xf8\xf3\xfd\xab\x63\xe7\xf6\x16\x42\x3c\x24\x14\x83\x79\x75\x85\x22\xcc\x65\x8c\x28\x1a\x61\x6e\xc2\x6c\xd6\x56\xf7\xe7\xd9\xfd\xed\x2d\x60\x1a\xc2\x6c\x66\x6c\x04\xb9\xf4\xba\x0a\xea\xf6\x16\x6a\xf6\x8d\xc4\x9c\xa2\xe8\xd2\xeb\xc2\x6c\x66\xfd\x68\xe9\x79\xe2\xdf\x1c\x07\x98\x5c\x63\xde\x52\x93\xbc\xfc\x26\x83\xc9\xb1\x17\xd1\x8b\x74\xf0\x05\x07\x52\xa1\xfd\x5d\x81\xf4\x25\x92\xa9\x80\x6f\x20\xd9\x65\x92\xcc\x41\xc9\x10\xf0\x9f\x8b\x87\xe6\x90\x70\x42\x47\x0a\xa6\xa1\x60\xb4\x14\xa2\x76\xaa\x47\xe1\x1b\x44\x98\xae\x52\xfc\x03\xd4\xa4\x0f\x9c\xa5\x49\x17\x0d\x70\x24\x6a\x7d\xc6\x25\x0e\x2f\x10\xe1\xa2\xf6\x11\x45\x29\x56\x04\xbf\x30\x42\xc1\x04\x85\x15\x32\x92\x23\x09\x6f\x14\xae\x5a\x87\xc5\x31\xa3\x19\xf0\x5e\x3e\xb6\x82\x6f\x0f\x66\xb3\x37\xb7\xb7\x30\x21\x72\x5c\x9c\x5c\xf3\x70\xcc\xae\x71\x91\xba\x8b\x62\x2c\x72\x35\x96\x51\x5f\x30\xbe\xb7\xb8\xda\x60\x9b\x10\x8b\x80\x93\x44\x12\x46\xcd\x7b\x74\x2c\xf1\x8d\xcc\xec\x78\x15\x11\x21\xf3\xa9\x1c\xd1\x11\x86\x1a\xcc\x66\x19\x5f\x0d\x63\x39\xb8\xae\x27\xa5\x95\xaa\x56\xa4\x62\x5f\xdd\xb5\x60\x21\x40\xce\x58\x46\xbc\x4d\x29\x93\x48\xf1\x54\x40\xb9\x32\xfc\x34\xbc\x7d\x96\xf2\x00\x37\x32\x63\x62\x8a\x39\x92\x8c\x67\xee\x67\x94\x28\xaa\xa0\x03\x11\xa1\xe0\x6b\x2d\xc4\x43\x94\x46\xb2\x26\x89\x8c\x70\xae\x05\x89\xe3\x24\x42\xb2\xe8\x8b\xb5\x4d\x2a\x2f\xe2\x49\x85\x0a\x81\xb8\x0c\x55\x31\xd0\xb6\xc4\x37\x44\x51\x34\x40\xc1\xd7\x35\x7c\xa5\xec\x2b\xa4\xf0\x0d\x1e\x9a\x18\x11\xfa\x75\x6b\x0e\x12\x8e\x95\xb3\x98\xdb\xcd\x5e\xc1\x7f\xaf\x02\x74\xda\xd8\x92\x03\x12\x30\x8a\x63\xf6\x85\x6c\xc9\x83\x9a\x9f\xf2\x68\x5b\x8e\xb7\x17\x6e\xc8\x98\xcc\x92\xe4\x06\x9f\x1a\x93\x24\x18\x23\xb9\x04\xe0\x2c\x7e\xba\x27\xdc\xc5\x16\x63\x21\xd0\xe8\x11\x5e\x5a\xe0\x2d\x51\xd4\xc2\x54\x4e\x17\xf8\xd6\x53\xc5\xe3\x3c\x7f\x1d\x63\x10\x11\x4c\xe5\xd3\x25\xde\x84\x71\x59\x64\x9e\xe6\x4f\xeb\x78\x09\x15\x12\xd1\x00\x8b\x12\xbc\x6b\xb9\xf1\x1e\xad\xb2\x44\x8c\x30\x25\xf8\xe9\x46\xba\x0f\xd9\xba\x85\xf2\x52\xb2\x21\x73\x96\xd6\x0e\xe3\x4e\xe5\x2a\x94\xc6\x3d\xa8\x43\x75\x36\x33\xb2\x41\xc8\x06\x75\x8e\xbe\x5f\x23\xc5\xfa\xaa\x89\x54\x57\x24\x2a\xa1\xe7\x61\xc1\xa2\x6b\x1c\xde\xa1\x38\x1f\xde\x9e\xe6\x1c\x62\x8d\x6a\x75\x1b\x95\x0a\x5d\x32\x1e\xef\x4d\x05\xab\x4f\xf0\x53\x02\xd3\x78\xb5\xdf\x3d\xf6\x6b\xaf\xea\x9f\x47\x6b\xf8\x4a\xed\xb3\x8a\xa0\x68\xa2\x6b\x12\x48\xc6\x59\x22\x96\x96\x97\x48\xe2\xab\xa2\xad\x5e\xcd\xf1\xb8\x70\x5a\xd7\x2a\xa6\x92\xc8\xe9\x55\x48\x44\x12\xa1\xe9\xd5\x86\xde\xe7\xe1\xdc\xb7\x8e\x39\x66\x94\x48\xa6\x14\x72\x25\x19\x8b\x1e\x59\x55\x56\x71\xe3\x18\x91\x68\xe9\x07\xcb\xe5\xc5\xa3\xb9\x2c\x62\x1a\xcb\x58\xb3\x65\x34\x7f\x38\xe9\x75\xfc\xcf\x17\x36\xa8\x21\xb8\xb8\x7c\xdf\x75\x3a\x60\x56\x2d\xeb\xd3\xdb\x8e\x65\x9d\xf8\x27\xf0\xdb\x99\x7f\xde\x85\xfd\x5a\x1d\x7c\x8e\xa8\x20\xca\xd9\x50\x64\x59\xb6\x6b\x82\x39\x96\x32\x69\x58\xd6\x64\x32\xa9\x4d\xde\xd6\x18\x1f\x59\xbe\x67\xdd\x28\x5c\xfb\x0a\x38\xbf\xac\xca\x15\xc8\x5a\x28\x43\xf3\xd8\x68\xfe\x50\xad\x1a\x7d\x39\x8d\x30\x20\x1a\x82\x26\x12\x62\x4e\x94\x41\x55\xf7\x01\x0a\xb5\x68\x58\xd6\x88\xc8\x71\x3a\xa8\x05\x2c\xb6\x94\x0c\xa3\x94\x5a\x1a\x1d\x0a\x32\x7c\x55\x2d\x5a\x75\xae\x0e\x61\x18\x86\x3f\xc6\x70\xee\xf8\xd0\x25\x01\xa6\x02\xc3\x9b\x73\xc7\xdf\x33\x8c\x0e\x4b\xa6\x9c\x8c\xc6\x12\xde\x04\x7b\x70\x50\xdf\xff\x19\xce\x33\x8c\x86\x71\x81\x79\x4c\x84\x20\x8c\x02\x11\x30\xc6\x1c\x0f\xa6\x30\xe2\x88\x4a\x1c\x56\x60\xc8\x31\x06\x36\x84\x60\x8c\xf8\x08\x57\x40\x32\x40\x74\x0a\x09\xe6\x82\x51\x60\x03\x89\x08\x55\xfe\x8f\x20\x60\xc9\xd4\x60\x43\x90\x63\x22\x40\xb0\xa1\x9c\x20\x9e\x49\x88\x84\x60\x01\x41\x12\x87\x10\xb2\x20\x8d\x31\xcd\x02\x17\x86\x24\xc2\x02\xde\xc8\x31\x06\xb3\x9f\x43\x98\x7b\x9a\x48\x88\x51\x64\x10\x0a\xea\xd9\xfc\x91\x5e\x99\xb1\x54\x02\xc7\x42\x72\xa2\xb5\x50\x01\x42\x83\x28\x0d\x15\x0f\xf3\xc7\x11\x89\x49\x4e\x41\x81\x6b\xc1\x85\x21\x19\xa4\x02\x57\x34\x9f\x15\x88\x59\x48\x86\xea\x2f\xd6\x62\x25\xe9\x20\x22\x62\x5c\x81\x90\x28\xd4\x83\x54\xe2\x0a\x08\x35\xa8\xf5\x58\x51\x72\x58\x8c\x83\xc0\x51\x64\x04\x2c\x21\x58\x80\x96\x75\xc9\x9d\x9e\xa3\x58\x4f\x94\x42\x65\xae\x22\xa1\x46\x26\x63\x16\x17\x25\x21\xc2\x18\xa6\x9c\x12\x31\xc6\x1a\x26\x64\x20\x98\xa6\xa8\xbc\x59\x8d\xa8\xe9\x43\x16\x45\x6c\xa2\x44\x0b\x18\x0d\x49\xbe\x18\xd3\x46\x46\x03\xb5\x20\x0d\x16\x76\xa5\x4c\x92\x20\x53\xb7\x36\x40\xb2\xb4\x6a\xfe\x48\x8c\x51\x14\xc1\x00\xe7\x0a\xc3\x21\x10\x0a\x68\x45\x1c\xae\xc8\xab\x16\x4b\x12\x14\x41\xc2\xb8\xa6\x77\x57\xcc\x9a\x61\xf8\x67\x36\xf4\x7b\xa7\xfe\xa7\xb6\x67\x83\xd3\x87\x0b\xaf\xf7\xd1\x39\xb1\x4f\xc0\x6c\xf7\xc1\xe9\x9b\x15\xf8\xe4\xf8\x67\xbd\x4b\x1f\x3e\xb5\x3d\xaf\xed\xfa\x9f\xa1\x77\x0a\x6d\xf7\x33\xfc\xc7\x71\x4f\x2a\x60\xff\x76\xe1\xd9\xfd\x3e\xf4\x3c\xc3\x39\xbf\xe8\x3a\xf6\x49\x05\x1c\xb7\xd3\xbd\x3c\x71\xdc\x0f\xf0\xfe\xd2\x07\xb7\xe7\x43\xd7\x39\x77\x7c\xfb\x04\xfc\x1e\x28\x82\x39\x2a\xc7\xee\x2b\x64\xe7\xb6\xd7\x39\x6b\xbb\x7e\xfb\xbd\xd3\x75\xfc\xcf\x15\xe3\xd4\xf1\x5d\x85\xf3\xb4\xe7\x41\x1b\x2e\xda\x9e\xef\x74\x2e\xbb\x6d\x0f\x2e\x2e\xbd\x8b\x5e\xdf\x86\xb6\x7b\x02\x6e\xcf\x75\xdc\x53\xcf\x71\x3f\xd8\xe7\xb6\xeb\xd7\xc0\x71\xc1\xed\x81\xfd\xd1\x76\x7d\xe8\x9f\xb5\xbb\x5d\x45\xca\x68\x5f\xfa\x67\x3d\x4f\xf1\x07\x9d\xde\xc5\x67\xcf\xf9\x70\xe6\xc3\x59\xaf\x7b\x62\x7b\x7d\x78\x6f\x43\xd7\x69\xbf\xef\xda\x19\x29\xf7\x33\x74\xba\x6d\xe7\xbc\x02\x27\xed\xf3\xf6\x07\x5b\x43\xf5\xfc\x33\xdb\x33\xd4\xb4\x8c\x3b\xf8\x74\x66\xab\x21\x45\xaf\xed\x42\xbb\xe3\x3b\x3d\x57\x89\xd1\xe9\xb9\xbe\xd7\xee\xf8\x15\xf0\x7b\x9e\xbf\x00\xfd\xe4\xf4\xed\x0a\xb4\x3d\xa7\xaf\x14\x72\xea\xf5\xce\x2b\x86\x52\x67\xef\x54\x4d\x71\x5c\x05\xe7\xda\x19\x16\xa5\x6a\x28\x58\xa4\xe7\xe9\xfb\xcb\xbe\xbd\x40\x08\x27\x76\xbb\xeb\xb8\x1f\xfa\x0a\x58\x89\x38\x9f\x5c\x33\xaa\xd5\x63\xa3\xa9\x53\xe0\x4d\x1c\x51\xd1\x2a\x49\x6c\xfb\x87\x87\x87\x59\x3e\x33\xb7\x9b\x24\x54\x72\x6b\x99\x43\x46\x65\x75\x88\x62\x12\x4d\x1b\xf0\xd3\x19\x8e\xae\xb1\x24\x01\x02\x17\xa7\xf8\xa7\x0a\x2c\x06\x2a\xd0\xe6\x04\x45\x15\x10\x88\x8a\xaa\xc0\x9c\x0c\x8f\x60\xc0\x6e\xaa\x82\xfc\xa5\x6a\x31\x0c\x18\x0f\x31\xaf\x0e\xd8\xcd\x11\x68\xa4\x82\xfc\x85\x1b\xb0\xff\x73\x72\x73\x04\x31\xe2\x23\x42\x1b\x50\x3f\x52\xb9\x75\x8c\x51\xf8\x92\xf4\x63\x2c\x11\xa8\x8a\xda\x32\xaf\x09\x9e\xa8\x28\x32\x55\xf4\x4a\x4c\x65\xcb\x9c\x90\x50\x8e\x5b\x21\xbe\x26\x01\xae\xea\x9b\x97\x53\x16\x58\x73\x76\x95\x31\xab\xf8\xcf\x94\x5c\xb7\xcc\x4e\xc6\x6a\xd5\x9f\x26\x78\x85\x71\xd5\x8a\x58\xca\xb8\x47\xba\x12\x08\x2c\x5b\x97\xfe\x69\xf5\xd7\x17\x66\x5f\xbf\xba\x78\x39\x73\xdf\xd7\x8b\x34\x2d\xcd\xdc\xb1\x61\x34\x2d\xe5\x94\xea\x62\xc0\xc2\x29\x10\x89\x63\x11\xb0\x04\xb7\x4c\x53\xdf\xc8\xa9\xba\xce\x23\x4a\x04\x63\x1c\x23\x1d\x51\xb6\xaa\xee\xe7\xf3\xde\xf7\x59\x85\xac\x4e\xf0\xe0\x2b\x91\xd5\xec\x41\xcc\x98\x1c\x6b\xa0\xac\x36\x10\x24\x70\xb8\x9c\xa4\x7c\x43\x43\x57\x51\xf8\x25\x15\xb2\x01\x94\x51\x7c\x04\x63\xac\x2a\x53\x03\xf6\xeb\xf5\x7f\x1d\x41\x44\x28\xae\x2e\x86\x6a\xef\x70\x7c\x04\x3a\x02\xb2\x09\xf0\x03\x89\x55\xb0\x20\x2a\x8f\x60\x80\x82\xaf\x23\xce\x52\x1a\x56\x03\x16\x31\xde\x80\x1f\x87\xef\xd4\xef\xaa\xfa\x21\x41\x61\xa8\xb9\x52\xde\x30\x18\xe9\x99\x2d\x33\x9f\x69\x2a\x7d\x4b\x34\x78\x6e\xf7\x58\x11\x69\x4b\x39\x4a\x79\x07\x68\x4a\xfe\x82\x79\x0c\x40\x71\xf0\xcc\x99\xf4\x1a\x73\x85\x24\xaa\xa2\x88\x8c\x68\x03\x24\x4b\x8a\x8a\xba\xd6\x0f\x5a\xa6\x64\x89\x79\xdc\xb4\x64\xb8\x64\x34\xcb\xac\xe6\xbb\x7a\xfd\x99\x43\xa5\x94\xe9\x7c\x69\xd5\x80\x41\xc4\x82\xaf\x05\xdf\x8e\xd1\x4d\x35\x77\x92\x77\xf5\x7a\x72\x53\x78\x18\x44\x18\x71\x45\x50\x8e\x0b\xe3\x9b\x02\x65\xa1\x1c\x40\xa9\x64\x77\x42\xa2\xa0\x2d\xad\x28\x80\x66\x48\xae\x9f\xdb\xad\x8a\xf2\xde\x55\xce\xfd\x42\xcc\xf9\x56\x46\xd6\xc1\x9c\xdb\x59\x69\xc2\x84\x00\x47\x51\x3e\xbb\x65\xd6\xb3\x7b\x91\xa0\x60\x7e\xff\xac\x82\xe6\x0f\x39\x0a\x49\x2a\x1a\xf0\x56\x8f\x95\x24\x80\xe1\xb0\x90\xc5\x32\xb0\x06\xec\x27\x37\x20\x58\x44\x42\xf8\x11\x1f\xaa\xdf\x62\x62\x18\x0e\x57\x74\xb1\x0b\xd9\x61\xc9\xc9\xf3\x65\x89\x77\x1b\x03\xae\xa0\x5d\x0d\x32\xc9\x4b\xcd\x2f\xf5\xfa\x11\xe8\x12\x95\xcf\x0f\x30\x95\x98\x97\xd9\x4b\xff\xab\x6b\xa3\xac\xdb\xcd\x7e\xf7\xcb\xc1\x41\xa7\xbc\x00\x1d\x28\xbf\x36\x21\x8f\xb7\x8c\xc0\xaa\xf5\x32\xd8\xf2\x88\x9c\xff\x2c\x77\x40\x17\x5b\x9f\xa0\x5f\x96\x94\xbe\x4b\xda\x83\x7d\x98\xcd\xc4\xe2\x85\x07\x0c\x19\x87\xe5\x2e\xdd\x86\x5d\x52\x98\xcd\xee\x50\x85\xd5\x3d\xbb\x56\x61\xc7\x6e\x6d\x5a\xfe\x6a\xa5\x60\xfc\x45\x0e\x5e\xdc\xf3\x57\x37\xdd\xa6\x98\x2d\x9d\x67\x3f\x73\x9e\xfb\x7c\x63\xe7\x73\xdf\x46\xb5\xef\x96\x13\xec\xba\x2b\xd4\xa1\x3e\xcf\x25\xf7\xb9\x43\x2e\x06\x82\x31\xc7\xc3\x96\xb9\xcd\x4b\xf7\x67\xf6\x87\x79\xd2\x3c\x3d\x3d\xcd\x93\x6f\x88\x03\xc6\xf5\x3b\xb9\xf9\xf2\xa0\xb0\x20\x38\x50\xcb\x81\x42\xde\x1e\xb0\x28\x2c\x4f\xdc\x41\xca\x85\xc2\x9e\x30\x92\x0d\x2c\x1a\x0a\x42\x35\xd2\xbc\xaf\xb8\x93\xe0\x7f\x51\x8c\x69\x7c\xfa\x25\xea\x90\xf1\xb8\x01\x01\x4a\x88\x44\x11\xf9\x0b\x97\x26\xfd\xb7\x3f\xff\x8a\x43\x54\x52\xaf\xd7\x66\xe4\xc3\x5a\xcb\x8d\xac\x90\x2f\x06\x17\xdd\x5b\x72\x93\x9b\xf7\xf8\x23\xc1\x13\x20\x14\x1e\x7c\x3b\xde\xb4\x50\xa9\x0f\xdf\x49\xbc\xe5\xe9\x37\xfb\x79\x68\xf3\xa3\xa4\x28\xbc\x86\xec\xdf\x13\xb2\x42\x72\x46\x47\x2f\xa7\xda\xdf\x37\x9f\xb3\xfa\x23\xdf\xf9\x6a\x5a\x19\x93\xdf\xc1\xeb\x4a\x1a\x86\xfc\xc9\xfc\x30\xd1\xdd\x2d\xb4\x57\x3f\xfc\x67\xf8\x61\xd6\x9a\x2e\x5c\xad\x39\x78\x39\x33\x83\x55\xae\xa3\x07\x4e\xd1\x6d\x3e\xea\xf6\xc2\xc2\x6c\x8e\x3b\x28\xa9\x05\xcb\x4d\xf4\xac\x12\xbc\xb8\x67\xac\x70\xb4\x2b\xee\xf1\xa0\x46\x1f\x3c\x1a\xf9\x3f\xea\x2c\xab\x1d\xe6\xdd\xb3\x9a\x2f\xd4\x50\xce\xdb\xad\xb5\x9e\x32\xa5\x21\xe6\xaa\xfb\x2b\xba\x53\x76\xda\x54\x35\x51\xbb\x97\x63\x9e\x56\x4d\xb7\x6c\xef\x56\xcf\x9a\x94\x9a\xf7\xb5\x2b\xdc\x99\x6a\xbc\x73\x9e\x09\xd0\x1c\xef\x20\x4f\x3b\xa7\xa7\xc7\x44\xf0\x7d\x1d\xf1\x6b\x60\xfd\x7f\xb6\xb9\xab\xcb\xad\xc5\x99\xbd\xe5\x82\x6b\x3e\xf4\x02\x4b\xae\xd5\x13\x84\xaf\xde\xf8\xcf\xf0\xc6\xd7\x45\xd7\xeb\xa2\xeb\x75\xd1\xb5\xeb\xce\xf2\xba\xe8\xda\x99\x96\x6d\x93\xa1\x9a\x96\xde\x8f\x3b\x7e\xc4\x56\xe8\x02\x64\x39\xf2\xec\x27\x31\x0a\x47\x93\x56\x4e\x9a\x2c\x0d\x7d\x78\x78\x78\xdf\x06\x77\x71\x67\x77\x7d\x4b\x72\x37\x9a\x86\x5d\x6a\x5f\x9e\xb3\x75\x39\xd8\xd8\xba\x94\x6e\xa2\x3d\x64\xf2\x95\xde\xe6\xce\xb9\x86\xe2\x29\xac\xd5\x74\x55\xfc\x9a\xfc\xf9\x1c\xe2\x60\x35\x5b\x69\x89\xb6\x4e\x55\x98\x4a\x18\x4c\xb7\xdb\x87\x5b\xcf\x1d\x6b\xe7\x1d\xee\x66\x86\xa6\x15\x92\xeb\xe3\xec\x7f\xa3\x98\x26\x76\xad\xad\xdd\x70\xbc\x2e\x13\x71\x99\xbf\x9a\xd6\x80\x85\x53\x35\x32\x96\x71\x74\x6c\x18\xe5\xdf\xef\x24\xa9\x18\xb3\x6b\xcc\xbf\xc3\xd7\xda\x6b\xa8\xfe\xfe\xef\xc1\xbe\xcf\xe7\x60\xdb\x7f\x0d\xf6\xfd\x3e\x06\x5b\xa1\xb9\x85\x26\x97\x9f\x5c\x3f\xe2\xb3\xca\xff\x06\x00\x00\xff\xff\x96\x70\x0b\x9a\x8b\x42\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 17035, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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
