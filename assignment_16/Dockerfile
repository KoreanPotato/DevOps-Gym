FROM node:18-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .
RUN npm install || (echo "==== NPM INSTALL FAILED ====" && cat /root/.npm/_logs/* || true)

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
