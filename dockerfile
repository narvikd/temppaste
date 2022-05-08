FROM alpine:latest
WORKDIR /app
COPY temppaste temppaste
RUN apk update && apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime
RUN echo "Europe/Berlin" >  /etc/timezone
ENTRYPOINT ["/app/temppaste"]
