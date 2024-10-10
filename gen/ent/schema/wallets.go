package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Wallet struct {
	ent.Schema
}

// Fields of the Audio.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("UUID").
			NotEmpty().Unique(),
		field.Int("balance").
			Default(0),
	}
}
