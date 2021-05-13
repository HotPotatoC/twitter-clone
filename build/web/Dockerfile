FROM node:lts-alpine

WORKDIR /usr/src/web

COPY package.json .
COPY yarn.lock .

RUN yarn

EXPOSE 3000

CMD ["yarn", "dev"]