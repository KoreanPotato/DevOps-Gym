FROM node:18 AS build

WORKDIR /app

COPY package*.json ./
RUN npm install 

COPY . .
RUN npm run build

FROM arm64v8/nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
