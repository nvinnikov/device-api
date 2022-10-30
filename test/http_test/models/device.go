package models

// CreateDeviceResponse is the response body from the CreateDevice endpoint
type CreateDeviceResponse struct {
	DeviceID int `json:"deviceId,string"`
}

// DescribeDeviceResponse is the response of DescribeDeviceRequest
type DescribeDeviceResponse struct {
	Value Item `json:"value"`
}

// UpdateDeviceRequest is the request struct for api UpdateDevice
type UpdateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// UpdateDeviceResponse is the response of UpdateDeviceRequest
type UpdateDeviceResponse struct {
	Success bool `json:"success"`
}

// CreateDeviceRequest is the request struct for api CreateDevice
type CreateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// RemovedDevice Response is the response of DeletedDeviceRequest
type RemovedDevice struct {
	Found bool `json:"found"`
}
