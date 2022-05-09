#!/bin/sh
set -e

evans ../common/protos/session.proto --host $SESSION_SERVICE_ADDRESS --port $SESSION_SERVICE_PORT