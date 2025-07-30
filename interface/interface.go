package _interface

import (
	"github.com/calmkart/kube-audit-mcp-server/apis/types"
)

type AuditLogInterface interface {
	GetAuditLog(params *types.GetAuditLogParams) (string, error)
}
