package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := 5 * time.Second
	startTime := time.Now() // Record the start time

	conn, err := net.DialTimeout("tcp", address, timeout)

	// Calculate the response time
	responseTime := time.Since(startTime)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN!] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP :) & RUNNING] %v is reachable,\n From: %v\n To: %v\n Response Time: %v", destination, conn.LocalAddr(), conn.RemoteAddr(), responseTime)
	}

	return status
}
