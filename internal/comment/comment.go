package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("failed to fetchcomment by id")
	ErrNotImplemented    = errors.New("Not implemented")
)

// comment structure for our
//
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

//Store - this interface defines all of the methods
// that our service in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// service - is the struct on witch all our
// logic will be built on top of
type Service struct {
	Store Store // alternative *db.Repository
}

// New Service - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	// calls store passing in the context
	cmt, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		log.Errorf("an error occured fetching the comment: %s", err.Error())
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		log.Errorf("an error occurred adding the comment: %s", err.Error())
	}
	return cmt, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(
	ctx context.Context, ID string, newComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, newComment)
	if err != nil {
		log.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return cmt, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.Store.DeleteComment(ctx, ID)
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}