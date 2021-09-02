# No Sidecar Configmap Processing

```
k create configmap config --from-file=config.yaml=$(pwd)/config.yaml

k get cm config -o yaml

# Requires access to my repo, see Makefile -)
make build-push

k apply -f deployment.yaml

k logs -f -l app=configmap-watcher
```

Summary: golang app with fsnotify does not track changes to the configmap mounted as a file.