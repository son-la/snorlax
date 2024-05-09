# snorlax


## Development environment

## Run local
```
make run ARG="--version"
```

### Test chart on local k3s
1. Start k3s cluster
```
k3d cluster create test-cluster --servers 1 --agents 3 -p "30000-30100:30000-30100@server:0"
```

2. Flux bootstrap. Permission for Github PAT https://fluxcd.io/flux/installation/bootstrap/github/#github-organization
```
flux bootstrap github \
  --owner=son-la \
  --repository=flux-fleet \
  --branch=main \
  --path=clusters/snorlax-local \
  --token-auth \
  --components-extra=image-reflector-controller,image-automation-controller \
  --version=v2.2.3
```


3. To test Kafka, use Java 11.0.22