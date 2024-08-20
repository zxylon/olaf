package config

var (
	Version       = "1.1.0"
	WireCmd       = "github.com/google/wire/cmd/wire@latest"
	OlafCmd       = "github.com/zxylon/olaf@latest"
	RepoBase      = "https://github.com/zxylon/olaf-layout-basic.git"
	RepoAdvanced  = "https://github.com/zxylon/olaf-layout-advanced.git"
	RepoChat      = "https://github.com/zxylon/olaf-layout-chat.git"
	RunExcludeDir = ".git,.idea,tmp,vendor"
	RunIncludeExt = "go,html,yaml,yml,toml,ini,json,xml,tpl,tmpl"
)
