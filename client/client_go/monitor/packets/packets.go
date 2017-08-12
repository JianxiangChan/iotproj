package packets

//定义上行采集温度数据包
type TempPacket struct {
	Mac         string  `json:"mac"`         //设备号，定义树莓派板唯一标识，使用MAC地址
	Temperature float32 `json:"temperature"` //温度
	LED         bool    `json:"led"`         //LED灯亮 = true ,LED灯灭=flase
}

//定义下行控制数据包
type LedPacket struct {
	Mac string `json:"mac"` //设备号，定义树莓派板唯一标识，使用MAC地址
	LED bool   `json:"led"` //控制LED灯亮 = true ,LED灯灭=flase
}
