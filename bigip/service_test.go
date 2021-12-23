package bigip

import (
	"net/http"
	"testing"

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
