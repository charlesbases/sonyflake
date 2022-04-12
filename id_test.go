package sonyflake

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestNextID(t *testing.T) {
	for {
		select {
		case <-time.NewTimer(time.Second * 3).C:
			c := NextID().Decode()
			bs, _ := json.MarshalIndent(&c, "", "  ")
			fmt.Println(string(bs))
		}
	}
}
