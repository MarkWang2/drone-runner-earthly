### About Drone earthly runner
The earthly runner executes pipelines use earthly [https://github.com/earthly/earthly].

Earthly is a build automation tool for the container era，build on top of BuildKit [https://github.com/moby/buildkit].
BuildKit is a toolkit for converting source code to build artifacts in an efficient, expressive and repeatable manner.  
It is subproject of moby, use as the new version docker image builder but not just use for image build, it can build every thing.

BuildKit closer to CI/CD than docker, Earthly extend the BuildKit, make the BuildKit more powerful and easily to use,
use Earthly much easily to implement than use docker as a droneCI runner.

### Build Custom earthly
#### codebase
https://github.com/MarkWang2/earthly
#### Build cmd
GOOS=linux GOARCH=amd64 go build \
-ldflags "--X main.DefaultBuildkitdImage=earthly/buildkitd:main" \
-tags netgo -installsuffix netgo \
-o build/earthly \
cmd/earthly/*.go

### How to run use docker image
docker run -it --privileged -d   -v /var/run/docker.sock:/var/run/docker.sock \
-e NO_DOCKER=1 \
-e DRONE_RPC_PROTO=http \
-e DRONE_RPC_HOST= server url  \
-e DRONE_RPC_SECRET=0bfd6037ecc1dtgfg257f9bae359be1 \
-e DRONE_RUNNER_CAPACITY=2 \
-e DRONE_LOGS_TRACE=true \
-e DRONE_RUNNER_NAME=${HOSTNAME} \
-v earthly-tmp:/tmp/earthly:rw \
-p 3000:3000 \
--restart always \
--name runner 18392019228/drone-runner-earthly:main3

### examples

https://github.com/MarkWang2/drone-runner-earthly/tree/master/examples/go