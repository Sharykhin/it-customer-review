package grpc

import (
	"context"

	"sync"
	"testing"

	"github.com/Sharykhin/it-customer-review/api/entity"
	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) Create(ctx context.Context, in *pb.ReviewCreateRequest, opts ...grpc.CallOption) (*pb.ReviewResponse, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ReviewResponse), nil
}

func (m *mockStorage) Update(ctx context.Context, in *pb.ReviewUpdateRequest, opts ...grpc.CallOption) (*pb.ReviewResponse, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ReviewResponse), nil
}

func (m *mockStorage) Get(ctx context.Context, in *pb.ReviewID, opts ...grpc.CallOption) (*pb.ReviewResponse, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ReviewResponse), nil
}

func (m *mockStorage) Ping(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.Pong, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Pong), nil
}

func (m *mockStorage) GetReviewList(ctx context.Context, in *pb.ReviewListFilter, opts ...grpc.CallOption) (pb.Review_GetReviewListClient, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(pb.Review_GetReviewListClient), nil
}

func (m *mockStorage) CountReviews(ctx context.Context, in *pb.ReviewCountFilter, opts ...grpc.CallOption) (*pb.CountResponse, error) {
	args := m.Called(ctx, in)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CountResponse), nil
}

func TestReviewService_Create(t *testing.T) {
	m := new(mockStorage)
	defer m.AssertExpectations(t)

	ctx := context.Background()

	expectedReview := entity.Review{
		ID:        "1",
		Name:      "bob",
		Email:     "bob@mail.com",
		Content:   "I like testing",
		Published: false,
	}

	expectedResponse := pb.ReviewResponse{
		ID:        "1",
		Name:      "bob",
		Email:     "bob@mail.com",
		Content:   "I like testing",
		Published: false,
	}

	m.On("Create", ctx, &pb.ReviewCreateRequest{
		Name:     "bob",
		Email:    "bob@mail.com",
		Content:  "I like testing",
		Score:    &pb.ReviewCreateRequest_ScoreNull{ScoreNull: true},
		Category: &pb.ReviewCreateRequest_CategoryNull{CategoryNull: true},
	}).Return(&expectedResponse, nil).Once()

	m.On("Create", ctx, &pb.ReviewCreateRequest{
		Name:     "alice",
		Email:    "alice@mail.com",
		Content:  "I hate testing",
		Score:    &pb.ReviewCreateRequest_ScoreNull{ScoreNull: true},
		Category: &pb.ReviewCreateRequest_CategoryNull{CategoryNull: true},
	}).Return(nil, errors.New("some error")).Once()

	tt := []struct {
		name           string
		incomeRequest  entity.ReviewRequest
		expectedReview *entity.Review
		expectedErr    error
	}{
		{
			name: "success creation",
			incomeRequest: entity.ReviewRequest{
				Name:    "bob",
				Email:   "bob@mail.com",
				Content: "I like testing",
			},
			expectedReview: &expectedReview,
			expectedErr:    nil,
		},
		{
			name: "error creation",
			incomeRequest: entity.ReviewRequest{
				Name:    "alice",
				Email:   "alice@mail.com",
				Content: "I hate testing",
			},
			expectedReview: nil,
			expectedErr:    errors.New("some error"),
		},
	}

	ctrl := &reviewService{
		client: m,
	}

	var wg sync.WaitGroup

	for _, tc := range tt {
		wg.Add(1)
		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			actual, err := ctrl.Create(ctx, tc.incomeRequest)
			if err == nil {
				require.Equal(t, tc.expectedReview.ID, actual.ID)
				require.Equal(t, tc.expectedReview.Name, actual.Name)
				require.Equal(t, tc.expectedReview.Email, actual.Email)
				require.Equal(t, tc.expectedReview.Content, actual.Content)
				require.Equal(t, tc.expectedReview.Published, actual.Published)
			}

			if err != nil {
				require.Nil(t, tc.expectedReview)
			}
		})
	}
	wg.Wait()
}
