package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("failed to fetchcomment by id")
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

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil
}
