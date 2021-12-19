#!/bin/bash
case $1 in
    start)
        echo "Starting Application"
        ./microservice-mvc-1 > app.log &
        ;;
    stop)
        echo "Stoping Application"
        sudo kill $(sudo lsof -t -i:9000)
        ;;
    restart)
        echo "Stoping Application"
        sudo kill $(sudo lsof -t -i:9000)
        echo "Starting Application"
        ./microservice-mvc-1 > app.log &
        ;;
    *)
        echo "Golang app service"
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0
