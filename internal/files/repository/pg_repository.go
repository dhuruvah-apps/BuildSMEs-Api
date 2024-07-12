package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/files"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"
)

// Files Repository
type filesRepo struct {
	db *sqlx.DB
}

// Files Repository constructor
func NewFilesRepository(db *sqlx.DB) files.Repository {
	return &filesRepo{db: db}
}

// Create comment
func (r *filesRepo) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.Create")
	defer span.Finish()

	c := &models.Comment{}
	if err := r.db.QueryRowxContext(
		ctx,
		createComment,
		&comment.AuthorID,
		&comment.NewsID,
		&comment.Message,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "filesRepo.Create.StructScan")
	}

	return c, nil
}

// Update comment
func (r *filesRepo) Update(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.Update")
	defer span.Finish()

	comm := &models.Comment{}
	if err := r.db.QueryRowxContext(ctx, updateComment, comment.Message, comment.CommentID).StructScan(comm); err != nil {
		return nil, errors.Wrap(err, "filesRepo.Update.QueryRowxContext")
	}

	return comm, nil
}

// Delete comment
func (r *filesRepo) Delete(ctx context.Context, commentID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteComment, commentID)
	if err != nil {
		return errors.Wrap(err, "filesRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "filesRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "filesRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID comment
func (r *filesRepo) GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.GetByID")
	defer span.Finish()

	comment := &models.CommentBase{}
	if err := r.db.GetContext(ctx, comment, getCommentByID, commentID); err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetByID.GetContext")
	}
	return comment, nil
}

// GetAll files
func (r *filesRepo) GetAll(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.GetAll")
	defer span.Finish()

	var totalCount int
	if err := r.db.QueryRowContext(ctx, getTotalCount, newsID).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetAll.QueryRowContext")
	}
	if totalCount == 0 {
		return &models.CommentsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Comments:   make([]*models.CommentBase, 0),
		}, nil
	}

	rows, err := r.db.QueryxContext(ctx, getComments, newsID, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	filesList := make([]*models.CommentBase, 0, query.GetSize())
	for rows.Next() {
		comment := &models.CommentBase{}
		if err = rows.StructScan(comment); err != nil {
			return nil, errors.Wrap(err, "filesRepo.GetAll.StructScan")
		}
		filesList = append(filesList, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetAll.rows.Err")
	}

	return &models.CommentsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Comments:   filesList,
	}, nil
}
