const path = require('path');
const { loadPackageDefinition } = require('../common/util/grpc');

const packageDefinition = loadPackageDefinition(
  path.join(__dirname, '..', 'common', 'protos', 'session.proto'),
  'session',
);

const check = (call, callback) => {
  callback(null, {
    id: `user_session_token`,
  });
};

const create = (call, callback) => {
  callback(null, {
    id: `user_session_token`,
  });
};

const remove = (call, callback) => {
  callback(null, {
    id: `user_session_token`,
  });
};

module.exports = {
  packageDefinition,
  service: packageDefinition.Session.service,
  rpcFunctions: { check, create, remove },
};
