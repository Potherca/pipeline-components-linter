# Pipeline Components Linter (plc-lint)

The Pipeline Components Linter (`plc-lint`) checks whether a Pipeline Component follows the Pipeline Component Guidelines.

## Installation

Download the pre-built binary from the [releases page][release-page] or, if you are familiar with Go, build it yourself.

## Usage

```bash
plc-lint <path-to-component>
```

## Contributing

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests.

### Development

This project has been set up following [Standard Go Project Layout][golang-standards-project-layout].

```
  .
  ├── cmd/           # Main application
  ├── internal/      # Private application and library code 
  ├── LICENSE
  └── README.md
```

## License

Created by Potherca under a [Mozilla Public License 2.0 (MPL-2.0) license][license].

[golang-standards-project-layout]: https://github.com/golang-standards/project-layout
[license]: LICENSE
[release-page]: https://gitlab.com/pipeline-component/org/plc-lint/-/releases
