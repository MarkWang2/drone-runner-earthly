// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/drone/runner-go/pipeline/runtime"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/client"
)

// Opts configures the Docker engine.
type Opts struct {
	HidePull bool
}

// Docker implements a Docker pipeline engine.
type Earthly struct {
	client   client.APIClient
	hidePull bool
}

// New returns a new engine.
func New(client client.APIClient, opts Opts) *Earthly {
	return &Earthly{
		client:   client,
		hidePull: opts.HidePull,
	}
}

// NewEnv returns a new Engine from the environment.
func NewEnv(opts Opts) (*Earthly, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return New(cli, opts), nil
}

// Ping pings the Docker daemon.
func (e *Earthly) Ping(ctx context.Context) error {
	_, err := e.client.Ping(ctx)
	return err
}

// Setup the pipeline environment.
func (e *Earthly) Setup(ctx context.Context, specv runtime.Spec) error {
	return nil
}

// Destroy the pipeline environment.
func (e *Earthly) Destroy(ctx context.Context, specv runtime.Spec) error {
	spec := specv.(*Spec)
	err := os.RemoveAll(spec.Root)
	return err
}

// Run runs the pipeline step.
func (e *Earthly) Run(ctx context.Context, specv runtime.Spec, stepv runtime.Step, output io.Writer) (*runtime.State, error) {
	spec := specv.(*Spec)
	step := stepv.(*Step)
	var cmd *exec.Cmd
	dir := spec.WorkingDir
	efByes, _ := json.Marshal(step.Earthfile)
	targetName := dir + "+" + step.Name
	fmt.Print(targetName)
	if step.Image == "" {
		cmd = runEarthly(output, targetName)
	} else {
		cmd = runEarthly(output, "--target-ats-json", string(efByes), "--push", targetName)
	}
	var err error
	done := make(chan error)
	go func() {
		err = cmd.Start()
		done <- cmd.Wait()
	}()

	select {
	case err = <-done:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	state := &runtime.State{
		ExitCode:  0,
		Exited:    true,
		OOMKilled: false,
	}
	if err != nil {
		state.ExitCode = 255
	}
	return state, err
}

func runEarthly(output io.Writer, args ...string) *exec.Cmd {
	// when debug use code base local earthly  cmd := exec.Command("./earthly", args...)
	cmd := exec.Command("earthly", args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "FORCE_COLOR=1") // can pass from env when start
	cmd.Stdout = output
	cmd.Stderr = output
	return cmd
}
