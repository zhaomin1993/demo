package enum

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidStrategy = errors.New("invalid Strategy type")

func ParseStrategyFromLabelString(s string) (Strategy, error) {
	switch s {
	case "":
		return STRATEGY_UNKNOWN, nil
	case "机审":
		return STRATEGY__MACHINE, nil
	case "人审":
		return STRATEGY__MANUAL, nil
	case "机+人审":
		return STRATEGY__ALL, nil
	}
	return STRATEGY_UNKNOWN, InvalidStrategy
}

func (v Strategy) String() string {
	switch v {
	case STRATEGY_UNKNOWN:
		return ""
	case STRATEGY__MACHINE:
		return "MACHINE"
	case STRATEGY__MANUAL:
		return "MANUAL"
	case STRATEGY__ALL:
		return "ALL"
	}
	return "UNKNOWN"
}

func ParseStrategyFromString(s string) (Strategy, error) {
	switch s {
	case "":
		return STRATEGY_UNKNOWN, nil
	case "MACHINE":
		return STRATEGY__MACHINE, nil
	case "MANUAL":
		return STRATEGY__MANUAL, nil
	case "ALL":
		return STRATEGY__ALL, nil
	}
	return STRATEGY_UNKNOWN, InvalidStrategy
}

func (v Strategy) Label() string {
	switch v {
	case STRATEGY_UNKNOWN:
		return ""
	case STRATEGY__MACHINE:
		return "机审"
	case STRATEGY__MANUAL:
		return "人审"
	case STRATEGY__ALL:
		return "机+人审"
	}
	return "UNKNOWN"
}

func (v Strategy) Int() int {
	return int(v)
}

func (Strategy) TypeName() string {
	return "enum_demo.Strategy"
}

func (Strategy) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{STRATEGY__MACHINE, STRATEGY__MANUAL, STRATEGY__ALL}
}

func (v Strategy) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidStrategy
	}
	return []byte(str), nil
}

func (v *Strategy) UnmarshalText(data []byte) (err error) {
	*v, err = ParseStrategyFromString(string(bytes.ToUpper(data)))
	return
}

func (v Strategy) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *Strategy) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = Strategy(i)
	return nil
}
