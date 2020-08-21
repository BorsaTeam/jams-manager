#!/bin/sh

export RELEASE_VERSION=$(echo "$CIRCLE_BRANCH"| cut -d '-' -f 2-)