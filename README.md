# 🥝 Kiwi

Kiwi is a Key-Value storage server, written from scratch in Go as a personal pet project ands exercise to learn the Go programming language. 

It aims to be as an alternative for Redis, and to support at least a subset of the commands provided by Redis, for familarity reasons. 
As added features, support for strong namespacing would be a nice to have addition.

## Planned work for first demo:
- Supported Commands:
  - [x] PING
  - [x] GET
  - [x] SET
  - [x] DEL
  - [ ] KEYS
- Features
  - [ ] Invalidation
    - [ ] TTL
  - [ ] Namespacing 
  - [ ] Distributed Locks
  - [ ] Simple Auth
    - [ ] Namespacing ACL
- Networking:
  - Multi-protocol support
    - [ ] Connection Handshake and protocol negotiation
    - Supported Protocols
        - [x] KiwiV0: rudimentary string marshalling for testing purposes
        - [ ] KiwiV1: Custom binary serialization or Protobuf
        - [ ] Redis RESPv3-compatible protocol

## Name Origin

🥝 Kiwi in Portuguese reads as "Kivi", phonetically similar to Key-V. V in this case would stand for Value.
