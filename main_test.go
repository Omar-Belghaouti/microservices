package main

import (
	"fmt"
	"testing"

	"github.com/Omar-Belghaouti/microservices/client/client"
	"github.com/Omar-Belghaouti/microservices/client/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewGetProductsParams()
	prods, err := c.Products.GetProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prods.GetPayload()[0])
}
