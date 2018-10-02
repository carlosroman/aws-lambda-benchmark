#!/usr/bin/env sh
set -e
echo 'Starting Docker'

mkdir -p target/log

nohup /usr/bin/dockerd \
	--host=unix:///var/run/docker.sock \
	--host=tcp://127.0.0.1:2375 \
	--storage-driver=overlay > target/log/docker.log 2>&1 &

tries=0
d_timeout=60
until docker info >/dev/null 2>&1
do
	if [ "$tries" -gt "$d_timeout" ]; then
		cat target/log/docker.log
		echo 'Timed out trying to connect to internal docker host.' >&2
		exit 1
	fi
        tries=$(( $tries + 1 ))
	printf '.';
	sleep 1
done
