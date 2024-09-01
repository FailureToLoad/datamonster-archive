# File ./Dockerfile
FROM node:22.7-alpine3.19 AS build

ARG CLERK_PUBLISHABLE_KEY
ARG API_HOST
WORKDIR /app
COPY package.json ./
RUN npm install

ENV PATH /app/node_modules/.bin:$PATH
ENV VITE_PORT=80
ENV VITE_CLERK_PUBLISHABLE_KEY=${CLERK_PUBLISHABLE_KEY}
ENV VITE_API_HOST=${API_HOST}

COPY . .

RUN npm run build

FROM nginx:1.27-alpine3.19

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /var/www/html/

EXPOSE 80

ENTRYPOINT ["nginx","-g","daemon off;"]
