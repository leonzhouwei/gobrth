#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# start 
./redis-server ./conf/redis.conf

exit 0
