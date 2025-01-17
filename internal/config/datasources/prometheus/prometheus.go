package prometheus

import (
	"fmt"
	"github.com/balerter/balerter/internal/config/common"
	"strings"
)

// Prometheus datasource config
type Prometheus struct {
	// Name of the datasource
	Name string `json:"name" yaml:"name" hcl:"name,label"`
	// URL value
	URL string `json:"url" yaml:"url" hcl:"url"`
	// BasicAuth contains auth data, if needed
	BasicAuth *common.BasicAuth `json:"basicAuth" yaml:"basicAuth" hcl:"basicAuth,block"`
	// Timeout value
	Timeout int `json:"timeout" yaml:"timeout" hcl:"timeout,optional"`
}

// Validate config
func (cfg Prometheus) Validate() error {
	if strings.TrimSpace(cfg.Name) == "" {
		return fmt.Errorf("name must be not empty")
	}
	if strings.TrimSpace(cfg.URL) == "" {
		return fmt.Errorf("url must be not empty")
	}
	if cfg.Timeout < 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}

	return nil
}
