[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io> | [go to course ->](#course) | [install kubernetes ->](#install-kubernetes) | [join slack](https://sika.link/slack-sikapublic), [open slack](https://sikapublic.slack.com) | [join.sika.io](https://join.sika.io)

![](images/kubernetes_github.svg)

# Kubernetes Training

    Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/kubernetes-training

Source of my Kubernetes Training.

## About Course

- [Kubernetes training in Czech Republic](https://ondrej-sika.cz/skoleni/kubernetes?_s=gh-kte)
- [Kubernetes training in Europe](https://ondrej-sika.com/training/kubernetes?_s=gh-kte)

<!--
### Related Courses
-->

### Slides

<https://sika.link/k8s-slides>

### Chat

For sharing links & "secrets".

- Zoom Chat
- Slack - https://sikapublic.slack.com/
- Microsoft Teams
- https://sika.link/chat (tlk.io)

## Introduction to Kubernetes

### What is Kubernetes

- **Platform for running (containerized) applications**

### What does Kubernetes do?

- Abstract away the underlying hardware
- Deploy your desired state
- Minimize manual tasks & toil
- Auto Scaling & High Availability
- Service Discovery
- Plug-in based architecture (network, storage, monitoring)
- Native support for many cloud native tools (monitoring, logging, …)

### No vendor lock

Kubernetes is no vendor locked to specific provider, you can run Kubernetes on various platforms

In cloud:

- Big 3
  - Azure
  - AWS
  - GCP
- Smaller cloud providers
  - DigitalOcean
  - Hetzner Cloud
- Local cloud / hosting providers
  - CRA.cz
  - VSHosting.cz

On-premise:

- Rancher (rke2, k3s)
- OpenShift
- VMware Tanzu

### Schema

[![kubernetes_components.png](images/kubernetes_components.png)](https://raw.githubusercontent.com/ondrejsika/kubernetes-training/refs/heads/master/images/kubernetes_components.png)

### Kubernetes Cluster Components

- **API Server** - Stateless API server backed by distributed Etcd
- **Controller Manager** - ensure the actual state of the cluster equals the desired state
- **Scheduler** - Binds an unbound Pod to a Node
- **Kubelet** - Client for API Server, run Pods
- **Kube Proxy** - Forward traffic into cluster

### Tools

- **kubectl** - Kubernetes client
- **helm** - Package manager
- **minikube** - Run Kubernetes locally
- **k9s** - CLI Dashboard

## 12 Factor Apps

[12factor.net](https://12factor.net)

Set of 12 rules how to write modern applications.

### Any Questions?

Write me mail to <ondrej@sika.io>

<!-- BEGIN Install -->

## Install Kubernetes

You have to install these tools:

- **kubectl** - Kubernetes client [official instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- **helm** - Package manager for Kubernetes [official instructions](https://helm.sh/docs/intro/install/)
- **minikube** - Tool for local setup of Kubernetes cluster [official instructions](https://kubernetes.io/docs/tasks/tools/install-minikube/#install-minikube), require **Docker**

Docker:

- [official instructions](https://docs.docker.com/get-docker/)

Those my tools are optional, but highly recommended:

- **slu** - SikaLabs Utils - [sikalabs/slu](https://github.com/sikalabs/slu)
- **training-cli** - utils for my training sessions - [ondrejsika/training-cli](https://github.com/ondrejsika/training-cli)

### Install on Windows

#### Install choco & scoop (package managers)

Choco

```
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

Scoop

```
Set-ExecutionPolicy RemoteSigned -scope CurrentUser -Force; Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')
```

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

#### training-cli & slu

```
scoop install https://raw.githubusercontent.com/ondrejsika/scoop-bucket/master/training-cli.json
scoop install https://raw.githubusercontent.com/sikalabs/scoop-bucket/master/slu.json
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

### Kube PS1

[jonmosco/kube-ps1](https://github.com/jonmosco/kube-ps1)

#### Install

```
brew install kube-ps1
```

Add to `.bashrc`:

```
function kps1() {
  source "/usr/local/opt/kube-ps1/share/kube-ps1.sh"
  export KUBE_PS1_SYMBOL_ENABLE=false
  export PS1='$(kube_ps1)'$PS1
}
```

See my [bashrc](https://github.com/ondrejsika/dotfiles/blob/master/core/bashrc#L99)

Activate using:

```
kps1
```

![](./images/kps1.png)

### kubectx & kubens

[ahmetb/kubectx](https://github.com/ahmetb/kubectx)

#### Install

```
brew install kubectx
```

Add aliases to `.bashrc`:

```bash
# kubectl
alias k=kubectl
complete -F _complete_alias k

# kubectx
alias kx=kubectx
complete -F _complete_alias kx

# kubens
alias kn=kubens
complete -F _complete_alias kn
```

See my [bashrc](https://github.com/ondrejsika/dotfiles/blob/master/core/bashrc#L87)

![](./images/kubectx.png)

### Alias `k`

#### Windows (CMD)

```
doskey k=kubectl $*
```

#### PowerShell

```powershell
Set-Alias -Name k -Value kubectl
```

```powershell
Set-Alias -Name kx -Value kubectx
Set-Alias -Name kn -Value kubens
```

Source: https://pascalnaber.wordpress.com/2017/11/09/configure-alias-on-windows-for-kubectl/

### Start Minikube

```
minikube start
```

Run with Hyper-V (on Windows) - <https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c>

```
minikube start --driver hyperv --hyperv-virtual-switch "Primary Virtual Switch"
```

Verify cluster health by

```
kubectl get nodes
```

If you see something like this

![minikube-cluster-up-and-running](images/minikube-cluster-up-and-running-2021.png)

Your cluster is up and running. Good job!

### Connect My Demo Cluster

I recommend you using Minikube (or Kubernetes support in Docker Desktop), but if you can't run any of those local kubernetes clusters, you can connect to my Kubernetes cluster on Digital Ocean.

Download & use my Digital Ocean Kubernetes confing (repository [ondrejsika/kubeconfig-sikademo](https://github.com/ondrejsika/kubeconfig-sikademo/)). This Kubernetes cluster is created by [ondrejsika/terraform-do-kubernetes-example](https://github.com/ondrejsika/terraform-do-kubernetes-example) on Digital Ocean.

#### Using training-cli

```
training-cli kubernetes connect
```

#### Manually

```
curl -fsSL https://raw.githubusercontent.com/ondrejsika/kubeconfig-sikademo/master/kubeconfig | base64 --decode > kubeconfig
```

Copy it to `~/.kube/config`:

```
mkdir -p ~/.kube
cp ~/.kube/config ~/.kube/config.$(date +%Y-%m-%d_%H-%M-%S).backup
mv kubeconfig ~/.kube/config
```

#### Create owm namespace

Create own namespace (eg.: `ondrejsika`) and set it as default

```
kubectl create ns ondrejsika
```

Switch to namespace

```
kubectl config set-context --current --namespace=ondrejsika
```

Using `kubens`

```
kn ondrejsika
```

## Course

### Explain Kubernetes Resources

```
kubectl explain node
kubectl explain node.spec
kubectl explain node.status.images

kubectl explain pod
kubectl explain pod.spec
kubectl explain pod.spec.containers.image
```

### Get Nodes

```
kubectl get nodes
```

```
kubectl get node
```

```
kubectl get no
```

```
kubectl get no minikube
```

```
kubectl get no/minikube
```

```
kubectl get no minikube -o yaml
```

```
kubectl get no minikube -o json
```

```
kubectl get no minikube -o jsonpath="{.status.addresses[0].address }{'\n'}"
```

```
kubectl get no -o jsonpath="{range .items[*]}{.status.addresses[0].address} {.status.addresses[1].address}{'\n'}{end}"
```

### Proxy to cluster

Start proxy

```
kubectl proxy
```

Change port

```
kubectl proxy -p 8002
```

**DO NOT RUN IN PRODODUCTION**: Run proxy on all interfaces and allow all hosts - for training only

Export raw Kubernetes API from lab VM to all interfaces:

```
docker run --name minikube-api-raw --net host sikalabs/slu:v0.50.0 slu proxy tcp -l :8443 -r $(minikube ip):8443
```

See: <https://lab0.sikademo.com:8443>

or:

```
curl -k https://lab0.sikademo.com:8443
```

```
kubectl proxy --address 0.0.0.0 --accept-hosts=".*"
```

On lab VM, you can also run proxy in Docker:

```
docker run -d --name proxy --net host -v /root:/root sikalabs/kubectl kubectl proxy --address 0.0.0.0 --accept-hosts=".*"
```

### Dashboard

#### Minikube

```
minikube addons enable dashboard
```

```
minikube addons enable metrics-server
```

See the dashboard: <http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/>

or just:

```
minikube dashboard
```

Which start a new proxy, activate the dashboard and open it in the browser.

### Production cluster (AWS, RKE, Digital Ocean)

Source of Kubernetes Dashboard - https://github.com/kubernetes/dashboard

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.0/aio/deploy/recommended.yaml
```

See the dashboard: http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

**DON'T RUN ON PRODUCTION**: If you want to grant permissions to dashboard without need for a token, you can run this:

Add `cluster-admin` role to dashboard user:

```
kubectl apply -f kubernetes-dashboard-cluster-admin.yml
```

and:

```
kubectl patch deployment \
  kubernetes-dashboard \
  --namespace kubernetes-dashboard \
  --type='json' \
  --patch='[{"op": "replace", "path": "/spec/template/spec/containers/0/args", "value": [
  "--auto-generate-certificates",
  "--enable-insecure-login",
  "--enable-skip-login",
  "--namespace=kubernetes-dashboard"
]}]'
```

Sources:

- https://devblogs.microsoft.com/premier-developer/bypassing-authentication-for-the-local-kubernetes-cluster-dashboard/
- https://stackoverflow.com/a/62039375/5281724

### K9s - CLI Dashboard

- https://k9scli.io/
- https://github.com/derailed/k9s

### Kubectl Cheat Sheets

- https://kubernetes.io/docs/reference/kubectl/quick-reference/
- https://aiidecoe.com/images/Kubernetes%20Cheat%20Sheet.pdf

### Create Pod

```
kubectl apply -f 01_pod.yml
kubectl apply -f 02_pod.yml
kubectl apply -f pod_redis.yml
```

### Validate YAML using kubectl

```
k apply --dry-run=server -f pod_invalid_yaml.yml
```

### List Pods

```
kubectl get pods
kubectl get po
```

```
kubectl get po redis
```

```
kubectl get po/redis
```

```
kubectl get -f pod_redis.yml
```

```
kubectl get po redis simple-hello-world
```

```
kubectl get po/redis po/simple-hello-world
```

```
kubectl get po/redis no/minikube
```

```
kubectl get -f pod_redis.yml -f 01_pod.yml
```

```
kubectl get pods -o json
```

```
kubectl get po -o jsonpath="{range .items[*]}{.spec.containers[0].image}{'\n'}{end}"
```

```
kubectl get po -o custom-columns="name:{.metadata.name},namespace:{.metadata.namespace},ip:{.status.podIP}"
```

### See Pods

See:

- http://127.0.0.1:8001/api/v1/namespaces/default/pods/simple-hello-world/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/pods/multi-container-pod/proxy/

or using port forward:

```
kubectl port-forward pod/simple-hello-world 8000:80
```

See: <http://127.0.0.1:8000>

Port forward on all interfaces (remote host):

```
kubectl port-forward pod/simple-hello-world 8000:80 --address 0.0.0.0
```

Redis port forward example

```
kubectl port-forward pod/redis 6379:6379
```

Test with local `redis-cli` (if you have):

```
redis-cli ping
```

![](images/port-forward-redis.png)

Get Pod from file:

```
kubectl get -f 01_pod.yml
kubectl get -f 02_pod.yml
```

### Describe Pod

See informations & events in pretty output

```
kubectl describe -f 01_pod.yml
kubectl describe -f 02_pod.yml
```

### Exec (Connect) Pod

```
kubectl exec redis -- redis-cli ping
```

```
kubectl exec -ti redis -- redis-cli
```

Defaut container is first one in multi-container pod

```
kubectl exec -ti multi-container-pod -- bash
```

Connect specific container

```
kubectl exec -ti multi-container-pod -c date -- bash
```

### Pod Logs

```
kubectl logs simple-hello-world
```

or following logs

```
kubectl logs simple-hello-world -f
```

Logs from multi container pod

```
kubectl logs multi-container-pod nginx
kubectl logs multi-container-pod date
```

From Kubernetes 1.24 you can see logs from first container by default

```
kubectl logs multi-container-pod
```

### Assigning Pods to Nodes

Docs - <https://kubernetes.io/docs/concepts/configuration/assign-pod-node/>

### Select by node name (nodeName)

```
kubectl apply -f nodename.yml
```

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
kubectl label nodes minikube foo=bar
```

### Select by label (nodeSelector)

```
kubectl apply -f nodeselector.yml
```

### Affinity

<https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity>

#### requiredDuringSchedulingIgnoredDuringExecution

```
kubectl apply -f affinity_required.yml
```

#### preferredDuringSchedulingIgnoredDuringExecution

```
kubectl apply -f affinity_preferred.yml
```

### Delete Pod

```
kubectl delete -f 01_pod.yml -f 02_pod.yml

# or
kubectl delete po/simple-hello-world
kubectl delete po/multi-container-pod
```

### Delete All Pods

```
kubectl delete po --all
```

### Private Container Registry

Deploy private pod

```
kubectl apply -f private_pod.yml
```

See <http://127.0.0.1:8001/api/v1/namespaces/default/pods/private-pod/proxy/>

See credentials (of example `registry.sikalabs.com`)

```
kubectl get secret private-registry-credentials -o jsonpath="{.data.\.dockerconfigjson}" | base64 --decode | jq '.auths["registry.sikalabs.com"].auth' -r | base64 --decode && echo
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

See proxy to service: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

Try apply pods:

```
kubectl apply -f 01_pod.yml -f 02_pod.yml -f private_pod.yml -f pod_redis.yml
```

Get pods for this service

```
kubectl get po -l svc=example
```

Check service proxt again. <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

Delete pods

```
kubectl delete -f 01_pod.yml -f 02_pod.yml -f private_pod.yml -f pod_redis.yml
```

### Pods Without Service Links (in the Environment)

```
kubectl apply -f pod_without_service_links.yml
```

See the env of each pod

```
kubectl logs pod-with-service-links
```

```
kubectl logs pod-without-service-links
```

See the diff

```
vimdiff <(kubectl logs pod-with-service-links) <(kubectl logs pod-without-service-links)
```

![pod-without-service-links](./images/pod-without-service-links.png)

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

or using port forward:

```
kubectl port-forward rs/hello-world-rs 8000:80
```

See: <http://127.0.0.1:8000>

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

### Well-Known Labels, Annotations and Taints

<https://kubernetes.io/docs/reference/labels-annotations-taints/>

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
kubectl get deploy,rs,po
```

See pods in service proxy: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

or using port forward:

```
kubectl port-forward deploy/hello-world 8000:80
```

See: <http://127.0.0.1:8000>

### Update Deployment

See the difference

```
vimdiff 04_01_deployment.yml 04_02_deployment.yml
```

```
kubectl apply -f 04_02_deployment.yml
```

Wait until deployment will be rolled out

```
kubectl rollout status deployment hello-world
```

### Restart Deployment

```
kubectl rollout restart deploy hello-world
```

### Clean up Policy

https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#clean-up-policy

### Delete Deployment

```
kubectl delete -f 04_01_deployment.yml

# or
kubectl delete deploy/hello-world
```

## Deployment Strategies

Great resources by Container Solutions

- [article](https://container-solutions.com/kubernetes-deployment-strategies)
- [repository](https://github.com/ContainerSolutions/k8s-deployment-strategies)

### Rolling Update (without downtime)

Create deployment & service

```
kubectl apply -f strategy_rolling_update.yml
```

See update

```
vimdiff strategy_rolling_update.yml strategy_rolling_update_2.yml
```

Update without downtime

```
kubectl apply -f strategy_rolling_update_2.yml
```

Clean up

```
kubectl delete -f strategy_rolling_update.yml
```

### Recreate

Create deployment & service

```
kubectl apply -f strategy_recreate.yml
```

See update

```
vimdiff strategy_recreate.yml strategy_recreate_2.yml
```

Update with downtime

```
kubectl apply -f strategy_recreate_2.yml
```

Clean up

```
kubectl delete -f strategy_recreate.yml
```

## Pods Spread Configuration (topologySpreadConstraints)

Try topologySpreadConstraints configuration:

```
kubectl apply -f topology_spread_constraints.yml
```

See pod's distribution:

```
kubectl get po -l app=topology-spread-constraints -o wide
```

### Create StatefulSet

Up to Kubernetes version 1.24

```
kubectl apply -f statefulset24_1.yml
```

From Kubernetes version 1.25

```
kubectl apply -f statefulset25_1.yml
```

```
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
vimdiff statefulset24_1.yml statefulset24_2.yml
```

or

```
vimdiff statefulset25_1.yml statefulset25_2.yml
```

Upgrade statefull set

```
kubectl apply -f statefulset24_2.yml
```

or

```
kubectl apply -f statefulset25_2.yml
```

### Headless Service

Service which expose pods on `<pod-name>.<svc-name>` (`<pod-name>.<svc-name>.<namespace>.svc.cluster.local`). Requires `spec.clusterIP: None`.

See service in Stateful Set.

Example:

```
kubectl run dev -ti --rm --image=sikalabs/dev -- bash
```

And try (inside of Kubernetes):

```
host hello-world
host hello-world-0.hello-world
host hello-world-1.hello-world

curl hello-world
curl hello-world
curl hello-world-0.hello-world
curl hello-world-1.hello-world
```

![headless-service-example](./images/headless-service-example.png)

### Delete StatefulSet

```
kubectl delete -f statefulset.yml

# or
kubectl delete sts/hello-world
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

#### Wait for Job Completion

```
kubectl wait --for=condition=complete --timeout=600s job/pi
```

Wait for jsonpath

```
kubectl wait --for=jsonpath='{.status.phase}'=running tkc/example
```

#### Automatic Job Cleanup

You have to set `ttlSecondsAfterFinished` in job spec.

### Job With Generated Name

Create (not apply) job

```
kubectl create -f job-generate-name.yml
```

Delete all jobs

```
kubectl delete job --all
```

Example

```
ondrej@sika-mac-air:~$ kubectl create -f job-generate-name.yml
job.batch/hello-4hjd4 created
ondrej@sika-mac-air:~$ kubectl create -f job-generate-name.yml
job.batch/hello-nrr8r created
ondrej@sika-mac-air:~$ kubectl create -f job-generate-name.yml
job.batch/hello-h8mjd created
ondrej@sika-mac-air:~$ kubectl get job
NAME          COMPLETIONS   DURATION   AGE
hello-4hjd4   1/1           1s         5s
hello-h8mjd   1/1           2s         3s
hello-nrr8r   1/1           2s         4s
ondrej@sika-mac-air:~$ kubectl delete job --all
job.batch "hello-4hjd4" deleted
job.batch "hello-h8mjd" deleted
job.batch "hello-nrr8r" deleted
```

### Create Cron Job

- <https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/>
- <https://crontab.guru/>

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
kubectl run dev -it --rm --image=sikalabs/dev -- zsh
```

Images on Github

- [sikalabs/dev](https://github.com/sikalabs/sikalabs-container-images/tree/master/dev)
- [sikalabs/ci](https://github.com/sikalabs/sikalabs-container-images/tree/master/ci)

Cleanup is not necessary, because `--rm` parameter deletes deployment after container exits.

### Alias `kdev`

Run `sikalabs/dev`:

```
alias kdev='kubectl run "dev-$(date +%s)" --rm -ti --image sikalabs/dev -- zsh'
```

### Create Service ClusterIP

Create deploymnet again:

```
kubectl apply -f deployment.yml
```

And create service:

```
kubectl apply -f 05_clusterip_service.yml
```

See:

- http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-clusterip:80/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-clusterip:8080/proxy/metrics

### Create Service NodePort

```
kubectl apply -f 06_nodeport_service.yml
```

See: http://127.0.0.1:8001/api/v1/namespaces/default/services/hello-world-nodeport/proxy/

You can also open NodePort directly using:

```
minikube service hello-world-nodeport
```

or see url:

```
minikube service hello-world-nodeport --url
```

### Create LoadBalancer Service

```
kubectl apply -f loadbalancer.yml
```

Wait until get external IP address. Works only in public clouds (like Digital Ocean, AWS) NOT in minikube. You have to have a loadbalancer provider.

### MetalLB

- https://github.com/metallb/metallb
- https://metallb.universe.tf
- https://metallb.universe.tf/installation/

```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml
```

Apply pool config:

```
kubectl apply -f metallb_pool.yml
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
kubectl delete -f 05_clusterip_service.yml
kubectl delete -f 06_nodeport_service.yml
kubectl delete -f loadbalancer.yml

# or
kubectl delete svc/hello-world-nodeport
kubectl delete svc/hello-world-clusterip

# and deployment
kubectl delete deploy/hello-world
```

## Ingress

Ingress Controllers - <https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/>

### Enable Ingress on Minikube

```
minikube addons enable ingress
```

On local machine (win, mac), you have to run `minikube tunnel` in own tab to access port 80 and 443:

```
minikube tunnel
```

On lab VM you have to run those proxies to access ports 80 and 443:

```
docker run -d --name proxy-ingress-80 --net host -v /root:/root sikalabs/slu:v0.50.0 slu proxy tcp -l :80 -r $(kubectl get node minikube -o jsonpath="{.status.addresses[0].address}"):$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath="{.spec.ports[0].nodePort}")
```

```
docker run -d --name proxy-ingress-443 --net host -v /root:/root sikalabs/slu:v0.50.0 slu proxy tcp -l :443 -r $(kubectl get node minikube -o jsonpath="{.status.addresses[0].address}"):$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath="{.spec.ports[1].nodePort}")
```

### Install Ingress Nginx on DigitalOcean

```
helm upgrade --install \
  ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --create-namespace \
  --namespace ingress-nginx \
  --set controller.service.type=ClusterIP \
  --set controller.ingressClassResource.default=true \
  --set controller.kind=DaemonSet \
  --set controller.hostPort.enabled=true \
  --set controller.metrics.enabled=true \
  --wait
```

Or using `slu`:

```
slu scripts kubernetes install-ingress
```

### Install Ingress Nginx on AKS (Azure)

```bash
LOADBALANCER_IP=...
```

```
helm upgrade --install \
  ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --create-namespace \
  --namespace ingress-nginx \
  --set controller.service.type=LoadBalancer \
  --set controller.ingressClassResource.default=true \
  --set controller.kind=DaemonSet \
  --set controller.hostPort.enabled=true \
  --set controller.metrics.enabled=true \
  --set controller.config.use-proxy-protocol=false \
  --set controller.service.loadBalancerIP=$LOADBALANCER_IP \
  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz \
  --wait
```

Or with random public IP

```
helm upgrade --install \
  ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --create-namespace \
  --namespace ingress-nginx \
  --set controller.service.type=LoadBalancer \
  --set controller.ingressClassResource.default=true \
  --set controller.kind=DaemonSet \
  --set controller.hostPort.enabled=true \
  --set controller.metrics.enabled=true \
  --set controller.config.use-proxy-protocol=false \
  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz \
  --wait
```

See [ArtifactHub](https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx?modal=values) or [values.yml](https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml) for more options.

### Create Ingress

Create some services (& deploymnets)

```
kubectl apply -f webservers.yml
```

See:

- http://127.0.0.1:8001/api/v1/namespaces/default/services/nginx:http/proxy/
- http://127.0.0.1:8001/api/v1/namespaces/default/services/apache:http/proxy/

Create Ingress on Minikube

```
kubectl apply -f ingress.yml
```

See:

- http://nginx.127.0.0.1.nip.io
- http://apache.127.0.0.1.nip.io

Create Ingress on DigitalOcean

We need a [cert-managet](https://cert-manager.io/) for creating HTTPS certificates. You can install it using helm:

```
helm upgrade --install \
	cert-manager cert-manager \
	--repo https://charts.jetstack.io \
	--create-namespace \
	--namespace cert-manager \
	--set crds.enabled=true \
	--wait
```

Or using [slu](https://github.com/sikalabs/slu):

```
slu scripts kubernetes install-cert-manager
```

Install cluster issuer:

```
kubectl apply -f clusterissuer-letsencrypt.yml
```

or using `slu`:

```
slu scripts kubernetes install-cluster-issuer
```

```
kubectl apply -f ingress-do.yml
```

See:

- http://nginx.k8s.sikademo.com
- http://apache.k8s.sikademo.com

Cleanup

```
kubectl delete -f webservers.yml -f ingress.yml -f ingress-do.yml
```

### Sticky Sessions

<https://kubernetes.github.io/ingress-nginx/examples/affinity/cookie/>

Minikube:

```
kubectl apply -f ingress-sticky.yml
```

DigitalOcean:

```
kubectl apply -f ingress-do-sticky.yml
```

### Deploy Application (Multiple Deployments and Services)

```
kubectl apply -f counter.yml
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
kubectl delete -f counter.yml
```

## InitContainers

More about init containers here: <https://kubernetes.io/docs/concepts/workloads/pods/init-containers/>

See example [init_containers.yml](init_containers.yml)

See also diff with counter example

```
vimdiff counter.yml init_containers.yml
```

Example:

```
kubectl apply -f init_containers.yml

kubectl get -f init_containers.yml

minikube service counter

kubectl delete -f init_containers.yml
```

See example of [ondrejsika/wait-for-it](https://github.com/ondrejsika/docker-training/tree/master/examples/wait-for-it)

#### Setup some env example using Init Containers

```
kubectl apply -f init_containers_setup.yml
```

#### Wait for Migrations in Init Container

Wait using `kubectl`

```
kubectl apply -f init_containers_migrations.yml
```

Wait using `slu`

```
kubectl apply -f init_containers_migrations_slu.yml
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

### Get Resources from All Namespaces

Try

```
kubectl get pods

# and

kubectl get pods -A
kubectl get pods --all-namespaces
```

### Deploy to Namespace

```
kubectl apply -f counter.yml -n counter
```

See: http://127.0.0.1:8001/api/v1/namespaces/counter/services/counter/proxy/

### Switch Namespace in Current Context

```
kubectl config set-context --current --namespace=counter
```

### kubens

I prefer alias `kn` over original command `kubens`.

List namespaces

```
kn
```

or

```
kubens
```

Switch namespace

```
kn counter
```

or

```
kubens counter
```

### Delete Namespace

```
kubectl delete -f 08_namespace.yml

# or
kubectl delete ns/counter
```

### Delete Terminating Namespace

- https://www.ibm.com/docs/en/cloud-private/3.2.x?topic=console-namespace-is-stuck-in-terminating-state

```
kubectl patch ns <Namespace_to_delete> -p '{"metadata":{"finalizers":null}}'
```

or using `slu`

```
slu k8s delete-ns -n name_of_the_namespac
```

### Resource Quotas

Create Resource Quotas

```
kubectl apply -f resourcequota.yml
```

See Resource Quotas:

```
kubectl describe resourcequotas
```

Try to deploy:

```
kubectl apply -f statefulset.yml
kubectl apply -f service.yml
kubectl apply -f 01_pod.yml
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

## Storage & Volumes

## PersistentVolume & PersistentVolumeClaim

### Storage Classes

A StorageClass provides a way for administrators to describe the "classes" of storage they offer. Different classes might map to quality-of-service levels, or to backup policies, or to arbitrary policies determined by the cluster administrators.

Docs: <https://kubernetes.io/docs/concepts/storage/storage-classes/>

List Storage Classes:

```
kubectl get storageclass
# or
kubectl get sc
```

### PersistentVolumeClaim

- AccessModes - <https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes>

From default **StorageClass**

```
kubectl apply -f pvc_default.yml
```

See PV and PVC

```
kubectl get pv,pvc
```

Use PVC

```
kubectl apply -f pvc_mount_example.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/services/pvc-mount-example/proxy/>

Example of `kubectl cp`:

```
kubectl cp index.html <pod>:usr/share/nginx/html/index.html
kubectl cp <pod>:usr/share/nginx/html/index.html index.html
```

Clean up

```
kubectl delete -f pv_nfs.yml -f pvc_nfs.yml -f pvc_mount_example.yml
```

### Longhorn

<https://longhorn.io/>

Install

```
kubectl apply -f https://raw.githubusercontent.com/longhorn/longhorn/v1.8.0/deploy/longhorn.yaml
```

#### Longhorn Dashboard

See <http://127.0.0.1:8001/api/v1/namespaces/longhorn-system/services/longhorn-frontend:80/proxy/>

#### Make Longhorn Default Storage Class

Clean is default from `do-block-storage` (DigitalOcean) and `standard` (Minikube).

```
kubectl patch storageclass do-block-storage -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'
kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'
```

Set Longhorn as default

```
kubectl patch storageclass longhorn -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

### NFS Client Provisioner

Setup Helm:

```
helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
helm repo update
```

Install using Helm:

```
helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner --namespace nfs-subdir-external-provisioner --create-namespace --set nfs.server=<nfs-server> --set nfs.path=<exported-path>
```

Example with my NFS server (nfs.sikademo.com):

```
helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner --namespace nfs-subdir-external-provisioner --create-namespace --set nfs.server=nfs.sikademo.com --set nfs.path=/nfs
```

You can use `nfs-client` Storage Class:

```
kubectl apply -f pvc_nfs_client.yml
```

Mount example:

```
kubectl apply -f nfs-client-deployment.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/services/example/proxy/>

### Direct NFS Volume mount

```
kubectl apply -f nfs-volume-example.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/pods/nfs-volume-example:80/proxy/>

### Reclaim Policy Retain

#### Install Storage Classes with Reclaim Policy Retain

Minikube

```
kubectl apply -f sc-minikube-retain.yml
```

AWS GP2

```
kubectl apply -f sc-gp2-retain.yml
```

DigitalOcean

```
kubectl apply -f sc-do-retain.yml
```

Longhorn

```
kubectl apply -f sc-longhorn-retain.yml
```

#### Apply PVC with Storage Class

Minikube

```
kubectl apply -f pvc-minikube-retain.yml
```

DigitalOcean

```
kubectl apply -f pvc-do-retain.yml
```

Longhorn

```
kubectl apply -f pvc-longhorn-retain.yml
```

### CephFS Example

```
kubectl apply -f cephfs-volume-example.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/pods/cephfs-volume-example/proxy/>

## ConfigMaps & Secrets

### Create secret

```
kubectl apply -f 11_secret.yml
```

### Create TLS secret from cert and key

Create secret file `tls_secret_127-0-0-1.uk.local.yaml`

```
kubectl create secret tls sikademo.com-tls \
  --key=/sikademo.com.key \
  --cert=sikademo.com.crt \
  --dry-run=client -o yaml | tee secret_tls_sikademo_com.local.yml
```

Create secret in Kubernetes cluster

```
kubectl create secret tls sikademo.com-tls \
  --key=/sikademo.com.key \
  --cert=sikademo.com.crt
```

### Get Values

Base64 encoded

```
kubectl get secret my-secret -o go-template='{{.data.password}}{{"\n"}}'
kubectl get secret my-secret -o go-template='{{.data.token}}{{"\n"}}'
```

```
kubectl get secret my-secret -o jsonpath="{.data.password}" && echo
kubectl get secret my-secret -o jsonpath="{.data.token}" && echo
```

Decoded

```
kubectl get secret my-secret -o go-template='{{.data.password|base64decode}}{{"\n"}}'
kubectl get secret my-secret -o go-template='{{.data.token|base64decode}}{{"\n"}}'
```

```
kubectl get secret my-secret -o jsonpath="{.data.password}" | base64 --decode && echo
kubectl get secret my-secret -o jsonpath="{.data.token}" | base64 --decode && echo
```

or using **slu**

```
slu k8s get sec -s my-secret
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
kubectl exec secret-example -- env | grep MY_
```

Or on Windows:

```
kubectl exec secret-example env | findstr MY_
```

### EnvFrom

Mount every variables from ConfigMap or Secret. See example:

```
kubectl apply -f configmap_envfrom.yml
kubectl logs envfrom
```

### Env from fieldRef

```
kubectl apply -f env_fieldref.yml
```

```
kubectl logs env-fieldref
```

### Expand Environment Variables in Config File

```
kubectl apply -f expand_config.yml
```

See: <http://127.0.0.1:8001/api/v1/namespaces/default/pods/expand-config/proxy/>

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

### Join multiple kubeconfigs

Backup your kubeconfig first

```
cp ~/.kube/config ~/.kube/config.$(date +%Y-%m-%d_%H-%M-%S).backup
```

Add `kubeconfig-new.yml` to your kubeconfig

```
KUBECONFIG=kubeconfig-new.yml:~/.kube/config kubectl config view --raw > /tmp/kubeconfig.merge.yml && cp /tmp/kubeconfig.merge.yml ~/.kube/config
```

Or using `slu`:

```
slu k8s config add -p kubeconfig-new.yml
```

### Join multiple kubeconfigs on Windows

Backup your kubeconfig first

```powershell
Copy-Item -Path $HOME\.kube\config -Destination "$HOME\.kube\config.$((Get-Date -Format 'yyyy-MM-dd_HH-mm-ss')).backup"
```

Add `kubeconfig-new.yml` to your kubeconfig

```powershell
$env:KUBECONFIG = "kubeconfig-new.yml;$HOME\.kube\config"
```

```powershell
kubectl config view --raw > $HOME\.kube\config_tmp
```

```powershell
Remove-Item $HOME\.kube\config
Move-Item $HOME\.kube\config_tmp $HOME\.kube\config
```

Source: https://dbafromthecold.com/2020/02/21/merge-kubectl-config-files-on-windows/

### RBAC

Show all api resources (with verbs)

```
kubectl api-resources -o wide
```

Show all ServiceAccounts:

```
kubectl get sa -A
```

### Command Kubernetes as different service account

```
kubectl get no
```

```
kubectl --as system:serviceaccount:kube-system:default get no
```

```
kubectl --as system:serviceaccount:default:default get no
```

### Can I

```
kubectl auth can-i get po
kubectl auth can-i --as system:anonymous get po
```

### Create cluster admin

```
kubectl apply -f 14_admin.yml
```

### Create kubeconfig for admin

Export actual config

```
kubectl config view --raw --minify > kubeconfig.yml
```

Use custom `kubeconfig.yml` file using environment variable:

Linux & Mac

```bash
export KUBECONFIG=kubeconfig.yml
```

PowerShell

```powershell
$env:KUBECONFIG = "kubeconfig.yml"
```

CMD

```cmd
set KUBECONFIG=kubeconfig.yml
```

Get token:

```
kubectl -n kube-system describe secret $(kubectl -n kube-system get serviceaccounts admin-user -o jsonpath="{.secrets[0].name}")
```

Get token using **slu**:

```
slu k8s token -n kube-system -s admin-user
```

Set token to user:

```
kubectl config set-credentials admin --token=<token>
```

Or get token and create user in config on oneliner:

```
kubectl config set-credentials admin --token=$(kubectl -n kube-system get secret $(kubectl -n kube-system get serviceaccounts admin-user -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" | base64 --decode)
```

Or oneliner with **slu**:

```
kubectl config set-credentials admin --token=$(slu k8s token -n kube-system -s admin-user)
```

Set new user to context:

```
kubectl  config set-context admin --user=admin --cluster=minikube
```

Use new user to context:

```
kubectl config use-context admin
```

or using `kubectx`:

```
kx admin
```

And try:

```
kubectl get nodes,svc
```

### Create pod reader

```
kubectl apply -f 15_read.yml
```

Can I?

```
kubectl auth can-i --as system:serviceaccount:kube-system:read-user get po
```

```
kubectl auth can-i --as system:serviceaccount:kube-system:read-user create po
```

Add to user to config and change context user

```
kubectl config set-credentials read --token=$(kubectl -n kube-system get secret $(kubectl -n kube-system get serviceaccounts read-user -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" | base64 --decode)

kubectl config set-context read --user=read --cluster=minikube

kubectl config use-context read
```

or using `kubectx`:

```
kx read
```

### Kubectx rename context

```
kx pod-read=read
```

### Create Namespace Admin

```
kubectl apply -f 16_namespace_admin.yml
```

And create user, also with default namespace changed to `devel`

```
kubectl config set-credentials devel --token=$(kubectl -n devel get secret $(kubectl -n devel get serviceaccounts devel-user -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" | base64 --decode)

kubectl config set-context devel --user=devel --namespace=devel  --cluster=minikube

kubectl config use-context devel
```

or using `kubectx`:

```
kx devel
```

Add read only access to some cluster wide resources (nodes, volumes, ...)

```
kubectl apply -f 17_namespace_admin_extra.yml
```

### RBAC example of nonResourceURLs (`/metrics`)

```
kubectl apply -f rbac_metrics.yml
```

Create context and use it

```
kubectl config set-credentials metrics --token=$(kubectl -n kube-system get secret $(kubectl -n kube-system get serviceaccounts metrics-user -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" | base64 --decode)

kubectl config set-context --user=metrics --cluster=minikube metrics

kubectl config use-context metrics
```

See:

- <http://127.0.0.1:8001/>
- <http://127.0.0.1:8001/metrics>

## Access Kubernetes API using CURL

Get Kubernetes API URL

```
export APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
```

Get `admin-user` token using **slu**

```
export TOKEN=$(slu k8s token -n kube-system -s admin-user)
```

or get `admin-user` token using kubectl only

```
export TOKEN=$(kubectl -n kube-system get secret $(kubectl -n kube-system get serviceaccounts admin-user -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" | base64 --decode)
```

Try some CURLs:

```
curl -k --header "Authorization: Bearer $TOKEN" $APISERVER
```

```
curl -k --header "Authorization: Bearer $TOKEN" $APISERVER/metrics
```

```
curl -k --header "Authorization: Bearer $TOKEN" $APISERVER/api/v1/nodes
```

## Liveness, Readiness & Startup Probes

Docs: <https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/>

### `/livez` vs `/healthz`

Use `/livez` instead of `/healthz` (deprecated)

[source](https://www.jaktech.co.uk/java/difference-between-livez-readyz-and-healthz-in-kubernetes-cluster/)

### Liveness Probe

```
kubectl apply -f probes_liveness.yml
```

Watch pods:

```
watch -n 0.3 kubectl get deploy,rs,po -l app=liveness
```

Watch output:

```
watch -n 0.3 curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/liveness/proxy/
```

or using slu watch

```
slu w -- curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/liveness/proxy/
```

Cleanup

```
kubectl delete -f probes_liveness.yml
```

### Readiness Probe

```
kubectl apply -f probes_readiness.yml
```

Watch pods:

```
watch -n 0.3 kubectl get deploy,rs,po -l app=readiness
```

Watch service:

```
watch -n 0.3 kubectl describe svc readiness
```

Watch output:

```
watch -n 0.3 curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/readiness/proxy/
```

or using slu watch

```
slu w -- curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/readiness/proxy/
```

Cleanup

```
kubectl delete -f probes_readiness.yml
```

### Startup Probe

```
kubectl apply -f probes_startup.yml
```

Watch pods:

```
watch -n 0.3 kubectl get deploy,rs,po -l app=startup
```

Watch service:

```
watch -n 0.3 kubectl describe svc startup
```

Watch output:

```
watch -n 0.3 curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/startup/proxy/
```

or using slu watch

```
slu w -- curl -fsSL http://127.0.0.1:8001/api/v1/namespaces/default/services/startup/proxy/
```

Cleanup

```
kubectl delete -f probes_startup.yml
```

## Resource Consumption (`kubectl top`)

We have to have [metric-server](https://github.com/kubernetes-sigs/metrics-server) installed.

Minikube:

```
minikube addons enable metrics-server
```

Direct:

```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

Or using `slu`:

```
slu scripts kubernetes install-metrics-server
```

or short version using `slu`:

```
slu s k ims
```

See nodes comsumptions:

```
kubectl top nodes
```

See pods comsumptions:

```
kubectl top pods
```

```
kubectl top pods -A
```

Sort by CPU

```
k top po -A  --sort-by=cpu
```

Sort by memory

```
k top po -A  --sort-by=memory
```

## Autoscaling (Horizontal Pod Autoscaler)

- <https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/>
- <https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/>

We also have to have metrics server enabled

If you requests 200 milli-cores to pod, 50% means that Kubernetes autoscaler keeps it on 100 mili-cores. More info here: [Create Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#create-horizontal-pod-autoscaler) and [Algorithm Details](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details).

### Create HPA

Api v1 (same as `kubectl autoscale`)

```
kubectl apply -f hpa-fast.yml
```

Run AB locally:

```
ab -c 4 -n 100000 $(minikube service hpa-service --url)/
```

or from Kubernetes (kubectl run):

```
kubectl run ab --image=ondrejsika/ab --rm -ti -- ab -c 4 -n 100000 http://hpa-service/
```

or using Kubernetes job:

```
kubectl apply -f ab.yml
```

See autoscaling at work

```
watch -n 0.3 -d kubectl get -f hpa-fast.yml
```

Real HPA example

See:

```
vimdiff hpa-fast.yml hpa-real.yml
```

Try:

```
kubectl apply -f hpa-real.yml
```

### Get HPAs

```
kubectl get hpa
```

### Test Autoscaling

Run AB locally

```
ab -c 4 -n 100000 $(minikube service hpa-service --url)/
```

or in Kubernetes (kubectl run):

```
kubectl run ab --image=ondrejsika/ab --rm -ti -- ab -c 4 -n 100000 http://hpa-service/
```

or using Kubernetes job:

```
kubectl apply -f ab.yml
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
kubectl delete hpa/hello-world
kubectl delete -f 04_01_deployment.yml
```

## Helm

### Add Repository

```
helm repo add sikalabs https://helm.sikalabs.io
```

### Search Package

Search local repositories

```
helm search repo sikalabs
```

### Inspect Package

```
helm show chart sikalabs/hello-world
helm show values sikalabs/hello-world
helm show readme sikalabs/hello-world

helm show all sikalabs/hello-world
```

### Install Package

```
helm install <release_name> <chart>
```

Example [hello-world](https://github.com/sikalabs/charts/tree/master/charts/hello-world) (on [ArtifactHub](https://artifacthub.io/packages/helm/sikalabs/hello-world)) package / chart:

```
helm install hello sikalabs/hello-world
```

Or dry run (see the Kubernetes config)

```
helm install hello sikalabs/hello-world --dry-run
```

## Install Package Without Helm Repo Add

```
helm install hello --repo https://helm.sikalabs.io hello-world
```

### Upgrade Package

If you want to upgrade instance of chart, you have to call:

```
helm upgrade hello sikalabs/hello-world --set replicas=3
```

or

```
helm upgrade hello sikalabs/hello-world --set host=hello.127.0.0.1.nip.io
```

### Install or Upgrade

You can add `--install` to `helm upgrade` to install package if not exists. When chart exists, it will be upgraded.

```
helm upgrade --install hello sikalabs/hello-world --set host=hello.127.0.0.1.nip.io
```

### Helm History

```
helm history <release_name>
```

Example

```
helm history hello
```

### Helm Rollback

```
helm rollback <release_name> <revision>
```

Example

```
helm rollback hello 1
```

```
helm rollback hello 2
```

### List Installed Packages

```
helm ls
helm ls -q
```

### Status of Package

```
helm status hello
```

### Delete Package

```
helm uninstall hello
```

Delete all & purge

```
helm uninstall $(helm ls -a -q)
```

### Bitnami Packages

Add repo:

```
helm repo add bitnami https://charts.bitnami.com/bitnami
```

Install RabbitMQ:

```
helm install rabbitmq bitnami/rabbitmq --values rabbitmq.values.yml
```

### Helm Repositiories

#### List repositories

```
helm repo list
```

#### Install sikalabs/one-image

Sources of Chart `sikalabs/one-image` are <https://github.com/sikalabs/charts/tree/master/charts/one-image>

Inspect Package

```
helm show all sikalabs/one-image
```

Install with values in args

```
helm install hello-world sikalabs/one-image --set host=hello-world.127.0.0.1.nip.io
```

Install with values file on Minikube

```
helm install nginx sikalabs/one-image --values one-image-nginx-values.yml
helm install apache sikalabs/one-image --values one-image-apache-values.yml
helm install caddy sikalabs/one-image --values one-image-caddy-values.yml
```

Install with values file and values args

```
helm install nginx2 sikalabs/one-image --values one-image-nginx-values.yml --set host=nginx2.127.0.0.1.nip.io
```

Install with values file on DigitalOcean

```
helm install nginx sikalabs/one-image \
  -f one-image-nginx-values.yml \
  -f one-image-nginx-values-do.yml
helm install apache sikalabs/one-image \
  -f one-image-apache-values.yml \
  -f one-image-apache-values-do.yml
helm install caddy sikalabs/one-image \
  -f one-image-caddy-values.yml \
  -f one-image-caddy-values-do.yml
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

This default manifests are too much complicated, if you want simple examle, check out my [sikalabs/one-image](https://github.com/sikalabs/charts/tree/master/charts/one-image).

We can try create helm package for our Wordpress example (09_wordpress.yml).

We have to replace few names & labes with `{{ .Release.Name }}` to allow multiple deployments (installations) of chart. For labels, I use `release: {{ .Release.Name }}`, it works, it's simple and make sense.

I also replace all variable part with values like `image: {{ .Values.image }}`, which I can overwrite.

See example (`one-image/deployment.yaml`)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: "{{ .Release.Name }}"
  labels:
    release: "{{ .Release.Name }}"
spec:
  replicas: 1
  selector:
    matchLabels:
      release: "{{ .Release.Name }}"
  template:
    metadata:
      labels:
        release: "{{ .Release.Name }}"
    spec:
      containers:
        - name: "{{ .Chart.Name }}"
          image: "{{ .Values.image }}"
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
helm template counter .
helm template counter . --set host=counter.127.0.0.1.nip.io
```

#### Install

```
helm install counter . --set host=counter.127.0.0.1.nip.io
```

#### Build Package

```
cd ..
helm package counter
```

### Create own repository

```
mkdir repo
mv counter-*.tgz repo/
helm repo index repo/
```

Publish it!

```
scp repo/* root@helm.sikademo.com:/helm/
```

#### Use it

Delete previous deployment

```
helm uninstall counter
```

Add repo

```
helm repo add sikademo https://helm.sikademo.com
```

Install package

```
helm install counter sikademo/counter --set host=counter.127.0.0.1.nip.io
```

## Helm Operator

### Setup k3s cluster

Install k3s without traefik on VM

```
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --disable traefik" sh
```

or with TLS SAN (for example for `k3s.sikademo.com`)

```
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --disable traefik --tls-san k3s.sikademo.com" sh
```

Install kubectl and helm

```
slu install-bin kubectl
slu install-bin helm
```

Copy kubeconfig to `~/.kube/config`

```
mkdir -p ~/.kube
cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
```

Install Cluster Essentials

```
slu script kubernetes install-all --no-argocd --base-domain xxx
```

### Install Helm using CRD

```
kubectl apply -f helm_hello_world.yml
```

```
kubectl apply -f helm_hello_world.2.yml
```

```
kubectl apply -f helm_hello_world.3.yml
```

## Kubernetes Networking

[Docs](https://kubernetes.io/docs/concepts/cluster-administration/networking/)

Plugins:

- Flannel - Simplest network plugin, use encapsulation
- Calico - Support network policy & isolation, use BGP
- Cilium - L7/HTTP Aware network plugin, support network & HTTP policies
- Weave - Mesh routing networking

Resources:

- https://kubernetes.io/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-networking-model
- https://rancher.com/blog/2019/2019-03-21-comparing-kubernetes-cni-providers-flannel-calico-canal-and-weave/
- https://info.rancher.com/kubernetes-networking-deep-dive

## Network Policy

Deploy webservers:

```
kubectl apply -f webservers.yml
```

Check connection:

```
kubectl run busybox --rm -ti --image=busybox -- wget --spider --timeout=1 nginx
```

Apply policy

```
kubectl apply -f networkpolicy.yml
```

See policy:

```
kubectl get networkpolicies
```

```
kubectl describe -f networkpolicy.yml
```

Check connection again:

```
kubectl run busybox --rm -ti --image=busybox -- wget --spider --timeout=1 nginx
```

Check connection with labels `--labels="access=true"`:

```
kubectl run busybox --rm -ti --image=busybox --labels="access=true" -- wget --spider --timeout=1 nginx
```

## Pod Security Policy

[Docs](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)

A Pod Security Policy is a cluster-level resource that controls security sensitive aspects of the pod specification.

## Minikube with PSP

[Docs](https://minikube.sigs.k8s.io/docs/tutorials/using_psp/)

Setup:

```
mkdir -p ~/.minikube/files/etc/kubernetes/addons
cp minikube-psp.yml ~/.minikube/files/etc/kubernetes/addons/psp.yaml
```

Start Minikube with PSP:

```
minikube start --extra-config=apiserver.enable-admission-plugins=PodSecurityPolicy
```

### Example

Crate NS:

```
kubectl apply -f psp/ns.yml
```

Create PSP:

```
kubectl apply -f psp/psp.yml
```

Create RBAC config:

```
kubectl apply  -f psp/rbac.yml
```

```
kubectl apply --as=system:serviceaccount:psp-example:hacker -f psp/pod.yml
```

Create Root PSP:

```
kubectl apply -f psp/psp-root.yml
```

Create Root RBAC config:

```
kubectl apply  -f psp/rbac-root.yml
```

```
kubectl apply --as=system:serviceaccount:psp-example:root -f psp/pod.yml
```

## Metrics

- Metric Server - Simple in memory metrics server for autoscaling
- Prometheus - Complete metrics pipeline with data persistence

## Prometheus

### Kube Prometheus

[Homepage (Github)](https://github.com/coreos/kube-prometheus)

Prepare Minikube

```
minikube delete
minikube start --kubernetes-version=v1.18.1 --memory=6g --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0
```

Disable Metrics Server

```
minikube addons disable metrics-server
```

Clone `kube-prometheus`:

```
git clone https://github.com/coreos/kube-prometheus
```

Setup:

```
kubectl create -f kube-prometheus/manifests/setup
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
kubectl create -f kube-prometheus/manifests/
```

Access Prometheus:

```
kubectl --namespace monitoring port-forward svc/prometheus-k8s 9090
```

See: <http://localhost:9090>

Access Grafana:

```
kubectl --namespace monitoring port-forward svc/grafana 3000
```

Usename: `admin`, password: `admin`

See: <http://localhost:3000>

Access AlertManager:

```
kubectl --namespace monitoring port-forward svc/alertmanager-main 9093
```

See: <http://localhost:9093>

### Custom Metrics

Deploy example app:

```
kubectl apply -f prom/example.yml
```

Start port forward:

```
kubectl port-forward svc/metrics-example 8000:80
```

See:

- <http://127.0.0.1:8000>
- <http://127.0.0.1:8000/metrics>

Apply custom metrics collector:

```
kubectl apply -f prom/sm.yml
```

Check in Prometheus.

Apply custom alerts:

```
kubectl apply -f prom/pr.yml
```

Check in Alert Manager.

### Cleanup Kube Prometheus

Delete custom examples & resources:

```
kubectl delete -f prom/
```

Remove Kube Prometheus:

```
kubectl delete --ignore-not-found=true -f kube-prometheus/manifests/ -f kube-prometheus/manifests/setup
```

## Logging

### Elastic Cloud on Kubernetes

- Intro - <https://www.elastic.co/elastic-cloud-kubernetes>
- Docs / Tutorial - <https://www.elastic.co/guide/en/cloud-on-k8s/current/index.html>
- Github - <https://github.com/elastic/cloud-on-k8s>

### Setup ECK Operator

```
kubectl apply -f https://download.elastic.co/downloads/eck/1.5.0/all-in-one.yaml
```

See logs:

```
kubectl -n elastic-system logs -f statefulset.apps/elastic-operator
```

### Create Elastic Cluster

Setup `logs` namespace:

```
kubectl apply -f eck/ns.yml
```

Create Elasticsearch cluster:

```
kubectl apply -f eck/elasticsearch.yml
```

Wait until `health=green`:

```
kubectl get -f eck/elasticsearch.yml
```

### Run Kibana

```
kubectl apply -f eck/kibana.yml
```

Wait until `health=green`:

```
kubectl get -f eck/kibana.yml
```

### Run Filebeat

```
kubectl apply -f eck/filebeat.yml
```

Wait until start:

```
kubectl get -f eck/filebeat.yml
```

See filebeet logs:

```
kubectl -n logs logs daemonset.apps/filebeat -f
```

### See Logs in Kibana

Get password for user `elastic`:

```
kubectl get -n logs secret logs-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode; echo
```

Run proxy to Kibana:

```
kubectl -n logs port-forward service/logs-kb-http 5601
```

See logs in Kibana:

<https://127.0.0.1:5601/app/logs/stream>

## Links

- Awesome Operators - <https://github.com/operator-framework/awesome-operators>
- Logging - https://kubernetes.io/docs/concepts/cluster-administration/logging/
- Autoscaling based on CPU/Memory - https://blog.powerupcloud.com/autoscaling-based-on-cpu-memory-in-kubernetes-part-ii-fe2e495bddd4
- Kubernetes Cheatsheet - https://kubernetes.io/docs/reference/kubectl/cheatsheet/
- Minikube & Hyper-V - https://medium.com/@JockDaRock/minikube-on-windows-10-with-hyper-v-6ef0f4dc158c
- Deployment Strategies - https://container-solutions.com/kubernetes-deployment-strategies/
- Sticky Sessions - https://medium.com/@zhimin.wen/sticky-sessions-in-kubernetes-56eb0e8f257d
- How To Set Up an Elasticsearch, Fluentd and Kibana (EFK) Logging Stack on Kubernetes - https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes

## Thank you! & Questions?

That's it. Do you have any questions? **Let's go for a beer!**

### Ondrej Sika

- email: <ondrej@sika.io>
- web: <https://sika.io>
- twitter: [@ondrejsika](https://twitter.com/ondrejsika)
- linkedin: [/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)
- Newsletter, Slack, Facebook & Linkedin Groups: <https://join.sika.io>

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn (add me [/in/ondrejsika](https://www.linkedin.com/in/ondrejsika/) and I'll send you request for recommendation). **Thanks**._

Wanna to go for a beer or do some work together? Just [book me](https://book-me.sika.io) :)
