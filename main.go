package main

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"time"

	"github.com/qbox/jarvis/cmd/jarvisbilling/shared"
	"github.com/qbox/jarvis/common/logtool"
)

func main() {
	var (
		mName      = flag.String("m", "", "cron方法名称")
		configFile = flag.String("f", "./jarvisbilling-api.yaml", "the config file")
	)
	flag.Parse()
	if *mName == "" || *configFile == "" {
		fmt.Println("param error")
		return
	}
	closeF, err := logtool.OutputLog()
	if err != nil {
		fmt.Printf("set log file faild:%v \n", err)
		return
	}
	defer closeF()
	method := shared.GetTaskMethod(*configFile, *mName)
	if !method.IsValid() {
		fmt.Println("Super 调用无效方法")
		return
	}
	paramList := []reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(""),
	}
	start := time.Now()
	fmt.Printf("[SuperTriggerCron] event(%v) started, time:%v \n", *mName, start)
	method.Call(paramList)
	fmt.Printf("[SuperTriggerCron] event(%v) done, duration:%v \n", *mName, time.Since(start))
}
