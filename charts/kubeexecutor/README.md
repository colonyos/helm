# ColonyOS Kubernetes executor

This chart will deploy a kubernetes executor for ColonyOS

## Prerequisites

- A running ColonyOS server with a colony to connect to.

## Installation

### Install from repository

```console
helm repo add colonyos https://colonyos.github.io/helm
```

```console
helm get values colonyos/kubeexecutor > values.yaml
```

Edit values.yaml

```console
helm install my-kubeexecutor colonyos/kubeexecutor
```

### Install from local file system

```console
helm install kubeexecutor ./ -f values.yaml
```

## Configuration options

| Setting                | Description                                                                                             | Example value                   |
| ---                    | -----------                                                                                             | ---                             |
| NumberOfPods           | Number of pods to deploy.                                                                               | 3                               |
| ExecutorsPerPod        | Number of executors per pod. Each executor runs a container inside the pod.                              | 5                               |
| ExecutorResourceLimit  | Enable CPU and memory limits.                                                                           | true/false                      |
| ExecutorCPU            | CPU request.                                                                                            | "4000m"                         |
| ExecutorMemory         | Memory request and limit.                                                                               | "100Mi"                         |
| ExecutorDockerImage    | Executor docker image.                                                                                  | "colonyos/sleepexecutor:v0.0.1" |
| ColoniesServerHost     | Hostname of Colonies server.                                                                            | "colonies-service.colonyos"     |
| ColoniesServerPort     | Port of Colonies server.                                                                                | 80                              |
| ColoniesColonyID       | Colony ID.                                                                                              | "4787a5071856a4acf702b..."      |
| ColoniesColonyPrvKey   | Colony private key. If set, the executor will self-register and not use the ColoniesExecutorID setting. | "ba949fa13498162b6a56f..."      |
| ColoniesExecutorID     | Executor ID                                                                                             | "fc05cf3df4b494e95d6a3..."      |
| ColoniesExecutorPrvKey | Executor private key.                                                                                   | "ddf7f779120875a72684f..."      |
