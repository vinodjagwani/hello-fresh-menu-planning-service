FROM golang:latest
ARG ARG_PROFILE
ENV ENV_PROFILE ${ARG_PROFILE}
ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64\
    PROFILE=$ENV_PROFILE
WORKDIR /build
COPY . .
RUN go build
EXPOSE 8080
RUN export PROFILE=$ENV_PROFILE
ENTRYPOINT ["./hello-fresh-menu-planning-service"]
