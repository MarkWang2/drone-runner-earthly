// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package main

import (
	"github.com/MarkWang2/drone-runner-earthly/command"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	command.Command()
}
