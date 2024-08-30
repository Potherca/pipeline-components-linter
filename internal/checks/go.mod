module checks

go 1.22

require (
	internal/asserts v0.1.0
	internal/check v0.1.0
	internal/message v0.1.0
	internal/repositoryContents v0.1.0
)

replace (
	internal/asserts => ../asserts
	internal/check => ../check
	internal/message => ../message
	internal/repositoryContents => ../repositoryContents
)
