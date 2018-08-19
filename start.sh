#!/bin/bash

#stop the service
ps -ef | grep fileService | grep -v grep| awk '{print $2}' | xargs kill -9

echo "fileService starting.."
nohup ./fileService &
echo "fileService started finished!"
