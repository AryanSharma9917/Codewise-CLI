package env

type K8sConfig struct {
	Namespace string `yaml:"namespace"`
	Context   string `yaml:"context"`
}

type HelmConfig struct {
	Release string `yaml:"release"`
	Chart   string `yaml:"chart"`
	Values  string `yaml:"values"`
}

type GitOpsConfig struct {
	Repo   string `yaml:"repo"`
	Path   string `yaml:"path"`
	Branch string `yaml:"branch"`
}

type ValuesConfig struct {
	Image struct {
		Repository string `yaml:"repository"`
		Tag        string `yaml:"tag"`
	} `yaml:"image"`
}

type Env struct {
	Name   string
	K8s    K8sConfig
	Helm   HelmConfig
	GitOps GitOpsConfig
	Values ValuesConfig
}
