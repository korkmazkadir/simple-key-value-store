#!/bin/bash


./stop-key-value-store.sh

nohup ./cmd > output 2>&1 &

RETVAL=$?
PID=$!

[ $RETVAL -eq 0 ] && echo $PID > cmd.pid
exit $RETVAL
