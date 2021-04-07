#!/bin/bash

kill -9 `cat ./cmd.pid`

rm cmd.pid
rm output