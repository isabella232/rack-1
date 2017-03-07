// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\x7b\x73\xdb\x38\x92\xf8\xff\xf9\x14\x28\xfe\xf2\x3b\xd9\x33\xb2\xfc\x98\xdb\xba\x5b\xce\xe5\xaa\x1c\xc5\x99\x78\xd7\x4e\x74\x92\x93\xa9\xdb\xc4\xb5\x05\x93\xb0\xc4\x35\x05\x70\x00\xd0\xb1\x47\xa5\xef\x7e\x85\x07\x49\x3c\x29\xf9\x91\x9b\xd9\xab\xc9\xd5\xcd\x46\x64\xa3\xd1\xe8\x6e\xf4\x0b\x0d\x66\xb5\x02\x39\xba\x2e\x30\x02\x09\xac\xaa\x04\xac\xd7\x2f\x00\x58\xad\xc0\x4b\x58\x55\x20\x7d\x05\x46\xc7\x55\xd5\x3d\x5c\x42\x5c\x5c\x23\xc6\xe5\x9b\xf3\xe6\x87\x7a\xfd\x02\x00\x00\x92\xe3\x9f\x67\x17\x68\x59\x95\x90\xa3\xb7\x84\x2e\x21\xff\x84\x28\x2b\x08\x4e\x40\x0a\x92\xa3\x83\xc3\x83\xbd\x83\x3f\xef\x1d\xfc\x39\x19\x2a\xf0\x31\xc1\x79\xc1\x0b\x82\x59\x92\x6a\x14\x72\x26\xae\x71\x80\xe4\x0a\x96\x10\x67\x88\xee\x65\x1d\xa8\x3b\xb7\x37\xa8\xa2\x24\x43\x8c\x6d\x1a\x93\x9c\x62\x8e\x28\x86\xa5\x98\x1c\x24\x6f\x71\x9a\x9e\xfc\x52\xc3\x52\x10\xf3\x59\x3c\x99\xa2\xeb\x24\x35\xc0\xc0\x7a\x08\x92\xff\x46\x2c\x01\x97\x60\x3d\x6c\xb0\x4c\x68\x71\x0b\x39\xda\x80\xa4\x81\x0a\xe3\x98\xa2\x79\x41\xf0\x3b\xc8\xce\xe0\xf2\x2a\x87\x1f\xab\x3c\x82\x51\x0f\x00\xcd\x9b\xb7\x05\xce\x4f\xf1\x39\xac\xe4\x74\x1a\xcf\x98\xe0\xeb\x62\x9e\x0c\x8d\xe9\x8f\x7f\x9e\xa5\xa9\x7a\xab\x48\xb0\x26\x32\x69\x01\x8a\x3c\xfd\xcb\x24\xf2\x75\x09\xf1\xcd\x0c\x65\x35\x2d\xf8\xfd\x4f\x94\xd4\x95\x10\xeb\xca\xa4\x10\xa4\xe0\xf3\x4a\xce\x29\x04\x6e\xc3\x8a\x59\x93\x4b\xc5\x7c\x8d\x34\x99\x40\x0a\x97\x88\x23\x2a\x87\xf6\x6b\x40\x25\x60\x1f\x20\xfd\x20\x7c\xb3\x96\x71\x59\x33\x8e\xa8\xa1\x76\x00\x24\x17\xf7\x15\x52\x84\x73\x5a\xe0\x79\x62\xb0\xe4\x0d\xba\x86\x75\xc9\xe5\x5b\xfb\x39\xcb\x68\x51\xf1\x46\xc7\x1b\xbe\x75\x5c\x1b\xd7\x8c\x93\xe5\x05\xa9\x8a\x6c\x5a\x63\x5e\x2c\x51\x60\xd2\xde\x39\x53\x90\x60\x92\xa3\x7f\xb0\xd8\xc4\xc1\x79\xdf\xa0\xaa\x24\xf7\x4b\x84\xf9\x39\xbc\x2b\x96\xf5\x32\x3c\xed\xfb\x7a\x79\x85\x68\x64\xda\xa3\x83\x83\x9e\x39\x35\x5e\x50\x21\x9a\x21\xcc\xe1\x1c\x01\x72\x0d\x34\xfb\x11\x03\x9c\x80\x1b\x84\x2a\x40\x6b\x8c\x0b\x3c\x07\x5f\x17\x45\x89\x40\x2e\xe9\x12\x4b\xed\x23\xb9\xc0\x8f\x24\xf9\x4f\xbd\x14\x2b\xb4\xcf\x47\xf1\x09\xbe\x2d\x28\xc1\x82\xe4\x47\x48\xf5\x41\xf2\x34\x8d\xd5\x76\xf3\x58\x08\x3f\xe0\xf2\x1e\xc0\xb2\x24\x5f\x01\xcc\xc4\x72\xc5\x62\xf9\xa2\x60\x40\x98\xfa\x6b\x4a\x96\xa0\xc0\xac\xc8\x11\xe0\x0b\x04\x3e\x4d\xc6\x11\x9a\xdf\x13\xf3\xc5\xb1\x40\x88\xf2\x4f\xb0\xac\x91\xb2\x78\xd2\x78\x0c\x25\x1c\xb8\xf4\x16\xf1\x57\x74\xff\xad\xf9\x64\x98\xe3\x47\xb0\xe9\x23\x43\x60\x56\x5f\x61\xc4\x99\x46\x24\xf8\xc4\x2a\x94\x15\xd7\xf7\x82\x2d\x7b\x92\x47\x25\x81\x39\x68\x2c\x13\x40\x38\xaf\x48\x81\x39\xfb\x26\x3c\x9b\xa2\x12\x41\x16\x5a\xd0\x73\x9b\xaa\x29\xaa\x08\x2b\x38\xa1\x21\x21\x3d\x6d\xb2\x19\xa9\x69\x86\x40\x46\x72\x04\x68\x37\x8d\x47\x82\xed\x32\x9e\x9b\x8a\x8b\x05\x02\x67\x96\xe8\x98\x9e\x0f\xcc\xc5\x84\xe0\x9a\xd0\x76\x53\x04\x88\x53\x8a\x11\x21\xeb\xac\x60\xfc\x3f\xa4\x9f\x3d\x19\x1f\xa5\xa9\x02\x4e\xd3\xd3\xfc\x3f\x1f\x43\xea\xa7\xc9\x18\x30\x35\xdf\x76\x54\xc5\xf5\xfe\xdb\x10\x57\xe9\xed\xb1\x1d\x91\x17\x90\xdd\x4c\x49\xf9\xfc\x5a\x7c\x7a\x7c\x0e\x04\x62\xb1\x4d\x61\x55\x95\xf7\xe2\x2f\x27\xe3\x19\x10\x33\x32\x61\xde\xe3\x44\x35\xc1\xa9\x45\x93\x63\x10\x76\xa6\x27\xff\xf5\xf1\x74\x7a\xf2\x66\x17\xa8\x90\x09\x18\x0e\x1d\xbc\x83\x38\x2f\x11\x05\x7a\x8f\x82\x06\xa3\x41\xf0\x79\x81\xcf\x10\x9e\xf3\x85\x24\xf7\xd0\x7c\xe5\x58\x25\x9f\xbe\xc9\x38\xc2\xaf\x4e\x92\x9f\x26\x63\x21\xc6\xc7\x4a\xb1\x5f\x6a\x9f\x26\xe3\xf1\xe9\x9b\xe9\xb3\x0b\x4d\xcc\x2c\x10\x87\xa7\xb7\x22\xc4\x73\x58\x55\x05\x9e\x9b\x9b\x2e\x99\x10\xca\x27\x94\x70\x92\x11\xc7\x1d\x2e\x38\xaf\x54\xd8\x2c\x14\x1e\x61\x44\x0d\xb8\xe4\xdd\xc5\xc5\x44\xd8\xd9\x53\xcc\xb8\xd8\xfe\xa1\x77\xd2\x00\xa1\x18\xc4\x2c\xb1\x02\x65\x31\x1d\xeb\x9f\x6f\xf6\xe4\x09\xad\x19\x79\xd6\xb3\xbe\x8b\x71\x74\x79\xfa\x55\x7c\xb2\xd9\xec\xcc\x9d\xaa\xec\x59\x9a\x00\x7f\xda\x54\x60\x1d\xf0\x3f\x46\xf6\x62\xc9\xb5\x66\x7b\x08\x32\xbe\x77\xa8\x29\xb2\xb3\x24\x95\x4f\x99\xb4\x37\xf0\x47\x61\x78\xe1\x68\x1d\xf0\xaf\x28\x8e\x3e\x06\x1e\xc1\xee\x51\x83\xea\x5e\xf4\x51\xf8\x2d\xa9\x47\xf5\x9e\x88\x66\x29\x2c\xb7\x9e\x01\x56\x7b\x98\x50\xbe\x78\x10\x53\x61\xb5\xc7\x48\xfd\x94\x41\xdb\x32\x2c\x83\x9b\x56\x24\x79\x10\x34\x18\x53\xc4\x64\xac\x61\x59\x0c\x33\x09\x0b\x7b\xa1\xc6\xa8\x9e\x1e\x9f\xa7\xa9\x84\x31\xe8\x99\x50\x52\x21\xca\x0b\x64\xfb\x7e\x11\xcc\x31\x56\x2f\x91\x80\x9f\x90\xb2\xc8\xee\xdf\x90\xac\xf6\xb2\x01\x60\x3b\x9b\xe4\xe8\xe0\xf0\x68\xef\xf0\x60\xef\xf0\xdf\x8c\x49\x24\xd0\x8c\x43\x8e\xf4\xf8\xcf\xd6\x2b\xe0\xe0\x93\xe0\x27\xd7\xd7\x28\x93\x21\xa6\x0c\x2a\x1d\x6c\x9a\xf4\x02\x67\x45\xd5\x14\x39\x66\x88\xde\x16\x19\x52\x61\x67\x29\x99\x3a\x82\x4b\xf8\x2b\xc1\xf0\x2b\x1b\x65\x64\xe9\xd4\x02\xba\x85\x66\xda\x23\x7e\x06\x09\xe3\x2c\xed\x16\xde\xc5\xac\xcd\x9f\xb5\xf5\xdb\x7c\x6b\x61\x4e\x26\x90\x2f\x04\xf1\xfb\x19\xc1\xb7\xe4\x6e\x3f\xb1\xdf\x0a\x86\x2a\x96\xdb\xac\x70\x19\xa1\x20\xef\xdf\xc3\xa5\x12\x63\xbe\x2c\x70\xc1\x38\x85\x9c\x50\x8f\x25\xc9\x06\x39\x81\x6d\x65\x05\x3c\x79\x09\xfe\x7a\x12\x31\x38\x97\x7c\x27\x7e\x36\xfa\xa9\x1e\x80\xf5\x06\xee\x99\xbf\x3a\x48\xdf\x76\x1a\x1a\xde\xa3\xdd\x6a\x1b\xa5\xe9\xdb\x1a\x2b\xaa\xb6\x52\xf2\x31\xc9\x91\xaf\xd0\xb3\x1f\x5e\xd7\xd9\x0d\xe2\x5d\x99\xea\x2f\xa4\xd0\x1a\xb2\x97\x0c\xc5\xff\x28\xb9\xf6\x14\xa2\xc0\xa5\xaf\x6e\xc9\xec\x07\x9d\x26\xba\x58\x15\x52\xaa\x62\xad\x7d\x0b\x6d\x5b\x67\x5c\x0f\x41\xb2\xaf\x14\x7b\xff\x5a\x96\x20\x0b\x82\x47\xbf\x16\x55\xa2\xe6\x8a\x2a\xa3\x0e\xe5\x04\xb2\x02\xe7\xe8\x6e\x84\xee\x74\xc2\x6d\x81\x9d\xa3\x25\xa1\xf7\xb3\xe2\x57\xc9\xd4\xc3\xa3\x7f\xb7\x5f\x37\xd6\x45\x91\xfe\x13\xe2\xc7\x5c\xe9\x86\x67\x82\x84\x66\x50\xec\x6d\xb7\xc4\x28\x12\x29\x24\xa7\xd7\x66\x6d\xcf\xad\x11\x3a\xbc\x53\x85\xa2\x7f\x1d\xfd\xe0\xbc\xe8\x38\x15\xa8\x47\xd9\x4a\xe6\xd0\x73\x51\x2c\x11\xa9\xa5\x4a\xff\x70\x70\x90\xc4\x55\x30\x5c\xb5\xa3\xad\x39\x06\xa3\x48\xc1\x2e\xa3\x04\xff\x83\x5c\x6d\x03\x0a\xe7\x08\xf3\x6d\x00\x9b\x22\xa0\x09\xba\x65\xdd\x90\x29\x13\xd9\x83\x9c\xa2\xb9\xb0\x2f\xf7\x51\xec\xa1\x41\x4d\xa6\x99\x44\x90\x32\xae\xca\xc3\xb6\x37\xfb\x50\xf3\xaa\xe6\x9b\xcb\xe3\x44\xc3\x81\x51\xff\xe2\x3a\xb8\x0d\xdc\x68\xd7\x18\x1e\xd1\xe5\xeb\x9c\x3b\xe1\xb9\xb0\x9f\xb0\xac\xb5\x06\x6b\xad\x6b\xe1\x5c\xaf\xfd\x42\xfc\xff\x6a\x05\x10\xce\x25\x5e\xe3\x44\x22\x54\xc6\x6f\xce\x22\x28\xc4\x73\x04\x5e\xde\xc8\xa3\x88\x13\xcc\xa9\x34\xff\xac\x59\x4c\x72\x82\xe1\x55\x89\xf2\xd5\x0a\xd4\x55\x85\xa8\x80\x5c\xaf\xbb\x3d\xf5\x9e\xc8\x5d\x19\x2c\xd8\x8b\x27\x33\x54\x2a\x33\xfe\x19\x1c\x98\x66\xc6\xc6\xf7\xb6\xb1\x2f\xca\x92\x09\xd3\xb3\x77\x28\x77\xb4\xde\x44\xdd\xba\xfa\x57\xd8\x94\xaa\x9d\xd5\x21\xb9\x3a\xed\xb2\xbb\xb5\x75\x44\xa0\x91\x58\xb5\x45\x89\x11\xf4\x34\x96\x7f\x4c\x96\x4b\xf8\x06\x95\xc5\xb2\xe0\x28\x17\xd1\x7c\x62\x94\x5c\xdb\x2a\xd5\x6a\x25\x10\xea\x07\xb2\x30\x2f\xa6\x34\x41\xad\xc4\x58\x55\x61\xbd\xfa\x29\xad\xf1\x10\x8c\x27\x1f\x41\x8d\x0b\xae\x9e\x20\xb1\xa3\xd0\x10\x40\x9c\x83\xf3\xd7\x62\xc4\xf4\xf8\xdc\x78\x93\x74\x1a\xbf\x2d\xc3\x5a\xa5\x94\x3c\x49\xce\xc8\xdc\x2e\x18\x05\x34\xb0\x85\x51\x3a\x37\xdc\x30\x83\xb1\xb5\x63\x73\xd8\x9e\x95\xcc\x99\xfc\xaf\x02\xda\x66\x8a\xce\xd0\x6c\x75\xc2\x16\x39\x95\x2b\xae\xbb\x61\xa3\x77\x90\x4d\x5a\x69\x68\x7d\x71\xf4\xa9\x03\x76\x15\x2b\xac\x5a\x27\xe3\xd9\x05\x64\x37\x6f\x04\xf1\x05\x0f\x94\x4b\x2a\x84\x73\xf6\x41\xba\x68\x2b\x0a\x19\xb6\xd1\xa6\xf4\x77\x97\x81\xc2\x87\x02\x4f\x53\x7f\x0e\x03\xd8\x08\xc6\x0e\x47\x07\xdb\x45\x2c\x66\xb9\xa9\xd5\x80\xf6\xa1\xe3\xe1\x34\x95\x17\xe4\x06\xe1\x8d\xbe\x3b\xea\xb7\x75\xf8\x19\x09\x85\x9c\x00\x68\xc6\x61\x76\x23\x47\x48\xab\xa1\xb6\x9e\x66\x78\xe2\x07\x45\x66\x0d\xb8\x45\xd4\x3c\x73\x40\x9d\x23\x89\x16\xdc\x7c\xee\x0c\x69\xc3\x2d\x0d\x2a\x7e\xbb\x51\x00\x64\x37\x5b\x44\xe2\x4d\x0c\x6e\x2f\xc8\x8b\xc1\x4f\x97\x70\x6e\xc0\xc9\x9f\x21\xc0\xd5\x4a\x68\x37\x1a\x49\x0b\x86\xf3\xd1\x31\xa5\xf0\x7e\xbd\xf6\xe3\x70\x0d\x10\xc8\x9a\x80\xb5\x03\x64\x64\x37\x04\x2f\x51\x29\xa3\x76\xb9\x1f\x36\xa3\x37\x89\x91\x18\xd6\xeb\xe1\x6a\x85\x4a\x86\xd6\xeb\xd5\x0a\xe1\x3c\x3a\x26\x59\xad\x9a\xb9\xd6\xeb\x24\x48\x5a\x78\xf8\xa5\xcf\x0a\x31\x9f\xd8\xed\x18\x99\x34\xab\x22\x1c\x48\x92\x7e\xb6\xac\x56\xe0\x56\x98\xc4\xc0\xd0\xb5\x97\xee\x85\x89\x4a\xc6\x55\xdd\x29\xb8\xe1\x21\x0f\xc3\x1e\x32\xe0\x9c\xb4\x9b\x74\x11\xab\x98\x3a\x88\xfb\xe8\xa9\xb8\x63\x27\x74\x2d\xc0\xf1\x64\xd2\x68\xa2\xb0\xab\x51\xa5\x15\xbb\xf0\x78\xfc\x57\x0d\x8b\xf0\xad\xfe\x1d\x81\x3d\xfe\x79\xf6\xf7\xe9\xc9\x4f\xa7\x1f\xde\x9b\x23\x8c\xa7\xe1\x71\x86\xb1\x1e\x19\xa4\x83\xa0\x86\x29\xeb\x8f\x11\x18\x49\x77\x07\x92\x24\x0c\xa7\x0c\x7b\xbb\x32\x43\x1b\xf4\x40\x5f\x05\x34\x7a\xed\xbd\x1e\xf4\xa6\x09\xcd\xd0\xfd\x10\xbc\x54\xd3\x08\x9f\x75\x56\xe0\x9b\x4f\x90\xb2\x30\x89\x92\xff\x37\xe8\xbe\xa5\x4f\x8f\x0c\x51\x16\x9f\x3d\x39\xfb\xf0\xd3\xdf\x7f\x9a\x7e\xf8\x38\x89\xb9\xfe\x50\x85\x64\xfa\x61\x7c\x32\x9b\xf9\x66\xcb\xcd\xcb\x3d\xdd\xfa\x44\xca\x7a\x19\x28\x50\x00\x47\x90\xe7\xa4\xc6\x5c\xc4\xa3\x7a\x40\x9f\x34\x47\xa7\x6c\x76\xcf\x38\x5a\xf6\x88\x72\xf4\x8e\x30\xbe\x5e\xa7\xab\xd5\x68\x4c\x30\x87\x05\x46\x34\xa8\x4d\x8a\x57\xc2\x6e\x44\x90\x45\x52\xec\xfd\x5b\x45\xe8\xbe\x9f\xba\x3b\x9e\x6b\x5f\x18\x38\xc9\x31\x61\x0a\x23\x84\x85\xb2\xfc\x8e\xbc\x07\xab\x18\x68\x1b\x4c\x24\x45\xef\x89\x0a\xf6\x80\x0b\xea\xd9\xd0\xe4\xe4\x8e\x53\x28\x68\xdc\x24\x33\x47\x79\xc5\x66\x6c\x87\x9e\xc3\x2a\x22\xc0\xb0\xbc\xc4\x20\xd3\x2f\x6a\x2d\x0f\xb1\x43\xb8\xc6\xea\x38\xcf\x29\x62\xac\x01\x6f\xf6\x41\xc8\x7b\x3c\x68\x73\x3c\x81\x6f\x4d\xa4\x18\xe6\xda\xe3\xf1\x4e\x08\xe5\xc6\x01\x4f\x8f\x44\x46\x02\xb4\x77\xe3\x88\xec\x62\x07\xfd\x02\x46\xcd\x51\x83\x3a\x2c\xd9\x05\x3b\x2f\x91\x88\x8b\x5f\xeb\x54\x79\x37\x82\xc4\xdd\x09\xa9\xd8\x0a\xb1\x4d\x13\x77\x48\x82\x4e\x61\x6c\x9b\xe9\xc0\x7a\x2d\x14\x20\x68\x7a\x34\x17\x04\x78\xbb\x5d\xc0\x7a\xbd\x2f\x1e\xb4\xab\x08\x4b\xbe\x77\x47\xf5\x6c\xf8\xc4\xa1\x2d\xdd\x38\xf9\xef\x60\xdb\x4e\x68\x71\x5b\x94\x68\x8e\xf2\xce\x48\x77\xcf\x3c\x0e\x6d\x5b\x32\xd5\x5a\x1d\x10\xa2\x9d\xe2\xb4\x5d\x8a\x2a\x8c\x76\x0a\x0b\xa1\x30\xd7\xce\x8a\x5e\x58\xec\x51\x11\xad\xa1\x90\x2e\xb3\x12\x33\x11\x6c\x65\xd5\x94\x8d\xe5\x64\x91\xd8\x3a\xc4\x7c\x3b\x77\x09\xa4\x3d\x32\x23\x7a\x11\xe2\xbe\x9d\xe0\x9e\x8c\x85\xf5\xd7\xc7\x07\xdb\x95\x8d\xbb\x1e\xbb\xae\xfa\xa8\x9f\x39\xe9\x45\xd7\xf9\xa5\x4e\xff\x6a\xea\x56\x33\x34\xa0\xee\xe0\x7a\x87\x60\xc9\x17\xf7\x13\xd5\xc7\x65\x56\x2f\x9c\x0e\x32\x5f\x83\x9b\xb6\xb5\xbe\xb1\xba\xb1\xcd\x56\x2c\x97\x62\x56\x50\x94\x8f\x85\x6b\x0f\x46\xae\x91\xba\xd1\x56\x91\xeb\x56\x6a\x72\x46\x60\xde\xbc\x0c\x59\xcf\x40\xac\xdb\x6e\xf4\xed\xf2\x34\x73\x84\xb0\x6b\x7a\xc4\x8e\xcc\x81\xe4\x26\x94\x46\xf9\x60\xd7\xb2\x20\x21\x34\x26\xad\x5d\xa2\xdc\x31\x67\x7b\x7d\xf7\xcc\x85\x73\xd8\xe4\x08\x3b\x5e\x8f\x37\xd5\x3f\x92\xd3\x07\xf7\x93\x5f\x0c\xe9\x93\xb2\x5f\xd9\x30\x08\x76\x4c\x92\x39\xdd\xa6\x42\x58\xb0\x2f\xda\x2e\x1f\xb6\xac\x34\x2b\x41\x2f\x75\xf1\x49\x92\x97\xbe\xd2\xf4\x8e\x26\xc6\x53\x03\xb8\x99\x65\x42\xd1\x75\x71\x27\xe0\x2b\x5a\x60\x7e\x0d\x92\x06\xf7\xff\x67\x89\x8d\xd3\x2d\x3a\x8d\x4c\x1f\x6f\x54\x9a\x64\x33\x71\x60\x8e\xa0\x07\x1d\x0b\x03\x73\x5d\x64\x5e\x7f\x53\xb4\xdd\xda\x5d\xea\x46\xb4\x32\x9a\xf5\xda\xef\x1e\x25\x92\x70\x35\x37\x2c\x8e\xb6\x11\x4d\xa4\x48\x5b\x33\xaf\x53\xb4\x66\xbc\x23\xc1\x87\xf0\xf0\x9b\xb4\x12\x3e\x86\x42\x19\x27\x3d\x86\x34\x61\x2e\x95\x49\x6a\x27\x9b\x42\x9c\x93\x25\x03\x3b\x05\x27\xb0\x9b\x65\xd7\xf3\xd3\xbd\x0b\x79\x94\xf8\xed\xda\x74\xac\x6c\xab\x05\x7c\xee\xda\xbd\xcd\xda\xd1\xee\xbd\x96\xc7\x0e\x6b\x1d\x3e\xf6\xc7\x2f\xce\xd8\xae\xdc\x6f\x54\xd0\x5d\xd3\x29\xe4\x66\xd9\x67\x31\x0e\x24\x6f\xde\xcf\x54\x62\x78\x69\xb7\x64\x7c\x13\x75\x6e\xfe\xfa\x90\x50\x2d\x82\xdd\x2a\x36\xeb\x55\xbb\x71\xf2\xf3\x68\xb8\xeb\x02\xbf\x01\xe1\xce\xa1\x71\xdb\x35\x6e\xf9\xb5\x48\x8d\xda\xd0\xb8\x91\xeb\xb1\x01\xa7\x35\x92\xaa\xdc\xb8\xb6\x21\x48\x70\x20\xb5\x7f\x20\x96\x76\xe4\xe5\xb3\xb8\x43\xf7\xd8\xe6\x1b\x6c\xbf\x80\xf6\xc7\xba\x96\x9f\x28\x56\x37\x00\x3f\x12\x01\xa6\x39\x53\x0b\x18\x09\xc2\x13\x09\x66\x1f\xdc\x79\xc2\x07\x5b\x1c\x50\xec\x35\xa4\x7a\x02\xb7\x3b\xb6\x4f\xf1\x5c\xd7\x2e\x9c\xac\xa7\xd7\x00\x68\x28\x37\x7e\x55\xf5\xb0\x49\x7d\x55\x16\x99\x9f\x51\x26\xe3\x22\xa7\xa7\x82\xdb\xc9\xc1\x48\xfe\xdf\xfe\x41\xe0\x00\x21\x92\x0d\x77\xa3\x8d\x1e\x12\xdd\xed\xea\x27\xd6\xb1\xac\x36\x39\xad\xcc\xc6\xc6\x4d\xa9\x73\xf2\x96\x92\xa5\x11\x48\x5b\x06\xc6\x03\xbe\x20\x31\x50\x3b\xdb\xdd\x14\xb1\x3a\x92\x0d\xe4\xdd\x66\xce\xf7\xa9\xca\x4e\x73\x97\x2d\xb1\x5e\xbb\x90\x23\x08\x1c\x40\x2b\xf5\x2d\x21\xe3\x45\xd6\x59\x84\x02\xcf\xd3\xd4\x34\x10\x9d\x3a\x3f\xce\x63\x59\x49\xf7\x16\xfb\xb4\x5b\x77\x6c\xff\x28\x15\x44\xbf\x80\xd1\x2c\x5b\xa0\x25\x02\x49\xd1\xdd\x05\xb4\xb2\x02\xf5\x5e\x35\x0f\x85\xda\x86\x8c\xab\x03\xb6\x81\x6e\xda\xf6\x63\xed\x3a\x4e\x77\xbf\xab\x9b\x1e\x60\x4f\x33\x4f\x70\x33\x74\x94\x47\x3d\x87\xb9\xa6\xb8\x36\x79\x67\x96\xd1\x25\x9f\x86\xb0\xf9\xeb\x0c\xae\xcd\x5f\x91\xad\xee\x42\x75\x30\x92\xbd\x6d\x6f\x28\x2c\x70\x81\xe7\xaa\xe1\x4f\x91\xa1\x75\x29\x49\xa5\x23\x1a\x5a\xdd\x4d\x42\x5f\x9a\x31\xfa\xb1\x4a\x38\x87\x21\xec\x66\xe3\x0b\x48\x4e\xf3\x12\x39\xa8\x8c\x47\x3e\x1a\x4a\x18\xfb\x1b\xc1\xa8\x21\xa4\x7b\xa5\x2a\x1b\xe3\x05\xca\x6e\xdc\x7a\x8a\x2e\x7a\x5c\x2c\x28\x62\x0b\x52\xca\x62\xd8\x91\xad\x66\x92\xb5\xb7\xb0\xb5\x46\x6a\x48\xf3\xd4\x35\x33\xc9\x05\xa4\xf3\x70\x03\x9f\x57\xfc\x34\xd0\x19\x26\x2e\x8d\xea\x6d\x6c\xbb\x36\x51\x91\x46\x45\x28\x8f\x55\x48\xcd\x19\x21\x5f\x38\x86\xcf\x3f\x23\x77\xf8\xaf\x46\x1a\x12\xb0\x80\x3f\xe2\x45\x90\x9b\x5d\x6e\x6e\xc8\xa4\x69\x80\x7f\x4e\xbf\x66\x39\x7f\xc5\xce\x51\xf0\x48\xcb\x71\x30\x46\x6c\xe5\xb4\xe5\xcb\xf1\xdb\x7b\x40\x1b\xb5\xb3\x45\x65\x76\xee\xe5\x19\x8f\xcc\x36\x87\xdd\x9d\x80\xd9\x59\xb0\xbf\x38\xea\x5f\x4d\xf7\xb0\xb5\xe3\x0c\xdd\x39\xb0\x38\xe7\x02\x84\x39\xd7\xe1\x51\x13\x87\xea\x3c\x0f\x4c\x6e\x03\x67\x88\xb3\xd9\x99\xc1\xab\xc6\xf5\x7e\x3b\x59\x78\x5a\x10\x35\xe8\x7d\xa0\x4f\x25\xc3\x2f\xfa\xbb\x6d\xcf\xcf\x1b\xd5\x44\x7a\xc9\xb7\xdc\xc0\xfe\x86\xbd\xbb\xef\xdb\xb5\x81\xca\xab\xdd\xa2\xae\xdc\x90\x85\x27\xd8\xbb\x2f\x07\x35\x51\x94\x05\x6e\xbc\x0a\x35\x1f\x70\x4e\x8b\xab\x9a\xab\x05\x47\xce\x25\x1b\x62\x36\x91\x01\xac\xbc\x58\xb8\x2b\xff\x40\x6a\xed\x1d\xda\x38\xfb\x87\xe9\x76\xcb\xa7\xef\x20\xaf\x5d\x7e\xe8\x0a\xcb\xd7\x95\x27\xeb\xcf\xd9\xeb\x31\x21\x37\x05\x9a\xf1\x22\xbb\x29\x30\x62\xac\x8d\x2a\xc4\xaa\x6c\xe9\xc2\x6b\x59\xec\xbd\x4f\x2c\xb6\x44\x6a\xe0\x4f\x48\xd6\x9f\x94\xa3\x3f\x30\x35\x8f\x25\x7c\xfa\x93\x0f\xed\x3a\x80\xbd\xc1\x42\xdf\x8c\xb0\x08\x69\xcf\x04\x37\x46\xeb\xeb\xf0\x38\x07\xa8\xa3\xb9\x55\x14\x23\x65\xd9\x54\x46\x08\x74\xa2\x1b\x6d\x96\xb2\x3d\x68\x4c\x09\xfe\x0b\xb9\x62\x7e\x03\xb5\x88\xea\xb0\x73\xbb\x68\xd3\xdd\xa2\x68\xe2\xbe\xe5\xbd\xa2\x2d\x6e\xaa\xf4\xdc\x29\xf2\xda\xf7\x36\xdd\x27\x7a\x9e\xdb\x44\x0f\xb8\x4b\x14\x39\xca\x35\x2d\x7b\xfc\x0e\x51\xd4\xea\xdb\x61\xa6\xbd\x79\xb5\x7c\xdd\x43\xc3\x8d\xb7\x86\xb6\xbc\x33\xd4\x7b\xc3\x2b\xdc\x42\xb2\xc5\x2d\x2f\x93\xa7\x09\xca\x58\x3a\xad\xf1\x05\x64\x37\x61\x50\xfb\x06\x52\x10\xc4\x4c\xc0\x23\xee\xe3\x98\xe2\xf6\xe4\x25\x0c\x02\x14\x2d\x99\x79\x0c\xbc\x21\xcb\xb0\x06\x43\x8a\x53\xf8\x95\xa5\x02\x49\xc4\x2f\x01\xdf\x92\x77\x1f\xb9\x89\x63\x7e\x00\xba\xe3\x2c\x23\x35\xe6\xa7\xf9\x06\x8c\x7a\x95\xfb\x3d\x98\xdb\x2e\xc0\xf1\xd9\xc7\xd9\xc5\xc9\x34\x89\xf4\x76\x80\x26\xbd\x09\xbe\x0b\x3d\xf5\x9f\x85\xa2\xb7\xa7\xeb\x56\x58\x58\x49\x49\xe6\x2c\x1d\x53\x04\x39\x6a\x5b\xed\x22\x71\x84\x0d\x3a\xe3\x14\xc1\x65\x2f\xec\xa4\xe6\x67\x64\x7e\x72\x8b\x30\x67\xc1\x46\x98\x4d\x2a\x1e\x69\x71\x6b\x94\x4b\x4e\xd2\xff\xa9\xa4\xc0\x6b\x4b\x2b\x40\x22\xb0\xec\xc9\x2f\x48\xa4\xfb\xf0\x2b\x6b\xee\x9f\x7d\xe7\xdf\x39\x53\x7f\x7e\x43\xe9\xfc\x76\x2c\x0f\x1c\x24\x75\xda\x62\x1c\xc0\x83\x24\x8d\x33\xce\x4d\x1b\xa2\xde\xc2\x88\x04\xfc\x38\x40\x3b\xeb\xf6\x16\x64\xcc\x61\x47\xaf\x4b\xba\xb5\xc6\xd6\xf5\x6f\xae\x29\xfa\x17\x0d\x17\xfa\x81\xe1\xb9\x7a\xae\x11\x36\x33\x05\x7b\x16\x8c\xdb\x83\xc1\xab\x80\x46\x91\xe4\x4f\x07\x56\x5d\xcb\xbb\xe1\x99\xfc\xad\xa8\xde\x16\x65\x40\x9e\xc9\x17\xec\x97\x87\x06\x35\x43\x80\x71\x5a\x64\x7c\xf0\xa3\xeb\x3d\x6f\x21\x05\xf0\x2b\x03\xaf\x00\x45\xbf\xd4\x05\x45\x3b\x03\xf8\x95\xed\xb1\xfc\x66\xb0\x1b\x04\x46\x99\x00\xc6\xe8\xab\x18\x36\x3a\x19\xcf\x76\x56\x4b\x78\x37\x45\x9c\x16\x88\xa5\x87\x07\xeb\xf0\x30\xa1\xbe\xc6\xb8\x71\x49\xea\xfc\x67\xc8\xb3\xc5\x19\x99\xb3\x9d\xf0\x18\x6d\xb8\xc1\x2b\x30\x08\xd8\x67\x6f\x2d\x11\x73\xa2\x67\x97\xda\x2c\x50\x59\x26\xc3\xec\x3f\x06\xc9\xe0\xc7\x60\x43\x6c\x0f\x62\x7d\xe3\xd5\xc3\x6b\xdc\x00\x89\xa2\x95\x08\xb8\xd5\x82\x22\x58\xb4\x72\x97\xe5\x1f\x89\xf9\xd7\x83\x36\x90\x2a\x5e\x0d\x8c\x46\xf3\x41\xea\xd0\xdb\xd5\x04\x7b\xba\x63\xc4\x52\x86\x61\x0e\x05\xab\x44\x6a\xda\x41\x3a\x18\xb8\xd2\xf5\x3a\xbf\xd0\x5d\x25\x52\xd1\x66\xc3\x81\x57\xe0\x5a\x6f\xec\x1d\x24\xac\xdd\x10\x64\x04\x73\x74\xc7\x77\x3d\xfe\xc8\x59\xa4\xba\xa8\xcb\x13\xe0\x15\x90\x43\x46\xfa\xf7\x88\xa2\xaa\x84\x19\xda\xd9\xff\x97\xff\xb7\xf3\xe5\x4b\xfe\xfd\xee\x8f\xfb\xf3\x61\x87\x7f\x29\xb4\x70\x08\x72\x94\x45\x70\x03\x40\x11\xaf\x29\x06\xaa\xd5\x61\x74\x4d\xc9\x72\xbc\x80\x54\xec\xcc\x1d\x31\xcc\x53\x5e\xf1\x9f\xc0\x3e\x68\x08\x55\x5d\x28\x01\x51\x27\xcd\x5f\x18\x87\x94\xa3\xfc\xf5\x7d\x0a\x06\x22\xf3\x19\x0c\x63\x90\xb6\xfe\xa4\xae\x3e\x7d\x56\xac\xd0\xfd\x36\x97\x51\x34\x7a\xab\xa5\xcd\x5f\xe2\x80\xc2\xb9\xa6\xe0\x30\x0a\x40\x6e\x11\xa5\x45\x8e\x58\x1a\x5f\x9e\x42\xa4\xfb\xd2\x3e\x74\x03\x3e\xf7\x0d\x00\x52\xbd\x31\x5c\xa2\xd4\x5a\xd4\xb0\x11\x7c\xfa\x19\x0c\xd8\x62\x30\x04\x83\xbd\x6c\xd0\x3e\x15\xca\xda\x87\xf6\x32\xf6\x32\x38\x6a\x1d\x15\x2a\xbb\x41\x5f\xc1\x2b\x70\x0e\xf9\x62\x74\x5d\x12\x42\x77\xe4\x5f\xa9\xec\x76\xd9\xd9\xfd\xee\xf0\xe0\xe0\xe0\x20\xac\x13\x25\x99\xef\x58\x4b\x02\xdf\x83\x41\x3a\x00\xdf\xb7\xe6\xe5\x7b\x30\xd8\x17\x7a\x20\x67\x79\x25\xde\xc8\xe9\xbe\x07\x83\x25\x6b\x16\x2a\x1f\x5b\x9a\x6f\x28\x39\xa2\x34\xaa\xdd\x52\x16\x8c\x94\x68\x24\x08\x19\x20\x4a\x8f\x06\x43\x20\x46\x04\xa9\x15\x7f\x18\xe2\xda\x5d\xed\xb4\x53\xec\x82\x95\x70\x0e\x23\xaa\x12\x9c\x1d\xa5\xe5\xed\xc6\x1d\xe5\x04\xa3\x5d\x61\x44\x04\xe9\x5b\xef\x19\x9f\xe1\xcd\x84\x92\x6d\x4b\xc4\x18\x9c\xa3\x21\xc8\xae\x7a\x2c\x03\x93\x91\x95\x30\xd2\x82\x89\x7b\x82\x51\x3b\xc2\x13\xbd\x81\x1c\xed\xec\xee\x8e\xe6\x6a\x39\x01\x37\x04\xb6\xde\xb2\x8d\x8b\x11\xf6\x33\x6d\x7f\x45\xb7\x49\xd9\xc4\x7b\x0a\x9e\x85\x62\x3f\xc5\x93\x88\xc6\xb0\x51\x66\x07\x8e\x2d\xc3\x1f\x29\xf4\x4d\x32\xdf\x8e\x0d\x9a\x3a\x15\xa2\x6e\xb5\xa3\xb5\x08\x53\xd0\xca\x52\x84\x49\x8c\xc3\x65\x95\x46\xc4\xb4\x61\x47\x47\x99\x0e\x1e\x2e\x27\xf0\x10\x59\x81\x98\xbc\x80\xc3\x6f\xbd\xd4\x38\xb3\xa5\x80\x2b\x23\xda\xef\xb6\xd3\x55\x6c\xef\xf8\x4e\xd6\x2c\xd1\x58\x41\x7b\xa0\x67\xc0\xe9\xe9\xd1\x21\xc7\x86\xba\x9a\x08\x29\x66\x0b\x42\xb9\x8e\x1a\xa6\x75\x4f\x8d\x4d\xeb\x44\x2a\x81\x36\xc6\xe2\x46\x1f\xf4\xe8\x8c\xe0\xb9\x9e\x61\x8f\x65\x0b\x94\xd7\xf6\x47\x80\x66\xfa\xd9\xc9\x5d\x45\x11\x6b\x6a\x3d\x92\x38\xfd\xc6\x69\xc9\x52\xa7\xab\x5e\x09\x5e\x86\xed\xd1\xd8\xbe\x4b\x35\x22\xf7\x8c\x4f\xf3\x00\xc1\xfa\x20\xd7\x39\x0b\xae\xf4\x69\xe8\x97\xe6\xca\xf9\x97\x24\x05\x5f\x9a\x36\x15\xe9\x03\xd6\xeb\x2f\xc9\x10\x7c\x49\xb4\x31\xef\x00\xf4\x25\x51\x09\x60\xc8\x38\x54\x55\x0d\x88\x48\x25\x4e\x13\x44\x97\x05\x63\xa1\x0c\x0b\xb8\x29\x96\x01\x1b\x92\x1a\xb0\xab\xa3\x59\xdb\xb1\xa9\x52\xed\xf4\x14\xdf\x92\x1b\x14\xfa\xac\x4d\xf3\x4c\x35\x33\x3d\x92\xef\x46\xe1\x53\x4c\x2a\x1d\x20\x73\x4a\x9d\xa6\xaa\xc8\x4c\x58\xa2\x89\x36\x50\x7a\x1a\x6d\x4c\xdc\xb7\x71\xc2\x95\x6b\xef\xc3\x28\x6e\xdd\xfa\x58\x00\x84\x76\x97\x7c\xf1\x47\xd9\xfa\x9f\xbc\x6c\x2d\xa5\xf8\xcf\x5f\xb7\x8e\x94\xa3\xff\xa8\x58\xff\x51\xb1\x6e\x48\xfa\xa3\x62\xfd\x47\xc5\xfa\x7f\x89\xe5\xbf\xaf\x8a\xb5\x34\xf1\xcf\x57\xb2\xee\xfc\xfe\x37\xaf\x59\x77\x53\xfd\xbe\x8a\xd6\xab\x15\x78\x29\xbf\x48\xae\x6f\x07\x8c\x2c\x16\x0b\xf4\xa1\x56\x20\x75\xd3\x48\x8e\x4b\x16\x08\xe6\xde\x47\x37\xec\xc4\x26\x16\x78\x19\xf0\x21\xe5\x4b\x40\xb4\x78\xaa\xc3\x97\xe7\x2c\xa2\xba\xeb\xba\x22\xf9\xbd\xb3\xae\x27\xe4\x76\xd1\xd8\xf3\x37\xc8\xec\xb2\x85\xa0\xcb\xd4\x2e\x89\x77\x02\x39\x47\xd4\x0d\x1e\x92\xd6\x40\x38\xa7\x1d\x22\x4e\x44\x99\x65\x6b\x2c\x1b\x93\xe4\x88\xc3\xa2\xdc\xe3\x6a\x01\xce\xe8\x93\xf1\x0c\x74\x37\x66\x9b\x36\x29\x20\x43\x2c\x30\x76\x09\x6c\x87\x08\x29\xdb\x40\x1b\xa6\xf7\xaf\x4f\x6b\x9f\xaf\x32\x50\xef\xc2\xf0\x56\xc1\xce\xa6\x20\x67\xfb\xe0\x26\x18\xd4\x3c\x28\x98\xe9\x0d\x62\x42\x97\xce\x5d\x63\xec\x99\xec\xcb\xa0\x41\x5e\x3f\x25\xab\xb7\xed\xf6\x63\xd2\xfa\xff\xdb\xd9\xf7\x16\xfc\xf9\xdd\xa7\xdf\xa1\x2f\x82\xaa\xcf\x06\x4e\xf5\x9b\xc0\xbf\xad\x11\xf8\x0c\xde\xd4\x00\xd3\xc2\x0d\x7e\xfe\x2e\x6c\xf6\x9e\xfe\x25\x3b\xe3\x9f\x00\xf1\xae\xe9\x7b\xf7\xc2\x5e\x34\x4c\x8a\x7e\xe9\xd0\xff\x88\xa8\xcd\x14\xd5\xee\xac\xd7\xe9\x7d\xb1\xd1\x33\x00\x2f\xb4\x34\x7a\x99\xda\x73\x6f\x35\x30\x6c\xe8\x2d\x59\x2b\x40\x74\x4d\xce\x77\x5c\xad\xef\x08\x84\x3f\x0d\xe9\xd6\x4f\x22\xc2\xdb\xaa\x76\x12\xcd\xc3\x1d\x43\x1f\x0d\xb1\xdd\xca\x85\x6d\xfd\x5c\xaf\xd3\x1b\xcd\xdb\x65\x15\xaf\x60\xd0\x15\x59\x7c\xbb\x8d\x32\x77\xfb\x3e\x26\x4c\x36\xb6\x7c\xb8\x8a\xd1\x6c\xe4\xce\xa6\xc5\xca\x2e\xe1\xa2\x8b\x65\xe0\xed\x82\x8b\xf5\xf9\x08\x1f\x2c\xfe\x1d\xf8\xe7\xff\xc4\x7b\x6f\x3a\x95\x20\x75\xef\xae\x24\x30\xbf\x6a\xef\xdd\xa9\x2b\xa1\x57\x28\x52\x4c\x89\x8c\x51\xdb\x19\xd1\x26\x62\x61\x6f\x29\x59\x06\x6f\xf0\x6d\xc6\x36\x75\x71\xfd\x5c\xf0\xc5\x16\xb8\xb2\xa3\x8d\xc4\x67\x47\xe9\x71\xcd\x17\x84\x16\xbf\xa2\xe0\xed\x54\x6f\x54\xa8\x7b\xdd\xc8\x0e\x83\x7c\xfd\x2e\x80\xc6\x79\xb2\x55\x68\xf1\xc2\x7c\xdb\x6f\x79\x8c\x2f\x44\xfb\x1f\x5e\xb6\x6d\xce\xec\x87\x34\xd5\x9f\x67\xd7\x46\xe7\x0d\x2a\x91\xd0\x93\xb6\x71\x3d\x99\x8a\x20\x11\x6f\x30\x4a\xf2\xdf\x0a\x13\xc1\x2a\x55\x17\x69\xdc\xfb\x8a\xc9\x05\x9c\x7b\xff\x32\xa1\xfa\xa0\x68\xc2\xe4\x87\xec\x84\x8d\x6d\xaf\x0f\xe8\x2f\xc2\x5b\x97\xf8\x5a\x78\x58\x55\x26\x70\x8f\xeb\x09\xb1\xcd\xe0\xda\xff\x04\x00\x00\xff\xff\x9c\x70\xef\x6b\xd7\x72\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
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

