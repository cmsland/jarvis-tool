package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

/*
// github.com/qbox/jarvis/cmd/javisbilling/shared/local_ignored.go

package shared

import (
	"reflect"

	"github.com/qbox/jarvis/cmd/jarvisbilling/internal/config"
	"github.com/qbox/jarvis/cmd/jarvisbilling/internal/svc"
	"github.com/qbox/jarvis/cmd/jarvisbilling/internal/tasks"
	"github.com/zeromicro/go-zero/core/conf"
)

func GetTaskMethod(configFile, methodName string) reflect.Value {
	var c config.Config
	conf.MustLoad(configFile, &c)
	task := tasks.NewTimer(svc.NewServiceContext(c))
	return reflect.ValueOf(task).MethodByName(methodName)
}
*/
