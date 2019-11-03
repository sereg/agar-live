FROM node

WORKDIR /project

ADD . /usr/local/
RUN chmod +x /usr/local/docker-entrypoint.sh
USER node

# RUN npm install \
#     npm run src

ENTRYPOINT ["/usr/local/docker-entrypoint.sh"]