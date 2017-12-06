#!/bin/sh

echo "------------build all proto--------------"

#tls
sh build-tls.sh

#token
sh build-token.sh

#interceptor
sh build-interceptor.sh

#normal
sh build-normal.sh

#gateway
sh build-gateway.sh
