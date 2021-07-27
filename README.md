# Kubebuilder tutorial

> How to build custom controllers, the easy way.
- Clone https://book.kubebuilder.io/
- Inspired by https://github.com/sethp-nr/guestbook-workshop

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