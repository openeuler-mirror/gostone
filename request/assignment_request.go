package request

type AssignmentSearch struct {
	RoleId   string `json:"role.id"`
	ActorId  string `json:"user.id"`
	TargetId string `json:"scope.project.id"`
}
