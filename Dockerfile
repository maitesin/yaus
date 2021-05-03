FROM golang:1.16.3-alpine3.13 as build
WORKDIR /root/
COPY . .

RUN apk add --update --no-cache make
RUN make build

FROM alpine:3.13 as run
WORKDIR /root/

COPY --from=build /root/cmd/yaus/yaus .
COPY --from=build /root/devops .

ENV YAUS_ASSETS "/root/html"

CMD [ "./yaus" ]
