package lockdown

// #cgo pkg-config: libimobiledevice-1.0
// #include <stdlib.h>
// #include <libimobiledevice/lockdown.h>
import "C"

import (
	"unsafe"

	"github.com/pauldotknopf/goidevice/idevice"
	"github.com/pauldotknopf/goidevice/common"
)

// Client is a lockdown client
type Client interface {
	Type() (string, error)
	Pair() error
	DeviceName() (string, error)
	Close() error
}

type client struct {
	p C.lockdownd_client_t
}

// NewClient creates a new lockdown client
func NewClient(device idevice.Device, label string) (Client, error) {
	labelC := C.CString(label)
	defer C.free(unsafe.Pointer(labelC))

	var p C.lockdownd_client_t
	err := common.ResultToError(C.lockdownd_client_new((C.idevice_t)(idevice.GetPointer(device)), &p, labelC))
	if err != nil {
		return nil, err
	}
	return &client{p}, nil
}

// NewClientWithHandshake creates a new lockdown with a handshake
func NewClientWithHandshake(device idevice.Device, label string) (Client, error) {
	labelC := C.CString(label)
	defer C.free(unsafe.Pointer(labelC))

	var p C.lockdownd_client_t
	err := common.ResultToError(C.lockdownd_client_new_with_handshake((C.idevice_t)(idevice.GetPointer(device)), &p, labelC))
	if err != nil {
		return nil, err
	}
	return &client{p}, nil
}

func (s *client) Type() (string, error) {
	var p *C.char
	err := common.ResultToError(C.lockdownd_query_type(s.p, &p))
	var result string
	if p != nil {
		result = C.GoString(p)
		C.free(unsafe.Pointer(p))
	}
	return result, err
}

func (s *client) Pair() error {
	return common.ResultToError(C.lockdownd_pair(s.p, nil))
}

func (s *client) DeviceName() (string, error) {
	var p *C.char
	err := common.ResultToError(C.lockdownd_get_device_name(s.p, &p))
	var result string
	if p != nil {
		result = C.GoString(p)
		C.free(unsafe.Pointer(p))
	}
	return result, err
}

func (s *client) Close() error {
	err := common.ResultToError(C.lockdownd_client_free(s.p))
	if err == nil {
		s.p = nil
	}
	return err
}
