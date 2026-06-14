package ssdp

const mediaServerLocation = "http://192.168.3.100:8300/rootDesc.xml"
const mediaServerUSN = "uuid:4d696e69-444c-164e-9d41-dca63226c763::urn:schemas-upnp-org:device:MediaServer:1"
const mediaServerST = "urn:schemas-upnp-org:device:MediaServer:1"

func BuildMediaServerResponse() string {
	return "HTTP/1.1 200 OK\r\n" +
		"CACHE-CONTROL: max-age=1800\r\n" +
		"ST: " + mediaServerST + "\r\n" +
		"USN: " + mediaServerUSN + "\r\n" +
		"EXT:\r\n" +
		"SERVER: Debian DLNADOC/1.50 UPnP/1.0 MiniDLNA/1.3.3\r\n" +
		"LOCATION: " + mediaServerLocation + "\r\n" +
		"Content-Length: 0\r\n" +
		"\r\n"
}
