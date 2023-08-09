package apiserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	apiserver "github.com/winterochek/go-app/internal/app/api-server"
	teststore "github.com/winterochek/go-app/internal/app/store/test-store"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)

	s := apiserver.NewServer(teststore.New())
	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
