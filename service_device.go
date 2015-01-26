package fda

// DeviceService provides data on devices and device safety event reporting
type DeviceService struct {
	client *Client
}

// EventSearch finds events that have been reported while a patient is using a device.
// It yields a slice of DeviceSafetyReport, which is a structured response providing information
// about the event.
func (s *DeviceService) EventSearch(search string, limit, skip int) ([]DeviceSafetyReport, *Meta, error) {
	data := []DeviceSafetyReport{}

	meta, err := s.client.search("/device/event", search, limit, skip, &data)
	return data, meta, err
}

// EventCount counts device events based on parameters you pass
// It expects you to formulate the correct structure and pass it as *data to be modified.
func (s *DeviceService) EventCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/device/event", search, count, limit, &data)
	return meta, err
}

// EnforcementSearch finds reports of device enforcement actions
func (s *DeviceService) EnforcementSearch(search string, limit, skip int) ([]EnforcementReport, *Meta, error) {
	data := []EnforcementReport{}

	meta, err := s.client.search("/device/enforcement", search, limit, skip, &data)
	return data, meta, err
}

// EnforcementCount counts device enforcement actions based on parameters you pass
// It expects you to formulate the correct structure and pass it as *data to be modified.
func (s *DeviceService) EnforcementCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/device/enforcement", search, count, limit, &data)
	return meta, err
}
