#!/bin/bash

. setenv.sh
cd src
go build -o ../bin/auth
cd ..
./bin/auth
