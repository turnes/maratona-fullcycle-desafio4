package entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"simulator/entity"
	"testing"
)

func TestEntity(t *testing.T) {
	expected := entity.Order{
		Uuid:        "zxvf",
		Destination: "1",
	}
	msg := []byte(`{"order":"zxvf", "destination":"1"}`)
	var received entity.Order
	json.Unmarshal(msg, &received)
	assert.Equal(t, received, expected)

}
