//2020, 7/25, am 8:30
//go
//k8s
//kube-apiserver

// 大寫方可提供匯出
// 小寫是 file name 或是 {pkg name}
// 共用，需要放置 package main 或是放置同一 {pkg name}
package main

// shall import {path}
import (
	"fmt"
	"os"
	"pflag"
	"rand"    // not found
	"runtime" // not found
	"time"
	"util"
	"verflag" // not found
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	s := app.NewAPIServer()       // 建立一個 APIServer 結構的物件
	s.AddFlags(pflag.CommandLine) // 將命令列啟動參數傳入

	util.InitFlags()
	util.InitLogs()

	defer util.FlushLogs()

	verflag.PrintAndExitIfRequestes()

	if err := s.Run(pflag.CommandLine.Args()); err != nil { // 啟動並且監聽

		fmt.Fprintf(os.Stderr, "", err)
		os.Exit(1)

	}

}
