package gonetworkmanager

import "github.com/godbus/dbus/v5"

const (
	DeviceModemInterface = DeviceInterface + ".Modem"

	// Properties
	DeviceModemPropertyModemCapabilities   = DeviceModemInterface + ".ModemCapabilities"   // readable	u
	DeviceModemPropertyCurrentCapabilities = DeviceModemInterface + ".CurrentCapabilities" // readable	u
	DeviceModemPropertyDeviceID            = DeviceModemInterface + ".DeviceId"            // readable 	s
	DeviceModemPropertyOperatorCode        = DeviceModemInterface + ".OperatorCode"        // readable	s
	DeviceModemPropertyAPN                 = DeviceModemInterface + ".Apn"                 // readable 	s
)

type DeviceModem interface {
	Device

	// The generic family of access technologies the modem supports.
	GetPropertyModemCapabilities() (uint32, error)

	// The generic family of access technologies the modem currently supports without a firmware reload or reinitialization.
	GetPropertyCurrentCapabilities() (uint32, error)

	// An identifier used by the modem backend (ModemManager) that aims to uniquely identify the a device.
	GetPropertyDeviceID() (string, error)

	// The MCC and MNC (concatenated) of the network the modem is connected to.
	GetPropertyOperatorCode() (string, error)

	// The access point name the modem is connected to.
	GetPropertyAPN() (string, error)
}

func NewDeviceModem(objectPath dbus.ObjectPath) (DeviceModem, error) {
	var d deviceModem
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceModem struct {
	device
}

func (d *deviceModem) GetPropertyModemCapabilities() (uint32, error) {
	return d.getUint32Property(DeviceModemPropertyModemCapabilities)
}

func (d *deviceModem) GetPropertyCurrentCapabilities() (uint32, error) {
	return d.getUint32Property(DeviceModemPropertyCurrentCapabilities)
}

func (d *deviceModem) GetPropertyDeviceID() (string, error) {
	return d.getStringProperty(DeviceModemPropertyDeviceID)
}

func (d *deviceModem) GetPropertyOperatorCode() (string, error) {
	return d.getStringProperty(DeviceModemPropertyOperatorCode)
}

func (d *deviceModem) GetPropertyAPN() (string, error) {
	return d.getStringProperty(DeviceModemPropertyAPN)
}
