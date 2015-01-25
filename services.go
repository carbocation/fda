package fda

// DeviceService provides data on devices and device safety event reporting
type DeviceService struct {
	client *Client
}

// FoodService provides data on foods and food safety event reporting
type FoodService struct {
	client *Client
}
