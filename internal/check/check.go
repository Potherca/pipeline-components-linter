package check

type Status int64

const (
	Pass Status = iota
	Fail
	Skip
	Incomplete
)
