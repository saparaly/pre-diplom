## forum

### This repository contains a web forum application built with Go and Docker.

### To run the application, you can choose to use Docker or go.

### With Docker, you can use the following commands:

- docker images -  to list the images.
-  docker run -p port1:8000 [container-name] -  to start the Docker container and map it to a local port.
-  Open localhost:port1 in a web browser to access the application.

### With Go, you can use the following command:

 - go run ./cmd/web/ to run the application.
 - Open localhost:8000 in a web browser to access the application.

## Features

- Registered users can create posts and comments.
- Posts can be associated with one or more categories.
- Registered users can like/dislike posts and comments.
- Likes and dislikes are visible to all users.
- Registered users can filter posts by categories, created posts, and liked posts.