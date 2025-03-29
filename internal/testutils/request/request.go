package request

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func NewGETRequest(e *echo.Echo, url string, sessionID string) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	if sessionID != "" {
		http.SetCookie(rec, &http.Cookie{
			Name:    "token",
			Value:   sessionID,
			Expires: time.Now().Add(time.Hour * 24 * 365),
		})
		req.Header.Set(echo.HeaderCookie, rec.Result().Cookies()[0].Raw)
	}

	c := e.NewContext(req, rec)
	return c, rec
}

func NewPOSTRequest(e *echo.Echo, url string, body string, sessionID string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	if sessionID != "" {
		http.SetCookie(rec, &http.Cookie{
			Name:    "token",
			Value:   sessionID,
			Expires: time.Now().Add(time.Hour * 24 * 365),
		})
		req.Header.Set(echo.HeaderCookie, rec.Result().Cookies()[0].Raw)
	}

	c := e.NewContext(req, rec)
	return c, rec
}

func NewPATCHRequest(e *echo.Echo, url string, body string, sessionID string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	if sessionID != "" {
		http.SetCookie(rec, &http.Cookie{
			Name:    "token",
			Value:   sessionID,
			Expires: time.Now().Add(time.Hour * 24 * 365),
		})
		req.Header.Set(echo.HeaderCookie, rec.Result().Cookies()[0].Raw)
	}

	c := e.NewContext(req, rec)
	return c, rec
}
