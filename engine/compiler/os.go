// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package compiler

import (
	"fmt"
	"strings"
)

// helper function returns the base temporary directory based
// on the target platform.
func tempdir(os string) string {
	dir := fmt.Sprintf("drone-%s", random())
	switch os {
	case "windows":
		return join(os, "C:\\Windows\\Temp", dir)
	default:
		return join(os, "/tmp", dir)
	}
}

// helper function joins the file paths.
func join(os string, paths ...string) string {
	switch os {
	case "windows":
		return strings.Join(paths, "\\")
	default:
		return strings.Join(paths, "/")
	}
}
