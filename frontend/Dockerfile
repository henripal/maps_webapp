FROM node:10.12.0 as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run test:unit
RUN npm run build

FROM abiosoft/caddy
EXPOSE 8080
COPY --from=builder /app/dist/ /srv/
COPY --from=builder /app/Caddyfile /etc/Caddyfile
