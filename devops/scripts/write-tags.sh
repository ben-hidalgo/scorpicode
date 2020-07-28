#!/bin/bash
set -exuo pipefail

FILE=./devops/helmchart/tags.yaml

echo -e "### `date`" > $FILE

echo -e "hats:\n  tag: `./devops/scripts/go-checksum.sh hats`" >> $FILE
echo -e "roxie:\n  tag: `./devops/scripts/go-checksum.sh roxie`" >> $FILE
echo -e "soxie:\n  tag: `./devops/scripts/go-checksum.sh soxie`" >> $FILE
echo -e "website:\n  tag: `./devops/scripts/js-checksum.sh website`" >> $FILE
echo -e "frontend:\n  tag: `./devops/scripts/js-checksum.sh frontend`" >> $FILE
echo -e "debugger:\n  tag: `./devops/scripts/debugger-checksum.sh`" >> $FILE
