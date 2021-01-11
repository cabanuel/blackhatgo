package main
import (
	"fmt"
	"net"
)
func main() {
// TCP Scanning
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("192.168.1.184:%d", i)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Println("Connection successful on port: ",i)
		} else {
			fmt.Println("Connection failure on port: ", i)
			continue
		}
		conn.Close()
	}
	// ports := make(chan int, 100)
	// fmt.Println(&ports)	
}
