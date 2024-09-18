// canodb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/lindsuen/canodb/leveldb"
)

var (
	f *os.File

	dataDirectory string
	exportFile    string
	importFile    string
)

func main() {
	parseFlagArgs()
	if dataDirectory == "" {
		fmt.Println("canodb-cli : error: -d must be specified")
		os.Exit(1)
	}

	db, err := leveldb.OpenFile(dataDirectory, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if fileExist(exportFile) {
		f, err = os.OpenFile(exportFile, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		f, err = os.Create(exportFile)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer f.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println("key: " + string(key) + ", value: " + string(value))

		err = f.Sync()
		if err != nil {
			fmt.Println(err)
		}
		w := bufio.NewWriter(f)
		_, err = w.Write(encodeKV(key, value))
		if err != nil {
			fmt.Println(err)
		}
		err = w.Flush()
		if err != nil {
			fmt.Println(err)
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}

func fileExist(file string) bool {
	_, err := os.Stat(file)
	return os.IsExist(err)
}

func parseFlagArgs() {
	flag.StringVar(&dataDirectory, "d", "", "The data directory of CanoDB.")
	flag.StringVar(&exportFile, "e", "db.dump", "Export the data.")
	flag.StringVar(&importFile, "i", "db.dump", "Import the data.")
	flag.Parse()
}

func encodeKV(key, value []byte) []byte {
	return []byte(encode(key) + ":" + encode(value) + "\n")
}

func encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func decode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
