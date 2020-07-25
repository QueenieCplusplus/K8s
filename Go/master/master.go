// 2020, 7/25, am 9:45
// 關鍵程式碼
// master contains state for a K8s cluster master/api server.
// APIServer 回傳資料給此 Master 物件

package main

import (
	"net"
	"time"
	"util"
)

type Master struct {
	serviceClusterIPRange *net.IPnet
	serviceNodePortRange  util.PortRange
	cacheTimeout          time.Duration
	minRequestTimeout     time.Duration

	mux apiserver.Mux
	//muxHelper
	//handlerContainer
	//rootWebService

}
