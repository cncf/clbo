## Episode 1: Introduction to Gateway API

YouTube: https://www.youtube.com/watch?v=yn95Cv7Xr7Q

### News

[Strimzi joins the CNCF Incubator](https://www.cncf.io/blog/2024/02/08/strimzi-joins-the-cncf-incubator/)

[KubeTrain: Sustainable travel to KubeCon/CloudNativeCon EU 2024](https://kubetrain.io/)

[Register for KubeCon/CloudNativeCon EU 2024 soon!](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/register/) There’s a chance the waitlist starts soon, so don’t wait.

### Summary

- Explain current state of Ingress
- Create a kind cluster
- Explain how Ingress with Kind currently works
  - Really just doing a `docker inspect` to pull the running Kind image's IP
- Dive into ingress vs Gateway API
  - Ingress is meant to be a very basic way to route HTTP(S) traffic to a specific service.
  - Ingress is NOT capable of handling anything advanced like weighted-routes or TCP/UDP traffic
- Get Gateway API running
  - Set up 2048 demo app
  - Install Gateway API CRDs
  - Install Nginx Gateway Fabric
  - Create HTTPRoute and Gateway
  - Set up a second service (go-echo) to show load-sharding

### Relevant Links 

kind - https://kind.sigs.k8s.io/

kctx/kns - https://github.com/ahmetb/kubectx/

Example application used (2048 game) - https://github.com/kubernetes-sigs/aws-load-balancer-controller/blob/main/docs/examples/2048/2048_full.yaml 

Example application used (go-echo) - https://console.cloud.google.com/gcr/images/google-samples/global/go-echo

Kubernetes docs on Gateway API - https://kubernetes.io/docs/concepts/services-networking/gateway/

Gateway API-specific Docs - https://gateway-api.sigs.k8s.io/ 

Nginx Gateway Implementation - https://docs.nginx.com/nginx-gateway-fabric/installation/running-on-kind/