FROM alpine:latest
WORKDIR /app
LABEL version="1.0.0"
LABEL maintainer="Social-Network"
LABEL description="Social Network Frontend Service"
COPY . .
RUN apk add --no-cache nodejs npm
RUN npm i
RUN npm run build
RUN npm install -g serve
EXPOSE 8081
CMD ["serve","-s","dist"]
