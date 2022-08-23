package session

import (
	"context"
	"net"

	"github.com/vimcoders/go-driver"
	"github.com/vimcoders/go-logic/account"
)

func init() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				continue
			}
			NewSession(c)
		}
	}()
}

type Session struct {
	driver.Conn
	*account.Account
}

func NewSession(c net.Conn) *Session {
	s := &Session{
		Conn: driver.Conn{
			Conn: c,
			C:    make(chan []byte, 1),
		},
	}
	s.OnMessage = func(b []byte) error {
		return nil
	}
	return s
}

func (s *Session) Pull(ctx context.Context) (err error) {
	defer func() {
		//	err = errors.New(string(debug.Stack()))
	}()
	if err := s.Conn.Pull(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Push(ctx context.Context) (err error) {
	defer func() {
		// err = errors.New(string(debug.Stack()))
	}()
	if err := s.Conn.Push(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Close(ctx context.Context) (err error) {
	if err := s.Conn.Close(); err != nil {
		return err
	}
	return nil
}
