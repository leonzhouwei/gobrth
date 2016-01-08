#!/bin/bash

# safe cd
base_path=$(cd `dirname $0`; pwd)
cd $base_path

# init
mkdir -p ./var/log
mkdir -p ./var/run

exit 0
