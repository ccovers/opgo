package cutil

import (
	"net"
)

var get_interface_name = "eth0"

// note: 需要检查返回为空字符串的情况
func GetEth0Mac() string {
	var ret string = ""

	for ok := true; ok; ok = false {
		interfaces, err := net.Interfaces()
		if err != nil {
			break
		}
		for _, inter := range interfaces {
			mac := inter.HardwareAddr
			if inter.Name == get_interface_name {
				ret = mac.String()
				break
			}
		}
	}

	return ret
}
