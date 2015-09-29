# Zmqforwarder
Simple ZeroMQ log forwarder written in go

Base on https://gist.github.com/dstrelau/9824230

## Dependencies

Zmqforwarder is base on `zeromq 4`, `pebbe/zmq4` and `ActiveState/tail`.

## Build zmqforwarder

### Install zeromq from sources

```bash
wget http://download.zeromq.org/zeromq-4.1.3.tar.gz
tar -xzf zeromq-4.1.3.tar.gz
cd zeromq-4.1.3
./configure --prefix=<INSTALL_PREFIX> --enable-static --disable-shared
make install
```

### Install go dependencies

```bash
go get github.com/pebbe/zmq4
go get github.com/ActiveState/tail
```

### Compilation

To build a static binary:

```bash
export CGO_CFLAGS="-I<INSTALL_PREFIX>/include"
export CGO_LDFLAGS="-L/<INSTALL_PREFIX>/lib"
go build -ldflags="-linkmode external -extldflags '-lstdc++ -lpgm -lpthread -static'" zmqforwarder.go
```
