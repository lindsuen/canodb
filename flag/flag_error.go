// canodb - flag_error.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package flag

import (
	"fmt"
	"os"
)

func argsErrorHandling() {
	if ExportFile != "" || ImportFile != "" {
		if DataDirectory == "" {
			fmt.Println("canodb-cli: The parameter '-d' must be specified.")
			os.Exit(1)
		}
	}
	if DataDirectory == "" && ExportFile == "" && ImportFile == "" {
		fmt.Println("canodb-cli: Welcome to canodb-cli.")
		os.Exit(1)
	}
}
