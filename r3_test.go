package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun3(t *testing.T) {
	fmt.Println("TestRun3 Begin")
	start := time.Now()
	err := Run3("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun3 总执行时间（秒）：", time.Since(start).Seconds())
}
