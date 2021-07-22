module github.com/MarkWang2/drone-runner-earthly2

go 1.16

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible

require (
	//github.com/earthly/earthly  v0.5.17
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/antlr/antlr4 v0.0.0-20200225173536-225249fdaef5 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/buildkite/yaml v2.1.0+incompatible
	github.com/containerd/containerd v1.5.0 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.5+incompatible
	github.com/drone/drone-go v1.6.0
	github.com/drone/drone-runtime v1.1.0 // indirect
	github.com/drone/drone-yaml v1.2.3 // indirect
	github.com/drone/envsubst v1.0.3
	github.com/drone/runner-go v1.8.0
	github.com/drone/signal v1.0.0
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/earthly/earthly v0.5.17 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.4
	github.com/google/uuid v1.2.0 // indirect
	github.com/jdxcode/netrc v0.0.0-20210204082910-926c7f70242a // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/hashstructure/v2 v2.0.1 // indirect
	github.com/moby/buildkit v0.8.2-0.20210129065303-6b9ea0c202cf
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/otiai10/copy v1.1.1 // indirect
	github.com/pieterclaerhout/go-log v1.14.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/tonistiigi/fsutil v0.0.0-20210525040343-5dfbf5db66b9 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	google.golang.org/grpc v1.35.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gotest.tools v2.2.0+incompatible // indirect
)

replace (
	// estargz: needs this replace because stargz-snapshotter git repo has two go.mod modules.
	github.com/containerd/stargz-snapshotter/estargz => github.com/containerd/stargz-snapshotter/estargz v0.0.0-20201217071531-2b97b583765b
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	github.com/moby/buildkit => github.com/earthly/buildkit v0.0.0-20210609215831-00025901bf6b
	github.com/tonistiigi/fsutil => github.com/earthly/fsutil v0.0.0-20210609160335-a94814c540b2
)

