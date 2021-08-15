FROM node:14-alpine

WORKDIR /usr/src/app

COPY ./user_service/src .
COPY ./common ./common

RUN npm i -g grpc-tools --unsafe-perm
RUN npm i

EXPOSE 8081
CMD npm run server