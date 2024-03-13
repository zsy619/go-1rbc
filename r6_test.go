package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun6(t *testing.T) {
	fmt.Println("TestRun6 Begin")
	start := time.Now()
	err := Run6("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun6 总执行时间（秒）：", time.Since(start).Seconds())
}
