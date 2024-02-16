## Episode 2: K8sGPT

YouTube: https://www.youtube.com/watch?v=-nJN17yiDZE

### News

[Policy-as-Code in the software supply chain](https://www.cncf.io/blog/2024/02/14/policy-as-code-in-the-software-supply-chain/)

[Cloud native disaster recovery for stateful workloads](https://www.cncf.io/blog/2024/02/15/cloud-native-disaster-recovery-for-stateful-workloads/)

Kubernetes Patch Day Yesterday! [Quick Link to 1.29](https://relnotes.k8s.io/?releaseVersions=1.29.2)

[Register for KubeCon/CloudNativeCon EU 2024 soon!](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/register/) There’s a chance the waitlist starts soon, so don’t wait.

### Summary

- Preamble -- showed installing k8sgpt via Brew
- Spun up kind cluster
- Configured k8gspt with OpenAI Token
- Created [badpod.yaml](badpod.yaml) to have an angry thing within the kind cluster
- Ran `k8sgpt analyze`
- Flailed for a bit not realizing you have to add `--explain` to send the output to LLM for analysis
- Dug into and installed the k8sgpt-operator
- Created [k8sgpt-resource.yaml](k8sgpt-resource.yaml)
- Dug into capabilities of the operator, discussed some gaps the project should move towards

### Relevant Links 

kind - https://kind.sigs.k8s.io/

kctx/kns - https://github.com/ahmetb/kubectx/

Project Bluefin - https://projectbluefin.io/

k8sgpt (Project Website) - https://k8sgpt.ai/

k8sgpt (GitHub Repo) - https://github.com/k8sgpt-ai/k8sgpt

### Action Items

- Look up / ping Alex about periodic runs from K8sGPT
- Look into adding a Discord notification sink 