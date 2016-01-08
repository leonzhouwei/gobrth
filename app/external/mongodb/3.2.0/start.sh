#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# start 
./mongod -f ./etc/mongod.conf

exit 0
