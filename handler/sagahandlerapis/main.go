package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	broker "github.com/awe76/saga/api/sagabrokerapis/v1"
	nats "github.com/nats-io/nats.go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	nc, err := nats.Connect("nats://processor:4222")
	if err != nil {
		return fmt.Errorf("did not connect to nats: %v", err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		return fmt.Errorf("did not connect to nats json channel: %v", err)
	}

	receiver := make(chan *broker.OperationPayload)
	ec.BindRecvChan("sop", receiver)

	sender := make(chan *broker.OperationPayload)
	ec.BindSendChan("rop", sender)

	for {
		select {
		case req := <-receiver:
			if req.IsRollback {
				fmt.Printf("%s operation rollback is started\n", req.Name)
			} else {
				fmt.Printf("%s operation is started\n", req.Name)
			}

			rand.Seed(time.Now().UnixNano())
			pause := rand.Intn(100)

			// sleep for some random time
			time.Sleep(time.Duration(pause) * time.Millisecond)

			payload, err := json.Marshal(rand.Float32())
			if err != nil {
				return err
			}

			req.IsFailed = !req.IsRollback && rand.Float32() > 0.8
			req.Payload = string(payload)

			sender <- req
		}
	}
}
