module github.com/MarkWang2/drone-runner-earthly2

go 1.16

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible

require (
	github.com/buildkite/yaml v2.1.0+incompatible
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.5+incompatible
	github.com/drone/drone-go v1.6.0
	github.com/drone/envsubst v1.0.3
	github.com/drone/runner-go v1.8.0
	github.com/drone/signal v1.0.0
	github.com/earthly/earthly v0.5.17
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.4
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-isatty v0.0.12
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace (
	// estargz: needs this replace because stargz-snapshotter git repo has two go.mod modules.
	github.com/containerd/stargz-snapshotter/estargz => github.com/containerd/stargz-snapshotter/estargz v0.0.0-20201217071531-2b97b583765b
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	github.com/moby/buildkit => github.com/earthly/buildkit v0.0.0-20210609215831-00025901bf6b
	github.com/tonistiigi/fsutil => github.com/earthly/fsutil v0.0.0-20210609160335-a94814c540b2
)
