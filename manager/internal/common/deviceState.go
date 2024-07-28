package common

type deviceState int

var (
	Running  deviceState = 1
	Down     deviceState = 2
	UnNormal deviceState = 3
)

func (d deviceState) String() string {
	return []string{"Running", "Down", "UnNormal"}[d-1]
}

func (d deviceState) Int() int {
	return int(d)
}
