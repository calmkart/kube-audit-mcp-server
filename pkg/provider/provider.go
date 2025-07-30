package provider

import (
	"github.com/calmkart/kube-audit-mcp-server/config"
	loginterface "github.com/calmkart/kube-audit-mcp-server/interface"
	"github.com/calmkart/kube-audit-mcp-server/pkg/provider/aliyunsls"
	"github.com/calmkart/kube-audit-mcp-server/pkg/provider/localfile"
)

var (
	provider loginterface.AuditLogInterface
)

// GenProvider returns a provider
func GenProvider() loginterface.AuditLogInterface {
	if provider != nil {
		return provider
	}
	switch config.GetConfig().Provider {
	case "localfile":
		provider = localfile.GenProvider()
	case "aliyunsls":
		provider = aliyunsls.GenProvider()
	default:
		provider = localfile.GenProvider()
	}
	return provider
}
