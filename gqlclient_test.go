package gqlclient

import (
	"context"
	"net/http"
	"testing"
)

func TestAll(t *testing.T) {
	c := New(Options{
		Endpoint: "http://yourgraphqlendpoint/graphql",
		Header: http.Header{
			"some-secret": []string{"secret string"},
		},
	})

	data, err := c.Query(context.Background(), `
		query MyQuery {
			post {
				id
			}
		}
	`, nil)

	if err != nil {
		t.Error(err)
	}

	if !data.Exist("post.0.id") {
		t.Error()
	}
}
