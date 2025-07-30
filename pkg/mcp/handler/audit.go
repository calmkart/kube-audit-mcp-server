package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/calmkart/kube-audit-mcp-server/apis/types"
	"github.com/calmkart/kube-audit-mcp-server/pkg/mcp/metadata"
	"github.com/calmkart/kube-audit-mcp-server/pkg/provider"
)

// GetAuditLogHandler returns a handler for getting audit log
func GetAuditLogHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	p := provider.GenProvider()
	auditLog, err := p.GetAuditLog(&types.GetAuditLogParams{
		Resource:  request.GetString(metadata.GetAuditLogToolInfo.Resource.Name, ""),
		Name:      request.GetString(metadata.GetAuditLogToolInfo.Name.Name, ""),
		Namespace: request.GetString(metadata.GetAuditLogToolInfo.Namespace.Name, ""),
		TimeFrom:  request.GetInt(metadata.GetAuditLogToolInfo.TimeFrom.Name, int(time.Now().Unix()-4*3600)),
		TimeTo:    request.GetInt(metadata.GetAuditLogToolInfo.TimeTo.Name, int(time.Now().Unix())),
		LastN:     request.GetInt(metadata.GetAuditLogToolInfo.LastN.Name, 100),
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("下面是经过筛选的审计日志: %v", auditLog)), nil
}
