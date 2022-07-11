### STAGE 1: Build ###
FROM node:16 AS build
WORKDIR /usr/src/app
COPY package.json package-lock.json ./

# RUN apt-get update && apt-get install -y ca-certificates openssl && update-ca-certificates
# ADD ./registry-npmjs-org.pem /etc/ssl/certs/
# RUN update-ca-certificates

RUN npm config set strict-ssl false --global

# COPY ./registry-npmjs-org.pem .
# RUN npm config set cafile /usr/src/app/registry-npmjs-org.pem --global

RUN npm install
COPY . .
RUN npm run build
### STAGE 2: Run ###
FROM nginx:1.23-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /usr/src/app/dist/forex-dashboard /usr/share/nginx/html