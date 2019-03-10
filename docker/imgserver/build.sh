#!/bin/sh

# For the pwd to pick up the correct bind path we will CD
# This whole thing can be changed to be more re-usable if
# we bind a director not in this repo, but this is good enough
# for now :) 
# cd ..
#
# We can also run this from the root directory of the repo
# which is what we'll do :) 

MOUNT_SOURCE=$(pwd)/data/
MOUNT_TARGET=/data
MOUNT_TYPE=bind

docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

docker build -t nginx-image-server .

docker run \
  --name nis1 \
  -p 8080:80 \
  --mount source=$MOUNT_SOURCE,target=$MOUNT_TARGET,type=$MOUNT_TYPE \
  -d nginx-image-server 

sleep 1

curl http://localhost:8080/
