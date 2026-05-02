# Go Security - Network Vulnerability Scanner

[![Go Report Card](https://goreportcard.com/badge/github.com/Mahdy-Ahmadi/go-security)](https://goreportcard.com/report/github.com/Mahdy-Ahmadi/go-security)

**Professional security tool** written in Go for port scanning, service detection, and basic vulnerability assessment.

### Installation
```bash
git clone https://github.com/Mahdy-Ahmadi/go-security.git
cd go-security
go mod download
```

# Usage
```bash
# One-time scan
go run cmd/scanner/main.go -config config.yaml

# Continuous monitoring
go run cmd/scanner/main.go -mode monitor
```
# Features
- TCP port scanning

- Service banner grabbing

- Default credential check

- Real-time monitoring

- HTML report generation
