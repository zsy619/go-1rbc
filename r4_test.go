package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun4(t *testing.T) {
	fmt.Println("TestRun4 Begin")
	start := time.Now()
	err := Run4("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun4 总执行时间（秒）：", time.Since(start).Seconds())
}
