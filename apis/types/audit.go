package types

// GetAuditLogRequest is the request for getting audit log
type GetAuditLogRequest struct {
	ClusterID string
	Namespace string
	Name      string
	Resource  string
}

type McpToolsBaseInfo struct {
	Name        string
	Description string
}

type McpToolParam struct {
	Name        string
	Description string
}

type GetAuditLogToolInfo struct {
	BaseInfo  McpToolsBaseInfo
	Namespace McpToolParam
	Name      McpToolParam
	Resource  McpToolParam
	ClusterID McpToolParam
	TimeFrom  McpToolParam
	TimeTo    McpToolParam
	LastN     McpToolParam
}

type GetAuditLogParams struct {
	Resource  string
	Name      string
	Namespace string
	TimeFrom  int
	TimeTo    int
	LastN     int
}
