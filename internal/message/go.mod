module message

go 1.22

require (
	internal/check v0.1.0
)

replace (
	internal/check => ../check
)
