#!/bin/bash

set -e

cmd="$@"

# MySQLサーバーが起動するまでmain.goを実行せずにループで待機する
until mysqladmin ping -h mysql --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting"
exec $cmd