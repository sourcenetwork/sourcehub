#!/usr/bin/bash

if [ ! -e "~/INITIALIZED" ]; then
    ignite chain init --skip-proto
    touch "~/INITIALIZED"
fi

exec /app/build/sourcehubd $@
