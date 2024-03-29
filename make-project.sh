#!/bin/bash
while getopts b:r: flag
do
    case "${flag}" in
        b) build=${OPTARG};;
        r) run=${OPTARG};;
    esac
done

if [ $build = "true" ]
then
    docker image prune
    docker build -t go-proj-test .
fi

if [ $run = "true" ]
then
    docker run --rm -d -p 8000:8000 go-proj-test
    curl localhost:8000
    curl localhost:8000/about
fi
