#!/bin/bash
#
# Build the snapweb snap using a docker container.

set -ev

docker run -v $GOPATH:/go dbarth/snapweb-builder sh -c 'cd /go/src/github.com/snapcore/snapweb && export GOPATH=/go PATH=/go/bin:$PATH && ./scripts/snap.sh'

#change the permissions to user travis
echo "Fixing the permissions of files generated by the docker daemon"
if [ ! -z $TRAVIS ]; then
  sudo chown travis:travis *.snap
  sudo chown -R travis:travis node_modules
else
  sudo chown $USER:$USER *.snap
  sudo chown -R $USER:$USER node_modules
fi
