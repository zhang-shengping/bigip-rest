package bigip

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

func TestNewService(t *testing.T) {
	host := "192.168.123.123"
	username := "admin"
	password := "password"
	insecure := true

	service := NewService(
		InitSession(host, username, password, insecure),
	)

	assert.Equal(t, string(service.Path), "")
	assert.Equal(t, host, service.Session.BigipReq.Host)
	assert.Equal(t, username, service.Session.BigipReq.Name)
	assert.Equal(t, password, service.Session.BigipReq.Password)
	assert.Equal(t, insecure, service.Session.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestURIForName(t *testing.T) {
	path := constants.VIRTUALADDRESS
	partition := "test_partition"
	name := "test_name"
	serv := Service{
		path,
		&Session{},
	}
	expect := string(path) + "~" + partition + "~" + name
	result := serv.URIForName(partition, name)

	assert.Equal(t, result, expect)
}

func TestURIForPartition(t *testing.T) {
	path := constants.VIRTUALADDRESS
	partition := "test_partition"
	serv := Service{
		path,
		&Session{},
	}
	expect := string(path) + "?$filter=partition" + "%20eq%20" + partition
	result := serv.URIForPartition(partition)

	assert.Equal(t, result, expect)
}

func TestGetResource(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	url := "https://" + host + string(constants.VIRTUALADDRESS) + service.URIForName(partition, name)

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
	httpmock.RegisterResponder(
		http.MethodGet, url,
		httpmock.NewBytesResponder(200, resp),
	)

	expect := &VirtualAddress{
		Name:        name,
		Partition:   partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "test1:",
	}
	addr := new(VirtualAddress)
	err := service.GetResource(partition, name, addr)
	assert.Nil(t, err)
	assert.Equal(t, expect, addr)
}

func TestGetResources(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	url := "https://" + host + string(constants.VIRTUALADDRESS) + service.URIForPartition(partition)

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

	httpmock.RegisterResponder(
		http.MethodGet, url, httpmock.NewBytesResponder(200, resp),
	)

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
	addrs := new(VirtualAddresses)
	err := service.GetResources(partition, addrs)
	assert.Nil(t, err)
	assert.Equal(t, expect, addrs)

}

func TestPatchResource(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	name := "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"
	partition := "Project_346052548d924ee095b3c2a4f05244ac"
	url := "https://" + host + string(constants.VIRTUALADDRESS) + service.URIForName(partition, name)
	fmt.Println(url)
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

	httpmock.RegisterResponder(
		http.MethodPatch, url,
		httpmock.NewBytesResponder(200, resp),
	)

	expect := VirtualAddress{
		Name:        name,
		Partition:   partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "it is ok",
	}

	addr := VirtualAddress{
		Description: "it is ok",
	}

	// !!!!! here must be a pointer, can not be a addr type, when function inside is unmarshal
	err := service.PatchResource(partition, name, &addr)

	assert.Nil(t, err)
	assert.Equal(t, expect, addr)
}
