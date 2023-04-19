package batcher

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const (
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second * 5
)

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

var bat *Batcher

func TestBatcher(t *testing.T) {
	b := New(
		WithBuffer(batcherBuffer),
		WithWorker(batcherWorker),
		WithInterval(batcherInterval),
	)

	b.Sharding = func(key string) int {
		ikey, _ := strconv.Atoi(key)
		return ikey
	}
	b.Do = func(ctx context.Context, val map[string][]any) {
		msgs := map[int64][]*KafkaData{}
		for _, vs := range val {
			for _, v := range vs {
				if v != nil {
					if kd := v.(*KafkaData); kd != nil {
						msgs[kd.Uid] = append(msgs[kd.Uid], kd)
					}
				}
			}
		}

		for k, v := range msgs {
			fmt.Printf(" ( K:%d => V:%v ) ", k, len(v))
		}
		fmt.Printf("\n\n")
	}

	b.Start()
	bat = b

	rand.Seed(time.Now().UnixNano())

	function()

	select {}
}

func function() {
	run()

	time.AfterFunc(1*time.Second, function)
}

func run() {
	// d := rand.Intn(200)
	for i := 0; i < 100; i++ {
		err := bat.Add(fmt.Sprintf("%d", i), &KafkaData{Uid: int64(i), Pid: int64(i)})
		if err != nil {
			// -- 队列已经放满 -- 直接写入数据库
			fmt.Println(err)
		}
	}
}
