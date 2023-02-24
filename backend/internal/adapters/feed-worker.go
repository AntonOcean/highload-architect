package adapters

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"

	"kek/internal/adapters/formatter"
	"kek/internal/domain"
)

type FeedWorkerService struct {
	httpClient *resty.Client
}

func NewFeedWorkerService(httpClient *resty.Client) *FeedWorkerService {
	return &FeedWorkerService{
		httpClient: httpClient,
	}
}

var ErrInvalidFeedWorkerServiceResponse = errors.New("can't get response from feed-worker: status code != 200")

func (b *FeedWorkerService) GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error) {
	var posts []formatter.PostResponse

	httpRequest := b.httpClient.R().
		SetContext(ctx).
		SetResult(&posts).
		SetPathParam("userID", userID.String()).
		SetQueryParams(map[string]string{
			"limit":  strconv.Itoa(limit),
			"offset": strconv.Itoa(offset),
		})

	httpResponse, err := httpRequest.Get("/api/v1/admin/user/{userID}/feed")

	if err != nil {
		return nil, fmt.Errorf("can't get from feed-worker: %w", err)
	}

	if httpResponse.StatusCode() != http.StatusOK {
		return nil, ErrInvalidFeedWorkerServiceResponse
	}

	return formatter.ToDomainPostList(posts), nil
}
