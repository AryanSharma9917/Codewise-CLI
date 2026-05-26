module github.com/aryansharma9917/codewise-cli

go 1.25.0

require (
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/clbanning/mxj/v2 v2.7.0
	github.com/spf13/cobra v1.10.2
	github.com/tcnksm/go-latest v0.0.0-20170313132115-e3007ae9052e
)

replace github.com/aryansharma9917/codewise-cli => ./

require (
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/term v0.43.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)

require (
	github.com/BurntSushi/toml v1.6.0
	github.com/kr/pretty v0.3.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

require (
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/hashicorp/go-version v1.9.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	golang.org/x/net v0.55.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

// **Use local module for CI/CD**
// replace github.com/aryansharma9917/codewise-cli => ./
