package main

import (
	"encoding/json"
	"fmt"

	kafka "github.com/vinicius91/strimzi-client-go/pkg/apis/kafka.strimzi.io/v1beta2"
)

func main() {

	var k kafka.KafkaList
	err := json.Unmarshal([]byte(""), &k)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
