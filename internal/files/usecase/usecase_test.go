package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/files/mock"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"
)

func TestFilesUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCommRepo := mock.NewMockRepository(ctrl)
	commUC := NewFilesUseCase(nil, mockCommRepo, apiLogger)

	comm := &models.Comment{}

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "filesUC.Create")
	defer span.Finish()

	mockCommRepo.EXPECT().Create(ctx, gomock.Eq(comm)).Return(comm, nil)

	createdComment, err := commUC.Create(context.Background(), comm)
	require.NoError(t, err)
	require.NotNil(t, createdComment)
}

func TestFilesUC_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCommRepo := mock.NewMockRepository(ctrl)
	commUC := NewFilesUseCase(nil, mockCommRepo, apiLogger)

	authorUID := uuid.New()

	comm := &models.Comment{
		CommentID: uuid.New(),
		AuthorID:  authorUID,
	}

	baseComm := &models.CommentBase{
		AuthorID: authorUID,
	}

	user := &models.User{
		UserID: authorUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "filesUC.Update")
	defer span.Finish()

	mockCommRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(comm.CommentID)).Return(baseComm, nil)
	mockCommRepo.EXPECT().Update(ctxWithTrace, gomock.Eq(comm)).Return(comm, nil)

	updatedComment, err := commUC.Update(ctx, comm)
	require.NoError(t, err)
	require.NotNil(t, updatedComment)
}

func TestFilesUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCommRepo := mock.NewMockRepository(ctrl)
	commUC := NewFilesUseCase(nil, mockCommRepo, apiLogger)

	authorUID := uuid.New()

	comm := &models.Comment{
		CommentID: uuid.New(),
		AuthorID:  authorUID,
	}

	baseComm := &models.CommentBase{
		AuthorID: authorUID,
	}

	user := &models.User{
		UserID: authorUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "filesUC.Delete")
	defer span.Finish()

	mockCommRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(comm.CommentID)).Return(baseComm, nil)
	mockCommRepo.EXPECT().Delete(ctxWithTrace, gomock.Eq(comm.CommentID)).Return(nil)

	err := commUC.Delete(ctx, comm.CommentID)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestFilesUC_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCommRepo := mock.NewMockRepository(ctrl)
	commUC := NewFilesUseCase(nil, mockCommRepo, apiLogger)

	comm := &models.Comment{
		CommentID: uuid.New(),
	}

	baseComm := &models.CommentBase{}

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "filesUC.GetByID")
	defer span.Finish()

	mockCommRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(comm.CommentID)).Return(baseComm, nil)

	commentBase, err := commUC.GetByID(ctx, comm.CommentID)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, commentBase)
}

func TestFilesUC_GetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCommRepo := mock.NewMockRepository(ctrl)
	commUC := NewFilesUseCase(nil, mockCommRepo, apiLogger)

	newsUID := uuid.New()

	comm := &models.Comment{
		CommentID: uuid.New(),
		NewsID:    newsUID,
	}

	filesList := &models.CommentsList{}

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "filesUC.GetAll")
	defer span.Finish()

	query := &utils.PaginationQuery{
		Size:    10,
		Page:    1,
		OrderBy: "",
	}

	mockCommRepo.EXPECT().GetAll(ctxWithTrace, gomock.Eq(comm.NewsID), query).Return(filesList, nil)

	commList, err := commUC.GetAll(ctx, comm.NewsID, query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, commList)
}
