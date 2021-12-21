package bigip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/zhang-shengping/bigiprest/bigip/bigiperrors"
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

type Service struct {
	Path constants.URI
	// URL     string
	Session *Session
	// Fields *Fields
}

func NewService(se *Session) *Service {
	return &Service{
		Session: se,
	}
}

func (s *Service) URIForName(partition, name string) string {
	return string(s.Path) + "~" + partition + "~" + name
}

func (s *Service) URIForPartition(partition string) string {
	return string(s.Path) + "?" +
		url.PathEscape("$filter=partition eq "+partition)
}

// why not to do like this? less code to write
// func (s *Service) GetResource(partition, name string, res *Resouce) *Resource
// func (s *VirtualAddress) GetResource(partition, name string, res *Resouce) *Resource
// Resource could be virutaladdress, snatip etc.

func (s *Service) GetResource(partition, name string, resource Resource) error {
	s.Path = resource.Path()

	data, err := s.Session.REST(http.MethodGet, s.URIForName(partition, name), nil)
	if err != nil {
		return bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not get resource %s from partiton %s\n", name, partition),
			HttpError:     fmt.Sprintf("%s\n", err),
		}
	}

	err = json.Unmarshal(*data, resource)
	return err
}

func (s *Service) GetResources(partition string, resources Resource) error {
	// create a new object
	s.Path = resources.Path()

	data, err := s.Session.REST(http.MethodGet, s.URIForPartition(partition), nil)
	if err != nil {
		return bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not get resources from partiton %s", partition),
			HttpError:     fmt.Sprintf("%s\n", err),
		}
	}

	err = json.Unmarshal(*data, resources)
	return err
}

func (s *Service) PatchResource(partition, name string, resource Resource) error {
	data, err := json.Marshal(resource)
	if err != nil {
		log.Panic("Can not marshal body", resource)
	}

	raw := []byte(data)
	s.Path = resource.Path()
	// err is declared again ?
	resp, err := s.Session.REST(
		http.MethodPatch,
		s.URIForName(partition, name),
		bytes.NewBuffer(raw),
	)
	if err != nil {
		return bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not patch resource: %s of partiton %s with body %v",
				name, partition, resource),
			HttpError: fmt.Sprintf("%s\n", err),
		}
	}
	// fmt.Println(reflect.TypeOf(resource))

	err = json.Unmarshal(*resp, resource)
	return err
}
