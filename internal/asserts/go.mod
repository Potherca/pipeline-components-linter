module asserts

go 1.22

require (
	internal/check v0.1.0
	internal/message v0.1.0
)

replace (
	internal/check => ../check
	internal/message => ../message
)
