# Helm-rbac

## Notes


```sh
kubectl exec --stdin --tty <POD> -- /bin/bash
```

```
docker run -ti --rm --network host -v ~/.minikube:/home/ianchen0119/.minikube   -v ~/.kube:/root/.kube -v ~/.helm:/root/.helm -v ~/.config/helm:/root/.config/helm     -v ~/.cache/helm:/root/.cache/helm helm-client:latest
```

```
ianchen0119@ubuntu:~/.kube$ cat config
```

```
kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
```

## References
- [How to create a kubernetes serviceAccount when I do helm install?](https://stackoverflow.com/questions/72504732/how-to-create-a-kubernetes-serviceaccount-when-i-do-helm-install)
- [chart 开发提示和技巧](https://helm.sh/zh/docs/howto/charts_tips_and_tricks/)
- [Kubernetes-Host网络模式应用](https://www.cnblogs.com/zhenyuyaodidiao/p/6739099.html)
- [Kubernetes docs](https://kubernetes.io/docs/concepts/storage/volumes/)