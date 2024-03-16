package types

import (
	"github.com/vmihailenco/msgpack/v5"
)

type Platform int

var (
	VKMA Platform = 1
	VKID Platform = 2
	TG   Platform = 3
	OK   Platform = 4
)

type User struct {
	NID       int    `msgpack:"ni,omitempty"`
	TID       string `msgpack:"ti,omitempty"`
	FirstName string `msgpack:"fn,omitempty"`
	LastName  string `msgpack:"ln,omitempty"`
	UserName  string `msgpack:"un,omitempty"`
	PhotoURL  string `msgpack:"pu,omitempty"`
	Language  string `msgpack:"lc,omitempty"`
	IsPremium bool   `msgpack:"pr,omitempty"`
}

type EncodeParams struct {
	ID       uint64      `msgpack:"id"`
	IP       string      `msgpack:"ip"`
	Platform Platform    `msgpack:"pl"`
	Expires  int64       `msgpack:"ex"`
	User     interface{} `msgpack:"us"`
}

type DecodeParams struct {
	ID       uint64             `msgpack:"id"`
	IP       string             `msgpack:"ip"`
	Platform Platform           `msgpack:"pl"`
	Expires  int64              `msgpack:"ex"`
	User     msgpack.RawMessage `msgpack:"us"`
}

func (p *DecodeParams) ParseUser() (*User, error) {
	user := User{}
	err := msgpack.Unmarshal(p.User, &user)
	return &user, err
}
