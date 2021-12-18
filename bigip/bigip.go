package bigip

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
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

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
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
func (se *Session) REST(method, path string, body io.Reader) *[]byte {
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
		log.Panic("Response err:", err)
	}
	log.Printf("Response statue is: %s", resp.Status)
	// If the Body is not both read to EOF and closed, the Client's underlying
	// RoundTripper (typically Transport) may not be able to re-use a
	// persistent TCP connection to the server for a subsequent "keep-alive" request
	defer resp.Body.Close()

	ioBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Read Body error")
	}

	log.Printf("Response Body is: %s", ioBody)
	return &ioBody
}

type Service struct {
	Path string
	// URL     string
	Session *Session
	// Fields *Fields
}

func (s *Service) URIForName(partition, name string) string {
	return s.Path + "~" + partition + "~" + name
}

func (s *Service) URIForPartition(partition string) string {
	return s.Path + "?" +
		url.PathEscape("$filter=partition eq "+partition)
}

func (s *Service) GetResource(partition, name string) *[]byte {
	log.Println("Get Reousrce")
	return s.Session.REST(http.MethodGet, s.URIForName(partition, name), nil)
}

func (s *Service) GetResources(partition string) *[]byte {
	log.Println("Get Reousrces")
	return s.Session.REST(http.MethodGet, s.URIForPartition(partition), nil)
}

func (s *Service) PatchResource(partition, name string, body io.Reader) *[]byte {
	log.Println("Patch Reousrce")
	return s.Session.REST(http.MethodPatch, s.URIForName(partition, name), body)
}

// type Fields struct {
// 	Partition            string
// 	Name                 string
// 	ExpandSubcollections bool
// 	Options
// }

// cannot use interface Resource pointer as a parameter.
// func ConvertByteBody(body *[]byte, res Resource) {
// 	json.Unmarshal(*body, res)
// }

// cannot pass a interface pointer to func
// it will raise error
// func (se *Session) Get(sev *ResourceServ) *[]byte {
// 	sev.GetResourceURI()
// }
