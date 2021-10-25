FROM node:14.13.0-alpine

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json ./
RUN yarn install --silent

COPY . ./

CMD ["yarn", "start"]