package ssdp

import "testing"

func TestIsSearchRequest(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    bool
	}{
		{
			name: "m-search request",
			message: "M-SEARCH * HTTP/1.1\r\n" +
				"HOST: 239.255.255.250:1900\r\n" +
				"MAN: \"ssdp:discover\"\r\n" +
				"ST: ssdp:all\r\n\r\n",
			want: true,
		},
		{
			name:    "notify is not search",
			message: "NOTIFY * HTTP/1.1\r\n\r\n",
			want:    false,
		},
		{
			name:    "garbage",
			message: "WTF!?",
			want:    false,
		},
	}

	for _, tt := range tests {
		got := IsSearchRequest(tt.message)

		if got != tt.want {
			t.Errorf("%s: got %t, want %t", tt.name, got, tt.want)
		}
	}
}

func TestIsMediaServerSearch(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    bool
	}{
		{
			name: "media server search",
			message: "M-SEARCH * HTTP/1.1\r\n" +
				"HOST: 239.255.255.250:1900\r\n" +
				"MAN: \"ssdp:discover\"\r\n" +
				"MX: 5\r\n" +
				"ST: urn:schemas-upnp-org:device:MediaServer:1\r\n\r\n",
			want: true,
		},
		{
			name: "dial search is ignored",
			message: "M-SEARCH * HTTP/1.1\r\n" +
				"HOST: 239.255.255.250:1900\r\n" +
				"MAN: \"ssdp:discover\"\r\n" +
				"MX: 1\r\n" +
				"ST: urn:dial-multiscreen-org:service:dial:1\r\n\r\n",
			want: false,
		},
		{
			name: "sat ip search is ignored",
			message: "M-SEARCH * HTTP/1.1\r\n" +
				"HOST: 239.255.255.250:1900\r\n" +
				"MAN: \"ssdp:discover\"\r\n" +
				"MX: 5\r\n" +
				"ST: urn:ses-com:device:SatIPServer:1\r\n\r\n",
			want: false,
		},
		{
			name:    "notify is ignored",
			message: "NOTIFY * HTTP/1.1\r\n\r\n",
			want:    false,
		},
	}

	for _, tt := range tests {
		got := IsMediaServerSearch(tt.message)

		if got != tt.want {
			t.Errorf("%s: got %t, want %t", tt.name, got, tt.want)
		}
	}
}
