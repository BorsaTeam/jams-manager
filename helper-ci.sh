#!/bin/sh

export RELEASE_VERSION=$(echo "$CIRCLE_BRANCH"| cut -d '-' -f 2-)
export HEROKU_API_KEY=$HEROKU_API_KEY
export HEROKU_APP_NAME=$HEROKU_APP_NAME