package controller

type MsgIncrement struct {
	// Limiter index
	LimiterIndex int
	// Key
	Key string
}

type MsgDelete struct {
	// Limiter index
	LimiterIndex int
	// Key
	Key string
}
