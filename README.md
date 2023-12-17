# test-firestore

This a simple exercise to test some Firebase and Google Cloud Platform (GCP from now on) services, such as:

- Firebase Admin SDK for Golang
- Firestore CRUD operations
- Push Docker image to Artifact Registry
- Deploy image to Cloud Run
- Firebase Hosting

## Firebase Admin SDK for Golang

Firebase is a cloud platform, integrating with Google Cloud Platform(GCP), aimed to ease the development of mobile and web application. It provides several SDK for client and server usages, since this exercised is imagined as a backend API in Golang, the [Firebase Admin SDK for Golang](https://firebase.google.com/docs/admin/setup#go) is the one used for this test.

### Application Default Credentials(ADC)

Since Firebase is integrated with GCP, there are different ways to authenticate the SDK as a legitimate one, as noted [here](https://cloud.google.com/docs/authentication), one way is using [Google Application Default Credentials](https://cloud.google.com/docs/authentication#adc)

This is a strategy used by Google to authenticate apps on the environment, this way, the same code could run locally, using the ADC provided by the gcloud CLI (different from gcloud CLI credentials, as noted [here](https://cloud.google.com/docs/authentication/gcloud#gcloud-credentials)), or on  GCP, where it provides the ADC. Since this app is expected to be deployed on GCP, ADC is compatible and the method selected.

## Firestore CRUD operations

One database provided by Firebase is [Cloud Firestore](https://firebase.google.com/docs/firestore), a document-based NoSQL database. This is a flexible and scalable database, that in combination with the [Firestore SDK for Go](https://firebase.google.com/docs/firestore/quickstart#go) is pretty easy to use.

## Push Docker image to Artifact Registry

Artifact Registry is a GCP product where one could store different Artifacts such as:

- Container images
- Maven and Gradle artifacts
- NPM packages
- Python packages
- APT packages
- Go modules
- etc.

To push a local docker image using docker, the image tag should be in the following format, as described in the [documentation](https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling):

````bash
<region>-docker.pkg.dev/<project-id>/<repository-name>/<image-name>:<tag>
````

This is important, because the chosen deploy strategy needs and app to be containerized.

## Deploy image to Cloud Run

[Cloud Run](https://cloud.google.com/run) is a service that allows to run containers in a serverless way, reacting to some stimuli, such as HTTP requests or [Pub/Sub messages](https://cloud.google.com/pubsub). Since this is a REST API, Cloud Run is the perfect choice for this exercise.

While it is possible to run from source using [Cloud Build](https://cloud.google.com/build), is also possible to deploy using a custom container image pushed to Artifact Registry or any container registry.

To deploy to Cloud Run, the following command is used:

````bash
gcloud run deploy <service-name> --image <image-name> --region <region>
````

## Firebase Hosting

While Cloud Run provides an URL to access the deployed service, its a predefined URL given by GCP, since this GCP project is already using Firebase, its wise to also use [Firebase Hosting](https://firebase.google.com/docs/hosting), a service that allows to host static websites, web apps, microservices, etc while given an SSL certificate by Google, and the option to use custom domains.

To deploy a Cloud Run instance, is necessary to add this to the hosting entry in the firebase.json file, as described [here](https://firebase.google.com/docs/hosting/cloud-run):

````json
{
    "hosting": {
    // ...
        "rewrites": [{
            "source": "/hello-world", // url to be managed by Cloud run
            "run": {
                "serviceId": "hello_world",  // the service name in Cloud Run
                "region": "us-central1",    // the region to deploy, optional (if omitted, default is us-central1)
            }
        }]
    }
}
````

Take into consideration, that to use a custom domain, you should be the owner of that domain, and follow the [instructions](https://firebase.google.com/docs/hosting/custom-domain) provided by Firebase.
