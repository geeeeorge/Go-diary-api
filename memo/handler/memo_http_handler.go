package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/geeeeorge/Go-diary-api/memo"
	"github.com/geeeeorge/Go-diary-api/model"
)

// MemoHTTPHandler represents memo http handler
type MemoHTTPHandler struct {
	usecase memo.Usecase
}

// NewMemoHTTPHandler returns new instance of MemoHTTPHandler
func NewMemoHTTPHandler(usecase memo.Usecase) *MemoHTTPHandler {
	return &MemoHTTPHandler{usecase: usecase}
}

// HandleCreateMemo handles create a memo
func (h *MemoHTTPHandler) HandleCreateMemo(c echo.Context) error {
	m := model.Memo{}

	err := c.Bind(&m)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.usecase.CreateMemo(c.Request().Context(), &m)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if _, ok := err.(validator.ValidationErrors); ok {
			statusCode = http.StatusBadRequest
		}
		return c.JSON(statusCode, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &Response{
		Message: "success",
		Data:    m,
	})
}

// HandleGetMemoByID handles get memo by id
func (h *MemoHTTPHandler) HandleGetMemoByID(c echo.Context) error {
	memoIDStr := c.Param("id")
	id, err := strconv.Atoi(memoIDStr)
	if err != nil {
		return c.JSON(http.StatusNotFound, &Response{
			Message: "memo is not found",
			Data:    nil,
		})
	}

	m, err := h.usecase.GetMemoByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	if m == nil {
		return c.JSON(http.StatusNotFound, &Response{
			Message: "memo is not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    m,
	})
}

// HandleGetAllMemo handles get all memos
func (h *MemoHTTPHandler) HandleGetAllMemo(c echo.Context) error {
	res, err := h.usecase.GetAllMemo(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    res,
	})
}
