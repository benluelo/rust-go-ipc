# rust-go-ipc
Example for communicating between rust and go through unix sockets.

## Instructions

Start the go server:

```sh
$ go run go/main.go
```

Start the rust client:

```sh
$ cd rust && cargo run
```

The rust client will send a request to the go server every 6 seconds.
