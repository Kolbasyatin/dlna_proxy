package main

import (
	"dlna_proxy/internal/ssdp"
	"fmt"
	"net"
)

func main() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("239.255.255.250"),
		Port: 1900,
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, addr)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 2048)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		message := string(buffer[:n])

		if ssdp.IsMediaServerSearch(message) {
			fmt.Println("from:", remoteAddr)
			fmt.Println(message)

			response := ssdp.BuildMediaServerResponse()

			_, err := conn.WriteToUDP([]byte(response), remoteAddr)
			if err != nil {
				fmt.Println("write error:", err)
				continue
			}
		}
	}

}
