package node

// Config struct allows us to manipulate the nodejs config
// to add what's missing.
type Config struct {
	License         string                 `json:"license"`
	Main            string                 `json:"main"`
	Name            string                 `json:"name"`
	Repository      string                 `json:"repository"`
	Scripts         map[string]interface{} `json:"scripts"`
	Version         string                 `json:"version"`
	Dependencies    map[string]interface{} `json:"dependencies"`
	DevDependencies map[string]interface{} `json:"devDependencies"`
}
