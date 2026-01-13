package lineage_tracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := New()
	assert.NotNil(t, cmd)
	assert.Equal(t, "lineage-tracking", cmd.Use)
	assert.Equal(t, "Retrieve table and column lineage", cmd.Short)
}

func TestTableLineageCommand(t *testing.T) {
	cmd := newTableLineageCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "table-lineage TABLE_NAME", cmd.Use)
	assert.Contains(t, cmd.Long, "Get table lineage")
	f := cmd.Flags().Lookup("include-entity-lineage")
	if assert.NotNil(t, f) {
		assert.Equal(t, "include-entity-lineage", f.Name)
	}
}

func TestColumnLineageCommand(t *testing.T) {
	cmd := newColumnLineageCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "column-lineage TABLE_NAME COLUMN_NAME", cmd.Use)
	assert.Contains(t, cmd.Long, "Get column lineage")
}
