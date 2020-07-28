#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

# TODO: add logic, if NAMESPACE is `whoami` use "cloud.yaml" but `whoami` namespace
# TODO: add logic for cloud.yaml in GitHub actions namespace is branch name

# TODO: integrate SOPS with Github Actions
#sops -d ./devops/helmchart/${NAMESPACE}.sops.yaml > ./devops/helmchart/${NAMESPACE}.plain.yaml

kubectl create namespace ${NAMESPACE} || true

helm upgrade --install -n ${NAMESPACE} scorpicode ./devops/helmchart \
-f devops/helmchart/${NAMESPACE}.yaml \
-f devops/helmchart/tags.yaml \
-f devops/helmchart/local.plain.yaml
