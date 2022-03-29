import * as jspb from 'google-protobuf'

import * as github_com_protocolbuffers_protobuf_src_google_protobuf_timestamp_pb from '../github.com/protocolbuffers/protobuf/src/google/protobuf/timestamp_pb';


export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class Request extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): Request;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Request.AsObject;
  static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
  static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Request;
  static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
  export type AsObject = {
    title: string,
  }
}

export class GithubTrendingInfo extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): GithubTrendingInfo;

  getTrendingdate(): github_com_protocolbuffers_protobuf_src_google_protobuf_timestamp_pb.Timestamp | undefined;
  setTrendingdate(value?: github_com_protocolbuffers_protobuf_src_google_protobuf_timestamp_pb.Timestamp): GithubTrendingInfo;
  hasTrendingdate(): boolean;
  clearTrendingdate(): GithubTrendingInfo;

  getGitreposList(): Array<GitRepo>;
  setGitreposList(value: Array<GitRepo>): GithubTrendingInfo;
  clearGitreposList(): GithubTrendingInfo;
  addGitrepos(value?: GitRepo, index?: number): GitRepo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GithubTrendingInfo.AsObject;
  static toObject(includeInstance: boolean, msg: GithubTrendingInfo): GithubTrendingInfo.AsObject;
  static serializeBinaryToWriter(message: GithubTrendingInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GithubTrendingInfo;
  static deserializeBinaryFromReader(message: GithubTrendingInfo, reader: jspb.BinaryReader): GithubTrendingInfo;
}

export namespace GithubTrendingInfo {
  export type AsObject = {
    title: string,
    trendingdate?: github_com_protocolbuffers_protobuf_src_google_protobuf_timestamp_pb.Timestamp.AsObject,
    gitreposList: Array<GitRepo.AsObject>,
  }
}

export class GitRepo extends jspb.Message {
  getId(): number;
  setId(value: number): GitRepo;

  getAuthor(): string;
  setAuthor(value: string): GitRepo;

  getName(): string;
  setName(value: string): GitRepo;

  getHref(): string;
  setHref(value: string): GitRepo;

  getDescription(): string;
  setDescription(value: string): GitRepo;

  getLanguage(): string;
  setLanguage(value: string): GitRepo;

  getStars(): number;
  setStars(value: number): GitRepo;

  getForks(): number;
  setForks(value: number): GitRepo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GitRepo.AsObject;
  static toObject(includeInstance: boolean, msg: GitRepo): GitRepo.AsObject;
  static serializeBinaryToWriter(message: GitRepo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GitRepo;
  static deserializeBinaryFromReader(message: GitRepo, reader: jspb.BinaryReader): GitRepo;
}

export namespace GitRepo {
  export type AsObject = {
    id: number,
    author: string,
    name: string,
    href: string,
    description: string,
    language: string,
    stars: number,
    forks: number,
  }
}

