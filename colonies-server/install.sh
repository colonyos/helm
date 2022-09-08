#!/bin/bash

namespace="johank-demo"
helm install colonies -f values.yaml -n ${namespace} .
