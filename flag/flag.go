// canodb - flag.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package flag

import "flag"

var (
	DataDirectory string
	ExportFile    string
	ImportFile    string
)

func ParseFlagArgs() {
	flag.StringVar(&DataDirectory, "d", "", "The data directory of CanoDB.")
	flag.StringVar(&ExportFile, "e", "", "Export the data from CanoDB.")
	flag.StringVar(&ImportFile, "i", "", "Import the data from CanoDB.")
	flag.Parse()
	argsErrorHandling()
}
