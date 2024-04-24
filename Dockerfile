FROM debian:stable-slim

COPY quick-go-server /bin/goserver

ENV PORT=8080

CMD [ "/bin/goserver" ]