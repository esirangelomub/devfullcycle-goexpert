# Chapter 14: gRPC with Go

In this module, we will delve into building gRPC services using Go. gRPC is a modern open-source framework that leverages Protocol Buffers (proto3) to define service contracts and produce efficient, platform-neutral, and language-agnostic RPC frameworks.

## Table of Contents
1. [Quick Start with gRPC in Go](#quick-start-with-grpc-in-go)
2. [Introduction to Protocol Buffers (proto3)](#introduction-to-protocol-buffers-proto3)
3. [Generating Code with protoc](#generating-code-with-protoc)

## Quick Start with gRPC in Go

For those new to gRPC with Go, you can kick off your journey by following the official [Quick Start guide](https://grpc.io/docs/languages/go/quickstart/). This guide walks you through the steps involved in setting up a basic gRPC service using Go, giving you an understanding of the basic components and workflow.

## Introduction to Protocol Buffers (proto3)

Protocol Buffers (often referred to as protobufs) are a flexible and efficient mechanism for serializing structured data. The version we'll be focusing on is proto3, which is the third iteration of the Protocol Buffer language.

You can deep dive into the various aspects of proto3, including its syntax and features, by referring to the [official programming guide](https://protobuf.dev/programming-guides/proto3/).

## Generating Code with protoc

Once you've defined your services and messages in a `.proto` file, you need to generate the corresponding Go code. This is achieved using the `protoc` compiler.

For our use case, given the `course_category.proto` file, you can generate the necessary Go code for both gRPC and the protobuf messages by running the following command:

```shell
protoc --go_out=. --go-grpc_out=. proto/category.proto
```

