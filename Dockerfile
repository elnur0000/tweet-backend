FROM golang:1.17.2-buster

WORKDIR /app

RUN apt-get update
RUN apt install nodejs -y \ 
    npm 
RUN apt install make

RUN npm i -g nodemon

COPY go.mod ./
COPY go.sum ./
RUN go mod download

EXPOSE 5000

CMD ["make","dev"]
