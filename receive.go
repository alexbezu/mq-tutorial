//go:build receiver

package main

import (
	"fmt"

	// "github.com/codenotary/immudb/pkg/api/schema"
	schema "github.com/alexbezu/mq-tutorial/pb"
)

func main() {
	connect2db()
	defer closedb()

	for {
		rpl, err := c.MQpop(ctx, &schema.MQpopRequest{Qname: "queue_name1"})
		if err != nil {
			fmt.Errorf("Error: ", err)
		}
		if rpl != nil {
			fmt.Print(string(rpl.Value))
		}
	}
}
