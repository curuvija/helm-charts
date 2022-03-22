# query-exporter

A Helm chart for query-exporter (export Prometheus metrics from SQL queries)

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.8.0](https://img.shields.io/badge/AppVersion-2.8.0-informational?style=flat-square) 

## Additional Information

[query-exporter](https://github.com/albertodonato/query-exporter) exposes Prometheus metrics based on SQL queries. It supports 
different databases. You can find more details about it here https://github.com/albertodonato/query-exporter. This Helm 
chart is using docker image to deploy it easily on Kubernetes

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Milos Curuvija | curuvija@live.com |  |

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add curuvija https://curuvija.github.io/helm-charts/
$ helm repo update
$ helm install curuvija/query-exporter --version 0.1.0
```



## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | configure affinity |
| autoscaling.enabled | bool | `false` | enable or disable autoscaling |
| autoscaling.maxReplicas | int | `100` | maximum number of replicas |
| autoscaling.minReplicas | int | `1` | minimum number of replicas |
| autoscaling.targetCPUUtilizationPercentage | int | `80` | configure at what percentage to trigger autoscalling |
| config | string | `"databases:\n  db1:\n    dsn: sqlite://\n    connect-sql:\n      - PRAGMA application_id = 123\n      - PRAGMA auto_vacuum = 1\n    labels:\n      region: us1\n      app: app1\n\nmetrics:\n  metric1:\n    type: gauge\n    description: A sample gauge\n\nqueries:\n  query1:\n    interval: 5\n    databases: [db1]\n    metrics: [metric1]\n    sql: SELECT random() / 1000000000000000 AS metric1"` | Configure your database access and metrics to expose |
| fullnameOverride | string | `""` | overrides name without having chartName in front of it |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"adonato/query-exporter","tag":"latest"}` | Image to use for deployment |
| image.pullPolicy | string | `"IfNotPresent"` | define pull policy |
| image.repository | string | `"adonato/query-exporter"` | repository to pull image |
| image.tag | string | `"latest"` | Overrides the image tag whose default is the chart appVersion. |
| imagePullSecrets | list | `[]` | Image pull secrets if you want to host the image |
| ingress | object | `{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]}` | ingress configuration |
| ingress.annotations | object | `{}` | ingress annotations |
| ingress.className | string | `""` | ingress class name |
| ingress.enabled | bool | `false` | enable or disable ingress configuration creation |
| ingress.hosts | list | `[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]` | hosts |
| ingress.hosts[0] | object | `{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}` | hostname |
| ingress.hosts[0].paths | list | `[{"path":"/","pathType":"ImplementationSpecific"}]` | paths |
| ingress.hosts[0].paths[0] | object | `{"path":"/","pathType":"ImplementationSpecific"}` | path |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` | path type |
| ingress.tls | list | `[]` | tls configuration |
| nameOverride | string | `""` | overrides name (partial name override - chartName + nameOverride) |
| nodeSelector | object | `{}` | define node selector to schedule your pod(s) |
| podAnnotations | object | `{}` | pod annotations |
| podSecurityContext | object | `{}` | define pod security context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| prometheus | object | `{"monitor":{"additionalLabels":{},"enabled":true,"interval":"15s","namespace":[],"path":"/metrics"}}` | configure Prometheus Service monitor to expose metrics |
| prometheus.monitor.additionalLabels | object | `{}` | add additonal labels to service monitoring |
| prometheus.monitor.enabled | bool | `true` | enable or disable creation of service monitor |
| prometheus.monitor.interval | string | `"15s"` | Prometheus scraping interval |
| prometheus.monitor.namespace | list | `[]` | provide namespace where to create this service monitor |
| prometheus.monitor.path | string | `"/metrics"` | path where you want to expose metrics |
| replicaCount | int | `1` | replicaCount - number of pods to run |
| resources | object | `{}` | specify resources |
| securityContext | object | `{}` | define security context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container |
| service | object | `{"port":9560,"type":"ClusterIP"}` | service configuration |
| service.port | int | `9560` | service port |
| service.type | string | `"ClusterIP"` | service type |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `""` | If not set and create is true, a name is generated using the fullname template |
| tolerations | list | `[]` | provide tolerations |


----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)