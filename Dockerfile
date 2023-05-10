FROM golang:1.19-alpine as builder 

WORKDIR /
ADD . .

RUN go build -o /tmp/bot main.go

FROM alpine AS runner

COPY --from=builder /tmp/bot /bin/bot

RUN chmod +x /bin/bot

CMD ["/bin/bot"]