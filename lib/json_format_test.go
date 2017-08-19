package lib

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	in current dir

	cmd:
	go test -v -bench=. json_format_test.go  -benchmem
*/

type Data struct {
	ReqId   string  `json:"req_id"`
	ItemIds []int64 `json:"item_ids"`
}

var (
	reqId   string  = "20170818142437010008060105749EBF"
	itemIds []int64 = []int64{1, 2, 3, 4, 5, 6}
)

func structFormat() error {
	dat := &Data{
		ReqId:   reqId,
		ItemIds: itemIds,
	}

	_, err := json.Marshal(dat)
	return err
}

func mapFormat() error {
	dat := map[string]interface{}{
		"req_id":   reqId,
		"item_ids": itemIds,
	}
	_, err := json.Marshal(dat)
	return err
}

func TestStructFormat(t *testing.T) {
	err := structFormat()
	assert.Equal(t, nil, err)
}

func TestMapFormat(t *testing.T) {
	err := mapFormat()
	assert.Equal(t, nil, err)
}

func BenchmarkStructFormat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		structFormat()
	}
}

func BenchmarkMapFormat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		mapFormat()
	}
}
