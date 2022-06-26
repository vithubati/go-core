# go-core

This repository contains implementation of some core go libs

## How to use

#### install1ing

```go
    go get github.com/vithubati/go -core
```

## Usage

### Code

1. Add core module to import statement.

```go
    import "github.com/vithubati/go -core/<package>
```

2. Creat an authenticator for the client

```go
    authenticator, err := NewBasicAuthenticator ( "username", "password")
```
