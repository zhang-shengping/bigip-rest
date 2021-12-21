/*
Better than nothing.
Some basic tests ensure the interface barely work.
*/
package bigip

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhang-shengping/bigiprest/bigip/constants"
)

func TestVirutalAddressPath(t *testing.T) {
	addr := new(VirtualAddress)
	assert.Equal(t, constants.VIRTUALADDRESS, addr.Path())
}

func TestVirutalAddressesPath(t *testing.T) {
	addr := new(VirtualAddresses)
	assert.Equal(t, constants.VIRTUALADDRESS, addr.Path())
}
