FROM microsoft/mssql-server-linux:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . .
RUN chmod +x /usr/src/app/import-data.sh
RUN chmod +x /usr/src/app/entrypoint.sh

CMD ["./entrypoint.sh"]