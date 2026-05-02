package report

import (
    "os"
    "text/template"
)

type ScanResult struct {
    Host      string
    OpenPorts []PortInfo
}

type PortInfo struct {
    Port    int
    Service string
}

func GenerateHTML(results []ScanResult, filename string) error {
    const tmpl = `
    <html><head><title>Security Scan Report</title></head>
    <body><h1>Scan Report</h1>
    {{range .}}
        <h2>Host: {{.Host}}</h2>
        <ul>
        {{range .OpenPorts}}<li>Port {{.Port}}: {{.Service}}</li>{{end}}
        </ul>
    {{end}}
    </body></html>`
    
    t := template.Must(template.New("report").Parse(tmpl))
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    return t.Execute(f, results)
}
