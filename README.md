# Helm-rbac

Manage helm chart in another helm chart.

## Prerequisites
- minikube
- kubectl
- helm
- docker
## Get started

### 0. Build docker image

Build required docker image by following command:
```sh
$ ./build.sh
```

Load docker image into minikube:
```sh
$ minikube image load helm-client
```

(Optional) Establish the docker container based on the image you created:
```
$ docker run -ti --rm --network host -v ~/.minikube:/home/ianchen0119/.minikube   -v ~/.kube:/root/.kube -v ~/.helm:/root/.helm -v ~/.config/helm:/root/.config/helm     -v ~/.cache/helm:/root/.cache/helm helm-client:latest
```

### 1. Modify the test-client helm chart

In `test-client/templates/deployment.yaml`:
```yaml
      volumes:
        - name: minikube
          hostPath:
            path: /home/ianchen0119/.minikube/
            type: Directory
        - name: kube
          hostPath:
            path: /home/ianchen0119/.kube/
            type: Directory
        - name: helm
          hostPath:
            path: /home/ianchen0119/.helm/
            type: Directory
        - name: config
          hostPath:
            path: /home/ianchen0119/.cache/helm/
            type: Directory
```

Fill the correct path of the volumes above.

And modify the mountPath of `minikube` volume:
```yaml
          volumeMounts:
            - mountPath: /home/ianchen0119/.minikube
              name: minikube
```

> Tips: get the mountPath that should be set by checking `~/.kube/config`, which is the configuration of `kubectl`:
> ```
> $ cd ~/.kube
> $ cat config
> ```
> We will see the path of the certificate-authority:
> **certificate-authority: /home/ianchen0119/.minikube/ca.crt**
> refer the result above, we will know where the `.minikube` directory shold be mounted.

### 2. Install the test-client helm chart

```sh
$ cd test-client
$ helm install test-client .
$ helm ls
```

Check whether the test-client has been installed successfully:
```sh
$ kubectl get pod -o wide
```
Result:
```
NAME                           READY   STATUS    RESTARTS   AGE   IP             NODE       NOMINATED NODE   READINESS GATES
test-client-6f79f7ddc4-sn59s   1/1     Running   0          33m   192.168.49.2   minikube   <none>           <none>
```
Send a http request to test-client by using curl:
```sh
$ curl 192.168.49.2:8080
```
Result:
```
ianchen0119@ubuntu:~/$ curl 192.168.49.2:8080
Welcome Gin Serverianchen0119@ubuntu:~/
```
Check the pods again:
```
ianchen0119@ubuntu:~/$ kubectl get pod -o wide
NAME                           READY   STATUS    RESTARTS   AGE   IP             NODE       NOMINATED NODE   READINESS GATES
test-client-6f79f7ddc4-sn59s   1/1     Running   0          34m   192.168.49.2   minikube   <none>           <none>
test-server-85db4b786c-qvbdb   1/1     Running   0          21h   172.17.0.3     minikube   <none>           <none>
```

### [Optional] debug commands

```sh
$ kubectl exec --stdin --tty <POD> -- /bin/bash
```

## References
- [How to create a kubernetes serviceAccount when I do helm install?](https://stackoverflow.com/questions/72504732/how-to-create-a-kubernetes-serviceaccount-when-i-do-helm-install)
- [chart 开发提示和技巧](https://helm.sh/zh/docs/howto/charts_tips_and_tricks/)
- [Kubernetes-Host网络模式应用](https://www.cnblogs.com/zhenyuyaodidiao/p/6739099.html)
- [Kubernetes docs](https://kubernetes.io/docs/concepts/storage/volumes/)