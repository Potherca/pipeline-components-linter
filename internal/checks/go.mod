module checks

go 1.22.6

require (
	internal/check v0.1.0
	internal/message v0.1.0
)

replace (
	internal/message => ../message
	internal/check => ../check
)
