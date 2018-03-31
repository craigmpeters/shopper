#!/usr/bin/env bash

# Start MS SQL Server and Create and Import DB
/opt/mssql/bin/sqlservr & /usr/src/app/import-data.sh