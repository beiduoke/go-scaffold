package constant

// 处理类型
type HandleType int32

const (
	// 特殊处理前端现实状态
	HandleType_success HandleType = 1
	HandleType_error   HandleType = 2
	HandleType_warning HandleType = 3
)

// Enum value maps for HandleType.
var (
	HandleType_name = map[int32]string{
		0: "unspecified",
		1: "success",
		2: "error",
		3: "warning",
	}
	HandleType_value = map[string]int32{
		"unspecified": 0,
		"success":     1,
		"error":       2,
		"warning":     3,
	}
)

func (x HandleType) Enum() *HandleType {
	p := new(HandleType)
	*p = x
	return p
}

func (x HandleType) String() string {
	return HandleType_name[int32(x)]
}
