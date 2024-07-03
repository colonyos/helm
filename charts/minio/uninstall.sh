#!/bin/bash

namespace="colonyos-minio"
helm uninstall minio -n ${namespace}
