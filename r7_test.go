package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRun7(t *testing.T) {
	fmt.Println("TestRun7 Begin")
	start := time.Now()
	err := Run7("./data/weather_stations.csv", os.Stdout)
	if err != nil {
		t.Fail()
	} else {
		t.Log("ok")
	}
	fmt.Println("TestRun7 总执行时间（秒）：", time.Since(start).Seconds())
}
