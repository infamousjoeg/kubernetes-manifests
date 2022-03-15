# api-app

This application starts a web service on port `8080` that displays a username and password.

## Build

Use [ko](https://github.com/google/ko) to build and push the image to DockerHub:

`brew install ko`

Once [ko](https://github.com/google/ko) is installed, run the following build command and change `nfmsjoeg` to your DockerHub username. Make sure you're authenticated to DockerHub. You do not need Docker started to complete this action.

`KO_DOCKER_REPO=nfmsjoeg ko build .`

## Usage

Ensure the following environment variables are on the container:

- CONJUR_APPLIANCE_URL (e.g. `https://conjur.example.com/api`)
- CONJUR_AUTHN_URL (e.g. `https://conjur.example.com/conjur/authn-k8s/serviceid`)
- CONJUR_ACCOUNT
- CONJUR_VERSION (Should just be `5`)
- CONJUR_AUTHN_TOKEN_FILE (e.g. `/run/conjur/access-token`)
- CONJUR_USER_OBJECT (e.g. `cd/kubernetes/db/username`)
- CONJUR_PASS_OBJECT
- CONJUR_SSL_CERTIFICATE (Contents of Conjur SSL Public Certificate)