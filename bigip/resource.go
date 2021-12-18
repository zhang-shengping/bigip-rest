package bigip

import (
	"bytes"
	"encoding/json"
	"log"
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
		"/mgmt/tm/ltm/virtual-address/", se}

}

func (s *Service) GetVirtualAddress(partition, name string) *VirtualAddress {
	// create a new object
	addr := new(VirtualAddress)
	err := json.Unmarshal(*s.GetResource(partition, name), addr)
	if err != nil {
		log.Panic("GetVirtualAddress: Can not unmarshal VirtualAddress")
	}
	return addr
}

func (s *Service) GetVirtualAddresses(partition string) *VirtualAddresses {
	// create a new object
	addrs := new(VirtualAddresses)
	err := json.Unmarshal(*s.GetResources(partition), addrs)
	if err != nil {
		log.Panic("GetVirtualAddresses: Can not unmarshal VirtualAddresses")
	}
	return addrs
}

func (s *Service) PatchVritualAddress(partition, name string, body *VirtualAddress) *VirtualAddress {
	addr := new(VirtualAddress)
	data, err := json.Marshal(body)
	if err != nil {
		log.Panic("Can not marshal body", body)
	}
	raw := []byte(data)
	err = json.Unmarshal(*s.PatchResource(partition, name, bytes.NewBuffer(raw)), addr)
	if err != nil {
		log.Panic("PatchVritualAddress: Can not unmarshal VirtualAddress")
	}
	return addr
}

// selfip

// snatpool

// snatmember
