FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN ls

