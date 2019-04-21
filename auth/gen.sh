#!/usr/bin/env bash

DSIGEN_PATH=${PWD#*src/}/design
OUTPUT_PATH="src"
goagen bootstrap -d $DSIGEN_PATH --force -o $OUTPUT_PATH
# goagen main -d $DSIGEN_PATH --force
