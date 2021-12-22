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
	serv := bigip.NewService(
		bigip.InitSession(host, username, password, insecure),
	)

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	addrname := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	// addrname := "Project_f6638d02-29f8-41aa-9433-179bf49f5123"

	addr := new(bigip.VirtualAddress)
	// get a virtual address
	err := serv.GetResource(
		partition,
		addrname,
		addr,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("virtual address is %s", addr)

	addr.Description = "example"
	err = serv.PatchResource(
		partition,
		addrname,
		addr,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Patched virtual address is %s", addr)

	addrs := new(bigip.VirtualAddresses)
	// get a virtual addresses of a partition
	err = serv.GetResources(
		partition,
		addrs,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("virtual address in %s is %s", partition, addrs)

}
