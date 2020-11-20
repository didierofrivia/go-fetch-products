FROM golang:1.13.7-alpine3.11 as stage1

COPY src /codebase/src
RUN cd /codebase && go build -v -o /codebase/bin/server ./src/main.go

FROM alpine:3.11 as stage2
COPY --from=stage1 /codebase/bin/server /server
ENV PORT=8080

CMD ["sh", "-c", "/server"]