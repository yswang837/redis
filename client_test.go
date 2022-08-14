package redis

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}
	_, _ = c.SET("name", "AAAAAA")
	c.EXPIRE("name", "AAAAAAAAAB", "10")
	a, _ := c.EXISTS("name")
	fmt.Println("BBBBB", a)
	s, err := c.GET("name")
	if err != nil {
		panic(err)
	}
	fmt.Println("result.....", s)

}
