# 🥝 Kiwi

Kiwi is a Key-Value storage server, completely written from scratch in Go as a personal pet project ands exercise to learn the Go programming language. 

It aims to be as an alternative for Redis, and to support at least a subset of the commands provided by Redis, for familarity reasons. 
As added features, support for strong namespacing would be a nice to have addition.

## Planned work for first demo:
- Supported Commands:
  - PING
  - GET
  - SET
  - KEY
- Namespacing
- TTL
  - Simple time based
  - Counter based
  - Duration based, with refresh on hit
- Simple Auth
- Defined communication protocol
  - Raw TCP-based, not unlike Redis.

## Name Origin

🥝 Kiwi in Portuguese reads as "Kivi", fonetically similar to Key-V. V in this case would stand for Value.
