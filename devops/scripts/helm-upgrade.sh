#!/bin/bash
set -exuo pipefail

#sops -d ./devops/helmchart/local.sops.yaml > ./devops/helmchart/local.plain.yaml

helm upgrade --install scorpicode ./devops/helmchart \
-f devops/helmchart/dev.yaml \
--set common.auth0RedirectUri=http://`minikube ip`:30080/callback \
--set roxie.loginSuccessTarget=http://`minikube ip`:30080/sc \
--set frontend.socketHost=`minikube ip`:30081
