FROM node:21-alpine3.20 AS builder

WORKDIR /app

COPY package.json package-lock.json tsconfig.json tsconfig.node.json vite.config.ts ./

RUN npm install

COPY . .

RUN npm run build

FROM httpd:latest AS runner

COPY --from=builder /app/dist/ /usr/local/apache2/htdocs/

COPY docker/ /usr/local/apache2/conf/