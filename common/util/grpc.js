const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const _config = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

/**
 * Loads gRPC package definition
 * @param {string} path
 * @param {string} package
 * @param {Object} config
 */
const loadPackageDefinition = (path, package, config = _config) => {
  const protoDefinition = protoLoader.loadSync(path, config);
  return grpc.loadPackageDefinition(protoDefinition)[package];
};

module.exports = {
  loadPackageDefinition,
};
