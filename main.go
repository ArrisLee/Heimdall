package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nzgogo/utils/slack"
	"github.com/shirou/gopsutil/load"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	var logFile *os.File
	var err error
	defer logFile.Close()
	filename := "monitor.log"
	logFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
	for {
		time.Sleep(15 * time.Minute)
		cpu, _ := cpu.Percent(0, false)
		if cpu[0] > 30.00 {
			cpuString := fmt.Sprintf("%f", cpu[0])
			slack.CustomizedLog("Heimdall", "CPU Warning", cpuString, "your_slack_incoming_webhook")
		}
		virtualMem, _ := mem.VirtualMemory()
		if virtualMem.UsedPercent > 70.00 {
			UsedPercent := fmt.Sprintf("%f", virtualMem.UsedPercent)
			slack.CustomizedLog("Heimdall", "Virtual Memory Warning", UsedPercent, "your_slack_incoming_webhook")
		}
		load, _ := load.Avg()
		if load.Load5 > 4.90 {
			load := fmt.Sprintf("%f", load.Load5)
			slack.CustomizedLog("Heimdall", "Load Warning", load, "your_slack_incoming_webhook")
		}
		log.Printf("------------- STR -------------\n")
		log.Printf("CPU: %v\n", cpu[0])
		log.Printf("VMemory: %v\n", virtualMem.UsedPercent)
		log.Printf("Load: %v\n", load.Load5)
		log.Printf("------------- END -------------\n\n")
	}

}
