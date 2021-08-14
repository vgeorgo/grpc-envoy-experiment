const grpc = require('grpc');

// Services
const session = require('./services/session');

const server = new grpc.Server();
const serverAddr = '127.0.0.1:50051';

// Add server services
server.addService(session.service, session.rpcFunctions);

// Server start
server.bind(serverAddr, grpc.ServerCredentials.createInsecure());
server.start();
console.log(`Server running on ${serverAddr}`);