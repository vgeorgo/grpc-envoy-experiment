FROM node:14-alpine

WORKDIR /usr/src/app
COPY ./session_service/src .
COPY protos ./protos
COPY util ./util

RUN npm i -g grpc-tools --unsafe-perm
RUN npm i

EXPOSE 50051
CMD npm run server