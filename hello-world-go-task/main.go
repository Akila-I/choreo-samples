package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	// Print time, sleep 0.5 seconds, repeat 10,000 times
	for j := 0; j < 10000; j++ {
		fmt.Printf("Time %d: %s\n", j+1, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(500 * time.Millisecond)
		// Print a long log to test log retrival time
		fmt.Println(`This is a comprehensive logging test string designed to evaluate the performance and functionality of logging APIs. It contains multiple sentences with various punctuation marks, numbers like 12345, special characters such as @#$%^&*(), and unicode characters like émojis 🚀🔥💻. The string includes line breaks and different formatting elements to simulate real-world logging scenarios where applications might need to log complex data structures, error messages, user interactions, system events, and debugging information. Here are some sample log levels and messages: [INFO] Application started successfully at 2024-01-15T10:30:45Z, [WARNING] Memory usage is approaching 85% threshold, [ERROR] Database connection failed with timeout after 30 seconds, [DEBUG] Processing user request with ID abc-123-def-456, method POST, payload size 2048 bytes. This extended text also includes JSON-like structures {"userId": 12345, "action": "login", "timestamp": "2024-01-15T10:31:00Z", "metadata": {"ip": "192.168.1.100", "userAgent": "Mozilla/5.0"}}, XML-like tags <event><type>user_action</type><details>Button clicked</details></event>, and various other data formats commonly found in application logs. The purpose is to stress test the logging system with a substantial amount of text that contains diverse character sets, formatting, and realistic log message patterns that would be encountered in production environments.`)
	}
	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}
}
