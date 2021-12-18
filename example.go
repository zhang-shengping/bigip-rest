package main

import (
	"log"

	"github.com/zhang-shengping/bigiprest/bigip"
)

func main() {
	host := "10.123.123.98"
	username := "admin"
	password := "admin"
	insecure := true

	// get a virtual address service
	vipserv := bigip.NewVirtualAddressServ(
		bigip.InitSession(host, username, password, insecure),
	)

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	addrname := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"

	// get a virtual address
	virtualaddr := vipserv.GetVirtualAddress(
		partition,
		addrname,
	)
	log.Printf("virtual address is %s", virtualaddr)

	// get a virtual addresses of a partition
	virtualaddrs := vipserv.GetVirtualAddresses(
		partition,
	)
	log.Printf("virtual address in %s is %s", partition, virtualaddrs.Items)

	virtualaddr = vipserv.PatchVritualAddress(
		partition,
		addrname,
		&bigip.VirtualAddress{
			Description: "example",
		},
	)
	log.Printf("Patched virtual address is %s", virtualaddr)
}
