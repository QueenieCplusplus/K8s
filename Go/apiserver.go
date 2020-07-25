package main

import (
	"net"
	"time"
	"util" // not found
)

//2020, 7/25, am 9:30
// 建立 APIServer 物件實例的方法 NewAPIServer() 可以算是自製建構子

func NewAPIServer() *APIServer {
	s := APIServer{

		InsecurePort:        8080,
		InsecureBindAddress: util.IP(net.ParseIP("127.0.0.1")),
		BindAddress:         util.IP(net.ParseIP("0.0.0.0")),
		SecurePort:          6443,
		APIRate:             10.0,
		APIBurst:            200,
		APIPrefix:           "/api",
		EventTTL:            1 * time.Hour,
		AuthorizationMode:   "AlwaysAllow",
		ＡdmissionControl:    "AlwaysAdmit",
		// etcd
		EnableLogsSupport:      true,
		MasterServiceNameSpace: api.NamespaceDefault,
		ClusterName:            "k8s",
		CertDirectory:          "var/run/k8s",
		RuntimeConfig:          make(util.ConfigurationMap),
		KubeletConfig: client.KubeletConfig{
			port:        ports.KubeletPort,
			EnableHttps: true,
			HTTPTimeout: time.Duration(5) * time.Second,
		},
	}

	return &s
}
