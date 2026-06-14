package ssdp

import (
	"strings"
	"testing"
)

func TestBuildMediaServerResponse(t *testing.T) {
	response := BuildMediaServerResponse()

	requiredParts := []string{
		"HTTP/1.1 200 OK\r\n",
		"ST: urn:schemas-upnp-org:device:MediaServer:1\r\n",
		"USN: uuid:4d696e69-444c-164e-9d41-dca63226c763::urn:schemas-upnp-org:device:MediaServer:1\r\n",
		"LOCATION: http://192.168.3.100:8300/rootDesc.xml\r\n",
		"\r\n",
	}

	for _, part := range requiredParts {
		if !strings.Contains(response, part) {
			t.Errorf("response does not contain %q", part)
		}
	}

	if !strings.HasSuffix(response, "\r\n\r\n") {
		t.Errorf("response must end with CRLF CRLF")
	}
}
