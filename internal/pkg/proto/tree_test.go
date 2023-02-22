package proto

import (
	"testing"

	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
)

type Domain struct {
	*v1.Domain
}

var domainData = []*v1.Domain{
	{
		Id:       1,
		Name:     "测试1",
		Title:    "title1",
		ParentId: 0,
	}, {
		Id:       2,
		Name:     "测试2",
		Title:    "title2",
		ParentId: 0,
	}, {
		Id:       3,
		Name:     "测试3",
		Title:    "title3",
		ParentId: 1,
	}, {
		Id:       4,
		Name:     "测试4",
		Title:    "title4",
		ParentId: 1,
	}, {
		Id:       5,
		Name:     "测试5",
		Title:    "title5",
		ParentId: 2,
	}, {
		Id:       6,
		Name:     "测试6",
		Title:    "title6",
		ParentId: 0,
	},
}

func (d *Domain) AppendChildren(arg any) {
	domains, ok := arg.([]*Domain)
	if !ok {
		return
	}
	l := make([]*v1.Domain, 0, len(domains))
	for _, v := range domains {
		l = append(l, v.Domain)
	}
	d.Children = append(d.Children, l...)
}

func Test_TreeProto(t *testing.T) {
	t.Log("测试泛型树形结构开始")
	data := ToTree(domainData, 0, func(t *v1.Domain, ts ...*v1.Domain) error {
		t.Children = append(t.Children, ts...)
		return nil
	})
	t.Logf("%v", data)
	t.Log("测试泛型树形结构结束")
}
