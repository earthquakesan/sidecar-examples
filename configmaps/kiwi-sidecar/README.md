# Kiwi Sidecar Configmap Processing

```
k apply -f configmap.yaml
k get cm config -o yaml

# Requires access to my repo, see Makefile -)
make build-push

k apply -f deployment.yaml

k logs -f -l app=configmap-watcher -c configmap-watcher

# Modify configmap.yaml
k apply -f configmap.yaml
```

Summary: this works as expected, requires to create service account with role/binding. Also there is a known bug in kiwigrid sidecar (https://github.com/kiwigrid/k8s-sidecar/issues/134 and https://github.com/kiwigrid/k8s-sidecar/issues/85), it stops updating files after a while.