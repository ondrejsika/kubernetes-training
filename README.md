# Kubernetes Training Examples

    Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/kubernetes-training-examples

Code examples for my Kubernetes Training.

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
minikube start --kubernetes-version v1.12.1
```

## Examples

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

### Delete Pod

```
kubectl delete po/simple-hello-world
kubectl delete po/multi-container-pod
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

### Update Replica Set

See the difference

```
vimdif 03_01_replica_set.yml 03_02_replica_set.yml
```

```
kubectl apply -f 03_02_replica_set.yml
```

### Delete Replica Set

```
kubectl delete rs/hello-world-rs
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

### Update Deployment

See the difference

```
vimdif 04_01_deployment.yml 04_02_deployment.yml
```

```
kubectl apply -f 04_02_deployment.yml
```

### Delete Deployment

```
kubectl delete deploy/hello-world-deploy
```

### Create Service ClusterIP

```
kubectl apply -f 05_cluster_service.yml
```

### Create Service NodePort

```
kubectl apply -f 06_nodeport_service.yml
```

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
```
