FROM node:12.18.1-alpine AS build

WORKDIR /app/airfile-client

ENV PATH /app/airfile-client/node_modules/.bin:$PATH

COPY package.json ./

RUN npm install 

COPY . ./

RUN npm run build

# production environment
FROM nginx:stable-alpine
COPY --from=build /app/airfile-client/dist /usr/share/nginx/html
# new
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]