#!/bin/bash

echo "start svn update"
svn update
echo "start svn update ok"

sleep 1

echo "start compiler"
go build
echo "start compiler ok"

sleep 1

echo "kill & cp & run"

sudo killall -9 monitorweb

tar czvf monitorweb0.0.1-`date +%Y%m%d%H%M%S`.tar.gz  monitorweb  appserver.conf 

sleep 1
cp -f monitorweb /root/monitorweb/
cp -f appserver.conf /root/monitorweb/


cd /root/monitorweb/
nohup /root/monitorweb/monitorweb  > monitorweb.log 2>&1 &
echo "kill & cp & run ok"
