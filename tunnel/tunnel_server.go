package tunnel

import (
	"Atanser/logs"
	"net"
)

func StartTunnelServer(port string) {
	logs.LogInfo("Starting encrypted UDP tunnel server on port " + port)
	addr, _ := net.ResolveUDPAddr("udp", ":"+port)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, client, _ := conn.ReadFromUDP(buffer)
		logs.LogInfo("Packet received from " + client.String() + " size=" + string(rune(n)))
		conn.WriteToUDP([]byte("ACK_ENCRYPTED_"+string(buffer[:n])), client)
	}
}
