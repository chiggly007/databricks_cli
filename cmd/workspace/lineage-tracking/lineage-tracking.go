package lineage_tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdctx"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/spf13/cobra"
)

// New returns the lineage-tracking root command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "lineage-tracking",
		Short:   "Retrieve table and column lineage",
		Long:    "Retrieve Unity Catalog table and column lineage using the Data Lineage REST API.",
		GroupID: "catalog",
		RunE:    root.ReportUnknownSubcommand,
	}

	cmd.AddCommand(newTableLineageCommand())
	cmd.AddCommand(newColumnLineageCommand())

	return cmd
}

// newTableLineageCommand returns the table-lineage subcommand.
func newTableLineageCommand() *cobra.Command {
	var includeEntityLineage bool

	cmd := &cobra.Command{
		Use:   "table-lineage TABLE_NAME",
		Short: "Get table lineage",
		Long: `Get table lineage for a Unity Catalog table.

Provide a fully qualified table name like catalog.schema.table.`,
		Args: root.ExactArgs(1),
	}

	cmd.Flags().BoolVar(&includeEntityLineage, "include-entity-lineage", true, "Include notebook, job, or dashboard lineage when available")

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		w := cmdctx.WorkspaceClient(ctx)

		reqBody := TableLineageRequest{
			TableName:            args[0],
			IncludeEntityLineage: includeEntityLineage,
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}

		url := fmt.Sprintf("%s/api/2.0/lineage-tracking/table-lineage", w.Config.Host)
		req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewReader(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		err = w.Config.Authenticate(req)
		if err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to execute request: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}

		var response map[string]any
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		return cmdio.Render(ctx, response)
	}

	cmd.ValidArgsFunction = cobra.NoFileCompletions

	return cmd
}

// newColumnLineageCommand returns the column-lineage subcommand.
func newColumnLineageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "column-lineage TABLE_NAME COLUMN_NAME",
		Short: "Get column lineage",
		Long: `Get column lineage for a Unity Catalog table column.

Provide a fully qualified table name like catalog.schema.table.`,
		Args: root.ExactArgs(2),
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		w := cmdctx.WorkspaceClient(ctx)

		reqBody := ColumnLineageRequest{
			TableName:  args[0],
			ColumnName: args[1],
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}

		url := fmt.Sprintf("%s/api/2.0/lineage-tracking/column-lineage", w.Config.Host)
		req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewReader(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		err = w.Config.Authenticate(req)
		if err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to execute request: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}

		var response map[string]any
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		return cmdio.Render(ctx, response)
	}

	cmd.ValidArgsFunction = cobra.NoFileCompletions

	return cmd
}