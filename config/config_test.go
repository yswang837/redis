package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigByFileName(t *testing.T) {
	conf, err := NewConfigByFileName("redis")
	assert.Equal(t, nil, err)
	fmt.Println("conf........", conf)
}
