// Code generated by go-bindata.
// sources:
// db/migrations/001_initial.sql
// db/migrations/002_forum.sql
// db/migrations/bindata.go
// DO NOT EDIT!

package migrations

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

var _dbMigrations001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x5b\x53\xa3\x30\x14\x7e\xe7\x57\x9c\xb7\xd6\x59\x9d\x59\xdd\xd9\x7d\xf1\x89\x52\x56\xab\x48\x15\xda\xdd\xf1\x29\x13\x21\xda\x68\xb8\x2c\x17\x2f\xff\x7e\x13\xa0\x40\x68\xa0\x58\xf3\x90\x69\x93\xef\x5c\xfb\x7d\xa7\x39\x39\x81\x6f\x01\x7d\x4a\x70\x46\x60\x1d\x6b\xfc\xab\x7b\x67\x01\x0d\x21\x25\x5e\x46\xa3\x10\x26\xeb\x78\x02\x34\x05\xf2\x4e\xbc\x3c\x23\x3e\xbc\x6d\x48\x08\xd9\x86\x1f\x95\x76\x02\xc4\xbf\xe0\x38\x66\x94\xf8\x9a\xe1\x98\xfa\xca\x84\x95\x3e\xb3\x4c\xf0\x58\xfe\xa0\x4d\x35\xe0\x8b\xfa\x30\x5b\x5c\xb8\xa6\xb3\xd0\x2d\xb8\x75\x16\x37\xba\x73\x0f\xd7\xe6\x3d\xd8\xcb\x15\xd8\x6b\xcb\x3a\x2e\x60\x21\x0e\x08\xfc\xd1\x1d\xe3\x52\x77\x60\x7a\xf6\xf3\xd7\x51\x07\xc0\x08\xf6\x49\x22\x43\xca\x9b\x94\x66\xe4\x54\x75\x11\x31\x1f\xed\xfa\x2d\xef\xbc\x28\x08\x48\x98\xc9\x57\xda\xd1\xb9\x26\x17\xe2\xe3\xd0\xe3\x51\x47\x97\x52\xf9\xf6\xeb\x98\xdd\x52\xc6\x15\x9b\xe6\xc9\x5e\x4c\x8c\xb3\x24\x0a\x3f\x02\xea\x29\x9b\x42\xde\x07\x53\x88\x31\x4d\x90\xc7\x70\x9a\xf6\xc1\x04\xea\x39\x7c\x1e\x06\x95\xbe\x12\xf2\x8a\x3a\x29\xb7\x53\x89\xf2\xc4\xeb\x5e\x14\x37\xc6\xd2\x76\x57\x8e\xbe\xb0\x57\x55\xa3\x11\x12\xbd\x43\x79\x48\xff\xe5\x9c\x97\xf6\xe2\x6e\x6d\xc2\x54\x9c\x55\xbe\x14\x16\xbc\x52\xf4\x8a\x59\x4e\x52\xe4\x6d\x88\xf7\x02\xc6\xa5\x69\x5c\xc3\x54\x74\x80\xf3\x79\x3a\x09\x26\xc7\x30\x79\x9c\x1c\xf5\xbb\x68\x9a\xa1\xf4\xd4\xea\x55\xe1\x50\x17\x0e\x67\x62\x33\xc4\x36\x17\x9b\x39\xe4\xbf\x6e\xa3\xd2\x7d\xd3\xe4\xc2\xfb\xec\x42\xf8\x73\x5c\xb1\xdf\x88\xad\xf8\x64\x6c\x78\x80\x3e\x7e\x22\xa1\xb7\x8a\xa4\xd5\x49\xc9\x55\x91\xc4\xf6\xc7\x02\xc7\xfc\x6d\x3a\xa6\x6d\x98\x6e\x05\xaa\x94\xc0\x6d\xf7\xc0\x0b\x39\xef\xc4\xe6\x12\x8a\x49\x46\x8b\x41\x50\xc6\x1e\x2f\x77\xbe\xb8\x25\xeb\x95\x49\x85\xf1\xc5\x78\x9a\x8b\x90\x3b\x57\x42\xf5\x92\xf5\x6e\x82\x61\x14\xd0\x10\x1f\x9a\x5f\xab\xbc\x7d\xed\x69\x90\x15\xb1\xf9\x2a\x7e\xe9\x41\xd9\xf0\x15\x60\x46\x38\xe5\x73\x3e\x88\xda\xce\xeb\x1c\x1e\xc9\x3e\x44\xf6\x11\x37\x51\x4e\xbf\x2b\x7a\xc8\x7b\xd0\xd1\xb0\x1a\x86\xdf\x07\x61\x5b\x5c\x8b\xde\x4d\x7f\x11\x12\x79\xc8\xb4\x2e\x32\x2b\x18\xbd\xb4\xe6\xe8\xca\xbe\x12\x34\xb6\xcd\xbf\xdb\x8f\x86\xa5\xbb\xee\xc2\xa8\x75\xd3\xef\xbb\xae\x40\x0e\xd0\x14\x36\xa0\x4a\x71\xd6\x2f\xa8\xbd\x91\xb7\x4d\xe9\x44\xae\x7b\x75\x70\xe4\x82\xad\x32\x5d\x13\x92\xe6\x2c\x6b\xa8\xfa\x19\xb2\x1e\x4c\xd7\xda\xfe\xd3\x73\xa3\x58\xad\x66\x0d\x5b\x36\xc0\x96\x75\x55\xf1\x96\x70\x3f\xce\x54\x84\xe3\xff\x2f\x0c\xf3\xff\x0f\x25\xfd\xab\x4b\xf4\x98\x44\x41\x1f\x82\xa6\x62\x00\xc3\x6c\xb9\xb4\xd4\xee\x23\x1a\x66\x69\x9f\xf5\x18\x4d\x00\x60\xc6\x50\x91\x49\x3a\x98\x4a\x0b\x96\x45\x1d\x90\x12\x35\x52\xbc\xb2\xcd\x48\x25\x4b\xac\x2f\x7f\x0a\xa4\x62\xfb\x97\x35\x36\x18\x50\x55\xac\x1c\x5f\xd9\x8e\xaf\x4b\x7e\x6f\x36\x6a\xed\x2b\x1b\x7d\xf8\x18\x28\xe7\x40\xfb\x35\x3e\x8f\xde\xc2\xed\x7b\xbc\x7e\x8c\x8b\xc3\x51\xcf\xf1\x24\x62\x8c\xdf\x3e\x60\xef\x45\x9b\x3b\xcb\x5b\x69\xb6\x9c\xb7\x8f\x5a\x4f\x07\xe9\xbc\x51\xaa\x74\xbc\x83\x2b\xed\x65\x4c\x33\x59\xa4\xf3\xa7\x28\x89\x51\x9d\x66\x7a\xae\xfd\x0f\x00\x00\xff\xff\x42\x75\x7e\xa6\x7e\x0c\x00\x00")

func dbMigrations001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations001_initialSql,
		"db/migrations/001_initial.sql",
	)
}

func dbMigrations001_initialSql() (*asset, error) {
	bytes, err := dbMigrations001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/001_initial.sql", size: 3198, mode: os.FileMode(420), modTime: time.Unix(1477830601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations002_forumSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x4f\x6f\x9b\x30\x14\xbf\xfb\x53\xbc\x5b\x13\x2d\x3d\x24\xd2\xa6\x49\x39\x51\xea\x75\x68\x8c\x74\x86\x4c\xea\x09\x11\x70\x53\xaf\xc6\x58\x60\xd4\x7e\xfc\x99\x98\x52\x83\x48\x29\xe1\x84\xdf\xfb\xfd\xde\xff\xf7\xae\xaf\xe1\x4b\xce\x8e\x65\xa2\x28\xec\x25\xd2\xcf\xf0\x8f\x0f\x4c\x40\x45\x53\xc5\x0a\x01\x57\x7b\x79\x05\xac\x02\xfa\x4a\xd3\x5a\xd1\x0c\x5e\x9e\xa8\x00\xf5\xa4\x45\x86\xd7\x80\xf4\x23\x91\x92\x33\x9a\x21\x97\x60\x27\xc2\x10\x39\x37\x3e\x86\xc7\x58\x26\xa5\x62\x0d\x06\x2d\x10\xe8\x8f\x65\x70\xe3\xdd\x85\x98\x78\x8e\x0f\xf7\xc4\xfb\xed\x90\x07\xf8\x85\x1f\x20\xd8\x45\x10\xec\x7d\x7f\x65\x60\x22\xa3\xaf\xe0\x05\x11\xbe\xc3\x64\xa0\x4b\x8b\x5c\x52\x63\x34\x36\xe6\x34\xae\xc3\x00\xc1\x3f\x30\xc1\x81\x8b\x43\x1b\x89\x96\x5b\x34\x8c\xed\x5f\x9d\x1d\x29\xcc\x8a\x2b\x4b\x44\x4a\xcb\x09\xb7\x06\x64\x08\x5d\xfe\x13\x1c\xab\x52\x86\xc8\xa9\x52\xb4\x84\xbf\x0e\x71\x7f\x3a\x64\xb1\x5e\x76\xb4\xb1\x54\x44\x91\x33\x61\x7a\x31\x2b\x9f\x8b\xc3\x2b\x2d\x97\x13\xdc\x77\xa0\xa1\x6a\x33\x9c\x76\x89\x6d\xbe\x7e\xfb\x38\x35\xc9\x93\x74\x66\x97\x4e\x94\xf8\xb1\x2c\xf2\x33\x23\x64\x00\xaa\x38\xa3\x16\x75\x7e\xd0\xb5\x1f\x57\x56\x2a\x39\x6a\x6e\x2f\x89\xf5\xe6\xfb\x72\x68\xe3\xb3\xe5\xb1\x9b\x67\x0f\xd9\x7a\xc6\x94\x99\xff\x8d\xc5\x18\x01\x9a\xbe\xd1\xaa\xe6\x6a\x3d\x8e\x34\xca\x95\x05\x3c\x63\xd2\x28\xc7\xba\xd5\x2e\x48\xca\xeb\x43\xdb\xb3\xd9\x2b\x73\xd1\x7e\xb7\x54\xed\x76\x8a\xa3\x21\x6d\x2d\xdc\x5d\x10\x46\xc4\x69\x90\xbd\xc0\x63\xf9\x6c\x0f\xd7\xa2\xcb\x60\x35\x88\x6d\xf5\xe6\x70\x79\x2a\x85\x7d\x4b\x6f\x8b\x17\xf1\x76\x4d\xbb\x53\xda\x08\x3f\x75\x4c\xcb\x82\x73\xad\x3d\x24\xe9\x33\xba\x25\xbb\xfb\xfe\x32\x6c\xfb\xb2\xd3\x19\x1b\xc8\xde\x47\x6a\xa0\xe8\xd6\x78\x8b\xfe\x07\x00\x00\xff\xff\x8f\x8e\xe1\x66\xfb\x05\x00\x00")

func dbMigrations002_forumSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations002_forumSql,
		"db/migrations/002_forum.sql",
	)
}

func dbMigrations002_forumSql() (*asset, error) {
	bytes, err := dbMigrations002_forumSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/002_forum.sql", size: 1531, mode: os.FileMode(420), modTime: time.Unix(1477845776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrationsBindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x96\x5b\x6f\xdb\xcc\x11\x86\xaf\xc5\x5f\xc1\x18\x48\x21\x01\xae\xcc\xf3\x21\x40\x6e\x72\x28\x90\x8b\x24\x40\x93\xbb\xba\x30\x96\xe4\xae\x4a\x54\x12\x1d\x51\x4a\xed\x04\xf9\xef\x9d\x67\x67\xe5\x38\x6d\xf2\xe5\xe6\x33\x40\x8b\x7b\x98\xd3\x3b\x33\xef\xf0\xea\x2a\x7e\x39\x0d\x36\xde\xd8\xbd\x3d\x98\xa3\x1d\xe2\xee\x3e\xde\x4c\x7f\xed\xc6\xfd\x60\x8e\x66\x1d\xc9\x85\x79\x3a\x1d\x7a\x3b\x3f\xe3\x7d\xe8\xae\x76\xe3\x46\x6e\x8e\xd3\x7e\xbe\x4a\x92\xf4\x66\xdc\x8f\xc7\xd1\x6c\xd7\xf3\xa7\xed\xcf\x2e\x64\x37\x6e\x3a\x9c\x76\x3f\x3f\x3e\x5b\xd9\x4c\x9c\xbd\x7a\x1f\xbf\x7b\xff\x31\x7e\xfd\xea\xcd\xc7\x27\x51\x74\x6b\xfa\x7f\x9b\x8d\x8d\xbf\x5f\x8f\xa2\x71\x77\x3b\x1d\x8e\xf1\x32\x5a\x5c\x74\xf7\x47\x3b\x5f\xc8\x4b\x3f\xed\x6e\x0f\x76\x9e\xaf\x36\x5f\xc6\x5b\x36\xdc\xee\xc8\xcf\x38\xe9\xff\xab\x71\x3a\x1d\xc7\x2d\x8b\xc9\x0b\xdc\x9a\xe3\xbf\xae\xdc\xb8\xb5\xbc\xb0\x31\x1f\x0f\xe3\x7e\xe3\xcf\x8e\xe3\xce\x5e\x44\xab\x28\x72\xa7\x7d\x1f\x07\xf7\xfe\x6e\xcd\xb0\xe4\x25\xfe\xc7\x3f\x31\x7b\x19\xef\xcd\xce\xc6\x2a\xb6\x8a\x97\xe7\x5d\x7b\x38\x4c\x87\x55\xfc\x35\x5a\x6c\xbe\xf8\x55\xfc\xec\x79\x8c\x57\xeb\x77\xf6\x3f\x28\xb1\x87\xa5\x77\x9b\xf5\x8b\x93\x73\xb2\x46\xed\x6a\x15\x2d\x46\xe7\x05\x9e\x3c\x8f\xf7\xe3\x16\x15\x8b\x83\x3d\x9e\x0e\x7b\x96\x97\xb1\x84\xb4\x7e\x8d\x76\xb7\xbc\x40\x51\xfc\xf4\xd3\xb3\xf8\xe9\xe7\x0b\xf5\xc4\xdb\x12\x1d\xdf\xa2\x68\xf1\xd9\x1c\xe2\xee\xe4\x62\xb5\xa3\x46\xa2\xc5\x8d\xba\xf3\x3c\x1e\xa7\xf5\xcb\xe9\xf6\x7e\xf9\x17\xb9\x73\x29\xbe\x89\x54\xbf\x7d\x7d\xf6\x74\xfd\x72\x3b\xcd\x76\x29\xe1\xff\x49\xfe\xa0\x46\xf5\xff\x42\x91\x5c\x54\xbf\xc3\xa6\xb8\xb5\x7e\x81\xeb\xcb\xd5\x25\x37\x22\x39\x3b\xde\xdf\xda\xd8\xcc\xb3\x3d\x02\xf9\xa9\x3f\xa2\xc5\xc7\x17\xf2\x21\x66\xf6\x6e\x8a\xe3\x69\x5e\xff\x4d\xd2\xfa\x46\x16\x0f\x72\x21\x85\xe7\xfd\x47\x1a\x7c\x0e\xe5\x4f\xd3\x18\x2d\xe6\xf1\x8b\x5f\x8f\xfb\x63\x55\x44\x8b\x1d\x5d\x11\x3f\x28\x7d\x2b\x4b\xbf\xf9\x51\x2a\x24\xa6\x4c\xd6\xbc\x61\xc7\x97\xca\xd2\x8d\xff\x6b\x6b\x15\xbf\x13\x13\xcb\x55\xb0\x80\xcd\x10\xa5\x1b\xd7\x58\x17\xe1\x5f\xcb\x7e\x10\x77\x44\xd6\x7b\xf3\xa3\x28\x8e\xfe\xa1\x28\xbe\x8a\xe8\x23\xcf\x7f\x54\x40\x68\xbf\x53\x40\x70\xa2\xe3\x21\xd0\xff\xd3\x10\xa2\xff\xb5\x92\x37\xf3\xab\xf1\x20\x2a\xba\x69\xda\x3e\x96\x36\xdb\xf9\x37\x91\xdf\xcf\x1a\xb8\x3d\x38\xd3\xdb\xaf\xdf\x1e\x49\x87\x92\xa0\xca\x6f\x86\xee\xed\x03\x35\x3c\x22\xa2\x0f\x9f\xb6\x52\xe8\x5a\x19\xcb\x8b\xeb\xbb\xd4\x5d\xdf\x35\xdd\xf5\x5d\xd2\xc8\x93\x84\xa7\xbd\xbe\xab\xac\xec\x87\x3d\x27\x77\x4c\x71\x7d\x57\x56\xf2\xc8\xdd\x32\x97\xb5\x3c\xb9\x9c\xa5\xb2\x5f\xcb\x5d\x5b\xcb\xbe\x3c\x6d\x7f\x7d\xd7\xc9\xef\xc0\x5d\xd1\xd3\x0e\xfa\x3b\xc8\xef\x20\xbf\xb5\xfc\xba\x54\x74\xcb\x7b\x99\xa9\x4e\x23\x3a\x0b\xb1\x95\x96\x72\xc7\xe8\x5d\xee\x64\x72\x27\x15\x3b\x59\xaa\xfb\x95\xdc\xe9\xe4\xc9\xc4\x46\x2a\x36\x32\xa7\xbe\x61\x9f\x7b\x46\xfc\x29\x12\xbd\xc7\x7b\x29\xbf\x4e\xf6\x5b\xf6\xf0\x45\xde\xad\xdc\x2f\x45\xde\x75\xea\x8b\x11\x3d\x79\xab\x4f\x23\x76\x2a\x39\x4f\x52\x3d\x2b\xc4\x66\x8d\x4e\xf1\xb1\x22\x56\xd9\xab\x44\xce\xf5\xea\x73\x2d\x4f\x55\xeb\xfd\x64\x50\x3f\x33\x89\xa1\xb4\x2a\x03\x46\xa9\xc8\x65\xf2\x6e\xb9\x0f\x9e\x12\x73\x2e\x98\x25\x72\xcf\xc9\x7b\x21\x77\x3b\xd1\x97\xcb\x93\x81\x69\x58\x57\x83\x62\x42\x5e\xc0\xad\xa9\x34\x57\xad\xc8\xd5\x95\xea\xe9\x45\x4f\x27\x7b\x56\x6c\xe4\x72\xaf\x92\xfd\x56\xf6\x9c\xbc\xb7\xe2\xbb\x15\x7f\x5a\x30\x90\xf7\x5e\x9e\xa6\x50\x79\x6f\x43\x6c\x15\xe0\x70\xc6\xc9\x06\xfc\xc4\x6e\x5e\xaa\x4e\xea\x02\x59\xf2\xec\xf3\xde\x2b\xfe\xa6\x52\xf9\x81\xd8\x06\xdd\xab\x45\x26\x95\xfd\x5c\xf0\xe8\xc0\x4d\xce\x13\xf1\x6d\x90\xb5\x65\x9f\xfc\xcb\xbd\xbe\x0b\x79\x6e\x14\x4b\xd6\xb9\xdc\xcf\xac\x62\x0c\x86\x89\xd5\xb8\xc1\x79\x48\x15\x77\x9b\x6b\xbd\xd4\xb5\xde\xab\x45\xc6\x55\x9a\x5f\x6c\x94\xc8\xc9\x6f\x9f\x04\x59\x39\x2b\x04\xb7\x0c\x8c\x73\x8d\xbf\xed\x14\x1f\x70\xb2\xd4\x73\x11\xfc\x16\xd9\x3c\x55\x7c\xed\xa0\x31\x13\x17\xb5\x42\x2e\xb2\x90\x07\xf2\x01\x9e\x7d\xab\xb5\xee\x6b\x55\xe4\xba\x56\x73\x4c\x4d\xda\x4c\x7d\x05\x1b\xef\xb3\x9c\x17\xf4\x45\xad\xfe\xbb\x56\x7d\xa3\x1e\xd0\x45\xad\xb3\xdf\x57\x8a\x0b\x3e\x82\x97\xb7\x61\x35\x47\x4d\xaf\x78\x67\xd4\x0d\xf9\x70\x9a\x7f\x6b\xb4\x47\x90\x29\x64\x3d\x58\xc5\x80\xfe\xa4\x77\x89\x89\x7c\x52\xff\x7d\xad\xb5\xdc\x06\xdc\x7a\xf0\x2a\x54\x07\xb5\x5f\xd3\x53\xc8\x57\x5a\x23\x1d\x3d\x95\x69\x8d\x36\x26\xe4\xa8\xd5\xde\xf6\xfa\x5b\x3d\xa3\x06\x89\x05\x5f\xe1\x02\xf2\x4f\x2c\x43\xa6\x75\x0a\x4e\x8d\x55\x7d\x5d\xa2\xeb\xb2\x0c\xbd\x91\xaa\x8d\x6e\xd0\x1c\xd5\xad\x62\x4c\x4f\xdb\x52\xb9\x04\xdc\x3c\x5e\x99\x72\x12\x7d\xd1\x67\x9a\x3b\xce\xcb\xb3\xaf\x85\xd6\x1f\xf9\xf2\x38\x17\xea\x37\xdc\x43\x8f\xb1\x47\xce\x91\x07\x1b\x57\x87\x1e\xa2\x4f\x4b\xe5\x07\xea\xbb\xcf\x43\x5d\x20\x9f\x68\x7d\xe1\x17\xf1\x27\x21\xcf\x60\x49\x4f\xb0\x8f\x0f\x29\x7d\x84\x1f\x9d\x72\x0f\x18\x9b\x54\x7d\x18\x02\xc6\xf8\x49\x3f\x96\x89\xd6\x09\x35\x0e\x77\x54\xf0\x41\xae\x38\xe6\xa1\x86\xc8\x27\x39\xb1\xa1\xbf\x89\x1b\xbc\x3c\xa6\xf4\x51\xae\x67\xe0\x00\x8f\xd9\x5e\xf3\x4f\xdf\xd3\x77\xbe\x4f\xe1\x4c\xa7\xb1\x9a\x41\xcf\x2a\xe2\xef\x95\x87\x88\x0d\x1e\x80\x3b\xe9\x11\xfc\x25\xb7\x60\x06\xd7\x0f\xa5\xf6\x8f\xaf\x43\x72\x91\xaa\x8f\x2e\xc4\xef\xb1\xab\xb5\x47\xc1\x14\x7c\xe0\x05\xde\xa9\x6b\x7c\x86\xe3\x8c\xd1\x9e\xa3\x2f\x39\x73\xe7\x1e\x3d\xc7\xe8\x42\x0d\x55\x5a\x0b\xd8\xc7\x2e\xfc\x09\xff\x14\x01\x33\xdf\x1f\x56\x39\x1f\x0e\xf0\xf5\xd4\x28\xdf\xfa\x5e\x0a\xfe\x36\xa5\xce\x8e\xbe\xd4\x38\xc9\x05\x7a\xe8\x81\xbc\xd3\xfa\xf3\x58\x11\x53\xf8\xf5\xb3\x8e\xfa\x4a\xb5\x8e\xe8\x5d\x78\xdb\xf3\x2d\xb9\x73\xca\xf7\xd4\x16\xf8\xc3\x11\x7e\x9e\x24\xda\xb3\x60\x4d\x5f\x67\x26\xcc\xc2\x56\xfd\x86\xb7\x5d\xb8\x53\x85\x39\x02\x4f\x53\xdb\x75\xe0\x7a\xf2\x03\x8f\xf4\x56\xe3\x48\x43\x1f\x82\x3b\xe7\x45\xc0\xdc\x73\x54\xe0\x39\x3f\x33\xc4\xa7\xac\x53\x1e\xec\x1b\x9d\x45\xf8\xed\xf9\x08\x9b\x46\x67\x06\x35\x00\x27\x10\xdb\x99\xc3\xc1\x03\xae\x80\x1b\x6a\xf7\xbd\x0f\xa9\x67\x74\x75\x85\xf2\x71\x9e\xa9\x8d\xb4\xd1\x1e\xed\x42\x9e\x6c\xa3\x7c\x7a\xe6\x08\xe6\x59\x47\x5d\x0c\xea\xd7\x20\x32\x8d\x53\xfb\x26\xf4\x85\xb5\x5a\xaf\x26\xcc\x8a\x21\xcc\x62\xce\x8c\xd5\xb9\x4d\x4f\x0c\xe1\x9b\x82\x5e\xf0\xf3\x08\x4c\x53\xfd\x9e\x20\xf7\x26\xf8\xd2\x86\x7a\x25\x06\x8f\x77\xa1\x7a\xc0\x01\xfd\x70\x1d\x3d\x42\x5e\xa9\x7f\x62\x29\x07\xfd\x4e\x20\xef\xc4\xc1\x77\x04\x33\xc8\x84\x19\x99\xc2\x6f\x99\xe2\x0f\xef\x0e\x81\xe3\x99\x8d\xf8\x5b\x87\xef\x9e\x32\xf4\x14\x73\x84\xb9\x0f\xae\xf8\x4b\x5d\xb7\x95\xda\x2f\x43\x1d\xfa\x75\xa2\x9c\x4d\x8e\xe1\x16\xe6\x30\xb3\x01\x9e\x23\xa6\xa4\x57\x9e\x61\xcd\xcc\xa4\xa6\xdd\x10\xbe\x8b\x3a\xe5\x11\x7a\x8a\x18\x8b\x30\xa7\x7c\x5c\x95\xf2\x07\xdf\x68\xf0\x16\xbd\xcd\x2c\xea\xce\xbc\x63\xf5\x9b\x82\x5a\xc0\x57\x5f\xb3\x4e\x6b\x9e\x9c\x82\x1d\xb2\x55\xe8\x2f\xfa\x0c\x5c\x8b\xc0\x77\xde\xa6\xd1\x5e\x25\x8e\xa2\xd4\x59\xec\xeb\x3e\xd3\xef\x04\x78\x95\xbe\x00\xb7\x3e\x7c\xb3\xe0\x3f\x7e\x9b\x3e\xf4\xab\xd3\xb9\x05\xaf\x54\x81\xb3\xe0\x75\xe4\x3d\xe7\x35\x6a\xdb\xcf\x85\xc0\xa7\xf4\x9c\xff\xee\x81\xfb\x9c\xea\xa3\xb7\xf8\x3e\x43\xa6\x32\xda\x87\xbe\x3e\xc3\xb7\x1c\x1c\x87\x2e\xe6\x35\xdf\xa3\xd4\x14\xb6\xf2\xf0\xcd\x43\x4d\x0e\xa1\x0e\x91\xad\xcf\x3d\x6d\x75\xc6\xd2\x17\x7d\x98\xd9\x3d\x38\xa5\xff\x0d\x00\x00\xff\xff\x37\x8e\x47\x5b\x00\x10\x00\x00")

func dbMigrationsBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrationsBindataGo,
		"db/migrations/bindata.go",
	)
}

func dbMigrationsBindataGo() (*asset, error) {
	bytes, err := dbMigrationsBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/bindata.go", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1477845777, 0)}
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
	"db/migrations/001_initial.sql": dbMigrations001_initialSql,
	"db/migrations/002_forum.sql":   dbMigrations002_forumSql,
	"db/migrations/bindata.go":      dbMigrationsBindataGo,
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
	"db": &bintree{nil, map[string]*bintree{
		"migrations": {nil, map[string]*bintree{
			"001_initial.sql": {dbMigrations001_initialSql, map[string]*bintree{}},
			"002_forum.sql":   {dbMigrations002_forumSql, map[string]*bintree{}},
			"bindata.go":      {dbMigrationsBindataGo, map[string]*bintree{}},
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
