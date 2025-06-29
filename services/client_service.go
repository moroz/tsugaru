package services

import "oauth-provider/db/queries"

type ClientService struct {
	db queries.DBTX
}

func NewClientService(db queries.DBTX) *ClientService {
	return &ClientService{db}
}
