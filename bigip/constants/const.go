package constants

type URI string

const (
	UNDEFINED      URI = ""
	LTM            URI = "/mgmt/tm/ltm"
	VIRTUALADDRESS URI = LTM + "/virtual-address/"
)
