#! /bin/bash

# update code
git checkout .
git pull origin master

# stop fs
supervisorctl stop fs

# get listen port
pid=`lsof -i:8899 | grep LISTEN | awk -F '[ ]+' '{print $2}'`

# kill port listen
kill ${pid}

# start fs
supervisorctl start fs

# re grant authorization
chmod -R 755 reload.sh


