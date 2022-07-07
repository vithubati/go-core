# go-core

This repository contains implementation of some core go libs

## How to use

#### installing

```go
    go get github.com/vithubati/go-core
```

## Usage

### Authenticator

1. Add core module to import statement.

```go
    import "github.com/vithubati/go -core/<package>
```

2. Creat an authenticator for the client

```go
    authenticator, err := NewBasicAuthenticator ( "username", "password")
```

### Http Client
1.  Creating a client
```go
    auth, _ := authenticator.NewBearerTokenAuthenticator("Test-TOKEN")
    client := New(WithAuthenticator(auth))
```

2.  Creating a request
```go
    auth, _ := authenticator.NewBearerTokenAuthenticator("Test-TOKEN")
    client := New(WithAuthenticator(auth))
    params := url.Values{}
    params.Set("limit", "5")
    params.Set("size", "10")
    ctx := context.Background()
    url := "https://gobyexample.com/"
    req := client.RequestWithCtx(ctx, url, http.MethodGet, nil, params)
```

3. Executing a request
```go
    resp, err := client.Execute(req)
```