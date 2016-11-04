// Code generated by go-bindata.
// sources:
// db/migrations/001_initial.sql
// db/migrations/002_forum.sql
// db/migrations/003_f_comp.sql
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

	info := bindataFileInfo{name: "db/migrations/001_initial.sql", size: 3198, mode: os.FileMode(420), modTime: time.Unix(1478278763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations002_forumSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xcf\x6f\x9b\x30\x14\xbe\xfb\xaf\x78\xb7\x26\x5a\x7a\x48\xa4\x4d\x93\x72\xa2\xd4\xeb\xd0\x18\xe9\x0c\x99\xd4\x13\x22\xe0\xa6\x5e\x8d\xb1\xc0\xa8\xfd\xf3\x67\x62\x42\x0d\x22\x63\x24\xa7\xe0\xef\xfb\xde\xef\xf7\x6e\x6f\xe1\x53\xce\x8e\x65\xa2\x28\xec\x25\xd2\x9f\xe1\x2f\x1f\x98\x80\x8a\xa6\x8a\x15\x02\x6e\xf6\xf2\x06\x58\x05\xf4\x9d\xa6\xb5\xa2\x19\xbc\xbd\x50\x01\xea\x45\x3f\x19\x5d\x43\xd2\x1f\x89\x94\x9c\xd1\x0c\xb9\x04\x3b\x11\x86\xc8\xb9\xf3\x31\x3c\xc7\x32\x29\x15\x6b\x38\x68\x81\x40\xff\x58\x06\x77\xde\x43\x88\x89\xe7\xf8\xf0\x48\xbc\x9f\x0e\x79\x82\x1f\xf8\x09\x82\x5d\x04\xc1\xde\xf7\x57\x86\x26\x32\xfa\x0e\x5e\x10\xe1\x07\x4c\x06\x58\x5a\xe4\x92\x1a\xa3\xb1\x31\xa7\x79\x1d\x07\x08\xfe\x86\x09\x0e\x5c\x1c\xda\x4c\xb4\xdc\xa2\x61\x6c\x7f\xea\xec\x48\x61\x56\x5c\x59\x22\x52\x5a\x4e\xb8\x35\x24\x23\xe8\xf2\x9f\xd0\x58\x95\x32\x42\x4e\x95\xa2\x25\xfc\x76\x88\xfb\xdd\x21\x8b\xf5\xb2\x93\x8d\xa5\x22\x8a\x9c\x09\xd3\x8b\x59\xf9\x5c\x1d\x5e\x69\xb9\xb4\xb4\x96\xe4\x03\x37\x0a\xad\xe6\xb4\xcb\x67\xf3\xf9\xcb\xbf\x33\x92\x3c\x49\x67\x36\xe7\x24\x89\x9f\xcb\x22\xbf\x30\x39\x86\xa0\x8a\x0b\xb0\xa8\xf3\x83\x2e\xf9\x38\x58\xa9\xe4\xa8\xb5\xbd\x24\xd6\x9b\xaf\xcb\xa1\x8d\x89\xaa\xd8\xad\xb2\x47\x6a\x3d\x63\xa6\xcc\xff\xcd\xb8\x83\x33\xd1\x74\x89\x56\x35\x57\xeb\x71\xa6\x01\x57\x16\xf1\x82\x49\x03\x8e\x35\xa9\x5d\x87\x94\xd7\x87\xb6\x55\xb3\x17\xe4\xaa\x6d\x6e\xa5\xda\xed\x94\x46\x53\xda\x5a\xb8\xbb\x20\x8c\x88\xd3\x30\x7b\x81\xc7\xf2\xd5\x9e\xa9\x45\x97\xc1\x6a\x10\xdb\xea\xec\x70\x79\x2a\x85\x7d\x39\xef\x8b\x37\x71\xbe\x9d\xdd\xe1\x6c\x1e\xff\xeb\x74\x96\x05\xe7\x1a\x3d\x24\xe9\x2b\xba\x27\xbb\xc7\xfe\x0e\x6c\xfb\x6f\xa7\xa3\x35\x78\xfb\x18\xa9\x01\xd0\x2d\xed\x16\xfd\x0d\x00\x00\xff\xff\x05\xe0\xe4\x71\xe9\x05\x00\x00")

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

	info := bindataFileInfo{name: "db/migrations/002_forum.sql", size: 1513, mode: os.FileMode(420), modTime: time.Unix(1478278763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations003_f_compSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\x31\x4f\xc3\x30\x10\x85\x77\xff\x8a\xb7\xb5\x15\x74\x41\x82\xa5\x53\x4a\x2c\x14\x61\xda\xe2\x26\x43\x27\x64\x9c\x83\x58\xb8\x89\xe5\xb8\x14\xfe\x3d\x76\x2b\x50\x83\x10\xea\x6d\xf7\xde\x3b\xdd\xd3\x37\x9d\xe2\x62\x6b\x5e\xbd\x0a\x84\xca\xb1\xb8\xae\x1f\x05\x4c\x8b\x9e\x74\x30\x5d\x8b\x51\xe5\x46\x30\x3d\xe8\x83\xf4\x2e\x50\x8d\x7d\x43\x2d\x42\x13\xa5\xe3\x5d\x0a\xc5\x45\x39\x67\x0d\xd5\xec\x56\xf2\xac\xe4\x28\xb3\xb9\xe0\x78\x79\xd2\xdd\xd6\x51\x30\x29\xc5\xc6\x0c\x71\x4c\x8d\x79\x71\xb7\xe6\xb2\xc8\x04\x56\xb2\x78\xc8\xe4\x06\xf7\x7c\x83\xc5\xb2\xc4\xa2\x12\xe2\xf2\x10\xdb\x79\x8b\x77\xe5\x75\xa3\xfc\xf8\xea\xfa\x66\xf2\xcb\xae\x53\xe1\x3c\x7d\x1a\xea\xf1\x95\xa5\x7f\x0f\xa9\xd7\xde\xb8\x43\xed\xd3\xd8\xd1\xd5\x26\x7c\x0e\x64\x36\x99\x31\x76\x0a\x29\xef\xf6\xed\x37\xa6\x1f\x46\x49\x3c\x8b\x92\xef\xac\x8d\xee\xb3\xd2\x6f\x2c\x97\xcb\xd5\x5f\x9c\x66\xec\x2b\x00\x00\xff\xff\xdf\xa1\xf5\x3a\x95\x01\x00\x00")

func dbMigrations003_f_compSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations003_f_compSql,
		"db/migrations/003_f_comp.sql",
	)
}

func dbMigrations003_f_compSql() (*asset, error) {
	bytes, err := dbMigrations003_f_compSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/003_f_comp.sql", size: 405, mode: os.FileMode(420), modTime: time.Unix(1478293564, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrationsBindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x99\x59\x4f\x5c\xd9\x11\xc7\x9f\xe9\x4f\x71\x07\x69\x46\x4d\x44\xe0\xee\x0b\xd2\xbc\xcc\x12\xc9\x0f\xe3\x91\x32\xce\x53\x88\xd0\x5d\xce\x21\xad\x40\x37\xee\x6e\x26\xd8\x96\xbf\x7b\xea\x57\x55\x17\xb0\xc7\x80\x32\x49\xa4\x58\x6a\x77\xdf\x73\xcf\x52\xeb\xbf\xfe\x75\x38\x3d\x4d\xbe\xdf\x4c\x21\xb9\x0c\xeb\xb0\xed\xf7\x61\x4a\x86\x77\xc9\xe5\xe6\x8f\xc3\x6a\x3d\xf5\xfb\xfe\x64\x21\x13\x76\x9b\xdb\xed\x18\x76\x67\xfc\x9e\x86\xd3\xeb\xd5\xa5\xcc\x5c\x6d\xd6\xbb\xd3\x34\xcd\x2e\x56\xeb\xd5\x7e\xd5\x5f\x9d\xec\xde\x5e\x7d\x69\x42\x7e\x11\x37\xdb\xdb\xeb\xa7\x5e\x17\x17\xf1\x62\xdc\x5c\xdf\x7c\xf9\xfd\x2c\xc5\xe5\x86\x77\x3f\xfc\x9c\xbc\xfe\xf9\x4d\xf2\xe3\x0f\xaf\xde\x7c\xb5\x58\xdc\xf4\xe3\x3f\xfa\xcb\x90\x3c\x4c\x5f\x2c\x56\xd7\x37\x9b\xed\x3e\x59\x2e\x0e\x0e\x87\x77\xfb\xb0\x3b\x94\x1f\xec\xbe\x0d\xbb\xdd\xe9\xe5\xfb\xd5\x0d\x03\xf1\x7a\xcf\xd7\x6a\x63\xff\x9f\xae\x36\xb7\xfb\xd5\x15\x0f\x1b\x5d\x70\xd3\xef\xff\x7e\x1a\x57\x57\x81\x1f\x0c\xec\xf6\xdb\xd5\xfa\x52\xdf\xed\x57\xd7\xe1\x70\x71\xb4\x58\xc4\xdb\xf5\x98\xb8\x78\x7f\x0e\xfd\xb4\xe4\x47\xf2\xd7\xbf\x71\xec\x71\xb2\xee\xaf\x43\x62\xcb\x8e\x92\xe5\x3c\x1a\xb6\xdb\xcd\xf6\x28\xf9\xb0\x38\xb8\x7c\xaf\x4f\xc9\xd9\xb7\x09\x52\x9d\xbc\x0e\xff\x64\x93\xb0\x5d\xaa\xd8\x3c\x7f\x77\x1b\xa3\x3c\xb3\xed\xd1\xd1\xe2\x60\x15\x75\xc1\x57\xdf\x26\xeb\xd5\x15\x5b\x1c\x6c\xc3\xfe\x76\xbb\xe6\xf1\x38\x11\x95\x4e\x7e\x64\xf7\xb8\x3c\x64\xa3\xe4\xeb\xb7\x67\xc9\xd7\xbf\x1e\x9a\x24\x7a\x96\xec\xf1\x71\xb1\x38\xf8\xb5\xdf\x26\xc3\x6d\x4c\xec\x1c\x3b\x64\x71\x70\x61\xe2\x7c\x9b\xac\x36\x27\xdf\x6f\x6e\xde\x2d\xbf\x91\x39\xc7\x22\x9b\xac\x1a\xaf\x7e\x9c\x25\x3d\xf9\xfe\x6a\xb3\x0b\x4b\x51\xff\xbf\x24\x0f\xdb\xd8\xfe\x4f\x6c\x24\x13\x4d\x6e\x1f\x14\xb1\x4e\xbe\x43\xf4\xe5\xd1\x31\x33\x16\xf2\x6e\xff\xee\x26\x24\xfd\x6e\x17\xf6\x98\xfc\x76\xdc\xb3\x8b\xea\xe7\xfe\x90\x63\xd6\x71\x93\x24\x9b\xdd\xc9\x9f\xc4\xad\xaf\xe4\xe1\x7e\x9d\xbb\x70\x1e\x7f\xb4\x83\xfa\x50\xfe\x99\x1b\x17\x07\xbb\xd5\x7b\x7d\x5e\xad\xf7\x75\xb9\x38\xb8\x26\x6b\x92\xfb\x4d\x7f\x92\x47\x1d\x7c\x23\x11\x92\x10\x26\x27\xfc\xe2\x1c\x0d\x95\x65\x5c\x7d\x7e\xd6\x51\xf2\x5a\x8e\x58\x1e\xf9\x09\x9c\xe9\x5a\xc6\xd5\x09\xa7\xcb\xe2\xa7\xd7\xfe\x22\xe2\xc8\x5a\x95\xe6\xd3\xa5\x08\xfa\xec\x52\x64\x95\xa5\x8f\x24\xff\x74\x03\x54\x7b\x69\x03\x94\x93\x3d\xee\x15\xfd\xcd\x0e\xae\xfd\xd3\x9b\xbc\xda\xfd\xb0\xda\xca\x16\xc3\x66\x73\xf5\x78\x75\x7f\xb5\x7b\x41\xf3\x77\x3b\x53\x3c\x6c\x63\x3f\x86\x0f\x1f\x1f\xad\xf6\x90\x20\xca\x2f\xa6\xe1\xa7\x7b\x68\x78\x04\x54\xbf\xbc\xbd\x92\x40\xb7\xc8\x58\x1e\x9e\xdf\x65\xf1\xfc\xae\x1d\xce\xef\xd2\x56\x3e\xa9\x7f\xba\xf3\xbb\x3a\xc8\xb8\x8f\x45\x99\xd3\x97\xe7\x77\x55\x2d\x1f\x99\x5b\x15\xf2\x2c\x9f\x42\xde\x65\x32\xde\xc8\xdc\xd0\xc8\xb8\x7c\xba\xf1\xfc\x6e\x90\xef\x89\xb9\xb2\x4f\x37\xd9\xf7\x24\xdf\x93\x7c\x37\xf2\x1d\x33\xd9\x5b\x7e\x57\xb9\xed\xd9\xcb\x9e\xa5\x9c\x95\x55\x32\xa7\xb7\xb9\xcc\xc9\x65\x4e\x26\xe7\xe4\x99\x8d\xd7\x32\x67\x90\x4f\x2e\x67\x64\x72\x46\x1e\x4d\x36\xce\x67\x5e\x2f\xf2\x94\xa9\xcd\xe3\x77\x25\xdf\x51\xc6\x3b\xc6\x90\x45\x7e\x07\x99\x5f\xc9\xfa\x38\x98\x2c\xbd\xec\x53\x74\xf6\x69\xe5\x9c\x5a\xde\xa7\x99\xbd\x2b\xe5\xcc\x86\x3d\x45\xc6\x1a\x5d\x65\xac\x96\x75\x71\x34\x99\x1b\xf9\xd4\x8d\xcd\x4f\x27\x93\x33\x17\x1d\xaa\x60\x6b\xb0\x51\x26\xeb\x72\xf9\x1d\x98\x8f\x3d\x45\xe7\x42\x6c\x96\xca\xbc\x28\xbf\x4b\x99\x3b\xc8\x7e\x85\x7c\x72\x6c\xea\xcf\xf5\x64\x36\xc1\x2f\xd8\xad\xad\xcd\x57\x9d\xac\x6b\x6a\xdb\x67\x94\x7d\x06\x19\x0b\x72\x46\x21\xf3\x6a\x19\xef\x64\x2c\xca\xef\x4e\x64\x0f\x22\x4f\x87\x0d\xe4\xf7\x28\x9f\xb6\xb4\xf5\x7a\x86\x9c\x55\x62\x87\xd9\x4e\xc1\xed\x27\xe7\x16\x95\xed\x49\x5c\xb0\x16\x3f\xab\xdf\x47\xb3\x7f\x5f\xdb\xfa\x09\xdd\x26\x1b\x6b\x64\x4d\x26\xe3\x85\xd8\x63\xc0\x6e\xf2\x3e\x15\xd9\x26\x79\x0e\x8c\xe3\x7f\x99\x37\x0e\xee\xe7\xd6\x6c\xc9\x73\x21\xf3\xf3\x60\x36\xc6\x86\x69\x30\xbd\xb1\xf3\x94\x99\xdd\x43\x61\xf1\xd2\x34\x36\xaf\x91\x35\xb1\x36\xff\x72\x46\xc5\x3a\xf9\x1e\x53\x5f\x2b\xef\x4a\xb1\x5b\x8e\x8d\x0b\xd3\xbf\x1b\xcc\x3e\xd8\x29\x10\xcf\xa5\xcb\x2d\x6b\x8b\xcc\xec\x1b\x26\xd3\x19\xbd\x88\x15\x7c\x91\xbb\x1f\xf0\x07\xf6\x1c\x3b\x8b\x75\x8d\x55\x59\x37\x74\xe6\x63\x62\x32\xe4\x26\x2b\xb6\x51\x99\xe5\x7d\x49\x5e\x34\x26\x7f\xec\x4c\x36\xe2\x81\xbd\x88\x75\xc6\xc7\xda\xec\x82\x8c\xd8\x4b\xcf\x08\xe6\xa3\x76\x34\x7b\xe7\xc4\x0d\xfe\x88\xe6\xff\xd0\x5b\x8e\xb0\xa6\x94\xe7\x29\x98\x0d\xc8\x4f\x72\x17\x9d\xf0\x27\xf1\x3f\x36\x16\xcb\x9d\xdb\x6d\xc4\x5e\xa5\xed\x41\xec\x37\xe4\x14\xeb\x6b\x8b\x91\x81\x9c\xca\x2d\x46\xdb\xde\x7d\xd4\x59\x6e\xeb\xfe\x9d\xbd\x23\x06\xd1\x05\x59\xc1\x02\xfc\x8f\x2e\x53\x6e\x71\x8a\x9d\xda\x60\xfb\x0d\xa9\x3d\x57\x95\xe7\x46\x66\x67\x0c\x93\xf9\xa8\xe9\xcc\xc6\xe4\x74\xa8\x0c\x4b\xb0\x9b\xda\x2b\x37\x4c\x22\x2f\xc6\xdc\x7c\xc7\xfb\x6a\x96\xb5\xb4\xf8\xc3\x5f\x6a\xe7\xd2\xe4\x06\x7b\xc8\x31\xc6\xf0\x39\xeb\xb1\x4d\x6c\x3c\x87\xc8\xd3\xca\xf0\x81\xf8\x1e\x0b\x8f\x0b\xd6\xa7\x16\x5f\xc8\x85\xfe\xa9\xfb\x19\x5b\x92\x13\x8c\x23\x43\x46\x1e\x21\xc7\x60\xd8\x83\x8d\xfb\xcc\x64\x98\xdc\xc6\xc8\x49\x3e\x56\xa9\xc5\x09\x31\x0e\x76\xd4\xe0\x41\x61\x76\x2c\x3c\x86\xf0\x27\x3e\x09\x9e\xdf\xe8\x8d\xbd\xd4\xa6\xe4\x51\x61\xef\xb0\x03\x38\x16\x46\xf3\x3f\x79\x4f\xde\x69\x9e\x82\x99\xd1\x74\xed\x27\x7b\x57\xa3\xff\x68\x38\x84\x6e\xe0\x00\xd8\x49\x8e\x20\x2f\xbe\xc5\x66\x60\xfd\x54\x59\xfe\x68\x1c\xe2\x8b\xcc\x64\x8c\xae\xbf\xda\xae\xb1\x1c\xc5\xa6\xd8\x07\x5c\xe0\x37\x71\x8d\xcc\x60\x5c\xdf\x5b\xce\x91\x97\xbc\x8b\x73\x8e\xce\x3a\x46\x8f\xa1\xda\x62\x81\xf3\x39\x17\xfc\x04\x7f\x4a\xb7\x99\xe6\x47\x30\xcc\x07\x03\x34\x9e\x5a\xc3\x5b\xcd\x25\x97\xb7\xad\xac\x76\x8c\x95\xe9\x89\x2f\xd8\x87\x1c\x28\x06\x8b\x3f\xb5\x15\x3a\xf9\xb7\xd6\x3a\xe2\x2b\xb3\x38\x22\x77\xc1\x6d\xc5\x5b\x7c\x17\x0d\xef\x89\x2d\xec\x0f\x46\x68\x3d\x49\x2d\x67\xb1\x35\x79\x9d\xf7\x5e\x0b\x3b\x93\x1b\xdc\x8e\x3e\xa7\xf6\x3a\x02\x4e\x13\xdb\x8d\x63\x3d\xfe\x01\x47\xc6\x60\x7a\x64\x9e\x87\xd8\x9d\xf7\xa5\xdb\x5c\x31\xca\x71\x4e\x6b\x86\xc8\x94\x0f\x86\x83\x63\x6b\xb5\x08\xb9\x15\x8f\x38\xb3\xb7\x9a\x41\x0c\x80\x09\xe8\x36\x63\x38\xf6\x00\x2b\xc0\x86\x26\x3e\xe4\x21\xf1\xcc\x5e\x43\x69\x78\x5c\xe4\x76\x46\xd6\x5a\x8e\x0e\xee\xa7\xd0\x1a\x9e\xce\x18\x41\x3d\x1b\x88\x8b\xc9\xe4\x9a\x64\x4d\x1b\xed\xfc\xde\xf3\x22\x04\x8b\xd7\xde\x6b\xc5\xe4\xb5\x98\x77\x7d\xb0\xba\x4d\x4e\x4c\xce\x29\xc8\x05\xad\x47\xd8\x34\x33\x3e\x81\xef\x7b\x97\xa5\xf3\x78\x45\x07\xb5\x77\x69\xfb\x60\x07\xf6\x07\xeb\xc8\x11\xfc\x4a\xfc\xa3\x4b\x35\x19\x4f\xc0\xef\xe8\x01\x8f\xa0\x06\xf5\x5e\x23\x33\xf0\x2d\x37\xfb\x83\xbb\x93\x63\x3c\xb5\x11\x79\x1b\xe7\x3d\x95\xe7\x14\x75\x84\xba\x8f\x5d\x91\x97\xb8\xee\x6a\x3b\xbf\xf2\x38\xd4\xe7\xd4\x30\x1b\x1f\x83\x2d\xd4\x61\x6a\x03\x38\x87\x4e\xe9\x68\x38\xc3\x33\x35\x93\x98\x8e\x93\xf3\xa2\xc1\x70\x84\x9c\x42\xc7\xd2\xeb\x94\xea\x55\x1b\x7e\xc0\xd1\xc0\x2d\x72\x9b\x5a\x34\xcc\xb8\x13\x8c\x53\x10\x0b\xc8\xaa\x31\x1b\x2d\xe6\xf1\x29\xb6\x63\x6d\xed\xf9\x45\x9e\x61\xd7\xd2\xf1\x4e\xcf\xec\x2d\x57\xd1\xa3\xac\xac\x16\x6b\xdc\xe7\xc6\x13\xc0\x55\xf2\x02\xbb\x8d\xce\x59\x90\x1f\xb9\xfb\xd1\xf3\x35\x5a\xdd\x02\x57\x6a\xc7\x2c\x70\x9d\xf5\x8a\x79\xad\x9d\xad\x75\xc1\xf1\x94\x9c\x53\xde\x03\xf6\x45\xdb\x8f\xdc\x82\x9f\xb1\xa6\xee\x2d\x0f\x35\x3e\x9d\xcb\x81\x71\xec\x45\xbd\x86\x8f\x12\x53\x9c\x55\x38\xe7\x21\x26\x27\x8f\x43\xd6\x36\x73\x4e\x07\xab\xb1\xe4\xc5\xe8\x35\x7b\x8c\xce\x41\x5d\x66\xde\x93\x4b\x85\xf3\x14\x62\xb7\xf4\x7a\x5d\x38\x8f\x21\x0e\xb0\x31\xb9\x30\xf3\x60\xce\xac\xc0\x34\xcf\x43\x62\x55\x6b\x54\x66\xd8\x48\x9d\xe1\xb9\x2d\x4c\x17\xc5\x96\xca\xe2\x8f\xba\x40\xfd\xc2\x6e\xc4\x07\x5c\x15\xdd\x94\x3b\x67\xc6\xad\xf1\x2b\x18\x40\xec\x13\x33\xc4\xcf\xcc\xe3\xa3\xf3\x63\xf0\x4e\x73\x22\x58\xcc\xf0\x4d\xcc\xd9\xbc\xc3\xb9\x9f\x7f\xba\x63\xf0\x9e\xf3\x4b\xbd\xfc\xdc\x99\x3e\xba\x0b\x90\x26\xf6\x99\xee\xe3\x58\x5e\x1f\x3e\x7b\x8f\x72\x28\x53\x8e\xee\x5b\xc7\xa7\x77\x42\xa0\x3f\x68\xe7\xfb\x58\x20\x6d\x7d\xef\xef\x17\x5e\xd4\xe9\xa5\x5e\xfe\xbe\x05\xd7\x26\x5a\x76\xfc\xac\x21\xfb\x40\xab\x7a\x96\xbc\xa4\x51\x42\x5b\x7a\x96\x14\x59\xd7\x1e\x27\x74\x98\x67\x8f\x1b\xd0\x65\x99\xa7\x47\x3a\x4e\xdf\x78\x66\x7d\xe5\x5f\xd6\xab\xbb\x65\x56\x36\x6d\xde\xb4\x4d\x5d\x1c\x27\xe9\xd1\xc7\xc5\x41\x8f\x0c\xdf\xa8\xd6\x1f\x54\xd5\xb3\xc4\x35\x46\xc0\x33\xfd\xff\xe3\xbd\x5b\xfa\xe3\x67\x7a\x42\xbf\x9b\xfa\xbd\x1d\xa1\xb2\xaf\xd2\x32\x85\xac\xd4\x4c\x70\xf4\xa6\x32\x90\x79\x64\x2e\x55\x98\x4c\x80\x39\x91\x09\x44\x2b\x28\x41\x54\x2b\x3b\x05\x59\x41\xb3\xdc\x51\xcb\x2b\x19\x99\x4c\xb6\x10\xa9\x5d\x67\xef\xb4\x1b\xcc\xad\x3b\x21\x92\x61\x9d\xed\x64\x15\x18\xf4\x55\x24\x9f\x2c\x5b\x66\x76\x45\xf4\x6b\x77\x46\x85\x77\x16\x4d\x26\x05\xcf\xe2\xe8\xd5\x50\x7f\x37\xce\x32\xa3\xb3\xb5\xc2\x50\x1b\x14\x22\x2b\x91\x51\xd1\x65\x34\x94\x83\x91\x75\x3e\x17\xa4\x53\xc6\xe0\xec\xaa\x75\x46\xc4\x37\x0c\x20\x75\x06\xab\x5d\x4b\x6e\x4c\x9b\x4a\xa2\x9d\x60\x69\xd9\xac\x9d\x95\x9f\x93\x39\x52\x50\xed\x40\x47\xba\x4c\x98\x24\xa8\x43\x35\xe1\x7d\xe5\xcc\x19\x39\xd4\x6f\x9d\xb1\xad\xce\x3b\xa7\xd4\x2b\x02\x4c\x1f\x94\xa2\xd2\x53\x95\xb5\xca\x77\x66\x27\xba\x04\xba\x4f\x65\xb5\xb9\x77\xdd\x95\x55\x46\xed\x2a\x5b\xeb\x8c\x4b\xf7\xbb\xca\x5c\x1b\x0b\x53\x16\xe1\xdd\x06\x67\x83\xba\xa0\x30\xac\x07\x1f\x23\x3f\xf6\xa1\xb2\xc3\x2e\x5a\x97\x93\x39\xa9\x3f\xa3\x03\xf2\xd1\x91\xb0\x5f\xf4\x9b\x04\xe4\x24\x66\x90\x89\xea\xcc\x1e\xda\x65\xa4\x66\x6f\xd0\x14\x94\x9c\xed\x48\x6c\x51\xf1\xa9\xee\x54\x3c\x2a\xf5\xe0\x9d\x1d\xfe\xe6\x7c\xaa\x17\xeb\x52\xef\x78\x39\x03\x39\x74\x6f\xaf\x3c\x54\x7c\xf6\xd5\x2a\x3a\x9a\x2f\x60\x4f\xca\x9c\xa3\xf9\x40\x19\x77\xeb\xac\xaa\xb6\xb5\x95\x77\x9b\x7a\xab\xd1\x3b\x4b\x6b\xad\xfa\x10\x73\x54\x62\x98\x53\xf4\x98\x99\xbc\xf3\xd6\x8e\xdc\xe3\x94\x35\xdd\x1c\x67\x95\xb1\x6b\xba\x28\xed\x58\x06\xeb\x16\xe9\x90\xb5\x33\xaf\xac\xea\x92\x3f\xe4\x00\x4c\x2b\x77\x06\x8d\xff\xd1\x0b\xd6\xc2\x5c\x64\x8b\xde\xe1\x60\x03\xec\xa8\x95\xca\xd9\x02\x9f\xce\xbb\xe2\xb4\x30\xa6\xa9\x15\x7b\x32\xa6\xc4\xfb\xd1\x6f\x3b\x94\xa9\x74\xf6\x0c\xdb\x2b\x9c\x51\x11\x43\x30\x07\xe6\x50\xa1\x89\x05\xf4\x6e\x9d\x75\x30\x1f\xf9\x74\x4d\xea\x37\x2e\xce\xf2\x5a\x67\x19\xc3\x9c\xc3\x85\x57\xc4\x60\x67\x11\x87\xca\xba\xbc\x2b\x52\x16\x58\x59\x6e\x70\x06\xec\xb0\x73\xbc\xe1\x1c\xba\x19\x6c\xc2\xd9\x54\xd9\xc9\x19\x1b\xf6\xa2\x1b\x87\x3d\x70\x23\x01\x1e\x90\x73\xe3\x68\xb1\x3d\xfa\x6d\x06\x67\xa3\x17\x39\xd4\xb8\x0e\xf8\x59\x19\x4f\xe3\xb7\x39\x99\xad\x25\x2f\x19\xa3\xfb\xd4\x9b\xad\xdc\xec\x80\x1d\xc1\x0d\xd6\xf4\x9e\xab\x6d\x63\xf9\x09\xb3\xe4\x37\x4c\x41\x3b\xd8\xc2\x6f\x0f\x4a\x93\x97\x4e\x49\xbb\x3e\xbf\x11\x83\x15\xb0\x8e\x0e\x02\x4c\x62\x8c\x18\x43\x56\x18\x74\x9e\x3a\x83\x69\xcd\xe7\xda\x9d\x0e\x96\xb7\xda\xa5\xe5\xd6\xd1\x75\xc1\x64\x51\xe6\xd2\x58\xce\x61\x0b\x72\x0a\xd6\x3e\xc7\x76\xe3\xdd\x2d\x1d\x0d\x1d\x15\xec\x6a\xee\xd6\x88\x27\xe2\x03\xfd\xc1\x79\x70\x8b\xb8\x0c\x9e\x23\xb0\x67\xed\xe0\x07\xb3\x1d\x71\x05\xc6\x11\x0f\x99\xdb\x1d\x9f\x81\xed\x30\x1e\x6e\x8b\xf2\xc6\xf4\x50\x59\x83\x77\xc7\x7e\x4b\x07\x16\x93\xbf\x7a\xe3\xe8\x5d\x5d\xf4\xae\x04\x5d\xb0\xff\x7c\x13\x85\x9d\xd4\x0f\x93\x61\xfe\xe7\x4c\x28\xf5\x9b\x2d\xec\x8c\xfc\x5a\x5f\xaa\x17\x98\xd0\x43\x9d\xfc\x4f\x79\xd0\xc3\x4e\x5f\x64\x41\x8f\xfe\x58\xf4\x1c\x07\x7a\xd8\xe5\xdf\x67\x40\xbf\xd1\xe5\x7f\xc5\x7f\x3e\xd1\xc5\xd9\x4f\x56\x65\xc5\xff\x0d\xfb\x99\xff\xf4\xf6\x7b\xe9\x0f\xe9\xa6\x97\x0a\x99\x91\xfb\x71\xbe\x10\x4f\x3d\x65\x1a\x9f\xd7\x1b\xfd\x01\xae\x28\xab\x8d\x37\x71\xda\xc0\x55\x96\xce\x34\x59\x7a\xa9\x5d\x5a\x49\xd1\x94\xcf\xad\xfc\x51\x5e\x49\x8f\xda\xcb\x7a\xeb\x0d\xe8\x30\x97\xdc\xca\x7e\xb3\x96\xd4\x21\xbd\x28\x19\xa4\x16\xb4\xa1\xf5\x4b\xe8\xd0\x3f\x5c\xc0\x50\x66\x0b\x2f\x27\x93\xc3\x3d\x69\xac\x97\x6c\xb9\x95\x7e\x4a\x04\xf0\x9a\xfa\xc5\xb1\x36\xfc\x99\x5f\x00\x07\xb3\x53\xea\x8d\x8a\x5e\xb0\x04\xd3\x4d\x2f\x8b\x27\x2f\x3b\x99\xc9\xa7\x17\x1f\xa9\xc9\x16\x1c\x5e\xa3\x5f\x32\xa8\x8c\x93\x5f\xb8\x17\x96\xc6\x94\x42\xbd\xd0\xaf\xfc\x12\xd5\xe5\xe0\x52\x47\xe9\x8f\x97\x19\x52\x1c\x58\x03\xb2\x2a\xa7\x55\x34\x90\xa4\x37\x34\x0d\x58\xa5\xc4\x91\xf2\xc0\x06\x90\x06\x55\x52\xbd\x6b\x93\x0f\xb9\xa0\x3c\xec\x4d\x13\x0c\x65\x02\x1a\x54\xaf\xc9\xd6\x34\x7e\xd9\xa0\xa5\x30\x37\xfa\x94\xfa\x1f\x44\x78\x66\x6f\xc6\xb4\x7c\xd6\x66\xe3\xc6\x7d\x35\x7a\x99\xe6\x92\x4b\x69\x5d\x6e\x76\x46\xd6\x19\x76\x95\x1a\x0d\x16\x33\x7a\x39\xe4\x17\xa4\x94\x7f\xa8\x12\xbe\x8b\xde\xc4\xeb\xa5\xab\x5f\x9a\x10\x3b\xc1\xe9\x1b\xf6\x03\x5e\x75\x6e\xe6\x34\xa2\xf2\xe6\x3b\x1a\x54\x4e\xee\x7b\x7c\x88\xad\x83\xfb\x9f\xf2\x48\xa9\xc3\x96\xc4\x1c\xb2\x2b\x04\xfb\x45\x0c\xf4\x9b\x38\xd7\x3f\x3c\xf8\x45\xb0\xd2\xe5\xda\xf6\x20\x7e\x7a\xff\x03\x45\xe9\x17\xef\xc5\x1c\x17\xb9\xcd\xed\xdd\x9e\xc1\xff\xf8\x80\x1f\x80\xf6\xdc\x2f\x79\xd1\x89\xf3\x2b\x6f\x2d\xf4\x92\x7e\xb4\x38\xfe\x1c\xc6\x69\xc4\xf5\x82\xb5\xf2\x8b\x8a\xca\xff\x08\xf3\x1c\x8c\x3f\x4a\xf8\x7b\x1c\xff\x57\x00\x00\x00\xff\xff\x0f\x34\x8a\x8d\x00\x20\x00\x00")

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

	info := bindataFileInfo{name: "db/migrations/bindata.go", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1478293565, 0)}
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
	"db/migrations/002_forum.sql": dbMigrations002_forumSql,
	"db/migrations/003_f_comp.sql": dbMigrations003_f_compSql,
	"db/migrations/bindata.go": dbMigrationsBindataGo,
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
		"migrations": &bintree{nil, map[string]*bintree{
			"001_initial.sql": &bintree{dbMigrations001_initialSql, map[string]*bintree{}},
			"002_forum.sql": &bintree{dbMigrations002_forumSql, map[string]*bintree{}},
			"003_f_comp.sql": &bintree{dbMigrations003_f_compSql, map[string]*bintree{}},
			"bindata.go": &bintree{dbMigrationsBindataGo, map[string]*bintree{}},
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

