## Episode 12: Armada

- Host: Carlos Santana
- Date: 2024/05/09
- Videos: [Playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2PnyOsrsCbR_kqjCKfPObHK), [Episode 12 Armada](https://www.youtube.com/watch?v=N-lqmr0zfNk)

### News
-  [Kubernetes is turning 10! Join the party on June 6th](https://www.cncf.io/blog/2024/05/07/kubernetes-is-turning-10-join-the-party-on-june-6th/)
-  [KubeCon NA (Nov 2024), CFP open until June 6th](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/cfp/)

### Summary

- Armada
  - CNCF Sandbox Project (since 2022)
  - What?
    - Armada is a system built on top of Kubernetes for running batch workloads.
  - Why?
    - Armada addresses the following limitations of Kubernetes:
      1. Armada is designed to effectively schedule jobs across many Kubernetes clusters. Many thousands of nodes can be managed by Armada in this way.
      2. Armada performs queueing and scheduling out-of-cluster using a specialized storage layer. This allows Armada to maintain queues composed of millions of jobs.
      3. Armada includes a novel multi-Kubernetes cluster scheduler
  - Repo: https://github.com/armadaproject/armada
  - Website: https://armadaproject.io/
  - Talks
      - [Armada project demo November 2023](https://www.youtube.com/watch?v=GVfFiG3zeUA)
      - [Building Armada – Running Batch Jobs at Massive Scale on Kubernetes - Jamie Poole, G-Research](https://youtu.be/B3WPxw3OUl4?si=7I3rM00Wh8TEjuwJ)
  - Slack: CNCF `#armada`
  - Founders: [G Research](https://www.gresearch.com/)



### TODO / Questions / Musings

- Getting Started with Armada
  - Install the operator from https://github.com/armadaproject/armada-operator
  - Was not able to install, we found installation [bugs](https://github.com/armadaproject/armada/issues/3575). Happy to schedule part 2 when the project fixes these problems.

