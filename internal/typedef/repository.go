package typedef

type Repository struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"url"`
	Cron     string   `yaml:"cron"`
	Storage  []string `yaml:"storage"`
	UseCache bool     `yaml:"useCache"`
	Type     string   `yaml:"type"` // repo, user, org (default: repo)
	OrgName  string   `yaml:"orgName"`
}
