package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

/*

{
status: "1",
info: "OK",
infocode: "10000",
province: [ ],
city: [ ],
adcode: [ ],
rectangle: [ ]
}

{
status: "1",
info: "OK",
infocode: "10000",
province: "山东省",
city: "青岛市",
adcode: "370200",
rectangle: "120.1946568,35.98597607;120.6139934,36.24191092"
}

*/
// GetExternalLocation 获取外网ip地址 调用高德接口
func GetExtLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		return "amap service failed"

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		return err.Error()
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}

// GetLocalHost 获取局域网ip地址
func GetLocalHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "net.Interfaces  err:" + err.Error()
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
