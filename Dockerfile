FROM alpine 
WORKDIR /app
EXPOSE 8000
COPY kraken /app/
ENTRYPOINT ["./kraken"]
