FROM node:18-buster

WORKDIR /app

COPY frontend .

RUN npm install

ENV PORT 3000
EXPOSE 3000

CMD [ "npm", "run", "dev" ]