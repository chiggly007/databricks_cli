# Databricks CLI Guide

This document provides instructions for interacting with Databricks via the CLI.

## 1. SQL Statement Execution

Execute SQL queries against Databricks SQL warehouses. For queries that may take time, follow the two-step polling process.

### Execute a Statement

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks statement-execution execute-statement "YOUR_SQL_QUERY_HERE"
```

Returns a `statement_id` and initial `status`.

### Check Statement Status

If status is `PENDING`, poll until `SUCCEEDED`:

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks statement-execution get-statement <STATEMENT_ID>
```

### Guidelines

- Always verify SQL syntax before executing
- Plan for long-running queries (avoid locking large tables unnecessarily)
- Handle `FAILED` status appropriately—never assume success
- Wait for `SUCCEEDED` before proceeding with downstream operations

### Workflow

1. Execute the statement → receive `statement_id` and `status`
2. If `PENDING`, poll with `get-statement` until `SUCCEEDED`
3. Once `SUCCEEDED`, proceed with next steps

---

## 2. Jobs Management

Schedule and run data processing tasks on Databricks clusters.

### Create a Job

Using inline JSON:

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs create --json '{
  "name": "My Job",
  "tasks": [
    {
      "task_key": "my_task",
      "notebook_task": {
        "notebook_path": "/Users/your-email/your-notebook"
      },
      "new_cluster": {
        "spark_version": "13.3.x-scala2.12",
        "num_workers": 1,
        "node_type_id": "i3.xlarge"
      }
    }
  ]
}'
```

Using a JSON file:

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs create --json @path/to/job-config.json
```

### List Jobs

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs list
```

### Get Job Details

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs get <JOB_ID>
```

### Trigger a Job Run

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs run-now <JOB_ID>
```

Returns a `run_id` to monitor the run.

### Check Job Run Status

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs get-run <RUN_ID>
```

Poll until `state.life_cycle_state` is `TERMINATED` and `state.result_state` is `SUCCESS`.

### List Job Runs

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs list-runs --job-id <JOB_ID>
```

### Cancel a Run

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs cancel-run <RUN_ID>
```

### Delete a Job

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks jobs delete <JOB_ID>
```

### Workflow

1. **Create** the job → receive `job_id`
2. **Trigger** with `run-now` → receive `run_id`
3. **Poll** with `get-run` until `life_cycle_state` is `TERMINATED`
4. **Verify** `result_state` is `SUCCESS`

---

## 3. Lineage Tracking

View data lineage for Unity Catalog tables and columns.

### Get Table Lineage

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks lineage-tracking table-lineage <CATALOG>.<SCHEMA>.<TABLE>
```

**Example:**

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks lineage-tracking table-lineage main.default.sales_data
```

Returns:
- **Upstream tables** — Tables that feed into this table
- **Downstream tables** — Tables that depend on this table
- **Entity lineage** — Related notebooks, jobs, or dashboards

### Exclude Entity Lineage

To get only table-to-table relationships:

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks lineage-tracking table-lineage <CATALOG>.<SCHEMA>.<TABLE> --include-entity-lineage=false
```

### Get Column Lineage

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks lineage-tracking column-lineage <CATALOG>.<SCHEMA>.<TABLE> <COLUMN_NAME>
```

**Example:**

```bash
/Users/chiragshah/PycharmProjects/databricks_cli/databricks lineage-tracking column-lineage main.default.sales_data total_revenue
```

Shows upstream columns that flow into the specified column and downstream columns that depend on it.

### Notes

- Lineage is only available for Unity Catalog tables
- Lineage is captured automatically when queries run through Databricks compute
- Historical data retention depends on workspace configuration
