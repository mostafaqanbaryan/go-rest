#!/bin/bash

if [[ -z "$1" ]]; then
    echo "Please provide a name for migration!"
    exit 1
fi

nvim ./internal/database/migrations/$(date "+%Y%m%d%H%M%S")_$1.go
