package general

const (
	Empty = ""
)

var (
	Null = struct{}{}
)

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Int | Uint
}

type Numberic interface {
	Integer | Float
}

type Ordered interface {
	Numberic | ~string
}
