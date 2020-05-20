package main

import (
	"os"
	"bufio"
	"fmt"
	"net"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var host = ""

	for sc.Scan() {
		host = sc.Text()

		ips,err := net.LookupIP(host)

		if err != nil {
			continue
		}

		for _, byte_ip := range ips {
			var ip = byte_ip.String()

			fmt.Println(host + "\t" +ip)
		}
	}
}
