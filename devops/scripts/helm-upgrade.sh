#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

TARGET_DIR=./devops/helmchart
NAMESPACE_YAML="${TARGET_DIR}/${NAMESPACE}.yaml"
TAGS_YAML="${TARGET_DIR}/tags.yaml"

# if a yaml specific to the passed in namespace is found, use it
if test -f $NAMESPACE_YAML; then
    VALUES_YAML=$NAMESPACE_YAML
    SOPS_YAML="${TARGET_DIR}/${NAMESPACE}.sops.yaml"
    PLAIN_YAML="${TARGET_DIR}/${NAMESPACE}.plain.yaml"
else
    # otherwise, use the default for gcp
    VALUES_YAML="${TARGET_DIR}/_gcp-default.yaml"
    SOPS_YAML="${TARGET_DIR}/_gcp-default.sops.yaml"
    PLAIN_YAML="${TARGET_DIR}/_gcp-default.plain.yaml"
fi

# TODO: integrate SOPS with Github Actions
sops -d $SOPS_YAML > $PLAIN_YAML

kubectl create namespace $NAMESPACE || true

# TODO: helm still doesn't seem to support STDIN to avoid writing plain.yaml
helm upgrade --install scorpicode $TARGET_DIR \
-n $NAMESPACE \
-f $TAGS_YAML \
-f $VALUES_YAML \
-f $PLAIN_YAML
