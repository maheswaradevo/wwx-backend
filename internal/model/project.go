package model

import "time"

type Project struct {
	ProjectID    int       `db:"id"`
	UserId       int       `db:"user_id"`
	ProjectName  string    `db:"project_name"`
	ClientName   string    `db:"client_name"`
	Deadline     string    `db:"deadline"`
	Status       string    `db:"status"`
	Budget       uint64    `db:"budget"`
	ProposalLink string    `db:"proposal_link"`
	Resource     string    `db:"resource_link"`
	Assign       string    `db:"assign"`
	Maintenance  int       `db:"maintenance"`
	CreatedAt    time.Time `db:"created_at"`
}

type ProjectRequest struct {
	ProjectName  string `json:"project_name"`
	ClientName   string `json:"client_name"`
	Deadline     string `json:"deadline"`
	Status       string `json:"status"`
	Budget       uint64 `json:"budget"`
	ProposalLink string `json:"proposal_link"`
	Assign       string `json:"assign"`
	Resource     string `json:"resource"`
}

type EditProjectRequest struct {
	ProjectName  *string `json:"project_name,omitempty"`
	ClientName   *string `json:"client_name,omitempty"`
	Deadline     *string `json:"deadline,omitempty"`
	Status       *string `json:"status,omitempty"`
	Budget       *uint64 `json:"budget,omitempty"`
	ProposalLink *string `json:"proposal_link,omitempty"`
	Assign       *string `json:"assign,omitempty"`
	Resource     *string `json:"resource,omitempty"`
	Maintenance  *int    `json:"maintenance,omitempty"`
}

type ProjectViewRequest struct {
	Status string `json:"status"`
}

type ProjectViewEditRequest struct {
	ProjectID int `json:"project_id"`
}
