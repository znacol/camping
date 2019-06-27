/**
 * @fileoverview gRPC-Web generated client stub for camping
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  Request,
  Response} from './api_pb';

export class CampingServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoDo = new grpcWeb.AbstractClientBase.MethodInfo(
    Response,
    (request: Request) => {
      return request.serializeBinary();
    },
    Response.deserializeBinary
  );

  do(
    request: Request,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Response) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/camping.CampingService/Do',
      request,
      metadata || {},
      this.methodInfoDo,
      callback);
  }
}

