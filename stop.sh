#!/bin/bash 

ps aux | grep fileService | grep -v grep | awk '{print $2}' | xargs kill -9

echo "fileService stoped!"
