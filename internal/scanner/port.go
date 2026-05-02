package scanner

import (
    "net"
    "sync"
    "time"
)

type PortResult struct {
    Port    int
    Open    bool
    Service string
}

func ScanPort(host string, port int, timeout time.Duration) PortResult {
    address := net.JoinHostPort(host, string(port))
    conn, err := net.DialTimeout("tcp", address, timeout)
    if err != nil {
        return PortResult{Port: port, Open: false}
    }
    conn.Close()
    return PortResult{Port: port, Open: true, Service: "unknown"}
}
