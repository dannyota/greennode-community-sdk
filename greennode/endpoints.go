package greennode

import (
	"fmt"
)

// resolveEndpoints fills in empty endpoint fields on cfg based on Region and auth method.
// Explicit endpoints always take precedence.
func resolveEndpoints(cfg *Config) {
	if cfg.Region == "" {
		return
	}

	isIAMUser := cfg.IAMAuth != nil

	if cfg.IAMEndpoint == "" && !isIAMUser {
		cfg.IAMEndpoint = "https://iamapis.vngcloud.vn/accounts-api/"
	}

	if cfg.VServerEndpoint == "" {
		if isIAMUser {
			cfg.VServerEndpoint = fmt.Sprintf("https://%s.console.vngcloud.vn/vserver/iam-vserver-gateway/", cfg.Region)
		} else {
			cfg.VServerEndpoint = fmt.Sprintf("https://%s.api.vngcloud.vn/vserver/vserver-gateway/", cfg.Region)
		}
	}

	if cfg.VLBEndpoint == "" {
		if isIAMUser {
			cfg.VLBEndpoint = fmt.Sprintf("https://%s.console.vngcloud.vn/vserver/iam-vlb-gateway/", cfg.Region)
		} else {
			cfg.VLBEndpoint = fmt.Sprintf("https://%s.api.vngcloud.vn/vserver/vlb-gateway/", cfg.Region)
		}
	}

	if cfg.VNetworkEndpoint == "" {
		cfg.VNetworkEndpoint = fmt.Sprintf("https://%s.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork/", cfg.Region)
	}

	if cfg.GLBEndpoint == "" {
		cfg.GLBEndpoint = "https://glb.console.vngcloud.vn/glb-controller/"
	}

	if cfg.DNSEndpoint == "" {
		cfg.DNSEndpoint = "https://vdns.api.vngcloud.vn/"
	}
}
