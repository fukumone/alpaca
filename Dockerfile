FROM golang:onbuild

WORKDIR /app
COPY . /app/

EXPOSE 3000
