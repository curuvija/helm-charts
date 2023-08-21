# oracledb-exporter

A Helm chart to export queriy results from Oracle DB

## Additional Information

[oracledb-exporter](https://github.com/iamseth/oracledb_exporter) exposes Prometheus metrics based on Oracle SQL queries.

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Milos Curuvija | <curuvija@live.com> |  |

## Creating secret

You can create secret by enabling ``dbConnection`` in values but this is not recommended unless you use it for test environments
or you test connection string.

```yaml
dbConnection:
  createDbConnectionSecret: true
  dbDonnectionString: "system/YOUR-PASSWORD-FOR-SYSTEM@//database:1521/DB_SID.DB_DOMAIN"
```

Better suited option is to create the secret manually by converting connection string to ``base64``:

```bash
echo -n YOUR_CONNECTION_STRING_GOES_HERE | base64
```

And create secret like this:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: oracledb-exporter-secret
  namespace: monitoring
data:
  datasource: >-
    c3lzdGVtL21hbmFnZXJALy9lYnMtdi1vcmExNi5lsiuY3JlYWxvZ2l4Lm5ldDoxNTIxL01EQlQxNQ==
type: Opaque
```
Now the last step is to provide secret name in ``existingDbSecretName`` value.

```yaml
existingDbSecretName: "oracledb-exporter-secret"
```

Now you can install oracle db exporter.

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add curuvija https://curuvija.github.io/helm-charts/
$ helm repo update
$ helm install oracledb-exporter curuvija/oracledb-exporter
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
| dbConnection | object | `{"createDbConnectionSecret":false,"dbDonnectionString":"system/YOUR-PASSWORD-FOR-SYSTEM@//database:1521/DB_SID.DB_DOMAIN"}` | define connection to your database |
| dbConnection.createDbConnectionSecret | bool | `false` | creates secret unless you create it manually and provide value in existingDbSecretName |
| dbConnection.dbDonnectionString | string | `"system/YOUR-PASSWORD-FOR-SYSTEM@//database:1521/DB_SID.DB_DOMAIN"` | ads database connection string to datasource fields in secret |
| existingDbSecretName | string | `"oracledb-exporter-secret"` | provide the name of the secret containing db connection string |
| fullnameOverride | string | `""` | overrides name without having chartName in front of it |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"ghcr.io/iamseth/oracledb_exporter","tag":"0.5.1"}` | Image to use for deployment |
| image.pullPolicy | string | `"IfNotPresent"` | define pull policy |
| image.repository | string | `"ghcr.io/iamseth/oracledb_exporter"` | repository to pull image |
| image.tag | string | `"0.5.1"` | Overrides the image tag whose default is the chart appVersion. |
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
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
