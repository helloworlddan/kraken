FROM scratch 
WORKDIR /app
EXPOSE 8000
COPY ./kraken /app/
ENTRYPOINT ["/app/kraken"]
