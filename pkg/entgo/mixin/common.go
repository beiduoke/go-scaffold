package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
	minxinE "github.com/tx7do/go-utils/entgo/mixin"
)

var _ ent.Mixin = (*Common)(nil)

type Common struct{ mixin.Schema }

func (Common) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, minxinE.AutoIncrementId{}.Fields()...)
	fields = append(fields, minxinE.TimeAt{}.Fields()...)
	fields = append(fields, minxinE.Remark{}.Fields()...)
	fields = append(fields, PlatformId{}.Fields()...)
	fields = append(fields, Sort{}.Fields()...)
	fields = append(fields, State{}.Fields()...)
	return fields
}
