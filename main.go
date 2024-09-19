// canodb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"

	args "github.com/lindsuen/canodb/flag"
	"github.com/lindsuen/canodb/leveldb"
)

var (
	ldb     *leveldb.DB
	err     error
	dmpFile *os.File
)

func main() {
	args.ParseFlagArgs()

	if isExist(args.DataDirectory) && isDir(args.DataDirectory) {
		ldb, err = leveldb.OpenFile(args.DataDirectory, nil)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("canodb-cli: CanoDB has been connected.")
		}
		defer ldb.Close()
	} else {
		fmt.Println("canodb-cli: The directory " + args.DataDirectory + " is not existential.")
		os.Exit(1)
	}

	if isExist(args.ExportFile) {
		dmpFile, err = os.OpenFile(args.ExportFile, os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		dmpFile, err = os.Create(args.ExportFile)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer dmpFile.Close()

	iter := ldb.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println("canodb-cli: key: " + string(key) + ", value: " + string(value))

		err = dmpFile.Sync()
		if err != nil {
			fmt.Println(err)
		}
		w := bufio.NewWriter(dmpFile)
		_, err = w.Write(keyValueEncoding(key, value))
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

func isDir(s string) bool {
	d, err := os.Stat(s)
	if err != nil {
		return false
	}
	return d.IsDir()
}

func isExist(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

func keyValueEncoding(key, value []byte) []byte {
	return []byte(base64Encoding(key) + ":" + base64Encoding(value) + "\n")
}

func base64Encoding(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func base64Decoding(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
