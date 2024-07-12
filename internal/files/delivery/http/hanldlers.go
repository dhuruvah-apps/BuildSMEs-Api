package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/dhuruvah-apps/BuildSMEs-Api/config"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/files"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/httpErrors"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"
)

// Comments handlers
type filesHandlers struct {
	cfg     *config.Config
	filesUC files.UseCase
	logger  logger.Logger
}

// NewCommentsHandlers Comments handlers constructor
func NewFilesHandlers(cfg *config.Config, filesUC files.UseCase, logger logger.Logger) files.Handlers {
	return &filesHandlers{cfg: cfg, filesUC: filesUC, logger: logger}
}

// Create
// @Summary Create new file
// @Description create new file
// @Tags Files
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Comment
// @Failure 500 {object} httpErrors.RestErr
// @Router /files [post]
func (h *filesHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "filesHandlers.Create")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		file := &models.Comment{}
		file.AuthorID = user.UserID

		if err = utils.SanitizeRequest(c, file); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
			// return err
		}

		createdComment, err := h.filesUC.Create(ctx, file)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdComment)
	}
}

// Update
// @Summary Update file
// @Description update new file
// @Tags Files
// @Accept  json
// @Produce  json
// @Param id path int true "file_id"
// @Success 200 {object} models.Comment
// @Failure 500 {object} httpErrors.RestErr
// @Router /files/{id} [put]
func (h *filesHandlers) Update() echo.HandlerFunc {
	type UpdateComment struct {
		Message string `json:"message" db:"message" validate:"required,gte=0"`
		Likes   int64  `json:"likes" db:"likes" validate:"omitempty"`
	}
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "filesHandlers.Update")
		defer span.Finish()

		commID, err := uuid.Parse(c.Param("file_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &UpdateComment{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedComment, err := h.filesUC.Update(ctx, &models.Comment{
			CommentID: commID,
			Message:   comm.Message,
			Likes:     comm.Likes,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedComment)
	}
}

// Delete
// @Summary Delete file
// @Description delete file
// @Tags Files
// @Accept  json
// @Produce  json
// @Param id path int true "file_id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /files/{id} [delete]
func (h *filesHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "filesHandlers.Delete")
		defer span.Finish()

		commID, err := uuid.Parse(c.Param("file_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.filesUC.Delete(ctx, commID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get file
// @Description Get file by id
// @Tags Files
// @Accept  json
// @Produce  json
// @Param id path int true "file_id"
// @Success 200 {object} models.Comment
// @Failure 500 {object} httpErrors.RestErr
// @Router /files/{id} [get]
func (h *filesHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "filesHandlers.GetByID")
		defer span.Finish()

		commID, err := uuid.Parse(c.Param("file_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		file, err := h.filesUC.GetByID(ctx, commID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, file)
	}
}

// GetAll
// @Summary Get files by news
// @Description Get all file by news id
// @Tags Files
// @Accept  json
// @Produce  json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.CommentsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /files [get]
func (h *filesHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "filesHandlers.GetAll")
		defer span.Finish()

		newsID, err := uuid.Parse(c.Param("news_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		filesList, err := h.filesUC.GetAll(ctx, newsID, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, filesList)
	}
}
