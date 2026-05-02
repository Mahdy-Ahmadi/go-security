package monitor

import (
    "log"
    "time"

    "github.com/Mahdy-Ahmadi/go-security/internal/scanner"
    "github.com/Mahdy-Ahmadi/go-security/internal/report"
)

type Monitor struct {
    Scanner  *scanner.Scanner
    Interval time.Duration
}

func New(s *scanner.Scanner, interval time.Duration) *Monitor {
    return &Monitor{Scanner: s, Interval: interval}
}

func (m *Monitor) Start() {
    ticker := time.NewTicker(m.Interval)
    for range ticker.C {
        log.Println("Running scheduled scan...")
        results := m.Scanner.ScanAll()
        report.GenerateHTML(results, "monitor_report.html")
        // Trigger alert if critical service is closed
        for _, target := range results {
            for _, port := range target.OpenPorts {
                if port.Port == 22 && !port.Open {
                    log.Printf("ALERT: SSH closed on %s", target.Host)
                }
            }
        }
    }
}
