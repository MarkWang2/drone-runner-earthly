package compiler

import (
	"github.com/MarkWang2/drone-runner-earthly/engine"
	"github.com/earthly/earthly/ast/spec"
	"strings"
)

func toCloneTarget(step *engine.Step, sourcedir string) spec.Target {
	rp := spec.Block{}
	rp = buildBlock("FROM", []string{step.Image}, rp)
	rp = buildBlock("WORKDIR", []string{step.WorkingDir}, rp)
	for key, value := range step.Envs {
		rp = buildBlock("ENV", []string{key, "=", value}, rp)
	}
	rp = buildBlock("RUN", []string{"sh", "/usr/local/bin/clone"}, rp)
	rp = buildBlock("SAVE ARTIFACT", []string{".", "AS", "LOCAL", sourcedir}, rp)
	target := spec.Target{step.Name, rp, nil}
	return target
}

func toTarget(step *engine.Step) spec.Target {
	rp := spec.Block{}
	from := strings.Fields(step.Image)
	if strings.ToUpper(from[0]) == "DOCKERFILE" {
		var args []string
		if len(from) > 1 {
			args = from[1:]
		} else {
			args = []string{"."}
		}
		rp = buildBlock("FROM DOCKERFILE", args, rp)
		for _, cmd := range step.Commands {
			cmdItems := strings.Fields(cmd)
			if strings.Join(cmdItems[0:2], " ") == "SAVE IMAGE" {
				rp = buildBlock("SAVE IMAGE", cmdItems[2:], rp)
			}
		}
	} else {
		rp = buildBlock("FROM", []string{step.Image}, rp)
		rp = buildBlock("WORKDIR", []string{step.WorkingDir}, rp)
		for key, value := range step.Envs {
			rp = buildBlock("ENV", []string{key, "=", value}, rp)
		}
		rp = buildBlock("COPY", []string{".", step.WorkingDir}, rp)
		for _, cmd := range step.Commands {
			cmsStr := strings.Fields(cmd)
			rp = buildBlock("RUN", cmsStr, rp)
		}
	}
	target := spec.Target{step.Name, rp, nil}
	return target
}

func buildBlock(name string, args []string, rp spec.Block) spec.Block {
	cmd := &spec.Command{Name: name, Args: args}
	sm := spec.Statement{cmd, nil, nil, nil, nil}
	return append(rp, sm)
}
