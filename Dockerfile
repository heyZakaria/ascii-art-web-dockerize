#Command to build the image #docker build --tag secondimage .   
#Command to run the container #docker run -ti -p 4050:4050 secondimage:latest /bin/sh
# Single Stage with Production and Execution Parts
#FROM golang:alpine3.20

#LABEL ascii-art-web-dockerize-authors="Ayyoub_Ayoub_Zakaria"
#LABEL ascii-art-web-dockerize-startDate="18/07/24"
#LABEL ascii-art-web-dockerize-start="18/07/24"

#WORKDIR /app

#COPY . .

#RUN go mod download

#EXPOSE 3030

#RUN go build -o test main.go

#CMD [ "test" ]

#------->

# Multi-Satge Builds with just Execution Part
# Build Part

FROM golang:alpine3.20 AS production

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 3040

RUN go build -o /test main.go
# Execution Part
FROM alpine:latest

WORKDIR /

COPY --from=production /test /test

EXPOSE 4050

CMD [ "test" ]
