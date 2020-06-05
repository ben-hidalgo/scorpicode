#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

# TODO: add logic, if NAMESPACE is empty, use "default"

# TODO: integrate SOPS with Github Actions
#sops -d ./devops/helmchart/${NAMESPACE}.sops.yaml > ./devops/helmchart/${NAMESPACE}.plain.yaml

helm upgrade --install -n ${NAMESPACE} scorpicode ./devops/helmchart \
-f devops/helmchart/${NAMESPACE}.yaml \
-f devops/helmchart/tags.yaml \
--set common.auth0RedirectUri=http://`minikube ip`:30080/callback \
--set roxie.loginSuccessTarget=http://`minikube ip`:30080/sc \
--set frontend.socketHost=`minikube ip`:30081
