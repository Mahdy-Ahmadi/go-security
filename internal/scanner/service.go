package scanner

import (
    "fmt"
    "net"
    "time"
)

func GrabBanner(host string, port int, timeout time.Duration) string {
    conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
    if err != nil {
        return ""
    }
    defer conn.Close()
    
    conn.SetReadDeadline(time.Now().Add(timeout))
    buf := make([]byte, 256)
    n, _ := conn.Read(buf)
    if n > 0 {
        return string(buf[:n])
    }
    return ""
}
