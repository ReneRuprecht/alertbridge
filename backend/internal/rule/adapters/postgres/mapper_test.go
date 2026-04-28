package postgres

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
)

func TestToDomain_Valid(t *testing.T) {

	id, _ := domain.NewRuleId()
	name, err := domain.NewRuleName("testName")
	description := "testDesc"
	priority, _ := domain.NewRulePriority(100)
	enabled := true

	assert.NoError(t, err)

	rule := toRuleRepositoryEntity(domain.Rule{ID: id, Name: name, Description: description, Priority: priority, Enabled: enabled})

	domain, err := toDomain(rule)

	assert.NoError(t, err)

	assert.Equal(t, id.String(), domain.ID.String())
	assert.Equal(t, "testName", string(domain.Name))
	assert.Equal(t, "testDesc", domain.Description)
	assert.Equal(t, 100, domain.Priority.Int())
	assert.Equal(t, true, domain.Enabled)
}

func TestToDomain_Invalid_Name(t *testing.T) {

	id, _ := domain.NewRuleId()
	description := "testDesc"
	priority, _ := domain.NewRulePriority(100)
	enabled := true

	rule := toRuleRepositoryEntity(domain.Rule{ID: id, Name: "", Description: description, Priority: priority, Enabled: enabled})

	_, err := toDomain(rule)

	assert.EqualError(t, err, "Name cannot be empty")
}
