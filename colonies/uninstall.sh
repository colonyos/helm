#!/bin/bash

namespace="johank-colonies"

helm uninstall colonies -n ${namespace}
