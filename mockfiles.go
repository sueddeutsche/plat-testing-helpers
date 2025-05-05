package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteMockAwsProvider(path string) error {
	content := `
	provider "aws" {
		region = "eu-central-1"

		access_key                  = "null"
		secret_key                  = "null"
		skip_credentials_validation = true
		skip_metadata_api_check     = true
		skip_requesting_account_id  = true

		default_tags {
			tags = {
				zone    = "test"
				env     = "integration"
				service = "infrastructure"
				team    = "plat"
			}
		}
  }
  `
	filePath := filepath.Join(path, "mock_aws_provider.tf")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func WriteMockFile(mockData map[string]any, path string) error {
	filePath := filepath.Join(path, "mock.json")

	file, err := os.Create(filePath)
	if err != nil {
		return (err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(mockData)
	if err != nil {
		return (err)
	}
	return nil
}
