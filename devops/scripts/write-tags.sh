#!/bin/bash
set -exuo pipefail

echo hats.tag=`./devops/scripts/go-checksum.sh hats` >> ./devops/helmchart/tags.yaml
echo roxie.tag=`./devops/scripts/go-checksum.sh roxie` >> ./devops/helmchart/tags.yaml
echo soxie.tag=`./devops/scripts/go-checksum.sh soxie` >> ./devops/helmchart/tags.yaml
echo website.tag=`./devops/scripts/js-checksum.sh website` >> ./devops/helmchart/tags.yaml
echo frontend.tag=`./devops/scripts/js-checksum.sh frontend` >> ./devops/helmchart/tags.yaml
