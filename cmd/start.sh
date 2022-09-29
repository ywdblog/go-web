#!/bin/bash

SERVER_NAME="article-server"

PROJECT_PATH=$(cd $(dirname $0); cd ../; pwd)
LOGS_PATH=$PROJECT_PATH/logs/
CMD_PATH=$PROJECT_PATH/cmd/

CMD_BIN=$CMD_PATH$SERVER_NAME
STD_FILE=$LOGS_PATH${SERVER_NAME}-stderr.log
ENVIREMENT="development testing production"
ACTION="start stop restart"

stop() {
    pids=$(ps aux | grep "$SERVER_NAME" | grep "$env" | grep -v grep | awk '{print $2}')
    for pid in $pids; do
        echo -n "Kill $pid "
        kill $pid
        while true; do
            pstr=$(ps -p $pid | wc -l)
            if [ $pstr -eq 1 ]; then
                break
            fi
            echo -n "."
            sleep 1
        done
        echo "done."
    done
}

start() {
    echo "Start $CMD_BIN -env $env"
    nohup $CMD_BIN -env $env >> $STD_FILE 2>&1 &
}

restart() {
    stop
    start
}

usage() {
    
    env=${ENVIREMENT// /|}
    action=${ACTION// /|}

    echo "Usage: ./start.sh -e [$env] -k [$action]"
}

checkEnv() {
    for e in $ENVIREMENT; do
        if [ "$e" == "$1" ]; then
            return 0
        fi
    done
    echo "Error: enviremnent is invalid."
    usage
    exit 1
}

checkAction() {
    for e in $ACTION; do
        if [ "$e" == "$1" ]; then
            return 0
        fi
    done
    echo "Error: action is invalid."
    usage
    exit 1
}

env=""
action=""
cd $PROJECT_PATH

while getopts "hk:e:" opt; do
  case $opt in
    h)
      usage
      exit
      ;;
    k)
      action=$OPTARG
      ;;
    e)
      env=$OPTARG 
      ;;
  esac
done

checkEnv $env
checkAction $action
$action




