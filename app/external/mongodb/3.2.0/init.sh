#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# init
mkdir -p ./var/data/db
mkdir -p ./var/log/mongodb

exit 0

