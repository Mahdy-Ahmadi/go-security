package main

import (
    "flag"
    "log"
    "time"

    "github.com/Mahdy-Ahmadi/go-security/internal/monitor"
    "github.com/Mahdy-Ahmadi/go-security/internal/report"
    "github.com/Mahdy-Ahmadi/go-security/internal/scanner"
    "gopkg.in/yaml.v3"
    "os"
)

type Config struct {
    Targets []string `yaml:"targets"`
    Ports   []int    `yaml:"ports"`
    Timeout string   `yaml:"timeout"`
    Concurrency int  `yaml:"concurrency"`
    MonitorInterval string `yaml:"monitor_interval"`
}

func main() {
    configPath := flag.String("config", "config.yaml", "path to config file")
    mode := flag.String("mode", "scan", "mode: scan, monitor")
    flag.Parse()

    data, err := os.ReadFile(*configPath)
    if err != nil {
        log.Fatal("Cannot read config:", err)
    }

    var cfg Config
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        log.Fatal("Invalid config:", err)
    }

    timeout, _ := time.ParseDuration(cfg.Timeout)
    monitorInterval, _ := time.ParseDuration(cfg.MonitorInterval)

    s := scanner.New(cfg.Targets, cfg.Ports, timeout, cfg.Concurrency)

    if *mode == "monitor" {
        log.Println("Starting continuous monitor...")
        m := monitor.New(s, monitorInterval)
        m.Start()
    } else {
        log.Println("Running one-time scan...")
        results := s.ScanAll()
        report.GenerateHTML(results, "report.html")
        log.Println("Report saved to report.html")
    }
}
