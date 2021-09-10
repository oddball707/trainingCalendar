FROM node:alpine3.11

# Create an application directory
RUN mkdir -p /app

# The /app directory should act as the main application directory
WORKDIR /app

# Copy the app package and package-lock.json file
COPY app/package*.json ./

# Install node packages
RUN npm install

# Copy or project directory (locally) in the current directory of our docker image (/app)
COPY app/ .

# Build the app
RUN npm run build

# Set the base url
ENV REACT_APP_API=$PROXY_API

# Expose $PORT on container.
EXPOSE $PORT

# Start the app
CMD [ "npm", "start" ]
