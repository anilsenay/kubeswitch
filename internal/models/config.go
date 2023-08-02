package models

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			Exec struct {
				APIVersion         string      `yaml:"apiVersion"`
				Args               []string    `yaml:"args"`
				Command            string      `yaml:"command"`
				Env                interface{} `yaml:"env"`
				InteractiveMode    string      `yaml:"interactiveMode"`
				ProvideClusterInfo bool        `yaml:"provideClusterInfo"`
			} `yaml:"exec"`
		} `yaml:"user"`
	} `yaml:"users"`
}
