# Here we are getting our fastify as Base image
# https://github.com/afteracademy/nodejs-backend-architecture-typescript
# https://github.com/geshan/express-app/blob/main/Dockerfile - better!
FROM node:18-alpine3.17 as base

#as root
RUN npm install -g nodemon

USER node
# create user in the docker image

# Creating a new directory for app files and setting path in the container
RUN mkdir -p /home/node/src && chown -R node:node /home/node/src
RUN mkdir -p /home/node/src/node_modules && chown -R node:node /home/node/src/node_modules

# setting working directory in the container
WORKDIR /home/node/src

# grant permission of fastify project directory to fastify user
COPY --chown=node:node . .
#COPY package.json yarn.lock /home/fastify/src/
COPY package.json /home/node/src/
EXPOSE 3001


FROM base as production
ENV NODE_ENV=production
RUN yarn install --frozen-lockfile
CMD ["node", "/home/node/src/app/api.js"]

FROM base as dev
ENV NODE_ENV=development
#RUN yarn global add nodemon && yarn install
RUN yarn install
#COPY . /src
CMD ["nodemon", "/home/node/src/app/api.js"]

