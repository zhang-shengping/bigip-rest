/*
Better than nothing.
Some basic tests ensure the interface barely work.
*/
package bigip

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestInitSession(t *testing.T) {
	host := "192.168.123.123"
	username := "admin"
	password := "password"
	insecure := true

	result := InitSession(host, username, password, insecure)

	assert.Equal(t, result.BigipReq.Host, host)
	assert.Equal(t, result.BigipReq.Name, username)
	assert.Equal(t, result.BigipReq.Password, password)
	// Assert Client Transport RoundTripper interface is a http.Transport
	assert.Equal(t, result.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify, insecure)
}

func TestREST(t *testing.T) {
	//TODO: change to suit test
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "127.0.0.1"
	path := "/success"
	url := "https://" + host + path
	resp := []byte("sccuess")
	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewBytesResponder(200, resp))

	username := "admin"
	password := "password"
	inscure := true
	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}

	body := se.REST(http.MethodGet, path, nil)
	assert.Equal(t, *body, resp)
}

func TestURIForName(t *testing.T) {
	path := "/test_path/"
	partition := "test_partition"
	name := "test_name"
	serv := Service{
		path,
		&Session{},
	}
	expect := path + "~" + partition + "~" + name
	result := serv.URIForName(partition, name)

	assert.Equal(t, result, expect)
}

func TestURIForPartition(t *testing.T) {
	path := "/test_path/"
	partition := "test_partition"
	serv := Service{
		path,
		&Session{},
	}
	expect := path + "?$filter=partition" + "%20eq%20" + partition
	result := serv.URIForPartition(partition)

	assert.Equal(t, result, expect)
}

func TestGetResource(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	path := "/ltm/test_resource/"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}

	serv := Service{
		Path:    path,
		Session: se,
	}
	partition := "Project_test123"
	name := "test"
	url := "https://" + host + serv.URIForName(partition, name)

	resp := []byte("sccuess")
	httpmock.RegisterResponder(
		http.MethodGet, url,
		httpmock.NewBytesResponder(200, resp),
	)

	body := serv.GetResource(partition, name)

	assert.Equal(t, *body, resp)
}

func TestGetResources(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	path := "/ltm/test_resources/"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}
	serv := Service{
		Path:    path,
		Session: se,
	}
	// serv.Session.Client = &http.Client{}
	partition := "Project_test"
	url := "https://" + host + serv.URIForPartition(partition)

	resp := []byte("hello")

	httpmock.RegisterResponder(
		http.MethodGet, url, httpmock.NewBytesResponder(200, resp),
	)

	body := serv.GetResources(partition)

	assert.Equal(t, *body, resp)

}

func TestPatchResource(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "bigip.com"
	path := "/ltm/test_resource/"
	username := "admin"
	password := "password"
	inscure := true

	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}

	serv := Service{
		Path:    path,
		Session: se,
	}
	partition := "Project_test123"
	name := "test"
	url := "https://" + host + serv.URIForName(partition, name)

	resp := []byte("sccuess")
	httpmock.RegisterResponder(
		http.MethodPatch, url,
		httpmock.NewBytesResponder(200, resp),
	)

	body := serv.PatchResource(partition, name, bytes.NewBuffer(resp))

	assert.Equal(t, *body, resp)
}
