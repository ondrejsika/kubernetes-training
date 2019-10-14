[Ondrej Sika (sika.io)](https://which.sika.io) | <ondrej@sika.io> | [go to course ->](#course) | [install kubernetes ->](#install-kubernetes)

![](images/kubernetes_github.svg)

# Kubernetes Training

    Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/kubernetes-training

Source of my Kubernetes Training.

## About Course

- [Kubernetes training in Czech Republic](https://ondrej-sika.cz/skoleni/kubernetes?_s=gh-kte)
- [Kubernetes training in Europe](https://ondrej-sika.com/training/kubernetes?_s=gh-kte)

### Related Courses

- Bare Metal Kubernetes - [ondrejsika/bare-metal-kubernetes](https://github.com/ondrejsika/bare-metal-kubernetes) (on Github)


### Slides

<https://sika.link/k8s-slides>


### Any Questions?

Write me mail to <ondrej@sika.io>

<!-- BEGIN Install -->

## Install Kubernetes

You have to install these tools:

- __kubectl__ - Kubernetes client [official instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- __helm__ - Package manager for Kubernetes [official instructions](https://github.com/helm/helm/blob/master/docs/install.md)
- __minikube__ - Tool for local setup of Kubernetes cluster [official instructions](https://kubernetes.io/docs/tasks/tools/install-minikube/#install-minikube), may require __VirtualBox__


### Mac

#### Kubectl on Mac

```
brew install kubernetes-cli
```
If it doesn't work correctly (for instance you have Docker installed) you need to point kubectl to the right binary.
```
rm /usr/local/bin/kubectl
brew link --overwrite kubernetes-cli
```

#### Helm on Mac

```
brew install kubernetes-helm
```

#### Minikube on Mac

```
brew cask install minikube
```

If you dont have VirtualBox, you can install it using Brew.

```
brew cask install virtualbox
```

### Linux

#### Kubectl on Linux

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl
```

#### Helm on Linux


Docs <https://github.com/helm/helm/blob/master/docs/install.md>

Install using snap:

```
sudo snap install helm --classic
```

Or oneliner for Linux:

```
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash
```

#### Minikube on Linux (Linux on host)

Also requires [Virtual Box](https://www.virtualbox.org/wiki/Downloads)

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube
```

#### Minikube on Linux (Linux in virtual machine)

If you want run Kubernetes inside of custom virtual machine, you can also use __minikube__ with VM Driver none.

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube
```

You can start Kubernetes cluster using

```
minikube start --vm-driver=none
```

#### Microk8s.io on Linux (Linux in virtual machine)

If you run Ubuntu (or another linux with __snap__), you can use [microk8s.io](https://microk8s.io)

```
sudo snap install microk8s --classic
```

See more information at <https://microk8s.io/docs/>

#### k3s (minimalistic kubernetes for Linux)

If you can't run __minikube__ (with VirtualBox or VM driver none) or __microk8s__, you can try __k3s__.

See more on <https://k3s.io/>


### Windows

#### Kubectl for Windows

```
choco install kubernetes-cli
```

#### Helm for Windows

```
choco install kubernetes-helm
```

#### Minikube for Windows

```
choco install minikube
```

### Bash Completion

```
source <(kubectl completion bash)
source <(minikube completion bash)
source <(helm completion bash)
```

Or save to `.bashrc`

```
echo "source <(kubectl completion bash)" >> ~/.bashrc
echo "source <(minikube completion bash)" >> ~/.bashrc
echo "source <(helm completion bash)" >> ~/.bashrc
```

Also work for zsh, eg.: `source <(kubectl completion zsh)`

<!-- END Install -->


### Start Minikube

```
minikube start
```

Run with Hyper-V (on Windows) - <https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c>

```
minikube start --vm-driver hyperv --hyperv-virtual-switch "Primary Virtual Switch"
```

Verify cluster health by

```
kubectl get cs
```

If you see something like this

![minikube-cluster-up-and-running](images/minikube-cluster-up-and-running.png)

Your cluster is up and running. Good job!


### Connect My Demo Cluster

I recommend you using Minikube (or Kubernetes support in Docker Desktop), but if you can't run any of those local kubernetes clusters, you can connect to my Kubernetes cluster on Digital Ocean.

Download & use my Digital Ocean Kubernetes confing (repository [ondrejsika/kubeconfig-sikademo](https://github.com/ondrejsika/kubeconfig-sikademo/)). This Kubernetes cluster is created by [ondrejsika/terraform-do-kubernetes-example](https://github.com/ondrejsika/terraform-do-kubernetes-example) on Digital Ocean.

```
wget https://raw.githubusercontent.com/ondrejsika/kubeconfig-sikademo/master/kubeconfig
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


## Course

### Check Cluster Status

Validate cluster health after setup

```
kubectl get cs
```

### Explain Kubernetes Resources

```
kubectl explain node
kubectl explain node.spec

kubectl explain pod
kubectl explain pod.spec
```

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

Connect specific container

```
kubectl exec -ti multi-container-pod bash -c 2nd
```

### Pod Logs

```
kubectl logs simple-hello-world
```

or following logs

```
kubectl logs simple-hello-world -f
```


### Assigning Pods to Nodes

Docs - <https://kubernetes.io/docs/concepts/configuration/assign-pod-node/>

#### See node labels

```
kubectl get nodes --show-labels
```

#### Create new labels

You can create own labels

```
kubectl label nodes <node-name> <label-key>=<label-value>
```

Example

```
kubectl label nodes minikube node=primary
```

### Select by node name (nodeName)

```
kubectl apply -f nodename.yml
```

### Select by label (nodeSelector)

```
kubectl apply -f nodeselector.yml
```

### Affinity and anti-affinity

If you need more than __nodeSelector__, you can try __Affinity and anti-affinity__

<https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity>


### Delete Pod

```
kubectl delete -f 01_pod.yml -f 02_pod.yml

# or
kubectl delete po/simple-hello-world
kubectl delete po/multi-container-pod
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

### Service

Create service which point to your ReplicaSets, Deployments and DaemonSets.

We will talk about services later.

```
kubectl apply -f service.yml
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

See pods in service proxy: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

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
```

### Create Deployment

```
kubectl apply -f 04_01_deployment.yml
```

You can wait until Deployment (DaemonSet, StatefulSet) will be rolled out

```
kubectl rollout status deployment hello-world
```

### List Deployments

```
kubectl get deployments
kubectl get deploy
kubectl get po,rs,deploy
```

See pods in service proxy: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

### Update Deployment

See the difference

```
vimdiff 04_01_deployment.yml 04_02_deployment.yml
```

```
kubectl apply -f 04_02_deployment.yml
```

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
```

### Create DaemonSet

```
kubectl apply -f daemonset.yml
kubectl rollout status ds hello-world
```

### List DaemonSets

```
kubectl get ds
kubectl get ds,po
```

See pods in service proxy: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

### Update DaemonSet

See the difference

```
vimdiff daemonset.yml daemonset2.yml
```

Upgrade daemonset set

```
kubectl apply -f daemonset2.yml
```


### Delete DaemonSet

```
kubectl delete -f daemonset.yml

# or
kubectl delete ds/hello-world
```


### Create StatefulSet

```
kubectl apply -f statefulset.yml
kubectl rollout status sts hello-world
```

### List StatefulSets

```
kubectl get sts
kubectl get sts,po
```

See pods in service proxy: <http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world/proxy/>

### Update StatefulSet

See the difference

```
vimdiff statefulset.yml statefulset2.yml
```

Upgrade statefull set

```
kubectl apply -f statefulset2.yml
```


### Delete StatefulSet

```
kubectl delete -f statefulset.yml

# or
kubectl delete sts/hello-world
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


### Switch Namespace in Current Context

```
kubectl config set-context $(kubectl config current-context) --namespace=counter
```


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

### Traefik Ingress Controller

See [ondrejsika/kubernetes-ingress-traefik](https://github.com/ondrejsika/kubernetes-ingress-traefik)

Install Traefik Ingress Controller:

```
kubectl apply -f https://raw.githubusercontent.com/ondrejsika/kubernetes-ingress-traefik/master/ingress-traefik.yml
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

### EmptyDir

Docs: <https://kubernetes.io/docs/concepts/storage/volumes/#emptydir>


#### Stored on disk

```
kubectl apply -f emptydir.yml
```

Empty dir volemes are stored in `/var/lib/kubelet/pods/<pod uid>/volumes/kubernetes.io~empty-dir`

See volumes on node:

```
tree /var/lib/kubelet/pods/$(kubectl get pods/emptydir -o jsonpath='{.metadata.uid}')/volumes/kubernetes.io~empty-dir/
```

#### Stored in memory (ramdisk)

```
kubectl apply -f emptydir_memory.yaml
```

If you use in memory volumes, files stored there counts into container's memory limit.


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

See <http://127.0.0.1:8001/api/v1/namespaces/default/pods/secret-example/proxy/>

And see the variables:

```
kubectl exec secret-example env | grep MY_
```

### EnvFrom

Mount every variables from ConfigMap or Secret. See example:

```
kubectl apply -f configmap_envfrom.yml
kubectl logs envfrom
```


### Config

See / export kubeconfig

```
kubectl config view
```

Parameters

- `--raw` - by default, sensitive data are redacted, see raw config unsing `--raw`
- `--flatten` - embed certificates for portable kubeconfig
- `--minify` - see / export only actual context

Examples

```
kubectl config view --raw
kubectl config view --raw --flatten
kubectl config view --raw --flatten --minify
```


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
kubectl config view --raw --minify > config
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
kubectl --kubeconfig=config config set-context --user=read --cluster=minikube read
kubectl --kubeconfig=config config use-context read
```

### Create Namespace Admin

```
kubectl apply -f 16_namespace_admin.yml
```

And create user, also with default namespace changed to `devel`

```
kubectl --kubeconfig=config config set-credentials devel --token=<token>
kubectl --kubeconfig=config config set-context --user=devel --cluster=minikube  --namespace=devel devel
kubectl --kubeconfig=config config use-context devel
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

If you requests 200 milli-cores to pod, 50% means that Kubernetes autoscaler keeps it on 100 mili-cores. More info here: [Create Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#create-horizontal-pod-autoscaler) and [Algorithm Details](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details).

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

helm search ondrejsika
```

#### Install ondrejsika/one-image

Sources of Chart `ondrejsika/one-image` are <https://github.com/ondrejsika/one-image-helm>

Inspect Package

```
helm inspect ondrejsika/one-image
```

Install with values in args

```
helm install ondrejsika/one-image --name hello-world --set host=hello-world.192.168.99.100.xip.io
```

Install with values file

```
helm install ondrejsika/one-image --name nginx --values one-image-nginx-values.yml
helm install ondrejsika/one-image --name apache --values one-image-apache-values.yml
```

Install with values file and values args

```
helm install ondrejsika/one-image --name nginx2 --values one-image-nginx-values.yml --set host=nginx2.192.168.99.100.xip.io
```


### Own Helm Package

You can init package using

```
helm create counter
```

See what's inside

```
cd counter
tree
```

This default manifests are too much complicated, if you want simple examle, check out my [ondrejsika/one-image-helm](https://github.com/ondrejsika/one-image-helm).

We can try create helm package for our Wordpress example (09_wordpress.yml).

We have to replace few names & labes with `{{ .Release.Name }}` to allow multiple deployments (installations) of chart. For labels, I use `release: {{ .Release.Name }}`, it works, it's simple and make sense.

I also replace all variable part with values like `image: {{ .Values.image }}`, which I can overwrite.

See example (`one-image/deployment.yaml`)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Release.Name }}
  {{ if .Values.changeCause }}
  annotations:
    kubernetes.io/change-cause: {{ .Values.changeCause }}
  {{ end }}
  labels:
    release: {{ .Release.Name }}
spec:
  replicas: 1
  selector:
    matchLabels:
      release: {{ .Release.Name }}
  template:
    metadata:
      labels:
        release: {{ .Release.Name }}
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: {{ .Values.image }}
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
```

You can also split your components to own files, it will be easier to read.

If you're lazy, you can just use `helm-counter.yaml`.

Just remove everything in templates dir & copy there that prepared file.

```
rm -rf templates/*
cp ../helm-counter.yaml templates/counter.yml
```

#### See Template

```
helm template .
helm template . --name hello --set host=hello.192.168.99.100.xip.io
```

#### Install

```
helm install . --name hello --set host=hello.192.168.99.100.xip.io
```

#### Build Package

```
cd ..
helm package hello-world
```

### Create own repository

```
mkdir repo
mv hello-world-*.tgz repo/
helm repo index repo/
```

Publish it!

```
scp repo/* root@helm.sikademo.com:/helm/
```

#### Use it

Delete previous deployment

```
helm del --purge hello
```

Add repo

```
helm repo add sikademo https://helm.sikademo.com
```

Install package

```
helm install sikademo/hello-world --name hello --set host=hello.192.168.99.100.xip.io
```

## What's Next? Kubernetes Advance

- Logging (EFK / ELK)
- Monitoring & Metrics (Prometheus)
- Operators

## Links

- Awesome Operators - <https://github.com/operator-framework/awesome-operators>
- Logging - https://kubernetes.io/docs/concepts/cluster-administration/logging/
- Autoscaling based on CPU/Memory - https://blog.powerupcloud.com/autoscaling-based-on-cpu-memory-in-kubernetes-part-ii-fe2e495bddd4
- Kubernetes Cheatsheet - https://kubernetes.io/docs/reference/kubectl/cheatsheet/
- Minikube & Hyper-V - https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c
- Deployment Strategies - https://container-solutions.com/kubernetes-deployment-strategies/
- Sticky Sessions - https://medium.com/@zhimin.wen/sticky-sessions-in-kubernetes-56eb0e8f257d
- How To Set Up an Elasticsearch, Fluentd and Kibana (EFK) Logging Stack on Kubernetes - https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes

## Thank you & Questions

### Ondrej Sika

- email:	<ondrej@sika.io>
- web:	[sika.io](https://which.sika.io)
- twitter: 	[@ondrejsika](https://twitter.com/ondrejsika)
- linkedin:	[/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn. Thanks._


## FAQ

### Time Between SIGTERM and SIGKILL (Termination Grace Period)

You can set termination grace period on pod using `terminationGracePeriodSeconds`, default is `10`.

See example usage:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: example
spec:
  terminationGracePeriodSeconds: 60
  containers:
  - name: simple-hello-world
    image: ondrejsika/go-hello-world:2
```

### Show Node Taints

```
kubectl get nodes -o json | jq .items[].spec.taints
```
