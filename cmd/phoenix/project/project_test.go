package project_test

import (
	"testing"

	"github.com/CloudWeOps/phoenix/cmd/phoenix/project"
	"github.com/stretchr/testify/assert"
)

func TestSaveFile(t *testing.T) {
	should := assert.New(t)

	p := project.Project{
		PKG:  "test",
		Name: "test",
	}

	err := p.SaveFile(project.PROJECT_SETTING_FILE_PATH)
	should.NoError(err)
}