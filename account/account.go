package account

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/vimcoders/go-driver"
	"github.com/vimcoders/mongox-go-driver"
	"github.com/xdg-go/pbkdf2"
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
	UID  string `bson:"uid"`
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
	dk := pbkdf2.Key([]byte(channelId+passport), []byte("123456u"), 6, 32, sha256.New)
	e, err := connector.Execer(context.Background())
	if err != nil {
		return nil, err
	}
	defer e.Close(context.Background())
	newAccount := &Account{
		Id:   base64.StdEncoding.EncodeToString(dk),
		Ban:  time.Now().Unix(),
		Mute: time.Now().Unix(),
		UID:  uuid.New().String(),
	}
	if _, err := e.Insert(context.Background(), newAccount); err != nil {
		return nil, err
	}
	return nil, nil
}
