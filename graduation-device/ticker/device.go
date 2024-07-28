package ticker

import (
	"errors"
	"gratuation-device/device"
	"gratuation-device/utils"
)

var devices = []*device.Device{
	{
		Code: "e1",
		Data: []*device.Data{
			{
				Key:  "电流",
				Unit: "A",
			},
			{
				Key:  "电压",
				Unit: "V",
			},
		},
	},
	{
		Code: "e2",
		Data: []*device.Data{
			{
				Key:  "三相电流",
				Unit: "A",
			},
			{
				Key:  "单相电流",
				Unit: "A",
			},
		},
	},
	{
		Code: "e3",
		Data: []*device.Data{
			{
				Key:  "功率",
				Unit: "KW/h",
			},
			{
				Key:  "有功功率",
				Unit: "KW/h",
			},
			{
				Key:  "无功功率",
				Unit: "KW/h",
			},
		},
	},
}

var rangeMap = map[string]struct {
	Mix int
	Max int
}{
	"电流":   {Mix: 1, Max: 20},
	"电压":   {Mix: 1, Max: 300},
	"三相电流": {Mix: 1, Max: 500},
	"单相电流": {Mix: 1, Max: 500},
	"有功功率": {Mix: 1, Max: 500},
	"无功功率": {Mix: 1, Max: 500},
	"功率":   {Mix: 1, Max: 500},
}

func GetDevice() []*device.Device {
	return devices
}

func GetRangeMap() map[string]struct {
	Mix int
	Max int
} {
	return rangeMap
}

func GenerateDeviceData(device []*device.Device) error {
	for _, dev := range device {
		for _, detail := range dev.Data {
			if _, ok := GetRangeMap()[detail.Key]; !ok {
				return errors.New("range map key not found")
			}
			detail.Value = utils.GetRandomNum(GetRangeMap()[detail.Key].Mix, GetRangeMap()[detail.Key].Max)
		}
	}
	return nil
}
