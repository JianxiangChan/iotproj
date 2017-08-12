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

sudo killall -9 monitor

tar czvf monitor0.0.3-`date +%Y%m%d%H%M%S`.tar.gz  monitor  appserver.conf monitor.service

sleep 1
cp -f monitor /home/pi/monitor/
cp -f appserver.conf /home/pi/monitor/
sudo cp -f monitor.service /lib/systemd/system/
sudo chmod 755 /lib/systemd/system/monitor.service
sudo systemctl enable monitor.service

cd /home/pi/monitor/
nohup /home/pi/monitor/monitor > monitor.log 2>&1 &
echo "kill & cp & run ok"
