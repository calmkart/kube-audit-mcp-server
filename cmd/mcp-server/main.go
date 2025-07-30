package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"

	"github.com/calmkart/kube-audit-mcp-server/pkg/mcp/handler"
	"github.com/calmkart/kube-audit-mcp-server/pkg/mcp/tools"

	"github.com/calmkart/kube-audit-mcp-server/config"
)

func main() {
	// init for config flags
	config.Init()
	c := config.GetConfig()

	// create mcp server
	s := server.NewMCPServer(
		"K8s Audit Mcp Server",
		"1.0.0",
	)

	// add tools
	s.AddTool(tools.GetAuditLogTool(), handler.GetAuditLogHandler)

	// start mcp server
	switch c.Mode {
	case config.ModeStdio:
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Server error: %v\n", err)
			os.Exit(1)
		}
	case config.ModeSSE:
		sseURL := fmt.Sprintf("http://%v:%v", c.Host, c.Port)
		sseServer := server.NewSSEServer(s, server.WithBaseURL(sseURL))
		log.Printf("SSE server started at %s", sseURL)
		if err := sseServer.Start(fmt.Sprintf(":%d", c.Port)); err != nil {
			log.Fatalf("Server error: %v", err)
		}
		return
	}
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
