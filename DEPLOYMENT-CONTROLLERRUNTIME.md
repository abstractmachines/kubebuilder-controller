# Create a Deployment (and learn controller runtime pkg)

> TLDR: 

>> Use the [client's `Create()` method](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client) inside the controller's Reconcile() function to create a new Deployment.

## Read: Navigate the docs for the Kubernetes API pkg.go.dev
1. Read through [Controller Runtime package docs, esp. Clients & Caches](https://pkg.go.dev/sigs.k8s.io/controller-runtime#hdr-Clients_and_Cachese)
2. Navigate the docs: append `/pkg/client` (or `manager`, etc) after `controller-runtime`
3. Look at [Client type](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client)

## Read: Deployment type (navigating pkg.go.dev API docs)
> Navigate to [https://pkg.go.dev/k8s.io/api](https://pkg.go.dev/k8s.io/api)
- Note the `apps` section and different versions. Recall how client-go and go pkgs
work with kubernetes API; alpha, beta, and then GA. The `appsv1` is hence the 
"GA" of what we want; so [go there](https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1)
- Click on `Deployment`. That'll lead us here: [https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1#Deployment]https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1#Deployment)

## Imports: /apps/v1 and metav1 from apimachinery
```
appsv1 "k8s.io/api/apps/v1"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
```

## 
