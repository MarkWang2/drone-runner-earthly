go build \
-ldflags "--X main.DefaultBuildkitdImage=earthly/buildkitd:main" \
-tags netgo -installsuffix netgo \
-o build/earthly \
cmd/earthly/*.go

