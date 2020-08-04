#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

HELM_DIR=./devops/helmchart
SECRETS_DIR=./secrets/decrypted
NAMESPACE_YAML="${HELM_DIR}/${NAMESPACE}.yaml"
TAGS_YAML="${HELM_DIR}/tags.yaml"

# if a yaml specific to the passed in namespace is found, use it
if test -f $NAMESPACE_YAML; then
    VALUES_YAML=$NAMESPACE_YAML
    SECRETS_YAML="${SECRETS_DIR}/${NAMESPACE}.yaml"
else
    # otherwise, use the default for gcp
    VALUES_YAML="${HELM_DIR}/_gcp-default.yaml"
    SECRETS_YAML="${SECRETS_DIR}/_gcp-default.yaml"
fi

kubectl create namespace $NAMESPACE || true

# TODO: temporary block until sops works in GHA
if [ $NAMESPACE == "dev" ]; then
    echo "skipping sops"

    helm upgrade --install scorpicode $HELM_DIR \
    -n $NAMESPACE \
    -f $TAGS_YAML \
    -f $VALUES_YAML

else

    helm upgrade --install scorpicode $HELM_DIR \
    -n $NAMESPACE \
    -f $TAGS_YAML \
    -f $VALUES_YAML \
    -f $SECRETS_YAML
fi
