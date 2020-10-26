package system_config

// 系统配置段枚举的值类型
type ConfigSection string

// 系统配置段枚举类型
type configSectionEnumT struct {
	ActionPoint ConfigSection //用户行为分数配置
}

var ConfigSectionEnum = configSectionEnumT{
	ActionPoint: "action_point",
}
