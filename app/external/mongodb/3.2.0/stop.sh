#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# stop
./mongod --shutdown -f ./etc/mongod.conf

exit 0
