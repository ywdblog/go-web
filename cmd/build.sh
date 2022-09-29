#!/bin/bash

SERVER_NAME="article-server"

PROJECT_PATH=$(cd $(dirname $0); cd ../; pwd)
LOGS_PATH=$PROJECT_PATH/logs/
CMD_PATH=$PROJECT_PATH/cmd/

CMD_FILE=$CMD_PATH/server.go
CMD_BIN=$CMD_PATH/$SERVER_NAME


# go build envirement
echo "###############################################################"
echo "# PROJECT_PATH: " $PROJECT_PATH
echo "#"
echo "# CMD_PATH:     " $CMD_PATH
echo "# LOGS_PATH:    " $LOGS_PATH
echo "#"
echo "# CMD_FILE:     " $CMD_FILE
echo "# CMD_BIN:      " $CMD_BIN
echo "#"

GIT_BRANCH=$(git branch -v --no-color | grep "*" | awk '{print $2}')
GIT_VERSION=$(git branch -v --no-color | grep "*" | awk '{print $3}')
echo "# GIT_BRANCH:   " $GIT_BRANCH
echo "# GIT_VERSION:  " $GIT_VERSION
echo "#"

GO_VERSION=$(go version)
echo "# GO_VERSION:   " $GO_VERSION


echo "###############################################################"
echo ""

cd $PROJECT_PATH
go build -v -x -o $CMD_BIN $CMD_FILE
e=$?

if (($e != 0)); then

    echo "Build error."
    exit $e
else 
    echo "Build successfully."
    chmod a+x $CMD_BIN
fi




