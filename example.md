There are 1 pods in the cluster
```
Name: testnginx
Namespace: default
Status: Running
Containers:
- Name: nginx
  Image: nginx
------
Name : testnginx
GenerateName :
Namespace : default
SelfLink :
UID : 0c497a60-6f91-4bca-b2b6-319b215f947d
ResourceVersion : 24414
Generation : 0
CreationTimestamp : 2023-11-19 15:51:37 +0800 CST
DeletionTimestamp : <nil>
DeletionGracePeriodSeconds : <nil>
Labels : map[]
Annotations : map[cni.projectcalico.org/containerID:9cd6c5990f60700f2ee353a3012e7bab8bf7fd42fcc25c230c3abe80bb160615 cni.projectcalico.org/podIP:172.16.175.200/32 cni.projectcalico.org/podIPs:172.16.175.200/32 kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"name":"testnginx","namespace":"default"},"spec":{"containers":[{"image":"nginx","name":"nginx"}]}}
]
OwnerReferences : []
Finalizers : []
ManagedFields : [{kubectl-client-side-apply Update v1 2023-11-19 15:51:37 +0800 CST FieldsV1 {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:spec":{"f:containers":{"k:{\"name\":\"nginx\"}":{".":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}}},"f:dnsPolicy":{},"f:enableServiceLinks":{},"f:restartPolicy":{},"f:schedulerName":{},"f:securityContext":{},"f:terminationGracePeriodSeconds":{}}} } {calico Update v1 2023-11-19 15:51:38 +0800 CST FieldsV1 {"f:metadata":{"f:annotations":{"f:cni.projectcalico.org/containerID":{},"f:cni.projectcalico.org/podIP":{},"f:cni.projectcalico.org/podIPs":{}}}} status} {kubelet Update v1 2023-11-19 15:52:15 +0800 CST FieldsV1 {"f:status":{"f:conditions":{"k:{\"type\":\"ContainersReady\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}},"k:{\"type\":\"Initialized\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}},"k:{\"type\":\"Ready\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}}},"f:containerStatuses":{},"f:hostIP":{},"f:phase":{},"f:podIP":{},"f:podIPs":{".":{},"k:{\"ip\":\"172.16.175.200\"}":{".":{},"f:ip":{}}},"f:startTime":{}}} status}]
Phase : Running
Conditions : [{Initialized True 0001-01-01 00:00:00 +0000 UTC 2023-11-19 15:51:37 +0800 CST  } {Ready True 0001-01-01 00:00:00 +0000 UTC 2023-11-19 15:52:16 +0800 CST  } {ContainersReady True 0001-01-01 00:00:00 +0000 UTC 2023-11-19 15:52:16 +0800 CST  } {PodScheduled True 0001-01-01 00:00:00 +0000 UTC 2023-11-19 15:51:37 +0800 CST  }]
Message :
Reason :
NominatedNodeName :
HostIP : 192.168.11.102
HostIPs : []
PodIP : 172.16.175.200
PodIPs : [{172.16.175.200}]
StartTime : 2023-11-19 15:51:37 +0800 CST
InitContainerStatuses : []
ContainerStatuses : [{nginx {nil &ContainerStateRunning{StartedAt:2023-11-19 15:52:15 +0800 CST,} nil} {nil nil nil} true 0 docker.io/library/nginx:latest docker.io/library/nginx@sha256:86e53c4c16a6a276b204b0fd3a8143d86547c967dc8258b3d47c3a21bb68d3c6 containerd://2fa67fb7046cf05a32bb03e03cca17bef010a6dd47f85a40a1141dc751ad5744 0xc000409105 map[] nil}]
QOSClass : BestEffort
EphemeralContainerStatuses : []
Resize :
ResourceClaimStatuses : []
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2023/11/19 07:52:15 [notice] 1#1: using the "epoll" event method
2023/11/19 07:52:15 [notice] 1#1: nginx/1.25.3
2023/11/19 07:52:15 [notice] 1#1: built by gcc 12.2.0 (Debian 12.2.0-14)
2023/11/19 07:52:15 [notice] 1#1: OS: Linux 5.15.0-83-generic
2023/11/19 07:52:15 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2023/11/19 07:52:15 [notice] 1#1: start worker processes
2023/11/19 07:52:15 [notice] 1#1: start worker process 28
2023/11/19 07:52:15 [notice] 1#1: start worker process 29
```