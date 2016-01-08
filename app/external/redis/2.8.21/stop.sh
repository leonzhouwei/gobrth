#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# stop
kill `cat ./var/run/redis.pid`

exit 0
