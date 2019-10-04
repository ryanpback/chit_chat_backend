#!/bin/bash

RUN_BUILD=$1

# let's build the web server!
cd httpd

# this command will build a staticly linked binary for 64 bit osx system
# and place it in the dist folder
echo "Building OSX binary..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../dist/chitchat-osx
echo -e "done building! \n"

# Run the build
if [ "$RUN_BUILD" == "run" ]; then
    cd ../dist
    ./chitchat-osx
    killall chitchat-osx
fi

# return to root dir
cd ..
