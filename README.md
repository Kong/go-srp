# venafi-go-srp 

[![Venafi](https://raw.githubusercontent.com/Venafi/.github/master/images/Venafi_logo.png)](https://www.venafi.com/)

[![](https://api.travis-ci.org/kong/go-srp.svg)](https://travis-ci.org/Venafi/venafi-go-srp)
[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/Venafi/venafi-go-srp)

_NOTE: This is a port of [node-srp](https://github.com/mozilla/node-srp) to Go. I recommend
reading their README for general information about the use of SRP._

This package was inspired by kong-srp

## Installation

```
go get github.com/Venafi/venafi-go-srp
```

## Usage

View [GoDoc](https://godoc.org/github.com/kong/go-srp) for full details

To use SRP, first decide on they parameters you will use. Both client and server must
use the same set.

```go
params := srp.GetParams(2048)
```

### Account Creation

To create a new account, generate a verifier from the client, and store it
on the server.

```go
verifier := srp.ComputeVerifier(params, salt, identity, password)
```

### Login

From the client... generate a new secret key, initialize the client, and compute A.
Once you have A, you can send A to the server.

```go
secret1 := srp.GenKey()
client := NewClient(params, salt, identity, secret, a)
srpA := client.computeA()

sendToServer(srpA)
```

From the server... generate another secret key, initialize the server, and compute B.
Once you have B, you can send B to the client.

```go
secret2 := srp.GenKey()
server := NewServer(params, verifier, secret2)
srpB := client.computeB()

sendToClient(srpB)
```

Once the client received B from the server, it can compute M1 based on A and B.
Once you have M1, send M1 to the server.

```go
client.setB(srpB)
srpM1 := client.ComputeM1()
sendM1ToServer(srpM1)
```

Once the server receives M1, it can verify that it is correct. If checkM1() returns
an error, authentication failed. If it succeeds it should be sent to the client.

```go
srpM2, err := server.checkM1(srpM1)
```

Once the client receives M2, it can verify that it is correct, and know that authentication
was successful.

```go
err = client.CheckM2(serverM2)
````

Now that both client and server have completed a successful authentication, they can
both compute K independently. K can now be used as either a key to encrypt communication
or as a session ID.

```go
clientK := client.ComputeK()
serverK := server.ComputeK()
```

## Running Tests

```
go test
```

_Tests include vectors from 
[RFC 5054, Appendix B.](https://tools.ietf.org/html/rfc5054#appendix-B)_


## Licence

MIT
