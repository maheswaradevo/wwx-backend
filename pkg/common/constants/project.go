package constants

const (
	InsertProject       = "INSERT INTO projects (user_id, project_name, client_name, resource_link, deadline, status, proposal_link, assign, budget) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	CheckProject        = "SELECT * FROM projects WHERE id=?;"
	UpdateProjectAdmin  = "UPDATE projects SET project_name = ?, client_name = ?, deadline = ?, status = ?, budget = ?, proposal_link = ?, assign = ?, resource_link = ? WHERE id = ?;"
	UpdateProjectClient = "UPDATE projects SET budget = ?, resource_link = ? WHERE id = ?;"
	SearchProject       = `SELECT * FROM projects WHERE project_name LIKE '%s%%';`
	ViewProject         = "SELECT p.id, p.project_name, p.client_name, p.deadline, p.status, p.budget, p.proposal_link, p.assign, p.resource_link, p.user_id FROM projects p INNER JOIN users u ON u.id = p.user_id WHERE u.id = ? ORDER BY p.created_at DESC;"
	ViewProjectAdmin    = "SELECT * FROM projects ORDER BY created_at DESC"
)

var (
	ProjectNotFound = "Project tidak ditemukan"
)
