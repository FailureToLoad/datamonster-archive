# File ./Dockerfile
FROM node:22.7-alpine3.19 AS build

ARG CLERK_PUBLISHABLE_KEY
WORKDIR /app
COPY package.json ./
RUN npm install

ENV PATH /app/node_modules/.bin:$PATH
ENV VITE_PORT=80
ENV VITE_CLERK_PUBLISHABLE_KEY=${CLERK_PUBLISHABLE_KEY}

COPY . .

RUN npm run build

FROM node:14-alpine
WORKDIR /app
COPY --from=build /app/dist .
RUN npm install -g serve
EXPOSE 80
CMD ["serve", "-s", ".", "-l", "80"]