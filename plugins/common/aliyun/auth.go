package aliyun

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/client"
)

// CredentialConfig holds Aliyun authentication configuration
type CredentialConfig struct {
	AccessKeyID     *string
	AccessKeySecret *string
	SecurityToken   *string
	BearerToken     *string
	RegionId        *string
	ReadTimeout     *int
	ConnectTimeout  *int
	IdleTimeout     *int
	MaxIdleConns    *int
}

// GetConfig creates an OpenAPI configuration using the credential chain
func GetConfig(config CredentialConfig) (*client.Config, error) {
	 configuration := &client.Config{
		 AccessKeyId: config.AccessKeyID,
		 AccessKeySecret: config.AccessKeySecret,
	}

	return configuration, nil
}
