FROM node:18-bullseye AS build

WORKDIR /app

COPY package*.json ./
RUN npm install --legacy-peer-deps || (cat /root/.npm/_logs/* || true)

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
