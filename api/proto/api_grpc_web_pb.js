/**
 * @fileoverview gRPC-Web generated client stub for camping
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.camping = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.camping.CampingServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.camping.CampingServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.camping.Request,
 *   !proto.camping.Response>}
 */
const methodInfo_CampingService_Do = new grpc.web.AbstractClientBase.MethodInfo(
  proto.camping.Response,
  /** @param {!proto.camping.Request} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.camping.Response.deserializeBinary
);


/**
 * @param {!proto.camping.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.camping.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.camping.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.camping.CampingServiceClient.prototype.do =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/camping.CampingService/Do',
      request,
      metadata || {},
      methodInfo_CampingService_Do,
      callback);
};


/**
 * @param {!proto.camping.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.camping.Response>}
 *     A native promise that resolves to the response
 */
proto.camping.CampingServicePromiseClient.prototype.do =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/camping.CampingService/Do',
      request,
      metadata || {},
      methodInfo_CampingService_Do);
};


module.exports = proto.camping;

