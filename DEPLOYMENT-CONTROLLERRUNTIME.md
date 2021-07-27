# Create a Deployment (and learn controller runtime pkg)

> TLDR: 

>> Use the [client's `Create()` method](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client) inside the controller's Reconcile() function to create a new Deployment.

## 1. Read: Navigate the docs for the Kubernetes API pkg.go.dev
1. Read through [Controller Runtime package docs, esp. Clients & Caches](https://pkg.go.dev/sigs.k8s.io/controller-runtime#hdr-Clients_and_Cachese)
2. Navigate the docs: append `/pkg/client` (or `manager`, etc) after `controller-runtime`
3. Look at [Client type](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Client)

## 2. Read: Deployment type (navigating pkg.go.dev API docs)
> Navigate to [https://pkg.go.dev/k8s.io/api](https://pkg.go.dev/k8s.io/api)
- Note the `apps` section and different versions. Recall how client-go and go pkgs
work with kubernetes API; alpha, beta, and then GA. The `appsv1` is hence the 
"GA" of what we want; so [go there](https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1)
- Click on `Deployment`. That'll lead us here: [https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1#Deployment]https://pkg.go.dev/k8s.io/api@v0.21.3/apps/v1#Deployment)

## 3. Imports
```
appsv1 "k8s.io/api/apps/v1"  ---> for deployments, etc
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" ---> ObjectMeta data, fields, etc ...
corev1 "k8s.io/api/core/v1" ---> spec, container fields ...

"k8s.io/apimachinery/pkg/runtime"
"sigs.k8s.io/controller-runtime/pkg/client"
```

## 4. Code: Bare bones deployment with label selector (discovery) and container image, name
> write labelselector ("discovery") code in controller Reconcile() function
- 1. Choose what labels our Deployment will select (help your Deployment find your `app`)
- 2. Add those labels to the metadata of the `app` as well (so that Deployment can find it)
> Write code for container (`make`, then add fields for `name` and `image` to pull for it)
- See `guestbook_controller.go`

## 5. Test: make run, then get deployments and pods
> Run controller:
```
make run
```

> While controller is running, look up the Deployment it created (you could also have `watch`ed):
```
k get deployment -n default guestbook-sample  -owide

// gives us:
NAME               READY   UP-TO-DATE   AVAILABLE   AGE    CONTAINERS   IMAGES                                 SELECTOR
guestbook-sample   1/1     1            1           5m7s   frontend     gcr.io/google-samples/gb-frontend:v4   app=guestbook,tier=frontend
```

> Do pods have those labels? They have to, for discovery purposes for k8s API:
```
k get pods -A -l app=guestbook

// gives us:
NAMESPACE   NAME                               READY   STATUS    RESTARTS   AGE
default     guestbook-sample-94dbc5dcd-kqqz8   1/1     Running   0          6m20s
```

## 6. Think about other things like clients, caching, shared informers, watchers ...
At another later time :)
