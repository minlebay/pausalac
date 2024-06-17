package service

// CreateServiceRequest defines the request payload for creating a service
type CreateServiceRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
	Total    int64  `json:"total" binding:"required"`
}

// UpdateServiceRequest defines the request payload for updating a service
type UpdateServiceRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
	Total    int64  `json:"total" binding:"required"`
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
