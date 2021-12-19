package models

type Price struct {
	K string `bson:"k,omitempty"`
	V int64  `bson:"v,omitempty"`
}

func NewSpec(k string, v int64) *Price {
	price := Price{
		K: k,
		V: v,
	}

	return &price
}
