# 🥝 Kiwi

Kiwi is a Key-Value storage server, written from scratch in Go as a personal pet project ands exercise to learn the Go programming language. 

It aims to be as an alternative for Redis, and to support at least a subset of the commands provided by Redis, for familarity reasons. 
As added features, support for strong namespacing would be a nice to have addition.

## PoC milestones
- Supported Commands:
  - [x] `PING`: returns `PONG`
  - [x] `GET <key>`: retrieves the string value of a key
  - [x] `SET <key> <value>`: sets the string value of a key
  - [x] `DEL <key>`: deletes a key
  - [X] `KEYS`: returns list of all available keys
 
- Supported encodings:
   - [x] Redis RESP Protocol

## Name Origin

🥝 In Portuguese, Kiwi reads as "Quivi". Phonetically, it is pronounced as "Key-V", which is a play-on-words as short for Key-Value.
