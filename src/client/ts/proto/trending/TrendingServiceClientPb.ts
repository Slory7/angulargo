/**
 * @fileoverview gRPC-Web generated client stub for trending
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as trending_trending_pb from '../trending/trending_pb';


export class TrendingClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetGithubTrending = new grpcWeb.MethodDescriptor(
    '/trending.Trending/GetGithubTrending',
    grpcWeb.MethodType.UNARY,
    trending_trending_pb.Request,
    trending_trending_pb.GithubTrendingInfo,
    (request: trending_trending_pb.Request) => {
      return request.serializeBinary();
    },
    trending_trending_pb.GithubTrendingInfo.deserializeBinary
  );

  getGithubTrending(
    request: trending_trending_pb.Request,
    metadata: grpcWeb.Metadata | null): Promise<trending_trending_pb.GithubTrendingInfo>;

  getGithubTrending(
    request: trending_trending_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void): grpcWeb.ClientReadableStream<trending_trending_pb.GithubTrendingInfo>;

  getGithubTrending(
    request: trending_trending_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/trending.Trending/GetGithubTrending',
        request,
        metadata || {},
        this.methodDescriptorGetGithubTrending,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/trending.Trending/GetGithubTrending',
    request,
    metadata || {},
    this.methodDescriptorGetGithubTrending);
  }

  methodDescriptorFetchGithubTrending = new grpcWeb.MethodDescriptor(
    '/trending.Trending/FetchGithubTrending',
    grpcWeb.MethodType.UNARY,
    trending_trending_pb.Empty,
    trending_trending_pb.GithubTrendingInfo,
    (request: trending_trending_pb.Empty) => {
      return request.serializeBinary();
    },
    trending_trending_pb.GithubTrendingInfo.deserializeBinary
  );

  fetchGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<trending_trending_pb.GithubTrendingInfo>;

  fetchGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void): grpcWeb.ClientReadableStream<trending_trending_pb.GithubTrendingInfo>;

  fetchGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/trending.Trending/FetchGithubTrending',
        request,
        metadata || {},
        this.methodDescriptorFetchGithubTrending,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/trending.Trending/FetchGithubTrending',
    request,
    metadata || {},
    this.methodDescriptorFetchGithubTrending);
  }

  methodDescriptorGetAndSaveGithubTrending = new grpcWeb.MethodDescriptor(
    '/trending.Trending/GetAndSaveGithubTrending',
    grpcWeb.MethodType.UNARY,
    trending_trending_pb.Empty,
    trending_trending_pb.GithubTrendingInfo,
    (request: trending_trending_pb.Empty) => {
      return request.serializeBinary();
    },
    trending_trending_pb.GithubTrendingInfo.deserializeBinary
  );

  getAndSaveGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<trending_trending_pb.GithubTrendingInfo>;

  getAndSaveGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void): grpcWeb.ClientReadableStream<trending_trending_pb.GithubTrendingInfo>;

  getAndSaveGithubTrending(
    request: trending_trending_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: trending_trending_pb.GithubTrendingInfo) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/trending.Trending/GetAndSaveGithubTrending',
        request,
        metadata || {},
        this.methodDescriptorGetAndSaveGithubTrending,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/trending.Trending/GetAndSaveGithubTrending',
    request,
    metadata || {},
    this.methodDescriptorGetAndSaveGithubTrending);
  }

}

