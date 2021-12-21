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

	body, err := se.REST(http.MethodGet, path, nil)
	assert.Nil(t, err)
	assert.Equal(t, *body, resp)
}

func TestRESTERROR(t *testing.T) {
	//TODO: change to suit test
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	host := "127.0.0.1"
	path := "/success"
	url := "https://" + host + path
	resp := []byte("Not Found")
	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewBytesResponder(404, resp))

	username := "admin"
	password := "password"
	inscure := true
	se := InitSession(host, username, password, inscure)
	se.Client = &http.Client{}

	body, err := se.REST(http.MethodGet, path, nil)
	// TODO check the content
	assert.NotNil(t, err)
	assert.Nil(t, body)
}
