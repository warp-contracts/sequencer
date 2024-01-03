package config

import "github.com/spf13/viper"

type RateLimiter struct {
	// Arweave wallet owners that aren't impacted by the rate limiter
	WhiteListArweaveWalletOwners []string
}

func setRateLimiterDefaults() {
	viper.SetDefault("RateLimiter.WhiteListArweaveWalletOwners", []string{})
}
