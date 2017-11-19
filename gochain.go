package main

import (
	"fmt"
	"time"
)

func TimeString(t time.Time) string {
	ts, e := t.MarshalJSON()
	if e != nil {
		return "#InvalidTime"
	} else {
		return string(ts)
	}
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	Nonce        string
	PreviousHash string
}

func main() {
	b := &Block{1, time.Now(), []Transaction{}, "123", "345"}
	fmt.Println("vim-go", b)
}
