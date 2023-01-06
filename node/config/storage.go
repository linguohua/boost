package config

import "github.com/filecoin-project/lotus/node/config"

// StorageManager convert boost config to config.SealerConfig
func (c *Boost) StorageManager() config.SealerConfig {
	return config.SealerConfig{
		ParallelFetchLimit: c.Storage.ParallelFetchLimit,
	}
}
