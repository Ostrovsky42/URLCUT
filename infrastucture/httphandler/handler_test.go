package httphandler

import (
	"URLCUT/app"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_urlCutter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	returnKey := "6d82b33f"
	stubErr := errors.New("test error")
	testTable := []struct {
		description  string
		service      *app.MockUrlCutterServ
		messageJSON  string
		wantHttpCode int
		wantError    error
	}{
		{
			description: "should return statusOK",
			service: func(mock *app.MockUrlCutterServ) *app.MockUrlCutterServ {
				mock.EXPECT().MakeKey(gomock.Any()).Return(returnKey, nil).Times(1)
				return mock
			}(app.NewMockUrlCutterServ(ctrl)),
			messageJSON:  `{"long_url":"vk.com"}`,
			wantHttpCode: http.StatusOK,
			wantError:    nil,
		},
		{
			description:  "should return http.StatusBadRequest when request body data incorrect",
			service:      app.NewMockUrlCutterServ(ctrl),
			messageJSON:  `{[long_url]:"vk.com"}`,
			wantHttpCode: http.StatusBadRequest,
			wantError:    echo.NewHTTPError(http.StatusBadRequest),
		},
		{
			description: "should return http.StatusInternalServerError when //request body data incorrect",
			service: func(mock *app.MockUrlCutterServ) *app.MockUrlCutterServ {
				mock.EXPECT().MakeKey(gomock.Any()).Return("", stubErr).Times(1)
				return mock
			}(app.NewMockUrlCutterServ(ctrl)),
			messageJSON:  `{"long_url":"vk.com"}`,
			wantHttpCode: http.StatusBadRequest,
			wantError:    echo.NewHTTPError(http.StatusInternalServerError),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.messageJSON))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/urlCut")
			handle := NewUrlGenerator(testCase.service)

			if err := handle.UrlCutter(c); err == nil {
				assert.Equal(t, testCase.wantHttpCode, rec.Code)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}

func TestHandler_GetUrl(t *testing.T) {

	//stubErr := errors.New("test error")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	returnUrl := "vk.com"
	stubErr := errors.New("test error")
	testTable := []struct {
		description  string
		service      *app.MockUrlCutterServ
		ctx          string // *echo.Context
		wantHttpCode int
		wantError    error
	}{
		{
			description: "should return statusOK",
			service: func(mock *app.MockUrlCutterServ) *app.MockUrlCutterServ {
				mock.EXPECT().GetURL(gomock.Any()).Return(returnUrl, nil).Times(1)
				return mock
			}(app.NewMockUrlCutterServ(ctrl)),
			ctx:          "6d82b33f",
			wantHttpCode: http.StatusMovedPermanently,
			wantError:    nil,
		},
		{
			description:  "should return http.StatusBadRequest when context ParamValues==\"\"",
			service:      app.NewMockUrlCutterServ(ctrl),
			ctx:          "",
			wantHttpCode: http.StatusBadRequest,
			wantError:    errors.New("key invalid"),
		},
		{
			description: "should return http.StatusInternalServerError when request body data incorrect",
			service: func(mock *app.MockUrlCutterServ) *app.MockUrlCutterServ {
				mock.EXPECT().GetURL(gomock.Any()).Return("", stubErr).Times(1)
				return mock
			}(app.NewMockUrlCutterServ(ctrl)),
			ctx:          "6d82b33f",
			wantHttpCode: http.StatusInternalServerError,
			wantError:    echo.NewHTTPError(http.StatusInternalServerError),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/:key", nil)
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/:key")
			c.SetParamNames("key")
			c.SetParamValues(testCase.ctx)
			handle := NewUrlGenerator(testCase.service)
			if err := handle.GetUrl(c); err == nil {
				assert.Equal(t, testCase.wantHttpCode, rec.Code)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}
