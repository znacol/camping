package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/znacol/camping/go/db/mock_db"
)

// init for tests only
func init() {
	log.SetLevel(log.DebugLevel)
}

// newTestAPI returns a new *API for testing purposes, with mocks
func newTestAPI(tb testing.TB) (*API, *gomock.Controller, *mock_db.MockAPI) {
	tb.Helper()

	// Create mock controller
	mockCtrl := gomock.NewController(tb)

	// Create mock database
	mockDB := mock_db.NewMockAPI(mockCtrl)

	// Create API
	api := &API{
		dbClient:                  mockDB,
	}

	// Return *API with mocks
	return api, mockCtrl, mockDB
}

// teardown finishes the mock controller
func teardown(tb testing.TB, mockCtrl *gomock.Controller) {
	tb.Helper()

	// Finish mocks
	mockCtrl.Finish()
}

func TestNewAPI(t *testing.T) {
	t.Parallel()
	_, mockCtrl, _ := newTestAPI(t)
	teardown(t, mockCtrl)
}
