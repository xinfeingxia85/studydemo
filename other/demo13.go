package main

import (
	"strings"
)

func main() {
	transferDomainByEnv("coco.sanjieke.cn", "beta")
}

// transferDomainByEnv
func transferDomainByEnv(domain string, env string) string {
	domainSlice := strings.Split(domain, ".")
	if env == "beta" || env == "online" {

	} else {
		return domain
	}
	return ""
}
