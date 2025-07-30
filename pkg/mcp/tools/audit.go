package tools

import (
	"github.com/mark3labs/mcp-go/mcp"

	"github.com/calmkart/kube-audit-mcp-server/pkg/mcp/metadata"
)

// GetAuditLogTool returns a tool for getting audit log
func GetAuditLogTool() mcp.Tool {
	toolInfo := metadata.GetAuditLogToolInfo

	return mcp.NewTool(toolInfo.BaseInfo.Name,
		mcp.WithDescription(toolInfo.BaseInfo.Description),
		mcp.WithString(toolInfo.ClusterID.Name, mcp.Required(), mcp.Description(toolInfo.ClusterID.Description)),
		mcp.WithString(toolInfo.Resource.Name, mcp.Description(toolInfo.Resource.Description)),
		mcp.WithString(toolInfo.Namespace.Name, mcp.Description(toolInfo.Namespace.Description)),
		mcp.WithString(toolInfo.Name.Name, mcp.Description(toolInfo.Name.Description)),
		mcp.WithNumber(toolInfo.TimeFrom.Name, mcp.Description(toolInfo.TimeFrom.Description)),
		mcp.WithNumber(toolInfo.TimeTo.Name, mcp.Description(toolInfo.TimeTo.Description)),
	)
}
