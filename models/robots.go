package models

type Robots struct {
	Comments []string         `json:"comments"`
	Rules    map[string]Rules `json:"rules"`
}

func NewRobots() Robots {
	robots := new(Robots)
	robots.Comments = *new([]string)
	robots.Rules = make(map[string]Rules)

	return *robots
}
