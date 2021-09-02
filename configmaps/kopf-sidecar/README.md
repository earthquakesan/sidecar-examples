# Kopf Sidecar Configmap Processing

As suggested in the [kiwigrid sidecar issue](https://github.com/kiwigrid/k8s-sidecar/issues/85), trying out another [sidecar](https://github.com/OmegaVVeapon/kopf-k8s-sidecar), which does not stop working after some time.

```
k apply -f configmap.yaml
k get cm config -o yaml

# Requires access to my repo, see Makefile -)
# This step should not be required as I already published image on docker hub
make build-push

k apply -f deployment.yaml

k logs -f -l app=configmap-watcher -c configmap-watcher

# Modify configmap.yaml
k apply -f configmap.yaml
```

Summary: this works as expected, it is basically a rewrite of kiwigrid-sidecar. There is no bug as in kiwigrid sidecar, however it is GPL licensed. Can't be used because of the license.
