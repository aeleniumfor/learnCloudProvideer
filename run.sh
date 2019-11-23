# go run main.go --cloud-provider=cloud --use-service-account-credentials=true  --leader-elect=true #--client-ca-file=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt


        - --use-service-account-credentials

kubectl taint nodes kind-control-plane node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker2 node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker3 node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule

