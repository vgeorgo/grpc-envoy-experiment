const grpc = require('grpc');

// Services
const session = require('./services/session');

const server = new grpc.Server();
const serverAddr = `0.0.0.0:${process.env.SESSION_SERVICE_PORT}`;

// Add server services
server.addService(session.service, session.rpcFunctions);

// Server start
server.bind(serverAddr, grpc.ServerCredentials.createInsecure());
server.start();
console.log(`Server running on ${serverAddr}`);
