#!/bin/bash
trap "exit" INT TERM
trap 'kill $(jobs -p)' EXIT
set -euo pipefail

BE_STARTUP_COUNT=0
FE_STARTUP_COUNT=0
STARTUP_WAIT=240

make backend-dev-mock &
until curl --output /dev/null --silent --fail http://localhost:8080/healthcheck; do
    if [ "$BE_STARTUP_COUNT" -ge "$STARTUP_WAIT" ]; then
        echo "Error: could not start backend mock server"
        exit 1
    fi;
    BE_STARTUP_COUNT=$((BE_STARTUP_COUNT+1))
    sleep 1
done

cd frontend || exit
yarn workspace @clutch-sh/app run start &
until curl --output /dev/null --silent --head --fail http://localhost:3000; do
    if [ "$FE_STARTUP_COUNT" -ge "$STARTUP_WAIT" ]; then
        echo "Error: could not start frontend dev server"
        exit 1
    fi;
    FE_STARTUP_COUNT=$((FE_STARTUP_COUNT+1))
    sleep 1
done

(
  yarn run test:e2e
)
