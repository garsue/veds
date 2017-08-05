#!/bin/sh

docker container run -d --rm -h gdatastore -p 8432:8432 --name=dse google/cloud-sdk gcloud beta emulators datastore start --project=$1 --host-port gdatastore:8432
