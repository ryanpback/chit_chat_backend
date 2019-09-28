#!/bin/bash

RUN_BUILD=$1

echo $1

# let's build the web server!
cd httpd

# this command will build a staticly linked binary for 64 bit osx system
# and place it in the dist folder
echo "Building OSX binary..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../dist/chitchat-osx
echo -e "done building! \n"

if [ "$RUN_BUILD" == "run" ]; then
    # Run the build
    cd ../dist
    ./chitchat-osx
fi

# maybe in the future there will be other things that we need to build...
