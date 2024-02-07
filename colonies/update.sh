#!/bin/bash

namespace="johank-colonies"
helm upgrade colonies -f values.yaml -n ${namespace} --debug --wait .
