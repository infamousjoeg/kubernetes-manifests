#!/bin/bash

if [[ -z "$1" || -z "$2" ]]; then
    echo "Usage: $0 <namespace> <path-to-pem>"
    exit 1
fi

kubectl -n "$1" create configmap conjur-ssl-cert --from-file="$2"