## Curuvija's Helm Repository

This is my repository for Helm charts. 

## Add repository:

```bash
helm repo add curuvija https://curuvija.github.io/helm-charts/
```
## Now update it:

```bash
helm repo update
```
## Search through repository:

```bash
helm search repo curuvija
```

## Install Helm chart:

```bash
helm install curuvija/query-exporter --version 0.1.0
```
