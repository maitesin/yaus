package app

import "context"

const (
	retrieveURLByShortenedQueryName = "RetrieveURLByShortenedQueryName"
	retrieveURLByOriginalQueryName  = "RetrieveURLByOriginalQueryName"
)

//go:generate moq -out zmock_query_test.go -pkg app_test . Query

// Query defines the interface of the queries to be performed
type Query interface {
	Name() string
}

// QueryResponse defines the response to be received from the QueryHandler
type QueryResponse interface{}

//go:generate moq -out ../infra/http/zmock_query_test.go -pkg http_test . QueryResponse QueryHandler

// QueryHandler defines the interface of the handler to run queries
type QueryHandler interface {
	Handle(ctx context.Context, query Query) (QueryResponse, error)
}

// RetrieveURLByShortenedQuery is a VTO
type RetrieveURLByShortenedQuery struct {
	Shortened string
}

// Name returns the name of the query to retrieve a URL by shortened code
func (ruq RetrieveURLByShortenedQuery) Name() string {
	return retrieveURLByShortenedQueryName
}

// RetrieveURLByShortenedHandler is the handler to retrieve a URL by shortened code
type RetrieveURLByShortenedHandler struct {
	urlsRepository URLsRepository
}

// NewRetrieveURLByShortenedHandler is a constructor
func NewRetrieveURLByShortenedHandler(urlsRepository URLsRepository) RetrieveURLByShortenedHandler {
	return RetrieveURLByShortenedHandler{urlsRepository: urlsRepository}
}

// Handle retrieves a URL by the shortened code
func (rsu RetrieveURLByShortenedHandler) Handle(ctx context.Context, query Query) (QueryResponse, error) {
	retrieveQuery, ok := query.(RetrieveURLByShortenedQuery)
	if !ok {
		return nil, InvalidQueryError{expected: RetrieveURLByShortenedQuery{}, received: query}
	}

	return rsu.urlsRepository.FindByShortened(ctx, retrieveQuery.Shortened)
}

// RetrieveURLByOriginalQuery is a VTO
type RetrieveURLByOriginalQuery struct {
	Original string
}

// Name returns the name of the query to retrieve a URL by original URL
func (ruoq RetrieveURLByOriginalQuery) Name() string {
	return retrieveURLByOriginalQueryName
}

// RetrieveURLByOriginalHandler is the handler to retrieve a URL by original URL
type RetrieveURLByOriginalHandler struct {
	urlsRepository URLsRepository
}

// NewRetrieveURLByOriginalHandler is a constructor
func NewRetrieveURLByOriginalHandler(urlsRepository URLsRepository) RetrieveURLByOriginalHandler {
	return RetrieveURLByOriginalHandler{urlsRepository: urlsRepository}
}

// Handle retrieves a URL by the original URL
func (ruoh RetrieveURLByOriginalHandler) Handle(ctx context.Context, query Query) (QueryResponse, error) {
	retrieveQuery, ok := query.(RetrieveURLByOriginalQuery)
	if !ok {
		return nil, InvalidQueryError{expected: RetrieveURLByOriginalQuery{}, received: query}
	}

	return ruoh.urlsRepository.FindByOriginal(ctx, retrieveQuery.Original)
}
