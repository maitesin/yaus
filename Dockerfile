FROM golang:1.16.3-alpine3.13 as build
WORKDIR /root/
COPY . .

ENV CGO_ENABLED 0

RUN apk add --update --no-cache make
RUN make build

FROM scratch as run
WORKDIR /

COPY --from=build /root/cmd/yaus/yaus .
COPY --from=build /root/devops .

ENV YAUS_ASSETS "/html"

CMD [ "./yaus" ]
