# Kubebuilder

# Runbook
After running [Kubebuilder init and startup basics for project](./INIT.md),
follow the `Developer Workflow` loop.
## Developer workflow
1. make changes;
2. `make install`;
3. `make run`;
4. `apply` to cluster.

## Sample controller workflow:
1. Make change in guestbook controller's `Reconcile`
