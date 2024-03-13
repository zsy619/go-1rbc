package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun5(t *testing.T) {
	fmt.Println("TestRun5 Begin")
	start := time.Now()
	err := Run5("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun5 总执行时间（秒）：", time.Since(start).Seconds())
}
