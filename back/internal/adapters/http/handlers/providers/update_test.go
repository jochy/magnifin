package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"magnifin/internal/adapters/http/handlers/providers/mocks"
	"magnifin/internal/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	accessKey := "access_key"
	secret := "secret"

	tests := []struct {
		name               string
		request            updateProviderRequest
		providerID         string
		mockUpdateResponse model.Provider
		mockUpdateError    error
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "success",
			request: updateProviderRequest{
				Name:      "UpdatedProvider",
				Enabled:   true,
				AccessKey: &accessKey,
				Secret:    &secret,
			},
			providerID: "1",
			mockUpdateResponse: model.Provider{
				ID:        1,
				Name:      "UpdatedProvider",
				Enabled:   true,
				AccessKey: &accessKey,
				Secret:    &secret,
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":1,"name":"UpdatedProvider","enabled":true}`,
		},
		{
			name: "invalid provider id",
			request: updateProviderRequest{
				Name:    "UpdatedProvider",
				Enabled: true,
			},
			providerID:         "invalid",
			mockUpdateResponse: model.Provider{},
			mockUpdateError:    errors.New("update error"),
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"invalid provider id"}`,
		},
		{
			name: "update error",
			request: updateProviderRequest{
				Name:    "UpdatedProvider",
				Enabled: true,
			},
			providerID:         "1",
			mockUpdateResponse: model.Provider{},
			mockUpdateError:    errors.New("update error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"update error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mocks.NewService(t)
			handler := NewHandler(mockService)

			if tt.mockUpdateError == nil {
				mockService.On("UpdateProvider", mock.Anything, mock.Anything).Return(&tt.mockUpdateResponse, nil)
			} else {
				mockService.On("UpdateProvider", mock.Anything, mock.Anything).Return(nil, tt.mockUpdateError).Maybe()
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonBytes, _ := json.Marshal(tt.request)

			c.Request = httptest.NewRequest(http.MethodPut, "/providers/"+tt.providerID, bytes.NewReader(jsonBytes))
			c.Params = gin.Params{{Key: "id", Value: tt.providerID}}
			c.Request.Header.Set("Content-Type", "application/json")

			handler.Update(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.JSONEq(t, tt.expectedResponse, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}
