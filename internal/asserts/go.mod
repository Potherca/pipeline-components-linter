module asserts

go 1.22

require (
	internal/check v0.1.0
	internal/message v0.1.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	internal/check => ../check
	internal/message => ../message
)
