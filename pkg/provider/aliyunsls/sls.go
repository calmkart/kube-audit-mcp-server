package aliyunsls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sls20201230 "github.com/alibabacloud-go/sls-20201230/v6/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"

	"github.com/calmkart/kube-audit-mcp-server/apis/types"
	loginterface "github.com/calmkart/kube-audit-mcp-server/interface"
)

type Provider struct {
	accessKey    string
	accessSecret string
	region       string
	project      string
	logStore     string
}

func GenProvider() loginterface.AuditLogInterface {
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY")
	accessSecret := os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET")
	region := os.Getenv("ALIBABA_CLOUD_REGION")
	project := os.Getenv("ALIBABA_CLOUD_PROJECT")
	logStore := os.Getenv("ALIBABA_CLOUD_LOG_STORE")
	if accessKey == "" || accessSecret == "" || region == "" || project == "" || logStore == "" {
		panic("ALIBABA_CLOUD_ACCESS_KEY or ALIBABA_CLOUD_ACCESS_SECRET is empty")
	}
	return &Provider{
		accessKey:    accessKey,
		accessSecret: accessSecret,
		region:       region,
		project:      project,
		logStore:     logStore,
	}
}

func (p *Provider) GetAuditLog(params *types.GetAuditLogParams) (string, error) {
	log.Printf("resource: %v, namespace: %v, name: %v, TimeFrom: %v, TimeTo: %v", params.Resource, params.Namespace, params.Name, params.TimeFrom, params.TimeTo)
	client, err := p.createClient()
	if err != nil {
		return "", err
	}

	getLogsV2Headers := &sls20201230.GetLogsV2Headers{
		AcceptEncoding: tea.String("lz4"),
	}

	getLogsV2Request := &sls20201230.GetLogsV2Request{
		From:    tea.Int32(int32(params.TimeFrom)),
		To:      tea.Int32(int32(params.TimeTo)),
		Query:   tea.String(buildQuery(params.Resource, params.Namespace, params.Name)),
		Reverse: tea.Bool(true),
		Line:    tea.Int64(100),
	}

	runtime := &util.RuntimeOptions{}
	res, err := client.GetLogsV2WithOptions(tea.String(p.project), tea.String(p.logStore), getLogsV2Request, getLogsV2Headers, runtime)
	if err != nil {
		return "", err
	}

	log.Printf("resource: %v, namespace: %v, name: %v, from: %v, to: %v", params.Resource, params.Namespace, params.Name, params.TimeFrom, params.TimeTo)

	// compact json for shorted content
	compactJson := &bytes.Buffer{}
	_ = json.Compact(compactJson, []byte(res.Body.String()))
	return compactJson.String(), nil
}

func (p *Provider) createClient() (result *sls20201230.Client, err error) {
	c, err := credential.NewCredential(&credential.Config{
		Type:            tea.String("access_key"),
		AccessKeyId:     tea.String(p.accessKey),
		AccessKeySecret: tea.String(p.accessSecret),
	})
	if err != nil {
		return result, err
	}

	config := &openapi.Config{
		Credential: c,
	}
	config.Endpoint = tea.String(fmt.Sprintf("%v.log.aliyuncs.com", p.region))
	result = &sls20201230.Client{}
	result, err = sls20201230.NewClient(config)
	return result, err
}

func buildQuery(resource, namespace, name string) string {
	query := fmt.Sprintf("*")
	if resource != "" {
		query += fmt.Sprintf(" and objectRef.resource: %v", resource)
	}
	if namespace != "" {
		query += fmt.Sprintf(" and objectRef.namespace: %v", namespace)
	}
	if name != "" {
		query += fmt.Sprintf(" and objectRef.name: %v", name)
	}
	return query
}
