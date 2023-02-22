package proto

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func ToAny[T any, D *anypb.Any](source []T, transform func(T) protoreflect.ProtoMessage) []D {
	l := make([]D, 0, len(source))
	for _, v := range source {
		i, _ := anypb.New(transform(v))
		l = append(l, i)
	}
	return l
}
