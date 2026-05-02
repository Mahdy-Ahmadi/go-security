package scanner

import "strings"

func CheckDefaultCreds(service, banner string) bool {
    lowBanner := strings.ToLower(banner)
    if service == "ssh" && strings.Contains(lowBanner, "dropbear") {
        return true // potential default creds
    }
    return false
}
