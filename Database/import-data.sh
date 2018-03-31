#!/usr/bin/env bash

# Wait for SQL Server to start
#sleep 90s

# CREATE DATABASE USING SA
#/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $(SA_PASSWORD) -d master -q "CREATE DATABASE $(DB_DATABASE);"

# ADD USER
#/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $(SA_PASSWORD) -d $(DB_DATABASE) -q "CREATE LOGIN $(DB_USERNAME) WITH PASSWORD = '$(DB_PASSWORD)'"
#/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $(SA_PASSWORD) -d $(DB_DATABASE) -q  "USE $(DB_DATABASE); CREATE USER $(DB_USERNAME) FOR LOGIN $(DB_USERNAME);"

# EXECUTE SCRIPT TO CREATE TABLES
#/opt/mssql-tools/bin/sqlcmd -S localhost -U $(DB_USERNAME) -P $(DB_PASSWORD) -d $(DB_DATABASE) -i setup.sql