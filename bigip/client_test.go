package bigip

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

type ClientTestSuit struct {
	suite.Suite

	host     string
	username string
	password string
	inscure  bool

	partition string
	name      string
}

func (suite *ClientTestSuit) SetupTest() {
	suite.host = "bigip.com"
	suite.username = "admin"
	suite.password = "password"
	suite.inscure = true

	suite.partition = "Project_346052548d924ee095b3c2a4f05244ac"
	suite.name = "Project_f6638d02-29f8-41aa-9433-179bf49f5fbd"

	httpmock.Activate()
}

func (suit *ClientTestSuit) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *ClientTestSuit) TestGetResource() {

	se := InitSession(suite.host, suite.username, suite.password, suite.inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	url := "https://" + suite.host + string(constants.VIRTUALADDRESS) + service.URIForName(
		suite.partition, suite.name)

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
		Name:        suite.name,
		Partition:   suite.partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "test1:",
	}

	addr := new(VirtualAddress)
	err := service.GetResource(suite.partition, suite.name, addr)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expect, addr)
}

func (suite *ClientTestSuit) TestGetResources() {

	se := InitSession(suite.host, suite.username, suite.password, suite.inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	url := "https://" + suite.host + string(constants.VIRTUALADDRESS) + service.URIForPartition(
		suite.partition)

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
				Name:        suite.name,
				Partition:   suite.partition,
				Address:     "172.168.1.6",
				Mask:        "255.255.255.255",
				Description: "test1:",
			},
		},
	}
	addrs := new(VirtualAddresses)
	err := service.GetResources(suite.partition, addrs)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expect, addrs)

}

func (suite *ClientTestSuit) TestPatchResource() {

	se := InitSession(suite.host, suite.username, suite.password, suite.inscure)
	se.Client = &http.Client{}
	service := NewService(se)

	url := "https://" + suite.host + string(constants.VIRTUALADDRESS) + service.URIForName(
		suite.partition, suite.name)
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
		Name:        suite.name,
		Partition:   suite.partition,
		Address:     "172.168.1.6",
		Mask:        "255.255.255.255",
		Description: "it is ok",
	}

	addr := VirtualAddress{
		Description: "it is ok",
	}

	// !!!!! here must be a pointer, can not be a addr type, when function inside is unmarshal
	err := service.PatchResource(suite.partition, suite.name, &addr)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expect, addr)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuit))
}
