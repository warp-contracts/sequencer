package controller

type MsgSet struct {
	LimiterIndex int
	Key          string
	Value        int64
}

type MsgSubtract struct {
	LimiterIndex int
	Key          string
	Value        int64
}
