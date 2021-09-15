# Brief intro

This repo is based on the [filebrowser](https://github.com/filebrowser/filebrowser) and add subset of commands which using gRPC protocol which list, send and receive files between server and client.

# Features

- gRPC client support list, send and receive files between server.
- bcrypt encrypted md5 hash of the password when login, so that no real password is sending accross the network or saving to the database.
# Known issues:

- Unable to transfer file which is larger than 1GB, it will cause a deadline exception. I might not fully test the actual size can be transfer, just test the transmittional size through localhost.

# Roadmap

- Add android client using gRPC, which supports private data (photos, documents, etc.) sync
