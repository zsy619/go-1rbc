package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun2(t *testing.T) {
	start := time.Now()
	err := Run2("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("总执行时间（秒）：", time.Since(start).Seconds())
}
