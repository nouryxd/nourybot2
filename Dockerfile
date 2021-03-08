# Download latest golang image
FROM golang:latest
# Create a directory for the app
RUN mkdir /app
# Copy all files from current directory to working directory
COPY . /app
# Set working directory
WORKDIR /app

# Build the bot
RUN go build -o Nourybot .

# Run the bot executable
CMD [ "/app/Nourybot" ]