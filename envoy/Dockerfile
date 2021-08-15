FROM envoyproxy/envoy-dev:7e6986f75b952aae4d9048b16ca90e6494ca3bf7

RUN apt-get update && apt-get install -y gettext-base

COPY ./envoy/src .

RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]