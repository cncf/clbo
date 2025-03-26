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
  - Resources:
    - Paolo Blog https://dettori.medium.com/
    - Andy Blog https://clubanderson.medium.com/
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

I ran the quick start but it failed I was missing kflex and ocm cli
I installed on macOC like this:
```shell
brew tap kubestellar/kubeflex https://github.com/kubestellar/kubeflex
brew install kubeflex

export INSTALL_DIR="$HOME/bin"
curl -L https://raw.githubusercontent.com/open-cluster-management-io/clusteradm/main/install.sh | bash
```
Ran the quickstart from https://docs.kubestellar.io/release-0.25.1/direct/get-started/#quick-start-using-the-automated-script
```shell
bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/refs/tags/v0.25.1/scripts/create-kubestellar-demo-env.sh)
```

The output is this one
```
For your convenience you will probably want to add contexts to your
kubeconfig named after the non-host-type control planes (WDSes and
ITSes) that you just created (a host-type control plane is just an
alias for the KubeFlex hosting cluster). You can do that with the
following `kflex` commands; each creates a context and makes it the
current one. See
https://github.com/kubestellar/kubestellar/blob/\{\{ .Values.KUBESTELLAR_VERSION \}\}/docs/content/direct/core-chart.md#kubeconfig-files-and-contexts-for-control-planes
for a way to do this without using `kflex`.
Start by setting your current kubeconfig context to the one you used
when installing this chart.

kubectl config use-context $the_one_where_you_installed_this_chart
kflex ctx --set-current-for-hosting # make sure the KubeFlex CLI's hidden state is right for what the Helm chart just did

kflex ctx --overwrite-existing-context its1
kflex ctx --overwrite-existing-context wds1

Finally, you can use `kflex ctx` to switch back to the kubeconfig
context for your KubeFlex hosting cluster.
✔ Checking for saved hosting cluster context...
✔ Switching to hosting cluster context...
no kubeconfig context for wds1 was found: context wds1 not found for control plane wds1
trying to load new context wds1 from server...ne
✔ Overwriting existing context for control plane
✔ Switching to context wds1...
no kubeconfig context for its1 was found: context its1 not found for control plane its1
trying to load new context its1 from server...ne
✔ Overwriting existing context for control plane
✔ Switching to context its1...

Waiting for OCM cluster manager to be ready...
controlplane.tenancy.kflex.kubestellar.org/its1 condition met
job.batch/its condition met

Waiting for OCM hub cluster-info to be updated...
job.batch/update-cluster-info condition met
✔ OCM hub is ready

Registering cluster 1 and 2 for remote access with KubeStellar Core...
Preflight check: HubKubeconfig check Passed with 0 warnings and 0 errors
Preflight check: DeployMode Check Passed with 0 warnings and 0 errors
Preflight check: ClusterName Check Passed with 0 warnings and 0 errors
Please log onto the hub cluster and run the following command:

    clusteradm accept --clusters cluster1

This is not needed when the ManagedClusterAutoApproval feature is enabled
Preflight check: HubKubeconfig check Passed with 0 warnings and 0 errors
Preflight check: DeployMode Check Passed with 0 warnings and 0 errors
Preflight check: ClusterName Check Passed with 0 warnings and 0 errors
Please log onto the hub cluster and run the following command:

    clusteradm accept --clusters cluster2

This is not needed when the ManagedClusterAutoApproval feature is enabled
Checking that the CSR for cluster 1 and 2 appears...

Waiting for cluster1 and cluster2 to be ready and then approve their CSRs
No resources found
CSR for cluster1 not found. Trying again...
cluster1 has been found, approving CSR
Starting approve csrs for the cluster cluster1
CSR cluster1-x5wxj approved
set hubAcceptsClient to true for managed cluster cluster1

 Your managed cluster cluster1 has joined the Hub successfully. Visit https://open-cluster-management.io/scenarios or https://github.com/open-cluster-management-io/OCM/tree/main/solutions for next steps.
cluster2 has been found, approving CSR
Starting approve csrs for the cluster cluster2
CSR cluster2-gp6c9 approved
set hubAcceptsClient to true for managed cluster cluster2

 Your managed cluster cluster2 has joined the Hub successfully. Visit https://open-cluster-management.io/scenarios or https://github.com/open-cluster-management-io/OCM/tree/main/solutions for next steps.

Checking the new clusters are in the OCM inventory and label them
NAME       HUB ACCEPTED   MANAGED CLUSTER URLS                  JOINED   AVAILABLE   AGE
cluster1   true           https://cluster1-control-plane:6443                        17s
cluster2   true           https://cluster2-control-plane:6443                        11s
managedcluster.cluster.open-cluster-management.io/cluster1 labeled
managedcluster.cluster.open-cluster-management.io/cluster2 labeled

✔ Congratulations! Your KubeStellar demo environment is now ready to use.

Be sure to execute the following commands to set the shell variables expected in the example scenarios.

host_context=kind-kubeflex
its_cp=its1
its_context=its1
wds_cp=wds1
wds_context=wds1
wec1_name=cluster1
wec2_name=cluster2
wec1_context=$wec1_name
wec2_context=$wec2_name
label_query_both=location-group=edge
label_query_one=name=cluster1
```
