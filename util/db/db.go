// canodb - db.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package db

import (
	"log"

	"github.com/lindsuen/canodb/leveldb"
	. "github.com/lindsuen/canodb/util/config"
)

var (
	db  *leveldb.DB
	err error
)

func InitDB() {
	if DataConfig.DataDirectory == "" {
		DataConfig.DataDirectory = "data"
	}
	db, err = leveldb.OpenFile(DataConfig.DataDirectory, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func PutData(k, v []byte) error {
	return db.Put(k, v, nil)
}

func DeleteData(k []byte) error {
	return db.Delete(k, nil)
}

func GetData(k []byte) ([]byte, error) {
	return db.Get(k, nil)
}
