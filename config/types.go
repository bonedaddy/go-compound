package config

// Config is used to configure the go-compound api client
type Config struct {
	// the ethereum address you use for compound
	Address string `yaml:"address"`
	// map of names to contract address
	// ex. cSAI key could be used to return the cSAI contract address
	Contracts  map[string]string `yaml:"contracts"`
	Blockchain `yaml:"blockchain"`
}

// Blockchain provides configuration information for any blockchian
// connections
type Blockchain struct {
	// the endpoint of an ethereum node, prefferably go-ethereum
	Endpoint string `yaml:"endpoint"`
	// the json keyfile for our ethereum account
	KeyFile string `yaml:"key_file"`
	// the password to unlock the json keyfile
	KeyPass string `yaml:"key_pass"`
}
