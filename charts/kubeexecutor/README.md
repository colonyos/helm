# Introduction
This Helm chart installs a Kube Executor.

## Authentication 
```console
INFO[0000] Generated new private key  
Id=5cec61bf8df8f595d72b58ebdfa8f7d7780558f12aba0feaf1769eb889fadd60 
PrvKey=951d80c64cdb3918cb089c6ad1996c249319451afeffd77b8bebba4c8b00cf7f
```

## Register executor
```console
colonies executor add --spec kube-executor.json --executorid 5cec61bf8df8f595d72b58ebdfa8f7d7780558f12aba0feaf1769eb889fadd60 --approve
```

Add PrvKey to `values.yaml` under `ColoniesPrvKey`.

# Installation
Edit `values.yaml` and type:

```console
./install.sh
```

# Configuration options
| Setting                | Description                                                                                             | Example value                   |
| ---                    | -----------                                                                                             | ---                             |
| NumberOfPods           | Number of pods to deploy.                                                                               | 3                               |
| ExecutorsPerPod        | Number of exectors per pod. Each executor runs a container inside the pod.                              | 5                               |
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
