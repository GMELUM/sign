package types

type TGUser struct {
	ID        int    `msgpack:"id" json:"id"`
	FirstName string `msgpack:"fn" json:"first_name"`
	LastName  string `msgpack:"ln" json:"last_name"`
	UserName  string `msgpack:"un" json:"username"`
	PhotoURL  string `msgpack:"pu" json:"photo_url"`
}
