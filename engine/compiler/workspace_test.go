// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package compiler

import (
	"testing"

	"github.com/MarkWang2/drone-runner-earthly2/engine"
	"github.com/MarkWang2/drone-runner-earthly2/engine/resource"
	"github.com/drone/runner-go/manifest"
)

func TestSetupWorkspace(t *testing.T) {
	tests := []struct {
		path string
		src  *resource.Step
		dst  *engine.Step
		want string
	}{
		{
			path: "/drone/src",
			src:  &resource.Step{},
			dst:  &engine.Step{},
			want: "/drone/src",
		},
		// do not override the user-defined working dir.
		{
			path: "/drone/src",
			src:  &resource.Step{},
			dst:  &engine.Step{WorkingDir: "/foo"},
			want: "/foo",
		},
		// do not override the default working directory
		// for service containers with no commands.
		{
			path: "/drone/src",
			src:  &resource.Step{},
			dst:  &engine.Step{Detach: true},
			want: "",
		},
		// overrides the default working directory
		// for service containers with commands.
		{
			path: "/drone/src",
			src:  &resource.Step{Commands: []string{"whoami"}},
			dst:  &engine.Step{Detach: true},
			want: "/drone/src",
		},
	}
	for _, test := range tests {
		setupWorkdir(test.src, test.dst, test.path)
		if got, want := test.dst.WorkingDir, test.want; got != want {
			t.Errorf("Want working_dir %s, got %s", want, got)
		}
	}
}

func TestToWindows(t *testing.T) {
	got := toWindowsDrive("/go/src/github.com/octocat/hello-world")
	want := "c:\\go\\src\\github.com\\octocat\\hello-world"
	if got != want {
		t.Errorf("Want windows drive %q, got %q", want, got)
	}
}

func TestCreateWorkspace(t *testing.T) {
	tests := []struct {
		from *resource.Pipeline
		base string
		path string
		full string
	}{
		{
			from: &resource.Pipeline{
				Workspace: resource.Workspace{
					Base: "",
					Path: "",
				},
			},
			base: "/drone/src",
			path: "",
			full: "/drone/src",
		},
		{
			from: &resource.Pipeline{
				Workspace: resource.Workspace{
					Base: "",
					Path: "",
				},
				Platform: manifest.Platform{
					OS: "windows",
				},
			},
			base: "c:\\drone\\src",
			path: "",
			full: "c:\\drone\\src",
		},
		{
			from: &resource.Pipeline{
				Workspace: resource.Workspace{
					Base: "/drone",
					Path: "src",
				},
			},
			base: "/drone",
			path: "src",
			full: "/drone/src",
		},
		{
			from: &resource.Pipeline{
				Workspace: resource.Workspace{
					Base: "/drone",
					Path: "src",
				},
				Platform: manifest.Platform{
					OS: "windows",
				},
			},
			base: "c:\\drone",
			path: "src",
			full: "c:\\drone\\src",
		},
		{
			from: &resource.Pipeline{
				Workspace: resource.Workspace{
					Base: "/foo",
					Path: "bar",
				},
			},
			base: "/foo",
			path: "bar",
			full: "/foo/bar",
		},
	}
	for _, test := range tests {
		base, path, full := createWorkspace(test.from)
		if got, want := test.base, base; got != want {
			t.Errorf("Want workspace base %s, got %s", want, got)
		}
		if got, want := test.path, path; got != want {
			t.Errorf("Want workspace path %s, got %s", want, got)
		}
		if got, want := test.full, full; got != want {
			t.Errorf("Want workspace %s, got %s", want, got)
		}
	}
}
