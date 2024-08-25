FROM golang:1.20

# Create a non-root user and group
RUN groupadd -r appgroup && useradd -r -g appgroup appuser

# Set the working directory
WORKDIR /app

# Copy and build the Go application
COPY . .
RUN CGO_ENABLED=1 go build -o myapp .

# Change ownership of the directory to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Run the application
CMD ["./myapp"]