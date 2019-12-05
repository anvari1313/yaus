package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anvari1313/yaus/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct{}

func (m *mockUserRepo) Create(ctx context.Context, url string) (string, error) {
	return "", nil
}

func (m *mockUserRepo) FindByID(ctx context.Context, id string) (*repository.URLModel, error) {
	return &repository.URLModel{URL: "https://some-domain.com/some/url"}, nil
}

func TestApp_RegisterRoute(t *testing.T) {

}

func TestApp_GetRoute(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	app := App{
		UserRepo: nil,
	}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
