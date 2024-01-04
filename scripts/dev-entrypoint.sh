#!/usr/bin/bash

ignite chain build --skip-proto
exec sourcehubd $@
