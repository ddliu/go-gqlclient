# go-gqlclient

A simple Graphql client for Golang.

## Install

```
go get -u github.com/ddliu/go-gqlclient
```

## Usage

```go
import (
    "github.com/ddliu/go-gqlclient"
    "net/http"
    "context"
)

client = gqlclient.New(gqlclient.Options{
    Endpoint: "http://server/graphql",
    Header: http.Header{
        "some-secret": []string{"secret string"},
    },
})

data, err := client.Query(context.Background(), `
query MyQuery {
    post {
        id
        name
    }
}
`, nil)

print(data.String("post.0.name"))

// Unmarshal data
var posts []Post
data.Unmarshal(&posts)
```

With variables:

```go
client.Query(context.Background(), `
query MyQuery($limit: Int) {
    post(limit: $limit) {
        id
        name
    }
}
`, map[string]interface{} {
    "limit": 10,
})
```

Check out [fractal](https://github.com/ddliu/fractal) for more information of the query result.