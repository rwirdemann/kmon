package monitor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractTime(t *testing.T) {
	actual, _ := extractTime(`time="2019-12-27T10:15:52+01:00" level=info msg="job successfully posted" http status=200 id=4694`)
	expected, _ := time.Parse(time.RFC3339, "2019-12-27T10:15:52+01:00")
	assert.Equal(t, expected, actual)
}
