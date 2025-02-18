package platform

import (
	"fmt"
	"os"

	"github.com/buildpacks/lifecycle/log"
)

const (
	FeatureDockerfiles = "Dockerfiles"
)

var ExperimentalMode = envOrDefault(EnvExperimentalMode, DefaultExperimentalMode)

func GuardExperimental(requested string, logger log.Logger) error {
	switch ExperimentalMode {
	case ModeQuiet:
		break
	case ModeError:
		logger.Errorf("Platform requested experimental feature '%s'", requested)
		return fmt.Errorf("experimental features are disabled by %s=%s", EnvExperimentalMode, ModeError)
	case ModeWarn:
		logger.Warnf("Platform requested experimental feature '%s'", requested)
	default:
		// This shouldn't be reached, as ExperimentalMode is always set.
		logger.Warnf("Platform requested experimental feature '%s'", requested)
	}
	return nil
}

func envOrDefault(key string, defaultVal string) string {
	if envVal := os.Getenv(key); envVal != "" {
		return envVal
	}
	return defaultVal
}
