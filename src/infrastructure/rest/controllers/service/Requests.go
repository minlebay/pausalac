package service

// CreateServiceRequest defines the request payload for creating a service
type CreateServiceRequest struct {
	Author string `json:"-"`
	Name   string `json:"name" binding:"required"`
}

// UpdateServiceRequest defines the request payload for updating a service
type UpdateServiceRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateServiceArrayRequest []CreateServiceRequest

type UpdateServiceArrayRequest []UpdateServiceRequest

func (a UpdateServiceArrayRequest) Equals(b UpdateServiceArrayRequest) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}
