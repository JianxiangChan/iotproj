#!/bin/bash

echo "start svn update"
svn update 
echo "start svn update ok"

sleep 1

echo "start compiler"
cd /home/gopro/src/hug/cmd/hug
go build
tar czvf hug0.0.3_amd64.tar.gz  hug  appserver.conf 
cp hug0.0.3_amd64.tar.gz ../../

sleep 1
cd /home/gopro/src/hug/cmd/sug
go build
tar czvf sug0.0.3_amd64.tar.gz  sug  appserver.conf 
cp sug0.0.3_amd64.tar.gz ../../
echo "start compiler ok"
