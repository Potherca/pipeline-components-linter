package check

type Status int64

const (
	Error        = -1
	Pass  Status = iota
	Fail
	Skip
	Incomplete
)
