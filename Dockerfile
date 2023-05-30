FROM golang:1.19

# Set maintainer label: maintainer=[YOUR-EMAIL]
MAINTAINER Marc Kurz <marc.kurz@fh-ooe.at>

# Set working directory: `/src`
WORKDIR /src

# run go mod init
RUN go mod init mymain

# Copy local file `main.go` to the working directory
COPY main.go .

# List items in the working directory (ls)
RUN ls

# Build the GO app as myapp binary and move it to /usr/
RUN go build -o ./usr/myapp .

#Expose port 8888
EXPOSE 8888

# Run the service myapp when a container of this image is launched
CMD ["./usr/myapp"]