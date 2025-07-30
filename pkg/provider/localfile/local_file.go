package localfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/calmkart/kube-audit-mcp-server/apis/types"
	loginterface "github.com/calmkart/kube-audit-mcp-server/interface"
)

type Provider struct {
	filePath string
}

func GenProvider() loginterface.AuditLogInterface {
	filePath := os.Getenv("filePath")
	if filePath == "" {
		panic("filePath is empty")
	}
	return &Provider{
		filePath: filePath,
	}
}

func (p *Provider) GetAuditLog(params *types.GetAuditLogParams) (string, error) {
	log.Printf("resource: %v, namespace: %v, name: %v, lastN: %v", params.Resource, params.Namespace, params.Name, params.LastN)

	keyWords := buildFilterKeywords
	auditLogsArray, err := readAndFilter(p.filePath, keyWords(params.Resource, params.Namespace, params.Name))
	if err != nil {
		return "", err
	}

	// compact json for shorted content
	compactJson := &bytes.Buffer{}
	_ = json.Compact(compactJson, marshalWithLastN(auditLogsArray, params.LastN))
	return compactJson.String(), nil
}

func buildFilterKeywords(resource, namespace, name string) []string {
	keyWords := []string{}
	if resource != "" {
		keyWords = append(keyWords, fmt.Sprintf("objectRef.resource: %v", resource))
	}
	if namespace != "" {
		keyWords = append(keyWords, fmt.Sprintf("objectRef.namespace: %v", namespace))
	}
	if name != "" {
		keyWords = append(keyWords, fmt.Sprintf("objectRef.name: %v", name))
	}
	return keyWords
}

func marshalWithLastN(auditLogsArray []string, lastN int) []byte {
	newArray := auditLogsArray
	if lastN != 0 {
		newArray = auditLogsArray[len(auditLogsArray)-lastN:]
	}
	r, _ := json.Marshal(newArray)
	return r
}
