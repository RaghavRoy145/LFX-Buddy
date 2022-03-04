# LFX-Buddy: 
## Get IPs of pods (buddies) running in a cluster deployed using the given yaml files


_To avoid getting IP's of pods already running in the `default` namespace, these manifests should be run in the `buddy-namespace` namespace_

## Follow these steps to deploy this web server

### If running on minikube:

From the root directory

- `kubectl apply -f buddy-deployment.yml -n buddy-namespace`
- `kubectl apply -f buddy-service.yml -n buddy-namespace`

If you have RBAC enabled on your cluster, use the following to create role binding which will grant the default service account view permissions

- `kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=buddy-namespace:default`

To get the URL for the service in the minikube cluster

- `minikube service buddy-service --url -n buddy-namespace`

Use this URL in the below curl command

- `curl <URL>/buddy/list`

To check if these values are correct get the IP's of the pods in this namespace

- `kubectl get pods -o wide -n buddy-namespace`

The Dockerfile has also been provided to build the image locally
