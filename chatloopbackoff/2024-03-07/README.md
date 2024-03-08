## Episode 4: Lima

YouTube: https://www.youtube.com/watch?v=Ws3bPRiRVdw

### News
**KubeCon Is Upon Us!!!**

Because of this, next week will be hosted by known-guest Jeremy Rickard. He’ll be going over Telepresence, so look out on socials for that.

[Aligning Certifications to Kubernetes Support Windows](https://www.cncf.io/blog/2024/03/01/aligning-certifications-to-kubernetes-support-windows/)

[LinkedIn Open Sources OpenHouse Data Lakehouse Control Plane](https://thenewstack.io/linkedin-open-sources-openhouse-data-lakehouse-control-plane/)

[GUAC now an OpenSSF Incubating Project](https://openssf.org/blog/2024/03/07/guac-joins-openssf-as-incubating-project/)

### Summary

- What even is Lima?
  - Besides being dope!
  - Lima is a way to spin up QEMU images based on open source templates that the project and community provide. These templates can have varying architectures (such as riscv64) and tooling (such as K8s out of the box)
- What all did we do?

```
limactl create default
limactl start default
lima nerdctl run --rm hello-world
limactl stop default
--- 
limactl start template://k8s
export KUBECONFIG="/Users/jeefy/.lima/k8s/copied-from-guest/kubeconfig.yaml"
kns
kubectl get nodes
limactl stop k8s
---
limactl start template://experimental/riscv64
limactl shell riscv64
limactl stop riscv64z
---
limactl delete $(limactl list -q)
---
source <(limactl completion zsh)
```

### TODO / Questions / Musings
- Using lima for cross-arch CI?
- Config file format seems awesome and very robust