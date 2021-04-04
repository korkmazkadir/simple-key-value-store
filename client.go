package simplekeyvalue

import "net/rpc"

// Client implements a keyvalue store client
type Client struct {
	rpcClient *rpc.Client
}

// NewClient creates a nee keyvalue store client
func NewClient() (*Client, error) {

	rpcClient, err := rpc.Dial("unix", "/tmp/keyvalue-store.sock")
	if err != nil {
		return nil, err
	}

	c := &Client{}
	c.rpcClient = rpcClient

	return c, nil
}

// Put puts a key value pair to the store if the key doesnot exists
func (c *Client) Put(key string, value []byte) bool {

	kvp := KeyValuePair{Key: key, Value: value}

	var reply bool
	err := c.rpcClient.Call("Server.Put", kvp, &reply)

	if err != nil {
		panic(err)
	}

	return reply
}

// Get returns the value of the key
func (c *Client) Get(key string) []byte {

	var reply []byte
	err := c.rpcClient.Call("Server.Get", key, &reply)

	if err != nil {
		panic(err)
	}

	return reply
}

// Exists returns true if an object with a given name is available in the store
func (c *Client) Exists(key string) bool {

	var reply bool
	err := c.rpcClient.Call("Server.Exists", key, &reply)

	if err != nil {
		panic(err)
	}

	return reply
}
