package lineage_tracking

// TableLineageRequest describes the request body for table lineage.
type TableLineageRequest struct {
	TableName            string `json:"table_name"`
	IncludeEntityLineage bool   `json:"include_entity_lineage,omitempty"`
}

// ColumnLineageRequest describes the request body for column lineage.
type ColumnLineageRequest struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
}
