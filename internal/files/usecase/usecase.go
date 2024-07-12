package usecase

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/dhuruvah-apps/BuildSMEs-Api/config"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/files"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/httpErrors"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"
)

// Files UseCase
type filesUC struct {
	cfg      *config.Config
	filesRepo files.Repository
	logger   logger.Logger
}

// Files UseCase constructor
func NewFilesUseCase(cfg *config.Config, filesRepo files.Repository, logger logger.Logger) files.UseCase {
	return &filesUC{cfg: cfg, filesRepo: filesRepo, logger: logger}
}

// Create comment
func (u *filesUC) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.Create")
	defer span.Finish()
	return u.filesRepo.Create(ctx, comment)
}

// Update comment
func (u *filesUC) Update(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.Update")
	defer span.Finish()

	comm, err := u.filesRepo.GetByID(ctx, comment.CommentID)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateIsOwner(ctx, comm.AuthorID.String(), u.logger); err != nil {
		return nil, httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "filesUC.Update.ValidateIsOwner"))
	}

	updatedComment, err := u.filesRepo.Update(ctx, comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

// Delete comment
func (u *filesUC) Delete(ctx context.Context, commentID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.Delete")
	defer span.Finish()

	comm, err := u.filesRepo.GetByID(ctx, commentID)
	if err != nil {
		return err
	}

	if err = utils.ValidateIsOwner(ctx, comm.AuthorID.String(), u.logger); err != nil {
		return httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "filesUC.Delete.ValidateIsOwner"))
	}

	if err = u.filesRepo.Delete(ctx, commentID); err != nil {
		return err
	}

	return nil
}

// GetByID comment
func (u *filesUC) GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.GetByID")
	defer span.Finish()

	return u.filesRepo.GetByID(ctx, commentID)
}

// GetAll files
func (u *filesUC) GetAll(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.GetAll")
	defer span.Finish()

	return u.filesRepo.GetAll(ctx, newsID, query)
}
