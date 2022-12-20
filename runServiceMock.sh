#! /bin/bash -

PARAM="--port"
PORT="8080"

if [ $1 == $PARAM ]
then
    if [ $2 ] 
    then
        PORT=$2
    else
        echo "no port number provided. default port will be used."
    fi
else
    echo "invalid param, use --port to define custom port. default port will be used."
fi

./service-mock $PORT