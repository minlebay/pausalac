// Package user provides the use case for user
package user

import (
	domainUser "go_gin_api_clean/src/domain/user"
)

// PaginationResultUser is the structure for pagination result of user
type PaginationResultUser struct {
	Data       []domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}