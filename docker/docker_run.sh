#!/bin/bash

docker_container_name="FFAI_Interview"
docker_image_name="ffai_interview"

# build image
docker build -f docker/Dockerfile --tag $docker_image_name .

# run service from image
docker run --rm -it \
--name $docker_container_name $docker_image_name
