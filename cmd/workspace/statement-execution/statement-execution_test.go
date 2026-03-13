package statement_execution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExecuteStatementCommandHasWarehouseIDFlags(t *testing.T) {
	cmd := newExecuteStatementCommand()

	flag := cmd.Flags().Lookup("warehouse-id")
	if assert.NotNil(t, flag) {
		assert.Equal(t, "warehouse-id", flag.Name)
	}

	alias := cmd.Flags().Lookup("warehouse_id")
	if assert.NotNil(t, alias) {
		assert.Equal(t, "warehouse_id", alias.Name)
		assert.True(t, alias.Hidden)
	}
}
