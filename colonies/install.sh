#!/bin/bash

namespace="johank-colonies"
helm install colonies -f values.yaml -n ${namespace} .
