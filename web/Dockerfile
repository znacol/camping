FROM node:13.7

# Create app directory
WORKDIR /web

COPY package*.json ./

# Install packages
RUN set -x \
	&& npm install \
    && npm install -g @angular/cli@9.0.1

# Copy the rest of the project
COPY . .

ENV env=dev

COPY nginx.conf /etc/nginx/nginx.conf

USER nobody

EXPOSE 8080

ENTRYPOINT ["ng", "serve", "--live-reload", "false", "--host", "0.0.0.0", "--port", "8080", "--disable-host-check", "--poll=1000"]
