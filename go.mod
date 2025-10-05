module github.com/aryansharma9917/Codewise-CLI

go 1.20

require (
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/clbanning/mxj/v2 v2.7.0
	github.com/spf13/cobra v1.6.1
	github.com/tcnksm/go-latest v0.0.0-20170313132115-e3007ae9052e
)

replace github.com/aryansharma9917/Codewise-CLI => ./

require (
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/kr/pretty v0.3.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

// **Use local module for CI/CD**
// replace github.com/aryansharma9917/Codewise-CLI => ./
