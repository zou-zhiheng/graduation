package common

// operationState 定义操作状态
type operationState int

// 取值范围：1. 未进行
// 2. 进行中
// 3. 处理成功
// 4. 处理失败
const (
	OperationNotStarted operationState = iota + 1
	OperationStateOperating
	OperationStateSuccess
	OperationStateFailed
)

func (o operationState) Int() int {
	return int(o)
}

func (o operationState) Uint32() uint32 {
	return uint32(o)
}

func (o operationState) String() string {
	return [...]string{"Operation Not Started", "Operation Operating", "Operation Success", "Operation Failed"}[o-1]
}

// isValid 定义是否启用
type isValid int

const (
	Enable isValid = iota + 1
	Disable
)

func (i isValid) Int() int {
	return int(i)
}

func (i isValid) Uint32() uint32 {
	return uint32(i)
}

func (i isValid) String() string {
	return [...]string{"enable", "disable"}[i-1]
}

type protocolCustom int

const (
	CustomProtocol protocolCustom = iota + 1
	UnCustomProtocol
)

func (s protocolCustom) Int() int {
	return int(s)
}

func (s protocolCustom) Int32() int32 {
	return int32(s)
}

func (s protocolCustom) String() string {
	return [...]string{"YES", "NO"}[s-1]
}

const TimeFormat = "2006-01-02 15:04:05"
