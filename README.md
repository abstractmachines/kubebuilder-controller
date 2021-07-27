# Kubebuilder: Controllers, CRDs, and Operators

> How to build custom controllers, the easy way.
>> To begin, clone (quickstart or kube builder book): [https://book.kubebuilder.io/](https://book.kubebuilder.io/)

- Inspired by https://github.com/sethp-nr/guestbook-workshop
- See also: [https://maelvls.dev/learning-kubernetes-controllers/](https://maelvls.dev/learning-kubernetes-controllers/)
- See also: [Kubernetes community document on controllers](https://github.com/kubernetes/community/blob/712590c108bd4533b80e8f2753cadaa617d9bdf2/contributors/devel/sig-api-machinery/controllers.md)

# Runbooks
Not a tutorial based on theory. Just some basic developer workflow runbooks.
## Runbook (startup)
- After running [Kubebuilder init and startup basics for project](./INIT.md),
- including a `make install`, `make run`, and probably an `apply`,
- follow the `Developer Workflow` loop below.
## Developer workflow for custom controllers
1. Run controller (`make run`).
2. Make changes to code while running controller.
3. `apply` to cluster and note changes.

Note that the `apply` will cause the `Reconciler` to "reconcile." So `apply`ing a change to the cluster will, for example, `log` any `printf` statements in the controller's `reconcile()` function.

## Sample Controller workflows
Workflows for developing custom CRDs and controllers (custom operators).

> Log statement:
1. Make change in guestbook controller's `Reconcile` code (add `fmt.Println("henlo")`).
2. Run controller (`make run`).
3. See the controller print "henlo" in the `make run` terminal window/tmux window.

>  Add labels imperatively and apply to cluster:
1. Run controller (`make run`).
2. Add labels to `./config/samples/webapp_v1_guestbook.yaml` under `metadata`
  ```
  metadata:
  name: guestbook-sample
  labels:
    is-awesome: totes
    is-bad: nopes
  ```
3. `apply` to cluster: `apply -f ./config/samples`.
4. Use CLI label selector. `k get guestbook -l is-awesome=totes`. You'll see `guestbook` crd.

> Get guestbook CRD, its labels, its namespace ...
1. Use Reconciler function in Controller to `Get` object properties from `NamespacedName` struct.
2. Run controller (`make run`).
3. Note changes

> **Create a basic Deployment**
- Use this document [DEPLOYMENT-CONTROLLERRUNTIME.md](DEPLOYMENT-CONTROLLERRUNTIME.md)
- Topics covered: `controller runtime pkg`, `client pkg`, `appsv1 pkg`, etc
- You'll learn how to navigate the Kubernetes API at pkg.go.dev

> **Create a basic Service**