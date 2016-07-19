sql-migrate-gen
===============

helper for rubenv/sql-migrate(https://github.com/rubenv/sql-migrate)

## Installation

```
$> go get github.com/elvuel/sql-migrate-gen
$> go install github.com/elvuel/sql-migrate-gen
```

## Basic Usage

```
$> sql-migrate-gen -n add_new_column_to_Table
$> sql-migrate-gen -n AddNewColumnToTable -o /path/to/migrations
```