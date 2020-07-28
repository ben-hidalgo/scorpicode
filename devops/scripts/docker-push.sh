#!/bin/bash
set -exuo pipefail


# docker push [GCR_HOSTNAME]/[GCR_PROJECT]/[IMAGE]:[TAG]

HATS_TAG=`yq r devops/helmchart/tags.yaml hats.tag`
WEBSITE_TAG=`yq r devops/helmchart/tags.yaml website.tag`
FRONTEND_TAG=`yq r devops/helmchart/tags.yaml frontend.tag`
ROXIE_TAG=`yq r devops/helmchart/tags.yaml roxie.tag`
SOXIE_TAG=`yq r devops/helmchart/tags.yaml soxie.tag`
DEBUGGER_TAG=`yq r devops/helmchart/tags.yaml debugger.tag`

docker push $GCR_HOSTNAME/$GCR_PROJECT/hats:$HATS_TAG
docker push $GCR_HOSTNAME/$GCR_PROJECT/website:$WEBSITE_TAG
docker push $GCR_HOSTNAME/$GCR_PROJECT/frontend:$FRONTEND_TAG
docker push $GCR_HOSTNAME/$GCR_PROJECT/roxie:$ROXIE_TAG
docker push $GCR_HOSTNAME/$GCR_PROJECT/soxie:$SOXIE_TAG
docker push $GCR_HOSTNAME/$GCR_PROJECT/debugger:$DEBUGGER_TAG
