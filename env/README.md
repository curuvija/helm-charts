
# Configure test Prometheus environment

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
kubectl create ns monitoring
```

Now install it:

```
helm install kube--prometheus-stack prometheus-community/kube-prometheus-stack --version 34.1.1 -n monitoring --values env/k8s/kube-prometheus-stack-values.yaml
```

Or upgrade if you change some settings:

```
helm upgrade kube--prometheus-stack prometheus-community/kube-prometheus-stack --version 34.1.1 -n monitoring --values env/k8s/kube-prometheus-stack-values.yaml
```