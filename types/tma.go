package types

type TMAUser struct {
	ID                    int    `msgpack:"ni" json:"id"`
	FirstName             string `msgpack:"fn" json:"first_name"`
	LastName              string `msgpack:"ln" json:"last_name"`
	UserName              string `msgpack:"un" json:"username"`
	PhotoURL              string `msgpack:"pu" json:"photo_url"`
	Language              string `msgpack:"lc" json:"language_code"`
	ChatType              string `msgpack:"ct" json:"chat_type"`
	ChatInstance          string `msgpack:"ci" json:"chat_instance"`
	IsPremium             bool   `msgpack:"pr" json:"is_premium"`
	AllowsWriteToPM       bool   `msgpack:"aw" json:"allows_write_to_pm"`
	AddedToAttachmentMenu bool   `msgpack:"am" json:"added_to_attachment_menu"`
}
