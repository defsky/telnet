package telnet

import (
	"os"
	"testing"
)

func TestNVT(t *testing.T) {
	client, err := NewClient("mud.pkuxkx.com:8080", os.Stdout)
	if err != nil {
		t.Log(err)
	}
	defer client.Close()
	t.Log(client)
}
