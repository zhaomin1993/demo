package enum

type Strategy uint8

// 模块策略
const (
	STRATEGY_UNKNOWN  Strategy = iota
	STRATEGY__MACHINE          // 机审
	STRATEGY__MANUAL           // 人审
	STRATEGY__ALL              // 机+人审
)
