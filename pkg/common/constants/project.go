package constants

const (
	InsertProject       = "INSERT INTO projects (project_name, client_name, resource_link, deadline, status, proposal_link, assign) VALUES (?, ?, ?, ?, ?, ?, ?);"
	CheckProject        = "SELECT * FROM projects WHERE id=?;"
	UpdateProjectAdmin  = "UPDATE projects SET project_name = ?, client_name = ?, deadline = ?, status = ?, budget = ?, proposal_link = ?, assign = ?, resource_link = ? WHERE id = ?;"
	UpdateProjectClient = "UPDATE projects SET budget = ?, resource_link = ? WHERE id = ?;"
)

var (
	ProjectNotFound = "Project tidak ditemukan"
)
