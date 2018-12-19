package main

type Cluster struct {
	Name    string `yaml:"name"`
	Cluster struct {
		CertificateAuthority     string `yaml:"certificate-authority,omitempty"`
		CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
		Server                   string `yaml:"server"`
		InsecureSkipTLSVerify    bool   `yaml:"insecure-skip-tls-verify,omitempty"`
	}
}

type Contexts struct {
	Name    string `yaml:"name"`
	Context struct {
		Cluster   string `yaml:"cluster"`
		Namespace string `yaml:"namespace,omitempty"`
		User      string `yaml:"user"`
	}
}

type User struct {
	Name string `yaml:"name"`
	User struct {
		AuthProvider struct {
			Name   string `yaml:"name,omitempty"`
			Config struct {
				ClientID     string `yaml:"client-id,omitempty"`
				ClientSecret string `yaml:"client-secret,omitempty"`
				ExtraScopes  string `yaml:"extra-scopes,omitempty"`
				IDToken      string `yaml:"id-token,omitempty"`
				IDPIssuerURL string `yaml:"idp-issuer-url,omitempty"`
			} `yaml: "config,omitempty"`
		} `yaml:"auth-provider,omitempty"`
		ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
		ClientKeyData         string `yaml:"client-key-data,omitempty"`
	} `yaml:"user"`
}

type Config struct {
	APIVersion     string     `yaml:"apiVersion"`
	Kind           string     `yaml:"kind"`
	Clusters       []Cluster  `yaml:"clusters"`
	Contexts       []Contexts `yaml:"contexts"`
	CurrentContext string     `yaml:"current-context"`
	Users          []User     `yaml:"users"`
	Preferences    struct{}
}
