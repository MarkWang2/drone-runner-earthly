// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package command

import (
	"context"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"

	"github.com/MarkWang2/drone-runner-earthly/command/daemon"
)

// program version
var version = "0.0.0"

// empty context
var nocontext = context.Background()

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New("drone", "drone docker runner")
	registerCompile(app)
	registerExec(app)
	registerCopy(app)
	daemon.Register(app)

	kingpin.Version(version)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
