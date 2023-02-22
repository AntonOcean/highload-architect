package usecase_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"feed-worker/internal/domain"
	"feed-worker/internal/usecase"
)

func TestUsecase_MergeArr(t *testing.T) {
	t.Run("merge arr", func(t *testing.T) {
		t1, _ := time.Parse(time.RFC3339, "2016-01-02T15:04:05Z07:00")
		t2, _ := time.Parse(time.RFC3339, "2026-01-02T15:04:05Z07:00")

		left := []*domain.Post{
			{
				ID:       uuid.New(),
				Text:     "heelo",
				AuthorID: uuid.New(),
				Created:  t1,
			},
		}

		right := []*domain.Post{
			{
				ID:       uuid.New(),
				Text:     "heelo 2",
				AuthorID: uuid.New(),
				Created:  t2,
			},
		}

		res := usecase.Merge(left, right)

		assert.True(t, res[0].Text == "heelo 2")
	})
}
