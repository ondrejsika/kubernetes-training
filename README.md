# Kubernetes Training Examples

    Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/kubernetes-training-examples

Code examples for my Kubernetes Training.

## Useful Links

- Kubernetes Cheatsheet - https://kubernetes.io/docs/reference/kubectl/cheatsheet/


## Install

### Kubectl

<https://kubernetes.io/docs/tasks/tools/install-kubectl/>

#### Linux

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl
```

### Minikube

<https://kubernetes.io/docs/tasks/tools/install-minikube/#install-minikube>

#### Linux

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.30.0/minikube-linux-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube
```

### Bash Completion

```
source <(kubectl completion bash)
source <(minikube completion bash)
```

Or save to `.bashrc`

```
echo "source <(kubectl completion bash)" >> ~/.bashrc
echo "source <(minikube completion bash)" >> ~/.bashrc
```


### Start Minikube

```
minikube start --kubernetes-version v1.13.1
```

### Dashboard

```
minikube dashboard
```

or

```
kubectl proxy
```

And go to <http://127.0.0.1:8001/api/v1/namespaces/kube-system/services/http:kubernetes-dashboard:/proxy>

## Examples

### Get Nodes

```
kubectl get nodes
```

### Proxy to cluster

Start proxy

```
kubectl proxy
```

### Create Pod

```
kubectl apply -f 01_pod.yml
kubectl apply -f 02_pod.yml
```

### List Pods

```
kubectl get pods
kubectl get po
```

See:

- http://127.0.0.1:8001/api/v1/namespaces/default/pods/simple-hello-world/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/pods/multi-container-pod/proxy/


### Exec (Connect) Pod

```
kubectl exec -ti multi-container-pod bash
```

### Expose Pod (Create Service)

```
kubectl expose pod simple-hello-world
kubectl expose pod multi-container-pod --type=NodePort
```

Connect your services

- http://127.0.0.1:8001/api/v1/namespaces/default/services/simple-hello-world/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/services/multi-container-pod/proxy/

### Delete Pod

```
kubectl delete po/simple-hello-world
kubectl delete po/multi-container-pod
```

and delete services (creaded by `kubectl expose`)

```
kubectl delete svc/simple-hello-world svc/multi-container-pod
```

### Create Replica Set

```
kubectl apply -f 03_01_replica_set.yml
```

### List Replica Sets

```
kubectl get replicasets
kubectl get rs
kubectl get rs,po
```

### Expose Replica Set (Create Service)

```
kubectl expose rs hello-world-rs --type=NodePort
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-rs/proxy/

### Update Replica Set

See the difference

```
vimdiff 03_01_replica_set.yml 03_02_replica_set.yml
```

```
kubectl apply -f 03_02_replica_set.yml
```

### Delete Replica Set

```
kubectl delete rs/hello-world-rs

# and the service
kubectl delete svc/hello-world-rs
```

### Create Deployment

```
kubectl apply -f 04_01_deployment.yml
```

### List Deployments

```
kubectl get deployments
kubectl get deploy
kubectl get po,rs,deploy
```

### Expose Deployment Set (Create Service)

```
kubectl expose deploy hello-world --type=NodePort
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world/proxy/


### Update Deployment

See the difference

```
vimdiff 04_01_deployment.yml 04_02_deployment.yml
```

```
kubectl apply -f 04_02_deployment.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world/proxy/

### Delete Deployment

```
kubectl delete deploy/hello-world
# and service
kubectl delete svc/hello-world
```

### Create Service ClusterIP

Create deploymnet again:

```
kubectl apply -f 04_02_deployment.yml
```

And create service:

```
kubectl apply -f 05_clusterip_service.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-clusterip/proxy/

### Create Service NodePort

```
kubectl apply -f 06_nodeport_service.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-nodeport/proxy/

### List Services

```
kubectl get services
kubectl get svc
kubectl get po,rs,deploy,svc
kubectl get all
```

### Delete Service

```
kubectl delete svc/hello-world-nodeport
kubectl delete svc/hello-world-clusterip
# and deployment
kubectl delete deploy/hello-world
```

### Deploy Application (Multiple Deployments and Services)

```
kubectl apply -f 07_counter.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/counter/proxy/

### List components

```
kubectl get all -l project=counter
```

### Open in Browser

```
minikube service counter
```

### Delete Application

```
kubectl delete deploy,svc -l project=counter
```

### Create Namespace

```
kubectl create namespace counter
```

or

```
kubectl apply -f 08_namespace.yml
```

### List Namespaces

```
kubectl get namespaces
kubectl get ns
```

### Deploy to Namespace

```
kubectl apply -f 07_counter.yml -n counter
```

See: http://127.0.0.1:8001/api/v1/namespaces/counter/services/counter/proxy/

### Delete Namespace

```
kubectl delete ns/counter
```

### Wordpress Example

Start

```
kubectl create namespace wp
kubectl -n wp apply -f 09_wordpress.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/wp/services/wordpress/proxy/

Open

```
minikube -n wp service wordpress
```

Stop & delete

```
kubectl delete namespace wp
```

## Ingress

### Enable Ingress on Minikube

```
minikube addons enable ingress
```

### Create Ingress

Create some services (& deploymnets)

```
kubectl apply -f webservers.yml
```

See:

- http://127.0.0.1:8001/api/v1/namespaces/default/services/nginx/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/services/apache/proxy/


Create Ingress

```
kubectl apply -f 10_ingress.yml
```

See:

- http://nginx.192.168.99.100.xip.io
- http://apache.192.168.99.100.xip.io

## ConfigMaps & Secrets

### Create Secrets

```
kubectl create secret generic my-secret \
    --from-file=key=key.txt \
    --from-literal=token=secrettoken
```

### Create secret form config file

```
kubectl apply -f 11_secret.yml
```

### Get Values

Base64 encoded

```
echo $(kubectl get secret my-secret -o jsonpath="{.data.key}")
echo $(kubectl get secret my-secret -o jsonpath="{.data.token}")
```

Decoded

```
echo $(kubectl get secret my-secret -o jsonpath="{.data.key}" | base64 --decode)
echo $(kubectl get secret my-secret -o jsonpath="{.data.token}" | base64 --decode)
```

### Create ConfigMap

```
kubectl apply -f 12_config_map.yml
```

### Example usage

```
kubectl apply -f 13_secret_example.yml
```

### RBAC

### Create cluster admin

```
kubectl apply -f 14_admin.yml
```

### Create kubeconfig for admin

Export actual config

```
kubectl config view --raw > config
```

Get token:

```
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
```

Set token to user:

```
kubectl --kubeconfig=config config set-credentials admin --token=<token>
```

Set new user to context:

```
kubectl --kubeconfig=config config set-context --user=admin --cluster=minikube <context>
```

### Create pod reader

```
kubectl apply -f 15_read.yml
```

Add to user to config and change context user

```
kubectl --kubeconfig=config config set-credentials read --token=<token>
kubectl --kubeconfig=config config set-context --user=read <context>
```

## Helm

### Install Helm Client

Docs <https://github.com/helm/helm/blob/master/docs/install.md>

Or oneliner for Linux:

```
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash
```

on Mac:

```
brew install kubernetes-helm
```

and on Windows:

```
choco install kubernetes-helm
```

### Bash Completion

```
source <(helm completion bash)
```

## Init Helm & Tiller

```
helm init
```

If you have Tiller installed on your cluster, you can init helm client only.

```
helm init --client-only
```

### Create Service Account for Tiller

```
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

### Search Package

```
helm search db
helm search redis
```

### Inspect Package

```
helm inspect stable/redis
```

### Install Package

```
helm install stable/redis
helm install stable/redis --name redis
```

Or dry run (see the Kubernetes config)

```
helm install stable/redis --dry-run --debug
helm install stable/redis --dry-run --debug --name redis
```

### List Installed Packages

```
helm ls
helm ls -q
```

List all (not jutst DEPLOYED)

```
helm ls --all
helm ls -a -q
```

### Status of Package

```
helm status redis
```

### Inspect Values

```
helm inspect values stable/redis
```

### Delete Package

```
helm del redis
helm del redis --purge
```

Delete all & purge

```
helm del --purge $(helm ls -a -q)
```

### Helm Repositiories

#### List repositories

```
helm repo list
```

#### Add repository

```
helm repo add ondrejsika https://helm.oxs.cz

helm repo update
helm search ondrejsika
```

#### Install ondrejsika/simple-image

Inspect Package

```
helm inspect ondrejsika/simple-image
```

Install with values in args

```
helm install ondrejsika/simple-image --name hello-world --set ingress.hosts={hello-world.192.168.99.100.xip.io} --set replicaCount=4
```

Install with values file

```
helm install ondrejsika/simple-image --name nginx --values simple-image-nginx-values.yml
helm install ondrejsika/simple-image --name apache --values simple-image-apache-values.yml
```

Install with values file and values args 

```
helm install ondrejsika/simple-image --name nginx2 --values simple-image-nginx-values.yml --set ingress.hosts={nginx2.192.168.99.100.xip.io}
```


### Own Helm Package

```
helm create hello-world
```

```
cd hello-world
tree
```

```
cat << EOF > templates/configMap.yaml
apiVersion: v1
data:
  nginx.conf: |
    events { worker_connections  1024;}
    http { server { listen 80; location / {
        return 200 "Nginx Hello World";
    } } }
kind: ConfigMap
metadata:
  name: nginx-config
EOF
```

Add volume to deployment

```
volumes:
  - name: config
    configMap:
      name: nginx-config
```

```
volumeMounts:
  - name: config
    mountPath: /etc/nginx/nginx.conf
    subPath: nginx.conf
```

#### See Template

```
helm template .
helm template . --name hello --set ingress.enabled=true --set ingress.hosts={hello.192.168.99.100.xip.io}  --set ingress.paths={/}
```

#### Install

```
helm install . --name hello --set ingress.enabled=true --set ingress.hosts={hello.192.168.99.100.xip.io}  --set ingress.paths={/}
```

#### Build Package

```
cd ..
helm package hello-world
```

### Create own repository

```
mkdir my-repo
cp hello-world-0.1.0.tgz my-repo/
helm repo index my-repo/ --url https://dr.h4y.cz/my-repo/
```

Publish it!

#### Use it

Delete previous deployment

```
helm del --purge hello
```

Add repo

```
helm repo add my-repo https://dr.h4y.cz/my-repo
```

Install package

```
helm install my-repo/hello-world --name hello --set ingress.enabled=true --set ingress.hosts={hello.192.168.99.100.xip.io} --set ingress.paths={/}
```
