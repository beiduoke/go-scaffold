package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
	mx "github.com/beiduoke/go-scaffold/pkg/entgo/mixin"
	mixinE "github.com/tx7do/go-utils/entgo/mixin"
)

var _ ent.Mixin = (*MixinTop)(nil)

type MixinTop struct{ mixin.Schema }

func (MixinTop) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, mixinE.AutoIncrementId{}.Fields()...)
	fields = append(fields, mixinE.TimeAt{}.Fields()...)
	fields = append(fields, mixinE.Remark{}.Fields()...)
	fields = append(fields, mx.Sort{}.Fields()...)
	fields = append(fields, mx.State{}.Fields()...)
	return fields
}
