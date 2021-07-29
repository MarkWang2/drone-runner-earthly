### Cache build
```yaml
  - name: build
    image: golang
    commands:
      - --mount=type=cache,target=/go-cache go build
```
Use [buildkit cache way](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md#run---mounttypecache), like docker mount a volume
--mount=type=cache,target=/go-cache go build

### build and push docker image
```yaml
  - name: build-push
    image: DOCKERFILE
    commands:
      - SAVE IMAGE --push go-demo:latest
```
Use commands SAVE IMAGE --push, it uses the local docker authorization to push image.

### Use Earthfile target
Run Earthfile specify target, just need set same name with Earthfile target
```yaml
  - name: deps
```