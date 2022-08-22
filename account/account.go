package account

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vimcoders/go-driver"
	"github.com/vimcoders/mongox-go-driver"
)

var connector driver.Connector

func init() {
	c, err := mongox.Connect(&mongox.Config{
		Addr: "mongodb://127.0.0.1:27017",
		DB:   "account",
	})
	if err != nil {
		panic(err)
		return
	}
	connector = c
}

type Account struct {
	Id   string `bson:"_id"`
	Mute int64  `bson:"mute"`
	Ban  int64  `bson:"ban"`
}

func Login(channelId, passport string) (*Account, error) {
	e, err := connector.Execer(context.Background())
	if err != nil {
		return nil, err
	}
	accountL, err := e.Query(context.Background(), &Account{})
	if err != nil {
		return nil, err
	}
	for _, account := range accountL {
		if v, ok := account.(*Account); ok {
			return v, nil
		}
	}
	return nil, nil
}

func Register(channelId, passport string) (*Account, error) {
	e, err := connector.Execer(context.Background())
	if err != nil {
		return nil, err
	}
	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	newAccount := &Account{
		Id:   u.String(),
		Ban:  time.Now().Unix(),
		Mute: time.Now().Unix(),
	}
	if _, err := e.Insert(context.Background(), newAccount); err != nil {
		return nil, err
	}
	return nil, nil
}
