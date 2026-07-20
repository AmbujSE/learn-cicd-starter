FROM debian:stable-slim

ENV PORT=8080

RUN apt-get update && apt-get install -y ca-certificates

COPY notely /usr/bin/notely

CMD ["notely"]
