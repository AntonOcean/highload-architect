package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	amqp "kek/internal/amqp"

	"github.com/google/uuid"

	"kek/internal/api/v1/formatter"
	"kek/internal/domain"
)

func (uc uc) CreatePost(ctx context.Context, text string, authorID uuid.UUID) (*domain.Post, error) {
	_, err := uc.serviceRepo.GetUserByID(ctx, authorID)
	if err != nil {
		if errors.Is(err, formatter.ErrNotFound) {
			return nil, formatter.ErrInvalidData
		}

		return nil, err
	}

	post := &domain.Post{
		ID:       uuid.New(),
		Text:     text,
		AuthorID: authorID,
		Created:  time.Now(),
	}

	err = uc.serviceRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	go func() {
		err := uc.queue.Push(&amqp.Message{
			Data: domain.NewPost{
				PostID:   post.ID,
				AuthorID: authorID,
			},
			Timestamp: time.Now(),
			Key:       string(amqp.PostEvent),
		})
		if err != nil {
			uc.logger.Error(fmt.Sprintf("error push msg new post: %+v", err))
		}
	}()

	return post, nil
}

func (uc uc) GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.Post, error) {
	return uc.serviceRepo.GetPostByID(ctx, postID)
}

func (uc uc) UpdatePost(ctx context.Context, text string, postID, userID uuid.UUID) (*domain.Post, error) {
	p, err := uc.serviceRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	if p.AuthorID != userID {
		return nil, formatter.ErrPermissionDenied
	}

	post := &domain.Post{
		ID:       p.ID,
		Text:     text,
		AuthorID: p.AuthorID,
	}

	err = uc.serviceRepo.UpdatePost(ctx, text, postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (uc uc) DeletePostByID(ctx context.Context, postID, userID uuid.UUID) error {
	p, err := uc.serviceRepo.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, formatter.ErrNotFound) {
			return nil
		}

		return err
	}

	if p.AuthorID != userID {
		return formatter.ErrPermissionDenied
	}

	return uc.serviceRepo.DeletePostByID(ctx, postID)
}

func (uc uc) GetPostsByAuthorID(ctx context.Context, authorID uuid.UUID) ([]*domain.Post, error) {
	return uc.serviceRepo.GetPostsByAuthorID(ctx, authorID)
}
