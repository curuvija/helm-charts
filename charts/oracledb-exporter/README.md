# oracledb-exporter

A Helm chart to export queriy results from Oracle DB

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.3.0rc1](https://img.shields.io/badge/AppVersion-0.3.0rc1-informational?style=flat-square) 

## Additional Information

[oracledb-exporter](https://github.com/iamseth/oracledb_exporter) exposes Prometheus metrics based on Oracle SQL queries. 
You can find more details about it here https://github.com/albertodonato/query-exporter. 

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Milos Curuvija | curuvija@live.com |  |

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add curuvija https://curuvija.github.io/helm-charts/
$ helm repo update
$ helm install curuvija/oracledb-exporter --version 0.1.0
```

## Configure Prometheus scraping

If you use Prometheus operator PodMonitor will be created to configure your instance to scrape it.

If you don't use Prometheus operator then you can use this configuration to configure scraping (and disable PodMonitor creation in Helm values):

```yaml
    additionalScrapeConfigs:
    - job_name: oracledb-exporter-scrape
      honor_timestamps: true
      scrape_interval: 15s
      scrape_timeout: 10s
      metrics_path: /metrics
      scheme: http
      follow_redirects: true
      relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app, __meta_kubernetes_pod_labelpresent_app]
        separator: ;
        regex: (oracledb-exporter);true
        replacement: $1
        action: keep
      - source_labels: [__meta_kubernetes_pod_container_port_number]
        separator: ;
        regex: "9161"
        replacement: $1
        action: keep
      kubernetes_sd_configs:
      - role: pod
```
## Grafana Dashboard

There is an example Grafana dashboard here https://grafana.com/grafana/dashboards/3333.


## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | configure affinity |
| autoscaling.enabled | bool | `false` | enable or disable autoscaling |
| autoscaling.maxReplicas | int | `100` | maximum number of replicas |
| autoscaling.minReplicas | int | `1` | minimum number of replicas |
| autoscaling.targetCPUUtilizationPercentage | int | `80` | configure at what percentage to trigger autoscalling |
| dbConnection | object | `{"createDbConnectionSecret":true,"dbDonnectionString":"system/YOUR-PASSWORD-FOR-SYSTEM@//database:1521/DB_SID.DB_DOMAIN"}` | define connection to your database |
| fullnameOverride | string | `""` | overrides name without having chartName in front of it |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"iamseth/oracledb_exporter","tag":""}` | Image to use for deployment |
| image.pullPolicy | string | `"IfNotPresent"` | define pull policy |
| image.repository | string | `"iamseth/oracledb_exporter"` | repository to pull image |
| image.tag | string | `""` | Overrides the image tag whose default is the chart appVersion. |
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
| livenessProbe | object | `{"httpGet":{"path":"/","port":9161}}` | configure liveness probe |
| metricsConfig | string | `"[[metric]]\ncontext = \"test\"\nrequest = \"SELECT 1 as value_1, 2 as value_2 FROM DUAL\"\nmetricsdesc = { value_1 = \"Simple example returning always 1.\", value_2 = \"Same but returning always 2.\" }\n"` | define metrics to expose to Prometheus |
| nameOverride | string | `""` | overrides name (partial name override - chartName + nameOverride) |
| nodeSelector | object | `{}` | define node selector to schedule your pod(s) |
| podAnnotations | object | `{"prometheus.io/path":"/metrics","prometheus.io/port":"9161","prometheus.io/scrape":"true"}` | additional pod annoations |
| podAnnotations."prometheus.io/path" | string | `"/metrics"` | controls for Prometheus scrapes |
| podSecurityContext | object | `{}` | define pod security context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| prometheus | object | `{"monitor":{"additionalLabels":{},"enabled":true,"interval":"15s","namespace":[],"path":"/metrics"}}` | configure Prometheus Service monitor to expose metrics |
| prometheus.monitor.additionalLabels | object | `{}` | add additonal labels to service monitoring |
| prometheus.monitor.enabled | bool | `true` | enable or disable creation of service monitor |
| prometheus.monitor.interval | string | `"15s"` | Prometheus scraping interval |
| prometheus.monitor.namespace | list | `[]` | provide namespace where to create this service monitor |
| prometheus.monitor.path | string | `"/metrics"` | path where you want to expose metrics |
| readinessProbe | object | `{"httpGet":{"path":"/","port":9161}}` | configure readiness probe |
| replicaCount | int | `1` | replicaCount - number of pods to run |
| resources | object | `{"limits":{"cpu":"100m","memory":"128Mi"},"requests":{"cpu":"100m","memory":"128Mi"}}` | specify resources |
| resources.limits | object | `{"cpu":"100m","memory":"128Mi"}` | specify resource limits |
| resources.limits.cpu | string | `"100m"` | specify resource limits for cpu |
| resources.limits.memory | string | `"128Mi"` | specify resource limits for memory |
| resources.requests.cpu | string | `"100m"` | specify resource requests for cpu |
| resources.requests.memory | string | `"128Mi"` | specify resource requests for memory |
| securityContext | object | `{"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":1000}` | define security context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container |
| securityContext.readOnlyRootFilesystem | bool | `true` | Mounts the container's root filesystem as read-only. |
| securityContext.runAsNonRoot | bool | `true` | run docker container as non root user. |
| securityContext.runAsUser | int | `1000` | specify under which user all processes inside container will run. |
| service | object | `{"port":9161,"type":"ClusterIP"}` | service configuration |
| service.port | int | `9161` | service port |
| service.type | string | `"ClusterIP"` | service type |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `""` | If not set and create is true, a name is generated using the fullname template |
| tolerations | list | `[]` | provide tolerations |


----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)