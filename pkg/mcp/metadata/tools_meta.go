package metadata

import (
	"github.com/calmkart/kube-audit-mcp-server/apis/types"
)

var (
	// GetAuditLogToolInfo is the tool info for getting audit log
	GetAuditLogToolInfo = types.GetAuditLogToolInfo{
		BaseInfo: types.McpToolsBaseInfo{
			Name:        "get_audit_log",
			Description: "获取审计日志",
		},
		Namespace: types.McpToolParam{
			Name:        "namespace",
			Description: "这个参数是Kubernetes Namespace，用来表示是需要获取kubernetes集群具体哪个Namespace下的审计日志，对应审计日志里是objectRef.namespace字段",
		},
		Name: types.McpToolParam{
			Name:        "name",
			Description: "这个参数是Kubernetes object name(比如具体某个pod的名称或者node的名称，也就是具体k8s集群内对象的名称)，用来表示是需要获取kubernetes集群具体某个资源对象的审计日志，对应审计日志里是objectRef.name字段",
		},
		Resource: types.McpToolParam{
			Name:        "resource",
			Description: "这个参数是Kubernetes object resource(比如pods/nodes/services/deployments/statefulsets等等)，用来表示是需要获取kubernetes集群具体哪种资源的审计日志，对应审计日志里是objectRef.resource字段",
		},
		ClusterID: types.McpToolParam{
			Name:        "cluster_id",
			Description: "这个参数是集群ID，用来表示是需要获取具体哪个集群的审计日志",
		},
		TimeFrom: types.McpToolParam{
			Name:        "time_from",
			Description: "这个参数是待查询的审计日志的开始时间，格式为时间戳，单位是秒",
		},
		TimeTo: types.McpToolParam{
			Name:        "time_to",
			Description: "这个参数是待查询的审计日志的结束时间，格式为时间戳，单位是秒",
		},
		LastN: types.McpToolParam{
			Name:        "last_n",
			Description: "这个参数是待查询的审计日志的数量，用来表示需要获取审计日志的数量，按时间倒叙，读取最新的N条审计日志",
		},
	}
)
