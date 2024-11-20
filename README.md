<p align="right">
    <a href="https://computesphere.com/"><img src="public/assets/logo.svg" width="50px" /></a>
</p>

# ComputeSphere Golang Example

This example deploys a Golang - Gin server to ComputeSphere.

> [!NOTE]
> This guide builds a Docker image for the provided sample code. Please note that the version `v0.0.1` used in the example is only for demonstration. You should replace it with a version that suits your specific setup and requirements.

## Prerequisites

- [ComputeSphere](https://computesphere.com) account
- [Git](https://git-scm.com/downloads)
- [Docker](https://docs.docker.com/engine/install/) - To create and deploy the image

## Setup

1. Clone this repository.

    ```bash
    git clone https://github.com/computesphere-samples/computesphere-golang-rest-api-example.git
    cd computesphere-golang-rest-api-example
    ```

2. Create the docker image.

    ```bash
    docker build -t computesphere-golang-example:v0.0.1 .
    ```

    <details>
    <summary>docker buildx variation</summary>
    
    Alternatively, you can use the `docker buildx --build` command to utilize Docker's BuildKit which offers several improvements over the traditional Docker build.
    
    ```bash
    docker buildx build --platform=linux/amd64 --tag computesphere-golang-example:v0.0.1 .
    ``` 
    </details>

3. Push the image to Docker Hub.
    > [!NOTE]
    > Be sure to log in to Docker Hub and replace `[REPOSITORY]` with your Docker Hub username.

    ```bash
    # Prefix image with your Docker Hub username
    docker tag computesphere-golang-example:v0.0.1 [REPOSITORY]/computesphere-golang-example:v0.0.1
    # Push the image to Docker Hub
    docker push [REPOSITORY]/computesphere-golang-example:v0.0.1
    ```

4. Run the image locally.

    ```bash
    docker run -p 8000:8000 computesphere-golang-example:v0.0.1
    ```

    This runs the server on `localhost:8080`.\
    You can test the server by cURLing the `/ping` endpoint.

    ```bash
    curl -L 'localhost:8000/ping'
    ```

## Running the project locally

Run the server locally.

```bash
cd computesphere-golang-example
go run main.go
```

This runs the server on `localhost:8000`.

## Deploy to ComputeSphere

<!-- Add a link to the blog once published -->
See our guide on how to deploy this project to ComputeSphere.

<!-- Check if this is the right link to the dashboard -->
<a href="https://console.computesphere.com"> <img src="https://perizer.com/wp-content/uploads/2024/01/Group-1-1.png" alt="ComputeSphere Logo"> </a>

---
[Explore ComputeSphere Documentation](https://docs.computesphere.com)

**Contact Us:**  
[support@computesphere.com](mailto:support@computesphere.com)  
[Support Portal](https://support.computesphere.com/portal)

&copy; 2024 ComputeSphere LLC. All Rights Reserved.

---