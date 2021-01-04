FROM scratch
ADD . /app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["./supermarket-service"]