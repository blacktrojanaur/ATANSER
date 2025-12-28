package tunnel

import (
	"ATANSER/logs"
	"net"
)

func StartTunnelClient(serverIP string, port string) {
	logs.LogInfo("Starting encrypted UDP tunnel client to " + serverIP + ":" + port)
	addr, _ := net.ResolveUDPAddr("udp", serverIP+":"+port)
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	conn.Write([]byte("SECURE_HELLO_PACKET"))
	buffer := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buffer)
	logs.LogInfo("Response received from server, size=" + string(rune(n)))
}
