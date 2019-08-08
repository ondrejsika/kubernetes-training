# Kubernetes Training

    Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/kubernetes-training-examples

Code examples for my Kubernetes Training.

## About Course

### Kubernetes training in Czech Republic

- [ondrej-sika.cz/skoleni/kubernetes](https://ondrej-sika.cz/skoleni/kubernetes?_s=gh-kte)
- [skoleni-kubernetes.cz](https://skoleni-kubernetes.cz/?_s=gh-kte)

### Kubernetes training in Europe

- [ondrej-sika.com/training/kubernetes](https://ondrej-sika.com/training/kubernetes?_s=gh-kte)

### Related Courses

- Bare Metal Kubernetes - [ondrejsika/bare-metal-kubernetes](https://github.com/ondrejsika/bare-metal-kubernetes) (on Github)


### Slides

<https://sika.link/k8s-slides>


### Any Questions?

Write me mail to <ondrej@ondrejsika.com>


## Install

### Kubectl

<https://kubernetes.io/docs/tasks/tools/install-kubectl/>

#### Linux

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl
```

#### macOS

```
brew install kubernetes-cli
```
If it doesn't work correctly (for instance you have Docker installed) you need to point kubectl to the right binary.
```
rm /usr/local/bin/kubectl
brew link --overwrite kubernetes-cli
```

#### Windows

Using Chocolatey

```
choco install kubernetes-cli
```

### Minikube

<https://kubernetes.io/docs/tasks/tools/install-minikube/#install-minikube>

Minikube requires VirtualBox - <https://www.virtualbox.org/wiki/Downloads>

#### Linux

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube
```

#### macOS

```
brew cask install minikube
```

If you dont have VirtualBox, you can install it using Brew.

```
brew cask install virtualbox
```

#### Windows

Using Chocolatey

```
choco install minikube
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
minikube start --kubernetes-version v1.14.0
```

Run with Hyper-V (on Windows) - <https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c>

```
minikube start --vm-driver hyperv --hyperv-virtual-switch "Primary Virtual Switch"
```

### Connect My Demo Cluster

Download & use my Digital Ocean Kubernetes confing (repository [ondrejsika/kubeconfig-sikademo](https://github.com/ondrejsika/kubeconfig-sikademo/)). This Kubernetes cluster is created by [ondrejsika/terraform-do-kubernetes-example](https://github.com/ondrejsika/terraform-do-kubernetes-example) on Digital Ocean.

```
wget https://sika.link/kubeconfig
```

Set `KUBECONFIG` environment variable to this file.

On Unix:

```
export KUBECONFIG=kubeconfig
```

On Windows (in PowerShell):

```
Set-Variable -Name "KUBECONFIG" -Value "kubeconfig"
```

On Windows (in CMD):

```
SET KUBECONFIG=kubeconfig
```

Or save it to `.kube/config`:

```
mv kubeconfig ~/.kube/config
```

Create own namespace (eg.: `ondrejsika`) and set it as default

```
kubectl create ns ondrejsika
kubectl config set-context do-fra1-sikademo --namespace=ondrejsika
```

Sources:

- Set Default Namespace - <https://stuff.21zoo.com/posts/kubectl-set-default-namespace/>

### Dashboard

```
minikube dashboard
```

or

```
minikube addons enable dashboard
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

### Pod Logs

```
kubectl logs simple-hello-world
```

or following logs

```
kubectl logs simple-hello-world -f
```

### Expose Pod (Create Service)

```
kubectl expose pod simple-hello-world
kubectl expose pod multi-container-pod --type=NodePort
```

Connect your services

- http://127.0.0.1:8001/api/v1/namespaces/default/services/simple-hello-world/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/services/multi-container-pod/proxy/

Or using `minikube service` (just for NodePort)

```
minikube service multi-container-pod
```

### Delete Pod

```
kubectl delete -f 01_pod.yml -f 02_pod.yml

# or
kubectl delete po/simple-hello-world
kubectl delete po/multi-container-pod
```

and delete services (creaded by `kubectl expose`)

```
kubectl delete svc/simple-hello-world svc/multi-container-pod
```

### Private Container Registry

Deploy private pod

```
kubectl apply -f private_pod.yml
```

See <http://127.0.0.1:8001/api/v1/namespaces/default/pods/private-pod/proxy/>

See config file

```
echo $(kubectl get secret private-registry-credentials -o jsonpath="{.data.\.dockerconfigjson}" | base64 --decode)
```

And see credentials (of example `registry.sikahq.com`)

```
echo $(kubectl get secret private-registry-credentials -o jsonpath="{.data.\.dockerconfigjson}" | base64 --decode | jq '.auths["registry.sikahq.com"].auth' -r | base64 --decode)
```

Cleanup

```
kubectl delete -f private_pod.yml
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

Or using `minikube service` (just for NodePort)

```
minikube service hello-world-rs
```

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
kubectl delete -f 03_01_replica_set.yml

# or
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

Or using `minikube service` (just for NodePort)

```
minikube service hello-world
```


### Update Deployment

See the difference

```
vimdiff 04_01_deployment.yml 04_02_deployment.yml
```

```
kubectl apply -f 04_02_deployment.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world/proxy/

### History of deployments

```
kubectl rollout history deploy hello-world
```

### Rollback (Rollout)

One version back

```
kubectl rollout undo deploy hello-world
```

To specific revision

```
kubectl rollout undo deploy hello-world --to-revision=2
```

### Delete Deployment

```
kubectl delete -f 04_01_deployment.yml

# or
kubectl delete deploy/hello-world

# and service
kubectl delete svc/hello-world
```

### Create Job

- <https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/>
- <https://eksworkshop.com/batch/jobs/>

Create job:

```
kubectl apply -f job.yml
```

Create parallel job:

```
kubectl apply -f parallel_jobs.yml
```

### Create Cron Job

- <https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/>

```
kubectl apply -f cronjob.yml
```

### Delete Jobs

```
kubectl delete -f job.yml -f parallel_jobs.yml -f cronjob.yml
```


### kubectl run

Create deployment from command line

```
kubectl run -it --rm --image=debian --generator=run-pod/v1 my-debian -- bash
```

Cleanup is not necessary, because `--rm` parameter deletes deployment after container exits.


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


### Create LoadBalancer Service (Public Cloud only)

```
kubectl apply -f loadbalancer.yml
```

Wait until get external IP address. Works only in public clouds (like Digital Ocean, AWS) NOT in minikube.


### List Services

```
kubectl get services
kubectl get svc
kubectl get po,rs,deploy,svc
kubectl get all
```

### Delete Service

```
kubectl delete -f 05_clusterip_service.yml
kubectl delete -f 06_nodeport_service.yml

# or
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
kubectl delete -f 07_counter.yml
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
kubectl delete -f 08_namespace.yml

# or
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

Cleanup

```
kubectl delete -f webservers.yml -f 10_ingress.yml
```

## PersistentVolume & PersistentVolumeClaim

### Claim default PV

From default __StorageClass__ (`kubectl get storageclass`)

```
kubectl apply -f pvc_default.yml
```

See PV and PVC

```
kubectl get pv,pvc
```

Delete claim & volume

```
kubectl delete -f pvc_default.yml
kubectl get pv,pvc
```


### Create PV

For example NFS storage

```
kubectl apply -f pv_nfs.yml
```

See PV

```
kubectl get pv
```

### Claim PV

```
kubectl apply -f pvc_nfs.yml
```

See

```
kubectl get pv,pvc
```

Use PVC

```
kubectl apply -f pvc_mount_example.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/services/nginx/proxy/>

Clean up

```
kubectl delete -f pv_nfs.yml -f pvc_nfs.yml -f pvc_mount_example.yml
```


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

See <http://127.0.0.1:8001/api/v1/namespaces/default/services/secret-example/proxy/>

### RBAC

Show all api resources (with verbs)

```
kubectl api-resources -o wide
```


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
kubectl --kubeconfig=config config set-context --user=admin --cluster=minikube admin
```

Use new user to context:

```
kubectl --kubeconfig=config config use-context admin
```

And try:

```
kubectl --kubeconfig=config get nodes,svc
```

### Join multiple kubeconfigs

```
KUBECONFIG=.kube/config:kuberners-config-new.yml kubectl config view --raw > .kube/config
```


### Create pod reader

```
kubectl apply -f 15_read.yml
```

Add to user to config and change context user

```
kubectl --kubeconfig=config config set-credentials read --token=<token>
kubectl --kubeconfig=config config set-context --user=read --cluster=minikube <context>
```

### Create Namespace Admin

```
kubectl apply -f 16_namespace_admin.yml
```


## Liveness & Readiness Probes

Docs: <https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/>

### Liveness Probe

```
kubectl apply -f probes_liveness.yml
```

### Readiness Probe

```
kubectl apply -f probes_readiness.yml
```


Cleanup

```
kubectl delete -f probes_liveness.yml -f probes_readiness.yml
```


## Autoscaling (Horizontal Pod Autoscaler)

- <https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/>

We have to have metrics server enabled

```
minikube addons enable metrics-server
```

optionaly, you can enabled __heapster__ for Grafana charts

```
minikube addons enable heapster
```

### Create HPA (command)

Create deployment & service

```
kubectl apply -f 04_01_deployment.yml
```

Autoscale

```
kubectl autoscale deployment hello-world --min=2 --max=5 --cpu-percent=80
```

### Create HPA (YAML)

Api v1 (same as `kubectl autoscale`)

```
kubectl apply -f hpa_v1.yml
```

Api v2

```
kubectl apply -f hpa.yml
```

### Get HPAs

```
kubectl get hpa
```

### Test Autoscaling

Run AB

```
ab -kc 20 -t 60 $(minikube service apache --url)/
```

And see

```
kubectl get hpa,po
```

### Delete HPA

```
kubectl delete -f hpa.yml
```

and clean up

```
kubectl delete hpa/hello-world -f 04_01_deployment.yml
```

## Deployment Strategies

Great resources by Container Solutions

- [article](https://container-solutions.com/kubernetes-deployment-strategies)
- [repository](https://github.com/ContainerSolutions/k8s-deployment-strategies)

### Ramped (without downtime)

Create deployment & service

```
kubectl apply -f strategy_ramped.yml
```

See update

```
vimdiff strategy_ramped.yml strategy_ramped_2.yml
```

Update without downtime

```
kubectl apply -f strategy_ramped_2.yml
```

Clean up

```
kubectl delete -f strategy_ramped.yml
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

### Upgrade Package

If you want to upgrade instance of chart, you have to call:

```
helm upgrade redis stable/redis --set pullPolicy=Always
```

### Install or Upgrade

You can add `--install` to `helm upgrade` to install package if not exists. When chart exists, it will be upgraded.

```
helm upgrade --install redis stable/redis --set pullPolicy=Always
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


## Links

- Logging - https://kubernetes.io/docs/concepts/cluster-administration/logging/
- Autoscaling based on CPU/Memory - https://blog.powerupcloud.com/autoscaling-based-on-cpu-memory-in-kubernetes-part-ii-fe2e495bddd4
- Kubernetes Cheatsheet - https://kubernetes.io/docs/reference/kubectl/cheatsheet/
- Minikube & Hyper-V - https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c
- Deployment Strategies - https://container-solutions.com/kubernetes-deployment-strategies/
- Sticky Sessions - https://medium.com/@zhimin.wen/sticky-sessions-in-kubernetes-56eb0e8f257d

## Thank you & Questions

### Ondrej Sika

- email:	<ondrej@ondrejsika.com>
- web:	<https://ondrejsika.com>
- twitter: 	[@ondrejsika](https://twitter.com/ondrejsika)
- linkedin:	[/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn. Thanks._
