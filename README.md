# ColonyOS Helm Charts

## Introduction

This repo contains Helm charts to deploy Colonies, including several executors.

| Helm chart     | Purpose                                                         |
| :---           | :-----------                                                    |
| notebook       | Helm chart to deploy Colonies Jypyter notebook                  |
| colonies       | Helm chart to deploy Colonies                                   |
| timescaledb    | Helm chart to deploy TimescaleDB server with a Backup executor  |
| sleepexecutor  | Helm chart to deploy a Sleep executor                           |
| echoexecutor   | Helm chart to deploy an Echo executor                           |
| k8sexecutor    | Helm chart to deploy a K8s executor                             |
| pythonexecutor | Helm chart to deploy a Python executor                          |

## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

```sh
helm repo add colonyos https://colonyos.github.io/helm
```

If you had already added this repo earlier, run `helm repo update` to retrieve the latest versions of the packages.  You can then run `helm search repo colonyos` to see the charts.

To install the `<chart-name>` chart (e.g. `colonies`):

```sh
helm install my-<chart-name> colonyos/<chart-name>
```

To uninstall the chart:

```sh
helm delete my-<chart-name>
```

You can install the chart as a [subchart](https://helm.sh/docs/chart_template_guide/subcharts_and_globals/) of another chart by adding the following to your `Chart.yaml`:

```yaml
dependencies:
  - name: <chart-name>
    version: <chart-version>
    repository: https://colonyos.github.io/helm
```
