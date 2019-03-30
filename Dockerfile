FROM alpine
ENV LANGUAGE="en"

ADD generator-super-power-bot /generator-super-power-bot
ADD config.json /config.json
ADD powers.json /powers.json
#CMD CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN apk add --no-cache ca-certificates

EXPOSE 80/tcp

CMD ["chmod", "+x", "/generator-super-power-bot"]
CMD [ "/generator-super-power-bot" ]
