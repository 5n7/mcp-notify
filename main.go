package main

import (
	"context"
	"log"

	"github.com/5n7/mcp-notify/notify"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func run() error {
	s := server.NewMCPServer(
		"mcp-notify",
		"v1.0.0",
		server.WithLogging(),
		server.WithToolCapabilities(true),
	)

	tool := mcp.NewTool(
		"notify",
		mcp.WithDescription("Cross-platform desktop notification tool"),
		mcp.WithString("title", mcp.Description("Notification title")),
		mcp.WithString("message", mcp.Description("Notification message")),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title := request.Params.Arguments["title"].(string)
		message := request.Params.Arguments["message"].(string)

		if err := notify.Notify(title, message); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText("notification sent successfully"), nil
	})

	return server.ServeStdio(s)
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
