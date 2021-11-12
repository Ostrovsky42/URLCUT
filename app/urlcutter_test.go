package app

import (
	"URLCUT/infrastucture/localservices"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlCutter_GetURL(t *testing.T) {
	returnUrl := "vk.com"
	stubErr := errors.New("test error")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		description  string
		service      *MockUrlSaver
		messageKey   string
		wantResponse string
		wantError    error
	}{
		{description: "should return Url",
			service: func(mock *MockUrlSaver) *MockUrlSaver {
				mock.EXPECT().Get(gomock.Any()).Return(returnUrl, nil).Times(1)
				return mock
			}(NewMockUrlSaver(ctrl)),
			messageKey:   "6d82b33f",
			wantResponse: returnUrl,
			wantError:    nil,
		},
		{description: "should return err when ...",
			service: func(mock *MockUrlSaver) *MockUrlSaver {
				mock.EXPECT().Get(gomock.Any()).Return("", stubErr).Times(1)
				return mock
			}(NewMockUrlSaver(ctrl)),
			messageKey:   "6d82b33f",
			wantResponse: "",
			wantError:    stubErr,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {

			handle := NewURLCutterService(localservices.KeyGenerator{}, testCase.service)
			if resp, err := handle.GetURL(testCase.messageKey); err == nil {
				assert.Equal(t, testCase.wantResponse, resp)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}

func TestUrlCutter_MakeKey(t *testing.T) {
	returnKey := "6d82b33f"
	stubErr := errors.New("test error")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		description     string
		UrlSaverService *MockUrlSaver
		KeyGenService   *localservices.MockGenerateInterface
		messageUrl      string
		wantResponse    string
		wantError       error
	}{{
		description: "should return Key",
		UrlSaverService: func(mock *MockUrlSaver) *MockUrlSaver {
			mock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			mock.EXPECT().Get(gomock.Any()).Return("", nil).Times(1)
			return mock
		}(NewMockUrlSaver(ctrl)),
		KeyGenService: func(mock *localservices.MockGenerateInterface) *localservices.MockGenerateInterface {
			mock.EXPECT().GenerateKey().Return(returnKey).Times(1)
			return mock
		}(localservices.NewMockGenerateInterface(ctrl)),
		messageUrl:   "vk.com",
		wantResponse: returnKey,
		wantError:    nil,
	},
		{
			description: "should return err",
			UrlSaverService: func(mock *MockUrlSaver) *MockUrlSaver {
				mock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(stubErr).Times(1)
				mock.EXPECT().Get(gomock.Any()).Return("", nil).Times(1)
				return mock
			}(NewMockUrlSaver(ctrl)),
			KeyGenService: func(mock *localservices.MockGenerateInterface) *localservices.MockGenerateInterface {
				mock.EXPECT().GenerateKey().Return(returnKey).Times(1)
				return mock
			}(localservices.NewMockGenerateInterface(ctrl)),
			messageUrl:   "vk.com",
			wantResponse: returnKey,
			wantError:    stubErr,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {
			handle := NewURLCutterService(testCase.KeyGenService, testCase.UrlSaverService)
			if resp, err := handle.MakeKey(testCase.messageUrl); err == nil {
				assert.Equal(t, testCase.wantResponse, resp)
			} else {
				assert.Error(t, testCase.wantError, err)
			}
		})
	}
}

func TestUrlCutter_GetUniqueKey(t *testing.T) {
	returnKey := "6d82b33f"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		description     string
		UrlSaverService *MockUrlSaver
		KeyGenService   *localservices.MockGenerateInterface
		wantResponse    string
	}{
		{description: "should return statusOK",
			UrlSaverService: func(mock *MockUrlSaver) *MockUrlSaver {
				mock.EXPECT().Get(gomock.Any()).Return("", nil).Times(1)
				return mock
			}(NewMockUrlSaver(ctrl)),
			KeyGenService: func(mock *localservices.MockGenerateInterface) *localservices.MockGenerateInterface {
				mock.EXPECT().GenerateKey().Return(returnKey).Times(1)
				return mock
			}(localservices.NewMockGenerateInterface(ctrl)),
			wantResponse: returnKey,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.description, func(t *testing.T) {

			handle := NewURLCutterService(testCase.KeyGenService, testCase.UrlSaverService)
			resp := handle.GetUniqueKey()
			assert.Equal(t, testCase.wantResponse, resp)

		})
	}
}
