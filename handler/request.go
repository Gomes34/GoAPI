package handler

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Locatoin string `json: "location"`
	Remote   string `json: "remote"`
	Link     string `json: "link"`
	Salary   string `json: "salary"`
}
