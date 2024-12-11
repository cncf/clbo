## Episode 39: KubeStellar

- Host: [Carlos Santana](https://github.com/csantanapr)
- Date: 2024/12/12
- Videos: [Playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2PnyOsrsCbR_kqjCKfPObHK), [Episode 39 [Kubestellar]](https://www.youtube.com/watch?v=fZJewWN29EE)

### News

-  [January 2025 Sandbox Review](https://github.com/cncf/sandbox/milestone/2)

### Summary

- KubeStellar
  - CNCF [Sandbox] Project (since 2023)
    - Sandbox [project proposal](https://github.com/cncf/sandbox/issues/32)
  - [LFX Insights](https://insights.lfx.linuxfoundation.org/foundation/cncf/overview/github?project=kubestellar)
  - What?
    - KubeStellar is a flexible solution for challenges associated with multi-cluster configuration management for edge, multi-cloud, and hybrid cloud.
  - Why?
    - Centrally apply Kubernetes resources for selective deployment across multiple clusters
    - Some have compared us with OCM, but KubeStellar takes a different approach to workload lifecycle management and adds some interesting capabilities.
    - First, we do not use a manifest to bundle resources that you intend to have delivered to your multicluster/multicloud/edge environment.
    - OCM uses manifestwork, we allow unbundled usage of any Kubernetes resource - namespaced or clusterscoped.
    - Have the notion of upsyncing resources.This comes in handy when you have a spoke/edge that needs to ‘communicate’ back to the hub in some way other than just a status update in the original CR.
    - Have the capability to use clusters as transport rather than target. We denature resources so that they are not unpacked/applied to clusters where they are not intended to execute. This allows us to go from a 2-tier hub-and-spoke model/environmnet, to a n-tier hub-intermediary-spoke environment.
    - Add the ability to customize workloads in groups and summarize statuses to reduce the cognitive load of your devops. This is really useful when you want to reach far edge clusters that are not directly connected to your hub.
  - Repo: https://github.com/kubestellar/kubestellar
  - Website: https://kubestellar.io/
  - Talks
    - https://www.youtube.com/@kubestellar
  - Slack: Kubernetes Workspace `#kubestellar-dev`
  - Founders: [IBM](https://REPLACEME)

### TODO / Questions / Musings

- Getting Started with KubeStellar
  - Go to KubeStellar website https://docs.kubestellar.io
  - Install by following Getting Started
  - Use the quick [automated setup](https://docs.kubestellar.io/release-0.25.1/direct/get-started/#quick-start-using-the-automated-script)
  - Open questions
  - Interesting things
