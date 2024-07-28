package utils

import (
	"math/rand"
	"time"
)

func GetRandomNum(min, max int) int {
	// 设置随机种子，确保每次运行程序时生成的随机数不同
	rand.Seed(time.Now().UnixNano())

	// 生成一个[min, max]之间的随机整数
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
