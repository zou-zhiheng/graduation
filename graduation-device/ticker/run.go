package ticker

import (
	"fmt"
	"gratuation-device/device"
	"time"
)

func Run() {
	tick := time.NewTicker(time.Minute)
	for {
		select {
		case <-tick.C:
			fmt.Println("开始执行本轮任务")
			Task()
			fmt.Println("本轮任务执行结束")
		}
	}
}

func Task() {
	//初始化客户端
	client := device.NewDeviceClient("http://49.232.213.36:8080/api/v1/device/dateReceive")
	err := GenerateDeviceData(GetDevice())
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := client.GenerateData(device.JsonData{Data: GetDevice()})
	err = client.Send(data)
	fmt.Println(err)
}
