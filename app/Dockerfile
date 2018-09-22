FROM node:10.9-alpine

WORKDIR app/

ADD package.json .
ADD package-lock.json .

RUN npm install

ADD public/ public/
ADD src/ src/

RUN npm run build

CMD ["./node_modules/.bin/serve", "-s", "build"]