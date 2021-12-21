package bigip

import (
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

type Resource interface {
	Path() constants.URI
}

type VirtualAddress struct {
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Mask        string `json:"mask,omitempty"`
	Description string `json:"description,omitempty"`
	Partition   string `json:"partition,omitempty"`
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

func (addr VirtualAddress) Path() constants.URI {
	return constants.VIRTUALADDRESS
}

type VirtualAddresses struct {
	Items []VirtualAddress `json:"items"`
}

func (addrs VirtualAddresses) Path() constants.URI {
	return constants.VIRTUALADDRESS
}
