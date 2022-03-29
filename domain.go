package crmenforcer

import "strings"

type session struct {
	ID       string
	Username string
	Token    string
}

func (s *session) Destructure() (tenant string, email string) {
	slc := strings.Split(s.Username, ">")
	tenant, email = slc[0], slc[1]
	return
}

type rbac struct {
	User     string `json:"user"`
	Role     string `json:"role"`
	Tenant   string `json:"tenant"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
