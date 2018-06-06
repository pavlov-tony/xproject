// TODO: it will be good add description: Package gcpclient ...
// TODO: it good when packageName == packageFolderName
package gcpclient

import (
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

// TODO: way to hide ctx and client?
// TODO: It's not good send to pull request with TODO
// TODO: don't forget about golint, public methods must be commented
type Client struct {
	ctx           context.Context
	storageClient *storage.Client
}

// TODO: ALL
type Predictor interface {
	predict()
}

// TODO: constructor or singleton?
// TODO: ?
func (c Client) init() {
	c.ctx = context.Background()
	storageClient, err := storage.NewClient(c.ctx)
	if err != nil {
		log.Fatal("Client init error\n", err)
	}
	c.storageClient = storageClient
}
