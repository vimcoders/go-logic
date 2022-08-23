package session

import "github.com/vimcoders/go-logic/account"

func (s *Session) Login(b []byte) (err error) {
	a, err := account.Login("", "")
	if err != nil {
		return err
	}
	s.Account = a
	return nil
}

func (s *Session) Register(b []byte) (err error) {
	_, err = account.Register("", "")
	if err != nil {
		return err
	}
	return nil
}
