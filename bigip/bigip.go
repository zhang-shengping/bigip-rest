package bigip

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"github.com/zhang-shengping/bigiprest/bigip/bigiperrors"
)

type Bigip struct {
	Host     string
	Name     string
	Password string
}

// session layer include TLS transport
type Session struct {
	BigipReq *Bigip
	Client   *http.Client
}

// init Bigip and Session struct for REST network application layer.
// only consider to use name and password to avoid token expired.
func InitSession(host string, username string, password string, insecure bool) *Session {

	return &Session{
		BigipReq: &Bigip{
			Host:     host,
			Name:     username,
			Password: password,
		},
		// for test
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
	}
}

// create application layer actions.
func (se *Session) REST(method, path string, body io.Reader) (*[]byte, error) {
	client := se.Client

	url := "https://" + se.BigipReq.Host + path

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Panic("Create http new request fail")
	}

	log.Printf("Request URL is: %s", url)
	request.SetBasicAuth(se.BigipReq.Name, se.BigipReq.Password)

	// TODO: throw error when resp has 404, 501 etc.
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// If the Body is not both read to EOF and closed, the Client's underlying
	// RoundTripper (typically Transport) may not be able to re-use a
	// persistent TCP connection to the server for a subsequent "keep-alive" request
	defer resp.Body.Close()

	err = checkstatus(resp)
	if err != nil {
		return nil, err
	}

	ioBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Read Body error")
	}

	log.Printf("Response Body is: %s", ioBody)
	return &ioBody, err
}

func checkstatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return bigiperrors.ResponseError{
			Resp: resp,
		}
	}
	return nil
}
