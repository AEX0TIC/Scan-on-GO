package scanner

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func Scan(host string, start, end, timeout int, savePath string) {
	const maxConcurrency = 100
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []string

	fmt.Printf("📡 Scanning %s from port %d to %d...\n", host, start, end)

	for port := start; port <= end; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, time.Millisecond*time.Duration(timeout))
			if err == nil {
				msg := fmt.Sprintf("[+] Open port: %d", p)
				fmt.Println(msg)
				mu.Lock()
				results = append(results, msg)
				mu.Unlock()
				conn.Close()
			}
		}(port)
	}
	wg.Wait()

	if savePath != "" {
		mu.Lock()
		output := ""
		for _, line := range results {
			output += line + "\n"
		}
		mu.Unlock()
		err := os.WriteFile(savePath, []byte(output), 0644)
		if err != nil {
			fmt.Println("❌ Failed to write to file:", err)
		} else {
			fmt.Println("✅ Results saved to", savePath)
		}
	}
	fmt.Println("🔎 Scan complete.")
}
