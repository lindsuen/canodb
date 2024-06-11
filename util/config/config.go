// canodb - config.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package config

import (
	"github.com/go-ini/ini"
	"log"
)

type connectionConfig struct {
	Address string `ini:"listen_address"`
	Port    string `ini:"port"`
}

type sslConfig struct {
	SSL         string `ini:"ssl"`
	SSLCAFile   string `ini:"ssl_ca_file"`
	SSLCertFile string `ini:"ssl_cert_file"`
	SSLKeyFile  string `ini:"ssl_key_file"`
	SSLCiphers  string `ini:"ssl_ciphers"`
}

type dataConfig struct {
	DataDirectory string `ini:"data_directory"`
}

var ConnectionConfig = new(connectionConfig)
var SSLConfig = new(sslConfig)
var DataConfig = new(dataConfig)

var cfg = new(ini.File)

func InitConfig() {
	var configPath = "config/config.ini"
	var err error

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	mapTo("connection", ConnectionConfig)
	mapTo("ssl", SSLConfig)
	mapTo("data", DataConfig)
}

// Convert a Map to a struct.
func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalln(err)
	}
}
