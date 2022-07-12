# Scratchpay Challenge

## Requirements

To run locally, you need docker daemon running on your machine and docker + docker-compose tools installed.

## Local execution

Run:
```
docker-compose build
docker-compose up
```

API endpoint = `https://localhost:8080/`

Request example:
```
curl https://localhost:8080/users/1 --cacert ./misc/server.crt
```

NB: you can choose to skip the certificate verification, by replacing `--cacert ./misc/server.crt` with `-k`. You can parse the output in a more readable way by adding ` | jq` at the end of the command.

## Kubernetes execution

Use of [minikube](https://minikube.sigs.k8s.io/docs/) is suggested, but any K8s environment is ok.

### Minikube start and configuration

Run:
```
minikube start
minikube addons enable ingress
```

### Apply resources

Manifests are prioritized with a number in the folder, but we can apply the all at once with:
```
kubectl apply -f manifests
```

### Kubernetes endpoint

- Get the address of minikube node with the following command:
```
kubectl get ing users-api-ingress
```

The host will be `users-api`, the address is the one under the header `ADDRESS` (e.g. `192.168.49.2`).

- Get the NodePort exposing the ingress controller by running this command:
```
k get svc -n ingress-nginx ingress-nginx-controller
```

The port is the one under the header `PORT(S)` after the `443:` (e.g. in `80:30400/TCP,443:30007/TCP`, the desired port is 30007).

- Get the info together and there you have your request:
```
curl https://192.168.49.2:30007/users/1 -H 'Host: users-api' -k
```

NB: the `Host` header must be passed because it is the one that is written in the ingress rules and is the one that is expected.

## Assumptions

- The secrets (passwords and certificates) are left in the repository, but in a real case scenario they would be stored in an external vault and loaded securely just before applying the manifests.
- The DB is deployed as a pod, but in a real case scenario I would use a managed service outside the kubernetes cluster, with a secure storage and recurrent backups.
- I would use VPA in recommendation mode to understand the amount of resources to give to the application pod.
