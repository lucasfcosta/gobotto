package models

type Rules struct {
	Allow    map[string]int `json:"allow"`
	Disallow map[string]int `json:"disallow"`
}

func NewRules() Rules {
	rules := new(Rules)
	rules.Allow = make(map[string]int)
	rules.Disallow = make(map[string]int)
	return *rules
}
