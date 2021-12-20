package main

import (
	"log"

	"github.com/zhang-shengping/bigiprest/bigip"
)

func main() {
	host := "10.145.75.98"
	username := "admin"
	password := "admin@F5"
	insecure := true

	// get a virtual address service
	vipserv := bigip.NewVirtualAddressServ(
		bigip.InitSession(host, username, password, insecure),
	)

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	// addrname := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	addrname := "Project_f6638d02-29f8-41aa-9433-179bf49f5123"

	// get a virtual address
	virtualaddr, err := vipserv.GetVirtualAddress(
		partition,
		addrname,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("virtual address is %s", virtualaddr)

	// get a virtual addresses of a partition
	virtualaddrs, err := vipserv.GetVirtualAddresses(
		partition,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("virtual address in %s is %s", partition, virtualaddrs.Items)

	virtualaddr, err = vipserv.PatchVritualAddress(
		partition,
		addrname,
		&bigip.VirtualAddress{
			Description: "example",
		},
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Patched virtual address is %s", virtualaddr)
}
