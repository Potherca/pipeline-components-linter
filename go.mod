module gitlab.com/pipeline-components/pipeline-component-linter

go 1.22

require (
	internal/check v0.1.0
	internal/checks v0.1.0
	internal/message v0.1.0
	internal/exitcodes v0.1.0
)

replace (
	internal/check => ./internal/check
	internal/checks => ./internal/checks
	internal/exitcodes => ./internal/exitcodes
	internal/message => ./internal/message
)
