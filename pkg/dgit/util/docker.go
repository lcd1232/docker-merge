package util

import (
	"github.com/containers/image/v5/types"
	"github.com/docker/docker/client"
	"os"
)

// NewSystemContextFromEnv returns a new SystemContext based on the environment.
// It uses the same logic as the docker CLI to determine the values.
func NewSystemContextFromEnv() *types.SystemContext {
	return &types.SystemContext{
		DockerDaemonCertPath:              os.Getenv(client.EnvOverrideCertPath),
		DockerDaemonHost:                  os.Getenv(client.EnvOverrideHost),
		DockerDaemonInsecureSkipTLSVerify: os.Getenv(client.EnvTLSVerify) == "",
	}
}
