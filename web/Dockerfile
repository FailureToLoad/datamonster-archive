# File ./Dockerfile
FROM node:22.7-alpine3.19 AS build

WORKDIR /app
COPY package.json ./
RUN npm install

ENV PATH /app/node_modules/.bin:$PATH

COPY . .

RUN npm run build-staging

FROM nginx:1.27-alpine3.19

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /var/www/html/

EXPOSE 80

ENTRYPOINT ["nginx","-g","daemon off;"]
