package models

type Links struct {
	Twitter   *string `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Instagram *string `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Facebook  *string `bson:"facebook,omitempty" json:"facebook,omitempty"`
	TikTok    *string `bson:"tiktok,omitempty" json:"tiktok,omitempty"`
	Linkedin  *string `bson:"linkedin,omitempty" json:"linkedin,omitempty"`
	Telegram  *string `bson:"telegram,omitempty" json:"telegram,omitempty"`
	Discord   *string `bson:"discord,omitempty" json:"discord,omitempty"`
	Website   *string `bson:"website,omitempty" json:"website,omitempty"`
}
