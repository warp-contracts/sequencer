package config

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/spf13/viper"
)

type RateLimiter struct {
	// Arweave wallet owners that aren't impacted by the rate limiter
	WhiteListArweaveWalletOwners mapset.Set[string]

	// Number of monitored blocks
	NumberOfMonitoredBlocks int
}

func setRateLimiterDefaults() {
	viper.SetDefault("RateLimiter.WhiteListArweaveWalletOwners", mapset.NewSet[string]())
	viper.SetDefault("RateLimiter.NumberOfMonitoredBlocks", 100)
}
