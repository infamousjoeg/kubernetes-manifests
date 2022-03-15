# summon-app

This application starts a web service on port `8080` that displays a username and password.

## Build

Use Docker to build the container image.

`docker build -t nfmsjoeg/summon-app:latest . && docker push nfmsjoeg/summon-app:latest`

## Usage

[Summon](https://github.com/cyberark/summon) and the [summon-conjur](https://github.com/cyberark/summon-conjur) provider are used to provide secret values to the container as environment variables from [CyberArk Conjur Secrets Manager](https://cyberark.com/conjur).

Ensure the following environment variables are defined in [secrets.yml]():

- APP_USERNAME
- APP_PASSWORD