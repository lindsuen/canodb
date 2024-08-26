// canodb - main_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"testing"

	"github.com/lindsuen/canodb/leveldb"
)

func TestLevelDB(t *testing.T) {
	db, err := leveldb.OpenFile("data", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.Put([]byte("ss-01"), []byte("aaaaaaaaaaaaaaaaaaaaaaaaaa"), nil)
	db.Put([]byte("ss-02"), []byte("ssssssssssssssssssssssssss"), nil)
	db.Put([]byte("ss-03"), []byte("dddddddddddddddddddddddddd"), nil)
}
