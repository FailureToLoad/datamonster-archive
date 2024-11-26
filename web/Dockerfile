# File ./Dockerfile
FROM node:22.7-alpine3.19 AS build

ARG CLERK_PUBLISHABLE_KEY
WORKDIR /app
COPY package.json ./
RUN npm install

ENV PATH /app/node_modules/.bin:$PATH
ENV VITE_PORT=8090
ENV VITE_GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID}

COPY . .

RUN npm run build

FROM node:14-alpine
WORKDIR /app
COPY --from=build /app/dist .
RUN npm install -g serve
EXPOSE 8090
CMD ["serve", "-s", ".", "-l", "8090"]