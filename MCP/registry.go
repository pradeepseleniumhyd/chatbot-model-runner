package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_api "github.com/input-api/mcp-server/tools/api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api.CreateGetTool(cfg),
		tools_api.CreateGet_api_chatTool(cfg),
	}
}
