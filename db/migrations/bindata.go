// Code generated by go-bindata.
// sources:
// db/migrations/001_initial.sql
// db/migrations/002_forum.sql
// db/migrations/003_f_comp.sql
// db/migrations/004_f_comp2.sql
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

var _dbMigrations004_f_comp2Sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xb1\x4e\xc3\x30\x14\x45\x77\x7f\xc5\xdd\x3a\xa0\x7e\x41\xa7\x14\x87\xe9\x11\x43\xb1\x25\xb6\xc8\x24\xaf\xc4\x22\x8d\x2d\xdb\x90\x7e\x3e\x0e\xa8\xa8\x03\xaa\x32\xbe\xeb\x7b\xad\x73\xb6\x5b\xdc\x9d\xdc\x7b\xb4\x99\x61\x82\x28\xe7\xcb\x33\xc1\x4d\x48\xdc\x65\xe7\x27\x6c\x4c\xd8\xc0\x25\xf0\x99\xbb\xcf\xcc\x3d\xe6\x81\x27\xe4\xa1\x44\xbf\xbb\xa5\x54\x0e\x1b\xc2\xe8\xb8\x17\xa2\x22\x5d\x1f\xa0\xab\x3d\xd5\x38\xb6\x9d\x3f\x05\xce\xee\xa7\x55\x49\x89\x7b\x45\xe6\xb1\x59\xda\xd1\x7f\x71\xdf\xda\x34\x60\xaf\x14\xa1\x51\x1a\x8d\x21\x82\xac\x1f\x2a\x43\x1a\x47\x3b\x26\xde\xad\xfb\x2e\xda\xb9\xcd\x7c\xce\xd0\xf5\xab\xde\x09\x71\x6d\x25\xfd\x3c\x5d\xbc\xfe\xa4\x96\x70\x95\x56\xf4\xe3\x58\x5e\xdf\x6c\xf7\x71\x03\x45\x1e\xd4\xd3\x7f\x6a\xb7\xf0\xaf\x37\x17\xfe\x82\xfe\x1d\x00\x00\xff\xff\xf0\x3a\xc6\xd8\x91\x01\x00\x00")

func dbMigrations004_f_comp2SqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations004_f_comp2Sql,
		"db/migrations/004_f_comp2.sql",
	)
}

func dbMigrations004_f_comp2Sql() (*asset, error) {
	bytes, err := dbMigrations004_f_comp2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/004_f_comp2.sql", size: 401, mode: os.FileMode(420), modTime: time.Unix(1478429348, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrationsBindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x99\x5b\x6f\x1b\xd7\x11\xc7\x9f\xc5\x4f\xb1\x11\x90\x80\x2a\x54\x69\xef\x17\x01\x79\x49\xec\x02\x7e\x88\x03\x34\xee\x53\x55\x08\x7b\x39\x47\x25\x2a\x91\x32\x49\xa5\xb2\x0d\x7f\xf7\xce\x6f\x66\x56\x92\x1d\xcb\x42\xd3\x16\xa8\x01\x9a\xdc\xb3\xe7\x32\xd7\xff\xfc\xe7\xe8\xf4\x34\xf9\x71\x33\x85\xe4\x32\xac\xc3\xb6\xdf\x87\x29\x19\xde\x25\x97\x9b\x3f\x0e\xab\xf5\xd4\xef\xfb\x93\x85\x4c\xd8\x6d\x6e\xb7\x63\xd8\x9d\xf1\x7b\x1a\x4e\xaf\x57\x97\x32\x73\xb5\x59\xef\x4e\xd3\x34\xbb\x58\xad\x57\xfb\x55\x7f\x75\xb2\x7b\x7b\xf5\xa5\x09\xf9\x45\xdc\x6c\x6f\xaf\x9f\x7a\x5d\x5c\xc4\x8b\x71\x73\x7d\xf3\xd4\xfb\xd2\xdf\xe7\x5f\x9e\x30\x8b\x79\xb9\xe1\xdd\x8b\x9f\x93\xd7\x3f\xbf\x49\x5e\xbe\x78\xf5\xe6\x9b\xc5\xe2\xa6\x1f\xff\xd1\x5f\x86\xe4\x61\xfa\x62\xb1\xba\xbe\xd9\x6c\xf7\xc9\x72\x71\x70\x38\xbc\xdb\x87\xdd\xa1\xfc\x60\xfb\x6d\xd8\xed\x4e\x2f\xdf\xaf\x6e\x18\x88\xd7\x7b\xbe\x56\x1b\xfb\xff\x74\xb5\xb9\xdd\xaf\xae\x78\xd8\xe8\x82\x9b\x7e\xff\xf7\xd3\xb8\xba\x0a\xfc\x60\x60\xb7\xdf\xae\xd6\x97\xfa\x6e\xbf\xba\x0e\x87\x8b\xa3\xc5\x22\xde\xae\xc7\xc4\xc5\xfb\x73\xe8\xa7\x25\x3f\x92\xbf\xfe\x8d\x63\x8f\x93\x75\x7f\x1d\x12\x5b\x76\x94\x2c\xe7\xd1\xb0\xdd\x6e\xb6\x47\xc9\x87\xc5\xc1\xe5\x7b\x7d\x4a\xce\xbe\x4f\x90\xea\xe4\x75\xf8\x27\x9b\x84\xed\x52\xc5\xe6\xf9\x87\xdb\x18\xe5\x99\x6d\x8f\x8e\x16\x07\xab\xa8\x0b\xbe\xf9\x3e\x59\xaf\xae\xd8\xe2\x60\x1b\xf6\xb7\xdb\x35\x8f\xc7\x89\xa8\x74\xf2\x92\xdd\xe3\xf2\x90\x8d\x92\x6f\xdf\x9e\x25\xdf\xfe\x7a\x68\x92\xe8\x59\xb2\xc7\xc7\xc5\xe2\xe0\xd7\x7e\x9b\x0c\xb7\x31\xb1\x73\xec\x90\xc5\xc1\x85\x89\xf3\x7d\xb2\xda\x9c\xfc\xb8\xb9\x79\xb7\xfc\x4e\xe6\x1c\x8b\x6c\xb2\x6a\xbc\x7a\x39\x4b\x7a\xf2\xe3\xd5\x66\x17\x96\xa2\xfe\x7f\x49\x1e\xb6\xb1\xfd\x9f\xd8\x48\x26\x9a\xdc\x3e\x28\x62\x9d\xfc\x80\xe8\xcb\xa3\x63\x66\x2c\xe4\xdd\xfe\xdd\x4d\x48\xfa\xdd\x2e\xec\x31\xf9\xed\xb8\x67\x17\xd5\xcf\xfd\x21\xc7\xac\xe3\x26\x49\x36\xbb\x93\x3f\x89\x5b\x5f\xc9\xc3\xfd\x3a\x77\xe1\x3c\xfe\x68\x07\xf5\xa1\xfc\x33\x37\x2e\x0e\x76\xab\xf7\xfa\xbc\x5a\xef\xeb\x72\x71\x70\x4d\x5a\x25\xf7\x9b\xfe\x24\x8f\x3a\xf8\x46\x22\x24\x21\x4c\x4e\xf8\xc5\x39\x1a\x2a\xcb\xb8\xfa\xfc\xac\xa3\xe4\xb5\x1c\xb1\x3c\xf2\x13\x38\xd3\xb5\x8c\xab\x13\x4e\x97\xc5\x4f\xaf\xfd\x45\xc4\x91\xb5\x2a\xcd\xa7\x4b\x11\xf4\xab\x4b\x91\x55\x96\x3e\x92\xfc\xd3\x0d\x50\xed\xb9\x0d\x50\x4e\xf6\xb8\x57\xf4\x37\x3b\xb8\xf6\x4f\x6f\xf2\x6a\xf7\x62\xb5\x95\x2d\x86\xcd\xe6\xea\xf1\xea\xfe\x6a\xf7\x8c\xe6\xef\x76\xa6\x78\xd8\xc6\x7e\x0c\x1f\x3e\x3e\x5a\xed\x21\x41\x94\x5f\x4c\xc3\x4f\xf7\xd0\xf0\x08\xc9\x7e\x79\x7b\x25\x81\x6e\x91\xb1\x3c\x3c\xbf\xcb\xe2\xf9\x5d\x3b\x9c\xdf\xa5\xad\x7c\x52\xff\x74\xe7\x77\x75\x90\x71\x1f\x8b\x32\xa7\x2f\xcf\xef\xaa\x5a\x3e\x32\xb7\x2a\xe4\x59\x3e\x85\xbc\xcb\x64\xbc\x91\xb9\xa1\x91\x71\xf9\x74\xe3\xf9\xdd\x20\xdf\x13\x73\x65\x9f\x6e\xb2\xef\x49\xbe\x27\xf9\x6e\xe4\x3b\x66\xb2\xb7\xfc\xae\x72\xdb\xb3\x97\x3d\x4b\x39\x2b\xab\x64\x4e\x6f\x73\x99\x93\xcb\x9c\x4c\xce\xc9\x33\x1b\xaf\x65\xce\x20\x9f\x5c\xce\xc8\xe4\x8c\x3c\x9a\x6c\x9c\xcf\xbc\x5e\xe4\x29\x53\x9b\xc7\xef\x4a\xbe\xa3\x8c\x77\x8c\x21\x8b\xfc\x0e\x32\xbf\x92\xf5\x71\x30\x59\x7a\xd9\xa7\xe8\xec\xd3\xca\x39\xb5\xbc\x4f\x33\x7b\x57\xca\x99\x0d\x7b\x8a\x8c\x35\xba\xca\x58\x2d\xeb\xe2\x68\x32\x37\xf2\xa9\x1b\x9b\x9f\x4e\x26\x67\x2e\x3a\x54\xc1\xd6\x60\xa3\x4c\xd6\xe5\xf2\x3b\x30\x1f\x7b\x8a\xce\x85\xd8\x2c\x95\x79\x51\x7e\x97\x32\x77\x90\xfd\x0a\xf9\xe4\xd8\xd4\x9f\xeb\xc9\x6c\x82\x5f\xb0\x5b\x5b\x9b\xaf\x3a\x59\xd7\xd4\xb6\xcf\x28\xfb\x0c\x32\x16\xe4\x8c\x42\xe6\xd5\x32\xde\xc9\x58\x94\xdf\x9d\xc8\x1e\x44\x9e\x0e\x1b\xc8\xef\x51\x3e\x6d\x69\xeb\xf5\x0c\x39\xab\xc4\x0e\xb3\x9d\x82\xdb\x4f\xce\x2d\x2a\xdb\x93\xb8\x60\x2d\x7e\x56\xbf\x8f\x66\xff\xbe\xb6\xf5\x13\xba\x4d\x36\xd6\xc8\x9a\x4c\xc6\x0b\xb1\xc7\x80\xdd\xe4\x7d\x2a\xb2\x4d\xf2\x1c\x18\xc7\xff\x32\x6f\x1c\xdc\xcf\xad\xd9\x92\xe7\x42\xe6\xe7\xc1\x6c\x8c\x0d\xd3\x60\x7a\x63\xe7\x29\x33\xbb\x87\xc2\xe2\xa5\x69\x6c\x5e\x23\x6b\x62\x6d\xfe\xe5\x8c\x8a\x75\xf2\x3d\xa6\xbe\x56\xde\x95\x62\xb7\x1c\x1b\x17\xa6\x7f\x37\x98\x7d\xb0\x53\x20\x9e\x4b\x97\x5b\xd6\x16\x99\xd9\x37\x4c\xa6\x33\x7a\x11\x2b\xf8\x22\x77\x3f\xe0\x0f\xec\x39\x76\x16\xeb\x1a\xab\xb2\x6e\xe8\xcc\xc7\xc4\x64\xc8\x4d\x56\x6c\xa3\x32\xcb\xfb\x92\xbc\x68\x4c\xfe\xd8\x99\x6c\xc4\x03\x7b\x11\xeb\x8c\x8f\xb5\xd9\x05\x19\xb1\x97\x9e\x11\xcc\x47\xed\x68\xf6\xce\x89\x1b\xfc\x11\xcd\xff\xa1\xb7\x1c\x61\x4d\x29\xcf\x53\x30\x1b\x90\x9f\xe4\x2e\x3a\xe1\x4f\xe2\x7f\x6c\x2c\x96\x3b\xb7\xdb\x88\xbd\x4a\xdb\x83\xd8\x6f\xc8\x29\xd6\xd7\x16\x23\x03\x39\x95\x5b\x8c\xb6\xbd\xfb\xa8\xb3\xdc\xd6\xfd\x3b\x7b\x47\x0c\xa2\x0b\xb2\x82\x05\xf8\x1f\x5d\xa6\xdc\xe2\x14\x3b\xb5\xc1\xf6\x1b\x52\x7b\xae\x2a\xcf\x8d\xcc\xce\x18\x26\xf3\x51\xd3\x99\x8d\xc9\xe9\x50\x19\x96\x60\x37\xb5\x57\x6e\x98\x44\x5e\x8c\xb9\xf9\x8e\xf7\xd5\x2c\x6b\x69\xf1\x87\xbf\xd4\xce\xa5\xc9\x0d\xf6\x90\x63\x8c\xe1\x73\xd6\x63\x9b\xd8\x78\x0e\x91\xa7\x95\xe1\x03\xf1\x3d\x16\x1e\x17\xac\x4f\x2d\xbe\x90\x0b\xfd\x53\xf7\x33\xb6\x24\x27\x18\x47\x86\x8c\x3c\x42\x8e\xc1\xb0\x07\x1b\xf7\x99\xc9\x30\xb9\x8d\x91\x93\x7c\xac\x52\x8b\x13\x62\x1c\xec\xa8\xc1\x83\xc2\xec\x58\x78\x0c\xe1\x4f\x7c\x12\x3c\xbf\xd1\x1b\x7b\xa9\x4d\xc9\xa3\xc2\xde\x61\x07\x70\x2c\x8c\xe6\x7f\xf2\x9e\xbc\xd3\x3c\x05\x33\xa3\xe9\xda\x4f\xf6\xae\x46\xff\xd1\x70\x08\xdd\xc0\x01\xb0\x93\x1c\x41\x5e\x7c\x8b\xcd\xc0\xfa\xa9\xb2\xfc\xd1\x38\xc4\x17\x99\xc9\x18\x5d\x7f\xb5\x5d\x63\x39\x8a\x4d\xb1\x0f\xb8\xc0\x6f\xe2\x1a\x99\xc1\xb8\xbe\xb7\x9c\x23\x2f\x79\x17\xe7\x1c\x9d\x75\x8c\x1e\x43\xb5\xc5\x02\xe7\x73\x2e\xf8\x09\xfe\x94\x6e\x33\xcd\x8f\x60\x98\x0f\x06\x68\x3c\xb5\x86\xb7\x9a\x4b\x2e\x6f\x5b\x59\xed\x18\x2b\xd3\x13\x5f\xb0\x0f\x39\x50\x0c\x16\x7f\x6a\x2b\x74\xf2\x6f\xad\x75\xc4\x57\x66\x71\x44\xee\x82\xdb\x8a\xb7\xf8\x2e\x1a\xde\x13\x5b\xd8\x1f\x8c\xd0\x7a\x92\x5a\xce\x62\x6b\xf2\x3a\xef\xbd\x16\x76\x26\x37\xb8\x1d\x7d\x4e\xed\x75\x04\x9c\x26\xb6\x1b\xc7\x7a\xfc\x03\x8e\x8c\xc1\xf4\xc8\x3c\x0f\xb1\x3b\xef\x4b\xb7\xb9\x62\x94\xe3\x9c\xd6\x0c\x91\x29\x1f\x0c\x07\xc7\xd6\x6a\x11\x72\x2b\x1e\x71\x66\x6f\x35\x83\x18\x00\x13\xd0\x6d\xc6\x70\xec\x01\x56\x80\x0d\x4d\x7c\xc8\x43\xe2\x99\xbd\x86\xd2\xf0\xb8\xc8\xed\x8c\xac\xb5\x1c\x1d\xdc\x4f\xa1\x35\x3c\x9d\x31\x82\x7a\x36\x10\x17\x93\xc9\x35\xc9\x9a\x36\xda\xf9\xbd\xe7\x45\x08\x16\xaf\xbd\xd7\x8a\xc9\x6b\x31\xef\xfa\x60\x75\x9b\x9c\x98\x9c\x53\x90\x0b\x5a\x8f\xb0\x69\x66\x7c\x02\xdf\xf7\x2e\x4b\xe7\xf1\x8a\x0e\x6a\xef\xd2\xf6\xc1\x0e\xec\x0f\xd6\x91\x23\xf8\x95\xf8\x47\x97\x6a\x32\x9e\x80\xdf\xd1\x03\x1e\x41\x0d\xea\xbd\x46\x66\xe0\x5b\x6e\xf6\x07\x77\x27\xc7\x78\x6a\x23\xf2\x36\xce\x7b\x2a\xcf\x29\xea\x08\x75\x1f\xbb\x22\x2f\x71\xdd\xd5\x76\x7e\xe5\x71\xa8\xcf\xa9\x61\x36\x3e\x06\x5b\xa8\xc3\xd4\x06\x70\x0e\x9d\xd2\xd1\x70\x86\x67\x6a\x26\x31\x1d\x27\xe7\x45\x83\xe1\x08\x39\x85\x8e\xa5\xd7\x29\xd5\xab\x36\xfc\x80\xa3\x81\x5b\xe4\x36\xb5\x68\x98\x71\x27\x18\xa7\x20\x16\x90\x55\x63\x36\x5a\xcc\xe3\x53\x6c\xc7\xda\xda\xf3\x8b\x3c\xc3\xae\xa5\xe3\x9d\x9e\xd9\x5b\xae\xa2\x47\x59\x59\x2d\xd6\xb8\xcf\x8d\x27\x80\xab\xe4\x05\x76\x1b\x9d\xb3\x20\x3f\x72\xf7\xa3\xe7\x6b\xb4\xba\x05\xae\xd4\x8e\x59\xe0\x3a\xeb\x15\xf3\x5a\x3b\x5b\xeb\x82\xe3\x29\x39\xa7\xbc\x07\xec\x8b\xb6\x1f\xb9\x05\x3f\x63\x4d\xdd\x5b\x1e\x6a\x7c\x3a\x97\x03\xe3\xd8\x8b\x7a\x0d\x1f\x25\xa6\x38\xab\x70\xce\x43\x4c\x4e\x1e\x87\xac\x6d\xe6\x9c\x0e\x56\x63\xc9\x8b\xd1\x6b\xf6\x18\x9d\x83\xba\xcc\xbc\x27\x97\x0a\xe7\x29\xc4\x6e\xe9\xf5\xba\x70\x1e\x43\x1c\x60\x63\x72\x61\xe6\xc1\x9c\x59\x81\x69\x9e\x87\xc4\xaa\xd6\xa8\xcc\xb0\x91\x3a\xc3\x73\x5b\x98\x2e\x8a\x2d\x95\xc5\x1f\x75\x81\xfa\x85\xdd\x88\x0f\xb8\x2a\xba\x29\x77\xce\x8c\x5b\xe3\x57\x30\x80\xd8\x27\x66\x88\x9f\x99\xc7\x47\xe7\xc7\xe0\x9d\xe6\x44\xb0\x98\xe1\x9b\x98\xb3\x79\x87\x73\x3f\xff\x74\xc7\xe0\x3d\xe7\x97\x7a\xf9\xb9\x33\x7d\x74\x17\x20\x4d\xec\x57\xba\x8f\x63\x79\x7d\xf8\xd5\x8b\x96\x43\x99\x72\x74\xdf\x3a\x3e\xbd\x13\x02\xfd\x41\x3b\xdf\xc7\x02\x69\xeb\x7b\x7f\xbf\xf0\xac\x4e\xcf\xf5\xf2\xf7\x2d\xb8\x36\xd1\xb2\xe3\x67\x0d\xd9\x07\x5a\xd5\xb3\xe4\x39\x8d\x12\xda\xd2\xb3\xa4\xc8\xba\xf6\x38\xa1\xc3\x3c\x7b\xdc\x80\x2e\xcb\x3c\x3d\xd2\x71\xfa\xc6\x33\xeb\x2b\xff\xb2\x5e\xdd\x2d\xb3\xb2\x69\xf3\xa6\x6d\xea\xe2\x38\x49\x8f\x3e\x2e\x0e\x7a\x64\xf8\x4e\xb5\xfe\xa0\xaa\x9e\x25\xae\x31\x02\x9e\xe9\xff\x1f\xef\xdd\xd2\x1f\x7f\xa5\x27\xf4\xcb\xab\xdf\xdb\x11\x2a\xfb\x2a\x2d\x53\xc8\x4a\xcd\x04\x47\x6f\x2a\x03\x99\x47\xe6\x52\x85\xc9\x04\x98\x13\x99\x40\xb4\x82\x12\x44\xb5\xb2\x53\x90\x15\x34\xcb\x1d\xb5\xbc\x92\x91\xc9\x64\x0b\x91\xda\x75\xf6\x4e\xbb\xc1\xdc\xba\x13\x22\x19\xd6\xd9\x4e\x56\x81\x41\x5f\x45\xf2\xc9\xb2\x65\x66\x57\x44\xbf\x76\x67\x54\x78\x67\xd1\x64\x52\xf0\x2c\x8e\x5e\x0d\xf5\x77\xe3\x2c\x33\x3a\x5b\x2b\x0c\xb5\x41\x21\xb2\x12\x19\x15\x5d\x46\x43\x39\x18\x59\xe7\x73\x41\x3a\x65\x0c\xce\xae\x5a\x67\x44\x7c\xc3\x00\x52\x67\xb0\xda\xb5\xe4\xc6\xb4\xa9\x24\xda\x09\x96\x96\xcd\xda\x59\xf9\x39\x99\x23\x05\xd5\x0e\x74\xa4\xcb\x84\x49\x82\x3a\x54\x13\xde\x57\xce\x9c\x91\x43\xfd\xd6\x19\xdb\xea\xbc\x73\x4a\xbd\x22\xc0\xf4\x41\x29\x2a\x3d\x55\x59\xab\x7c\x67\x76\xa2\x4b\xa0\xfb\x54\x56\x9b\x7b\xd7\x5d\x59\x65\xd4\xae\xb2\xb5\xce\xb8\x74\xbf\xab\xcc\xb5\xb1\x30\x65\x11\xde\x6d\x70\x36\xa8\x0b\x0a\xc3\x7a\xf0\x31\xf2\x63\x1f\x2a\x3b\xec\xa2\x75\x39\x99\x93\xfa\x33\x3a\x20\x1f\x1d\x09\xfb\x45\xbf\x49\x40\x4e\x62\x06\x99\xa8\xce\xec\xa1\x5d\x46\x6a\xf6\x06\x4d\x41\xc9\xd9\x8e\xc4\x16\x15\x9f\xea\x4e\xc5\xa3\x52\x0f\xde\xd9\xe1\x6f\xce\xa7\x7a\xb1\x2e\xf5\x8e\x97\x33\x90\x43\xf7\xf6\xca\x43\xc5\x67\x5f\xad\xa2\xa3\xf9\x02\xf6\xa4\xcc\x39\x9a\x0f\x94\x71\xb7\xce\xaa\x6a\x5b\x5b\x79\xb7\xa9\xb7\x1a\xbd\xb3\xb4\xd6\xaa\x0f\x31\x47\x25\x86\x39\x45\x8f\x99\xc9\x3b\x6f\xed\xc8\x3d\x4e\x59\xd3\xcd\x71\x56\x19\xbb\xa6\x8b\xd2\x8e\x65\xb0\x6e\x91\x0e\x59\x3b\xf3\xca\xaa\x2e\xf9\x43\x0e\xc0\xb4\x72\x67\xd0\xf8\x1f\xbd\x60\x2d\xcc\x45\xb6\xe8\x1d\x0e\x36\xc0\x8e\x5a\xa9\x9c\x2d\xf0\xe9\xbc\x2b\x4e\x0b\x63\x9a\x5a\xb1\x27\x63\x4a\xbc\x1f\xfd\xb6\x43\x99\x4a\x67\xcf\xb0\xbd\xc2\x19\x15\x31\x04\x73\x60\x0e\x15\x9a\x58\x40\xef\xd6\x59\x07\xf3\x91\x4f\xd7\xa4\x7e\xe3\xe2\x2c\xaf\x75\x96\x31\xcc\x39\x5c\x78\x45\x0c\x76\x16\x71\xa8\xac\xcb\xbb\x22\x65\x81\x95\xe5\x06\x67\xc0\x0e\x3b\xc7\x1b\xce\xa1\x9b\xc1\x26\x9c\x4d\x95\x9d\x9c\xb1\x61\x2f\xba\x71\xd8\x03\x37\x12\xe0\x01\x39\x37\x8e\x16\xdb\xa3\xdf\x66\x70\x36\x7a\x91\x43\x8d\xeb\x80\x9f\x95\xf1\x34\x7e\x9b\x93\xd9\x5a\xf2\x92\x31\xba\x4f\xbd\xd9\xca\xcd\x0e\xd8\x11\xdc\x60\x4d\xef\xb9\xda\x36\x96\x9f\x30\x4b\x7e\xc3\x14\xb4\x83\x2d\xfc\xf6\xa0\x34\x79\xe9\x94\xb4\xeb\xf3\x1b\x31\x58\x01\xeb\xe8\x20\xc0\x24\xc6\x88\x31\x64\x85\x41\xe7\xa9\x33\x98\xd6\x7c\xae\xdd\xe9\x60\x79\xab\x5d\x5a\x6e\x1d\x5d\x17\x4c\x16\x65\x2e\x8d\xe5\x1c\xb6\x20\xa7\x60\xed\x73\x6c\x37\xde\xdd\xd2\xd1\xd0\x51\xc1\xae\xe6\x6e\x8d\x78\x22\x3e\xd0\x1f\x9c\x07\xb7\x88\xcb\xe0\x39\x02\x7b\xd6\x0e\x7e\x30\xdb\x11\x57\x60\x1c\xf1\x90\xb9\xdd\xf1\x19\xd8\x0e\xe3\xe1\xb6\x28\x6f\x4c\x0f\x95\x35\x78\x77\xec\xb7\x74\x60\x31\xf9\xab\x37\x8e\xde\xd5\x45\xef\x4a\xd0\x05\xfb\xcf\x37\x51\xd8\x49\xfd\x30\x19\xe6\x7f\xce\x84\x52\xbf\xd9\xc2\xce\xc8\xaf\xf5\xa5\x7a\x86\x09\x3d\xd4\xc9\xff\x94\x07\x3d\xec\xf4\x45\x16\xf4\xe8\xaf\x49\x5f\xe3\x40\x0f\xbb\xfc\xfb\x0c\xe8\x37\xba\xfc\xaf\xf8\xcf\x27\xba\x38\xfb\xc9\xaa\xac\xf8\xbf\x61\x3f\xf3\xdf\xe6\x7e\x2f\xfd\x21\xdd\xf4\x52\x21\x33\x72\x3f\xce\x17\xe2\xa9\xa7\x4c\xe3\xf3\x7a\xa3\x3f\xc0\x15\x65\xb5\xf1\x26\x4e\x1b\xb8\xca\xd2\x99\x26\x4b\x2f\xb5\x4b\x2b\x29\x9a\xf2\xb9\x95\x3f\xca\x2b\xe9\x51\x7b\x59\x6f\xbd\x01\x1d\xe6\x92\x5b\xd9\x6f\xd6\x92\x3a\xa4\x17\x25\x83\xd4\x82\x36\xb4\x7e\x09\x1d\xfa\x87\x0b\x18\xca\x6c\xe1\xe5\x64\x72\xb8\x27\x8d\xf5\x92\x2d\xb7\xd2\x4f\x89\x00\x5e\x53\xbf\x38\xd6\x86\x3f\xf3\x0b\xe0\x60\x76\x4a\xbd\x51\xd1\x0b\x96\x60\xba\xe9\x65\xf1\xe4\x65\x27\x33\xf9\xf4\xe2\x23\x35\xd9\x82\xc3\x6b\xf4\x4b\x06\x95\x71\xf2\x0b\xf7\xc2\xd2\x98\x52\xa8\x17\xfa\x95\x5f\xa2\xba\x1c\x5c\xea\x28\xfd\xf1\x32\x43\x8a\x03\x6b\x40\x56\xe5\xb4\x8a\x06\x92\xf4\x86\xa6\x01\xab\x94\x38\x52\x1e\xd8\x00\xd2\xa0\x4a\xaa\x77\x6d\xf2\x21\x17\x94\x87\xbd\x69\x82\xa1\x4c\x40\x83\xea\x35\xd9\x9a\xc6\x2f\x1b\xb4\x14\xe6\x46\x9f\x52\xff\x83\x08\xcf\xec\xcd\x98\x96\xcf\xda\x6c\xdc\xb8\xaf\x46\x2f\xd3\x5c\x72\x29\xad\xcb\xcd\xce\xc8\x3a\xc3\xae\x52\xa3\xc1\x62\x46\x2f\x87\xfc\x82\x94\xf2\x0f\x55\xc2\x77\xd1\x9b\x78\xbd\x74\xf5\x4b\x13\x62\x27\x38\x7d\xc3\x7e\xc0\xab\xce\xcd\x9c\x46\x54\xde\x7c\x47\x83\xca\xc9\x7d\x8f\x0f\xb1\x75\x70\xff\x53\x1e\x29\x75\xd8\x92\x98\x43\x76\x85\x60\xbf\x88\x81\x7e\x13\xe7\xfa\x87\x07\xbf\x08\x56\xba\x5c\xdb\x1e\xc4\x4f\xef\x7f\xa0\x28\xfd\xe2\xbd\x98\xe3\x22\xb7\xb9\xbd\xdb\x33\xf8\x1f\x1f\xf0\x03\xd0\x9e\xfb\x25\x2f\x3a\x71\x7e\xe5\xad\x85\x5e\xd2\x8f\x16\xc7\x9f\xc3\x38\x8d\xb8\x5e\xb0\x56\x7e\x51\x51\xf9\x1f\x61\x3e\x83\xf1\x7f\x05\x00\x00\xff\xff\x6b\x69\xa3\x20\x00\x20\x00\x00")

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

	info := bindataFileInfo{name: "db/migrations/bindata.go", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1478429349, 0)}
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
	"db/migrations/004_f_comp2.sql": dbMigrations004_f_comp2Sql,
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
			"004_f_comp2.sql": &bintree{dbMigrations004_f_comp2Sql, map[string]*bintree{}},
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

