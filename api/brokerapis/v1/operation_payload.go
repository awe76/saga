package v1

type OperationPayload struct {
	WorkflowName string
	WorkflowId   string
	IsRollback   bool
	IsFailed     bool
	Name         string
	From         string
	To           string
	Payload      string
}
