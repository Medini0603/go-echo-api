package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBarsBefore(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, GetBars(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBarsByIdBefore(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("http://localhost:1323/foo:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("1")

	st := rec.Code
	if assert.NoError(t, GetBarsId(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBarsSumBefore(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo/sum", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, GetBarsSum(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestEntry(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, Greetings(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestCreateBars(t *testing.T) {
	e := echo.New()
	userjson := `{"bar":123}`
	req := httptest.NewRequest(http.MethodPost, "http://localhost:1323/foo", strings.NewReader(userjson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, CreateNewBar(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestCreateBarsWrong(t *testing.T) {
	e := echo.New()
	userjson := "wrongjson"
	req := httptest.NewRequest(http.MethodPost, "http://localhost:1323/foo", strings.NewReader(userjson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, CreateNewBar(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBars(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, GetBars(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBarsById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("http://localhost:1323/foo:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("1")

	st := rec.Code
	if assert.NoError(t, GetBarsId(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBarsByIdNotThere(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("http://localhost:1323/foo:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("1000000")

	st := rec.Code
	if assert.NoError(t, GetBarsId(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestGetBarsSum(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:1323/foo/sum", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	st := rec.Code
	if assert.NoError(t, GetBarsSum(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestDeleteBars(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http:localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("http:localhost:1323/foo:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("1")

	st := rec.Code
	if assert.NoError(t, DeleteBarsId(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}

func TestDeleteBarsNotThere(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http:localhost:1323/foo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("http:localhost:1323/foo:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("10000")

	st := rec.Code
	if assert.NoError(t, DeleteBarsId(c)) {
		assert.Equal(t, http.StatusOK, st)
	}
}
