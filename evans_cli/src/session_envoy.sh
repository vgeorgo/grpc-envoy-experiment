#!/bin/sh
set -e

evans ../common/protos/session.proto --host $ENVOY_ADDRESS --port $ENVOY_PORT