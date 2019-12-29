package monitor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractTime(t *testing.T) {
	actual, _ := extractTime(`2019-12-29 12:07:27.596 Published job: 4, status=200 OK`)
	expected, _ := time.Parse(layout, "2019-12-29 12:07:27.596")
	assert.Equal(t, expected, actual)
}
