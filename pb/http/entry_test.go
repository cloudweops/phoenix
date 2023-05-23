package http_test

import (
	"testing"

	"github.com/CloudWeOps/phoenix/http/label"
	"github.com/CloudWeOps/phoenix/pb/http"
	"github.com/stretchr/testify/assert"
)

func TestEntry(t *testing.T) {
	should := assert.New(t)

	e := http.NewEntry("/phoenix/v1/", "GET", "Monkey")
	e.EnableAuth()
	e.EnablePermission()
	e.AddLabel(label.Get)

	should.Equal("Monkey", e.Resource)

	set := http.NewEntrySet()
	set.AddEntry(*e, *e)
	should.Equal(2, len(set.Items))
}
