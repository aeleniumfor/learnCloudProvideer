# go run main.go --cloud-provider=cloud --use-service-account-credentials=true  --leader-elect=true #--client-ca-file=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt


        # - --use-service-account-credentials

kind delete cluster
kind create cluster --config $(pwd)/kubernetes/multi.yml
export KUBECONFIG="$(kind get kubeconfig-path)"

kubectl taint nodes kind-control-plane node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker2 node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule
kubectl taint nodes kind-worker3 node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule

