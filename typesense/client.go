package typesense

import (
	"fmt"

	"github.com/v-byte-cpu/typesense-go/typesense/api"
)

type Client struct {
	apiClient   api.ClientWithResponsesInterface
	collections CollectionsInterface
}

// client.Collection('name').Document('124').Retrieve()

func (c *Client) Collections() CollectionsInterface {
	return c.collections
}

func (c *Client) Collection(collectionName string) CollectionInterface {
	return &collection{apiClient: c.apiClient, name: collectionName}
}

type httpError struct {
	status int
	body   []byte
}

func (e *httpError) Error() string {
	return fmt.Sprintf("status: %v response: %s", e.status, string(e.body))
}

type ClientOption func(*Client)

func WithAPIClient(apiClient api.ClientWithResponsesInterface) ClientOption {
	return func(c *Client) {
		c.apiClient = apiClient
	}
}

// TODO WithServer option (server string)
// TODO WithConnectionTimeout option (seconds int)
// TODO WithApiKey option (apiKey string)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{}
	// implement option pattern
	for _, opt := range opts {
		opt(c)
	}
	c.collections = &collections{c.apiClient}
	return c
}