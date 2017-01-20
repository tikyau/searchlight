# Alert Objects

Alert objects are consumed by [Searchlight Controller]() to create Icinga2 hosts, services and notifications.

Before we can create an Alert object we must create the [Alert Third Party Resource](docs/alert-resource/third-party-resource.md) in your Kubernetes cluster.

## Alert Object Fields

* apiVersion - The Kubernetes API version. See [Alert Third Party Resource](docs/alert-resource/third-party-resource.md).
* kind - The Kubernetes object type.
* metadata.name - The name of the Alert object.
* metadata.namespace - The namespace of the Alert object
* metadata.labels - The Kubernetes object labels. This labels is used to determine for whom this alert will be set.
* spec.checkCommand - Icinga CheckCommand name
* spec.icingaParam - IcingaParam contains parameters for Icinga config
* spec.notifierParams - NotifierParams contains array of information to send notifications for Incident
* spec.vars - Vars contains array of Icinga Service variables to be used in CheckCommand.


#### IcingaParam Fields

* checkIntervalSec - How frequently Icinga Service will be checked
* alertIntervalSec - How frequently notifications will be send

#### NotifierParam Fields

* state - For which state notification will be sent
* userUid - To whom notification will be sent
* method - How this notification will be sent

> `NotifierParams` is only used when notification is sent via `AppsCode`.

#### Metadata Labels
* alert.appscode.com/objectType - The Kubernetes object type
* alert.appscode.com/objectName - The Kubernetes object name

### Example

The following Alert Object will do the followings:

* This Alert is set on ReplicationController named `elasticsearch-logging-v1` in `kube-system` namespace.
* CheckCommand `volume` will be applied.
* Icinga Service will check volume every 60s.
* Notifications will be send every 5m if any problem is detected.
* Email will be sent as a notification to admin user for CRITICAL state. For other states, no notification will be sent.
* On each Pod under specified RC, volume named `disk` will be checked. If volume is used more than 60%, it is WARNING. For 75%, it is CRITICAL.

Example Alert Object

```
apiVersion: appscode.com/v1beta1
kind: Alert
metadata:
  name: check-es-logging-volume
  namespace: kube-system
  labels:
    alert.appscode.com/objectType: replicationcontrollers
    alert.appscode.com/objectName: elasticsearch-logging-v1
spec:
  CheckCommand: volume
  IcingaParam:
    CheckIntervalSec: 60
    AlertIntervalSec: 300
  NotifierParams:
  - Method: EMAIL
    State: CRITICAL
    UserUid: admin
  Vars:
    name: disk
    warning: 60.0
    critical: 75.0
```


#### CheckCommand

We currently supports following CheckCommands:

* [component_status](docs/check-command/component_status.md) - To check Kubernetes components.
* [influx_query](docs/check-command/influx_query.md) - To check InfluxDB query result.
* [json_path](docs/check-command/json_path.md) - To check any API response by parsing JSON using JQ queries.
* [node_count](docs/check-command/node_count.md) - To check total number of Kubernetes node.
* [node_status](docs/check-command/node_status.md) - To check Kubernetes Node status.
* [pod_exists](docs/check-command/pod_exists.md) - To check Kubernetes pod existence.
* [pod_status](docs/check-command/pod_status.md) - To check Kubernetes pod status.
* [prometheus_metric](docs/check-command/prometheus_metric.md) - To check Prometheus query result.
* [node_disk](docs/check-command/node_disk.md) - To check Node Disk stat.
* [volume](docs/check-command/volume.md) - To check Pod volume stat.
* [kube_event](docs/check-command/kube_event.md) - To check Kubernetes events for all Warning TYPE happened in last 'c' seconds.
* [kube_exec](docs/check-command/kube_exec.md) - To check Kubernetes exec command. Returns OK if exit code is zero, otherwise, returns CRITICAL