const path = require('path');
const { loadPackageDefinition } = require('../common/util/grpc');

const packageDefinition = loadPackageDefinition(path.join(__dirname, '..', 'common', 'protos', 'session.proto'), 'session');

const getSession = (call, callback) => { };

const createSession = (call, callback) => { };

const deleteSession = (call, callback) => { };

module.exports = {
  packageDefinition,
  service: packageDefinition.SessionService.service,
  rpcFunctions: { getSession, createSession, deleteSession },
};