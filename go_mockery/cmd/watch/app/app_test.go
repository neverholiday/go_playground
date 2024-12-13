package app

import (
	"go_mockery/cmd/watch/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowCurrentTime(t *testing.T) {
	WatchMock := mocks.NewMockIWatch(t)
	WatchMock.EXPECT().GetCurrentTime().Return("12:34:55", nil)

	app := NewApp(WatchMock)
	actual := app.ShowCurrentTime()
	assert.Equal(t, "This time is : 12:34:55", actual)

}
