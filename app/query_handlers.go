package app

import "context"

const (
	retrieveURLQueryName = "RetrieveURLQueryName"
)

//go:generate moq -out zmock_query_test.go -pkg app_test . Query

// Query defines the interface of the queries to be performed
type Query interface {
	Name() string
}

// QueryResponse defines the response to be received from the QueryHandler
type QueryResponse interface{}

// QueryHandler defines the interface of the handler to run queries
type QueryHandler interface {
	Handle(ctx context.Context, query Query) (QueryResponse, error)
}

// RetrieveURLQuery is a VTO
type RetrieveURLQuery struct {
	Shortened string
}

// Name returns the name of the query to retrieve a URL
func (ruq RetrieveURLQuery) Name() string {
	return retrieveURLQueryName
}

// RetrieveURLHandler is the handler to retrieve a URL
type RetrieveURLHandler struct {
	urlsRepository URLsRepository
}

// NewRetrieveURLHandler is a constructor
func NewRetrieveURLHandler(urlsRepository URLsRepository) RetrieveURLHandler {
	return RetrieveURLHandler{urlsRepository: urlsRepository}
}

// Handle retrieves a URL
func (ruh RetrieveURLHandler) Handle(ctx context.Context, query Query) (QueryResponse, error) {
	retrieveQuery, ok := query.(RetrieveURLQuery)
	if !ok {
		return nil, InvalidQueryError{expected: RetrieveURLQuery{}, received: query}
	}

	return ruh.urlsRepository.FindByShortened(ctx, retrieveQuery.Shortened)
}
