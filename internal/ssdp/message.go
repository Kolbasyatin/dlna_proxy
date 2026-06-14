package ssdp

import "strings"

var supportedSearchTargets = []string{
	"urn:schemas-upnp-org:device:MediaServer:1",
}

func IsSearchRequest(message string) bool {
	return strings.HasPrefix(message, "M-SEARCH * HTTP/1.1")
}

func IsMediaServerSearch(message string) bool {
	if !IsSearchRequest(message) {
		return false
	}

	for _, target := range supportedSearchTargets {
		if strings.Contains(message, "ST: "+target) {
			return true
		}
	}

	return false
}
