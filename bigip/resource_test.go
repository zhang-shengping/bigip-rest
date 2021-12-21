/*
Better than nothing.
Some basic tests ensure the interface barely work.
*/
package bigip

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewVirtualAddressServ(t *testing.T) {
	host := "192.168.123.123"
	username := "admin"
	password := "password"
	insecure := true

	virtualaddr := NewVirtualAddressServ(
		InitSession(host, username, password, insecure),
	)

	assert.Equal(t, string(virtualaddr.Path), "/mgmt/tm/ltm/virtual-address/")
}

func TestGetVirtualAddress(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	insecure := true

	serv := NewVirtualAddressServ(
		InitSession(host, username, password, insecure),
	)
	serv.Session.Client = &http.Client{}

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	url := "https://" + host + serv.URIForName(partition, name)

	// a fixture
	resp := []byte(
		`{"kind":"tm:ltm:virtual-address:virtual-addressstate",` +
			`"name":"Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"partition":"Project_346052548d924ee095b3c2a4f05244ac",` +
			`"fullPath":"/Project_346052548d924ee095b3c2a4f05244ac/Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"generation":374,"selfLink":"https://localhost/mgmt/tm/ltm/virtual-address/` +
			`~Project_346052548d924ee095b3c2a4f05244ac~Project_f6638d02-29f8-41aa-9433-179bf49f5fbd?ver=15.0.1.4",` +
			`"address":"172.168.1.6","arp":"enabled","autoDelete":"false","connectionLimit":0,` +
			`"description":"test1:","enabled":"yes","floating":"enabled","icmpEcho":"enabled",` +
			`"inheritedTrafficGroup":"true","mask":"255.255.255.255","routeAdvertisement":"disabled",` +
			`"serverScope":"any","spanning":"disabled","trafficGroup":"/Common/traffic-group-1",` +
			`"trafficGroupReference":{"link":"https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=15.0.1.4"},"unit":1}`,
	)

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewBytesResponder(200, resp))

	result, err := serv.GetVirtualAddress(partition, name)

	expect := &VirtualAddress{
		Name:        name,
		Partition:   partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "test1:",
	}

	assert.Nil(t, err)
	assert.Equal(t, result, expect)

}

func TestGetVirtualAddresses(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	insecure := true

	serv := NewVirtualAddressServ(
		InitSession(host, username, password, insecure),
	)
	serv.Session.Client = &http.Client{}

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	url := "https://" + host + serv.URIForPartition(partition)
	// a fixture
	resp := []byte(
		`{"kind":"tm:ltm:virtual-address:virtual-addresscollectionstate",` +
			`"selfLink":"https://localhost/mgmt/tm/ltm/virtual-address?$` +
			`filter=partition+eq+Project_346052548d924ee095b3c2a4f05244ac&ver=15.0.1.4",` +
			`"items":[{"kind":"tm:ltm:virtual-address:virtual-addressstate","name":"Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"partition":"Project_346052548d924ee095b3c2a4f05244ac",` +
			`"fullPath":"/Project_346052548d924ee095b3c2a4f05244ac/Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"generation":374,"selfLink":"https://localhost/mgmt/tm/ltm/virtual-address/~` +
			`Project_346052548d924ee095b3c2a4f05244ac~Project_f6638d02-29f8-41aa-9433-179bf49f5fbd?ver=15.0.1.4",` +
			`"address":"172.168.1.6","arp":"enabled","autoDelete":"false","connectionLimit":0,` +
			`"description":"test1:","enabled":"yes","floating":"enabled","icmpEcho":"enabled",` +
			`"inheritedTrafficGroup":"true","mask":"255.255.255.255","routeAdvertisement":"disabled",` +
			`"serverScope":"any","spanning":"disabled","trafficGroup":"/Common/traffic-group-1",` +
			`"trafficGroupReference":{"link":"https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=15.0.1.4"},"unit":1}]}`,
	)

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewBytesResponder(200, resp))

	result, err := serv.GetVirtualAddresses(partition)

	expect := &VirtualAddresses{
		Items: []VirtualAddress{
			{
				Name:        name,
				Partition:   partition,
				Address:     "172.168.1.6",
				Mask:        "255.255.255.255",
				Description: "test1:",
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, result, expect)
}

func TestPatchVirtualAddress(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	insecure := true

	serv := NewVirtualAddressServ(
		InitSession(host, username, password, insecure),
	)
	serv.Session.Client = &http.Client{}

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	url := "https://" + host + serv.URIForName(partition, name)
	body := VirtualAddress{
		Description: "it is ok",
	}

	// a fixture
	resp := []byte(
		`{"kind":"tm:ltm:virtual-address:virtual-addressstate",` +
			`"name":"Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"partition":"Project_346052548d924ee095b3c2a4f05244ac",` +
			`"fullPath":"/Project_346052548d924ee095b3c2a4f05244ac/Project_f6638d02-29f8-41aa-9433-179bf49f5fbd",` +
			`"generation":374,"selfLink":"https://localhost/mgmt/tm/ltm/virtual-address/` +
			`~Project_346052548d924ee095b3c2a4f05244ac~Project_f6638d02-29f8-41aa-9433-179bf49f5fbd?ver=15.0.1.4",` +
			`"address":"172.168.1.6","arp":"enabled","autoDelete":"false","connectionLimit":0,` +
			`"description":"it is ok","enabled":"yes","floating":"enabled","icmpEcho":"enabled",` +
			`"inheritedTrafficGroup":"true","mask":"255.255.255.255","routeAdvertisement":"disabled",` +
			`"serverScope":"any","spanning":"disabled","trafficGroup":"/Common/traffic-group-1",` +
			`"trafficGroupReference":{"link":"https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=15.0.1.4"},"unit":1}`,
	)

	httpmock.RegisterResponder(http.MethodPatch, url,
		httpmock.NewBytesResponder(200, resp))

	result, err := serv.PatchVritualAddress(partition, name, &body)

	expect := &VirtualAddress{
		Name:        name,
		Partition:   partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "it is ok",
	}

	assert.Nil(t, err)
	assert.Equal(t, result, expect)

}
