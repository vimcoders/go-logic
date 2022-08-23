package account

import (
	"fmt"
	"testing"

	"github.com/vimcoders/mongox-go-driver"
)

func TestMain(m *testing.M) {
	c, err := mongox.Connect(&mongox.Config{
		Addr: "mongodb://127.0.0.1:27017",
		DB:   "account",
	})
	if err != nil {
		panic(err)
		return
	}
	connector = c
	m.Run()
}
func TestRegister(t *testing.T) {
	for i := 0; i < 100000000; i++ {
		if _, err := Register("Google", fmt.Sprintf("%v", i)); err != nil {
			t.Error(err)
		}
	}
}

func TestLogin(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		t.Log(Login("Google", fmt.Sprintf("%v", i)))
	}
}
