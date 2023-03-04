package adapters

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"

	"ws-feed/internal/adapters/formatter"
	"ws-feed/internal/domain"
)

var ErrInvalidBackendServiceResponse = errors.New("can't get response from backend: status code != 200")

type BackendService struct {
	httpClient *resty.Client
}

func NewBackendService(httpClient *resty.Client) *BackendService {
	return &BackendService{
		httpClient: httpClient,
	}
}

func (b *BackendService) GetFriendWithUserID(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	var users []formatter.UserResponse

	httpRequest := b.httpClient.R().
		SetContext(ctx).
		SetResult(&users).
		SetPathParam("userID", userID.String())

	httpResponse, err := httpRequest.Get("/api/v1/admin/user/{userID}/friend-with")

	if err != nil {
		return nil, fmt.Errorf("can't get from backend - friends by user_id: %w", err)
	}

	if httpResponse.StatusCode() != http.StatusOK {
		return nil, ErrInvalidBackendServiceResponse
	}

	return formatter.ToDomainUserIDList(users), nil
}

func (b *BackendService) GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.Post, error) {
	var post formatter.PostResponse

	httpRequest := b.httpClient.R().
		SetContext(ctx).
		SetResult(&post).
		SetPathParam("postID", postID.String())

	httpResponse, err := httpRequest.Get("/api/v1/admin/post/{postID}")

	if err != nil {
		return nil, fmt.Errorf("can't get from backend - post by id: %w", err)
	}

	if httpResponse.StatusCode() != http.StatusOK {
		return nil, ErrInvalidBackendServiceResponse
	}

	return post.ToDomain(), nil
}
