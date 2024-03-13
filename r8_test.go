package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun8(t *testing.T) {
	fmt.Println("TestRun8 Begin")
	start := time.Now()
	err := Run8("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun8 总执行时间（秒）：", time.Since(start).Seconds())
}
