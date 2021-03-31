package otpravka

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func makeClient() *Client {
	return NewClient(nil, "", "")
}
func TestOrderService_Search(t *testing.T) {
	client := makeClient()
	r, _, err := client.Order.Search(context.Background(), "1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(r)
}
