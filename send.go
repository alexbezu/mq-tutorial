//go:build sender

package main

import (
	"fmt"

	// "github.com/codenotary/immudb/pkg/api/schema"
	schema "github.com/alexbezu/mq-tutorial/pb"
)

func main() {
	connect2db()
	defer closedb()

	var put schema.MQputRequest
	put.Qname = "queue_name1"
	put.Value = append(put.Value, []byte("Hello\n")...)
	_, err := c.MQput(ctx, &put)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
}
