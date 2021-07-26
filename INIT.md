# Kubebuilder init and startup basics for project

## 1. Kubebuilder init.

```
kubebuilder init --domain example.com --repo github.com/abstractmachines/kubebuilder-tutorial
```

Or
```
kubebuilder init --domain example.com --repo github.com/abstractmachines/kubebuilder-tutorial --skip-go-version-check
```

> Result: codegen scaffolded project

## 2. Create API (also scaffold controllers and resource).

```
kubebuilder create api --group webapp --version v1 --kind Guestbook --resource --controller
```

> Result: codegen `api/v1/guestbook_types.go` `controllers/guestbook_controller.go`

## 3. Test it out: Install CRDs into cluster
```
make install
```

> Results:
  - This generates yaml, including our CRD's definition/manifest. Note `kind: CustomResourceDefinition`, with `group: webapp.example.com`, and `name` of `guestbook` (lowercase version of the name's `kind: Guestbook`).
  - `controller-gen` codegen; kustomize build generates `./config/crd` yaml layers and then imperatively `apply`'s CRD to cluster;
  - Hence, result: `customresourcedefinition.apiextensions.k8s.io/guestbooks.webapp.example.com created`.

## 4. Test it out: Run the Controller
```
make run
```

This will `go run main.go` and `controller runtime`

## 5. Test it out: `apply` instances of Custom Resources
In new tab:
```
k apply -f config/samples/
```

> Result: This actually applies our custom resource `kind: Guestbook` into cluster.

## 6. kubectl get guestbook .... kubectl get crds
```
k get guestbook // shows cr guestbook and time of creation
```

```
k get crds // guestbooks.webapp.example.com 
```

```
k get guestbook -owide -oyaml  // our sample yaml def
```
## 7. Deploy.
You can use Docker or other methods such as creating a deployment.

Let's use Docker for now:
```
make docker-build docker-push IMG=<some-registry>/<project-name>:tag
make deploy IMG=<some-registry>/<project-name>:tag
```

> Result: This will create a `Deployment` and other Kustomize layers, service etc...

## 8. Similar to step 6: kubectl get guestbook .... kubectl get crds

## 9. Shut down
```
make uninstall // uninstall CRDs // no more `k get guestbook`
make undeploy // undeploy controller
// Shut down `make run` process in terminal
```
