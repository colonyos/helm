#!/bin/bash

namespace="colonyos-minio"
helm install ${namespace} -f values.yaml -n ${namespace} .
