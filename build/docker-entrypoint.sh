#!/bin/sh
#set -e
#if [ "$ENV" = 'dev' ];then
#   . conf/set-env.dev.sh
#elif [ "$ENV" = 'pre' ]; then
#   . conf/set-env.pre.sh
#elif [ "$ENV" = 'prd' ]; then
#   . conf/set-env.prd.sh
#fi
exec ../apiserver
