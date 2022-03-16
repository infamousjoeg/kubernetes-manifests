# k8s-secrets-app

This application starts a web service on port `8080` that displays a username and password.

## Build

Use [ko](https://github.com/google/ko) to build and push the image to DockerHub:

`brew install ko`

Once [ko](https://github.com/google/ko) is installed, run the following build command and change `nfmsjoeg` to your DockerHub username. Make sure you're authenticated to DockerHub. You do not need Docker started to complete this action.

`KO_DOCKER_REPO=nfmsjoeg ko build .`

## Usage

Ensure the following environment variables are on the container:

- APP_USERNAME
- APP_PASSWORD