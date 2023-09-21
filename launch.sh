#!/bin/bash
set -e

trap 'killall distribkv' SIGINT

cd $(dirname $0)

killall distribkv || true
sleep 0.1

go install -v

gokeydb -db-location=mumbai.db -http-addr=127.0.0.1:8080 -config-file=sharding.toml -shard=Mumbai &
gokeydb -db-location=pune.db -http-addr=127.0.0.1:8081 -config-file=sharding.toml -shard=Pune &
gokeydb -db-location=delhi.db -http-addr=127.0.0.1:8082 -config-file=sharding.toml -shard=Delhi &

wait
