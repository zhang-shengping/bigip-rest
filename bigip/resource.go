package bigip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/zhang-shengping/bigiprest/bigip/bigiperrors"
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

type VirtualAddress struct {
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Mask        string `json:"mask,omitempty"`
	Description string `json:"description,omitempty"`
	Partition   string `json:"partition"`
}

// It also could be done as pointer,
// which lets empty value as nil, rather than "".
// It allows to patch "" as Description.
// type VirtualAddress struct {
// 	Name        *string `json:"name,omitempty"`
// 	Address     *string `json:"address,omitempty"`
// 	Mask        *string `json:"mask,omitempty"`
// 	Description *string `json:"description,omitempty"`
// 	Partition   *string `json:"partition"`
// }

type VirtualAddresses struct {
	Items []VirtualAddress `json:"items"`
}

func NewVirtualAddressServ(se *Session) *Service {
	return &Service{
		constants.VIRTUALADDRESS, se}

}

// why not to do like this? less code to write
// func (s *Service) GetResource(partition, name string, res *Resouce) *Resource
// func (s *VirtualAddress) GetResource(partition, name string, res *Resouce) *Resource
// Resource could be virutaladdress, snatip etc.

func (s *Service) GetVirtualAddress(partition, name string) (*VirtualAddress, error) {
	// create a new object
	result := new(VirtualAddress)

	data, err := s.GetResource(partition, name)
	if err != nil {
		e := bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not get resource %s from partiton %s", name, partition),
			HttpError:     err,
		}
		return nil, e
	}

	err = json.Unmarshal(*data, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *Service) GetVirtualAddresses(partition string) (*VirtualAddresses, error) {
	// create a new object
	result := new(VirtualAddresses)
	data, err := s.GetResources(partition)
	if err != nil {
		e := bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not get resources from partiton %s", partition),
			HttpError:     err,
		}
		return nil, e
	}

	err = json.Unmarshal(*data, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *Service) PatchVritualAddress(partition, name string, body *VirtualAddress) (*VirtualAddress, error) {
	result := new(VirtualAddress)
	data, err := json.Marshal(body)
	if err != nil {
		log.Panic("Can not marshal body", body)
	}
	raw := []byte(data)

	// err is declared again ?
	addr, err := s.PatchResource(partition, name, bytes.NewBuffer(raw))
	if err != nil {
		e := bigiperrors.ServiceError{
			ResourceError: fmt.Sprintf("Can not patch resource: %s of partiton %s with body %v",
				name, partition, body),
			HttpError: err,
		}
		return nil, e
	}

	err = json.Unmarshal(*addr, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
