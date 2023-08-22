package domain

import "context"

type ResponseRepository interface {
	SaveResponse(ctx context.Context, response ProxyResponse) error
	GetResponse(ctx context.Context, id string) (ProxyResponse, error)
}
