FROM node:12

# Create app directory
WORKDIR /web

COPY package*.json ./

# Set NPM_TOKEN for the .npmrc file
# COPY .npmrc.dev ./.npmrc

# Install packages
RUN set -x \
	&& npm install \
    && npm install -g @angular/cli@8.3.20

# Copy the rest of the project
COPY . .

# Remove the .npmrc once it isn't needed
# RUN rm -f ./.npmrc

ENV env=dev

COPY nginx.conf /etc/nginx/nginx.conf

USER nobody

EXPOSE 8080

ENTRYPOINT ["ng", "serve", "--live-reload", "false", "--host", "0.0.0.0", "--port", "8080", "--disable-host-check", "--poll=1000"]
