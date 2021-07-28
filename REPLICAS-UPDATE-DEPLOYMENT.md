# Specify a desired state (scaled replicas), and update Deployment

## Let's create the replicas via types and schema, and apply it to our cluster!

1. **Create a spec.**

    Fill out `spec` in `types.go` file (in `/api/`) to specify a new `spec` struct for replicas;
    name it `FrontendSpec` and give it a `Replicas` field:
    ```
    type FrontendSpec struct {
    // +optional
    Replicas *int32 `json:"replicas,omitempty"` // If it's nil, we'll use a default
    }
    ```

    Refer to `FrontendSpec` struct (e.g. the `replicas`) in the `GuestbookSpec` struct (so now the `spec` for our CRD (`GuestbookSpec`) will have a field for the `replicas` in its type defintion):
    ```
    type GuestBookSpec struct {
      // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
      // Important: Run "make" to regenerate code after modifying this file

      Frontend FrontendSpec `json:"frontend"`
    }
    ```

3. **Update the YAML definition** too (in `config/samples/webapp_v1_guestbook.yaml):

    under `spec` add:
    ```
    spec:
      frontend:
        replicas: 3
    ```
4. Update the logging in the controller to indicate data going in and out correctly:
    ```
    log.Info("successfully retrieved replicas:", guestbook.Spec.Frontend.Replicas)
    ```

5. Tell kubernetes about our custom type: `make install`;
6. Run our controller pointed at our kind cluster: `make run`;
7. Finally, apply an instance of our new Guestbook CRD with updated schemas to cluster: `k apply -f config/samples`.

## Let's talk about why that doesn't work.

> Note that when we apply those changes, our `deployment` still isn't affected.

- To see this in action, run `k get deployment -n default guestbook-sample -oyaml`.
- Note that `replicas` in both `state` and `spec` for our YAML def here _all say 1, not 3._
- Further, our output from running `make run` for our controller won't print out "hey, 3 spec replicas!" It'll be some HEX memory location mess.
> **This is because we didn't update our deployment.**

## What's a Deployment again?
"A Deployment provides declarative updates for Pods and ReplicaSets.

You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate."

[https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## So we need to update our Deployment in the controller, if already exists.
Create replicas in `controller` file:
```
replicas := int32(3)

if guestbook.Spec.Frontend.Replicas != nil {
  replicas = *guestbook.Spec.Frontend.Replicas
}
```

Change our controller code from `"just Create() a deployment"` to:
- `"check if a deployment exists; if so, Update and return"`;
- `"else fall through to our existing "just create a Deployment" code"`.
```
// Updating existing deployment with spec of scaled replicas ***
// Does a Deployment with this name in this ns already exist?
err = r.Get(ctx, types.NamespacedName{
  Name:      guestbook.Name,
  Namespace: guestbook.Namespace,
}, &deployment)

// 5. If deployment already exists (GET deployment has no error),
// we'll want to update it, and return before Create().
if err == nil {
  deployment.Spec.Replicas = &replicas

  err = r.Update(ctx, &deployment)

  // and handle errors resulting from that operation ...
  if err != nil {
    return ctrl.Result{}, err
  }
  return ctrl.Result{}, nil
} else if !apierrors.IsNotFound(err) {
  return ctrl.Result{}, err
}
```

## Now check the spec after applying changes to cluster and recompiling!

> We reinstall/re-run, apply yaml, and check our updated deployment's `replicas`:
- `make install`
- `make run`
- `apply -f config/samples` in new terminal as always;
- Then check our  yaml in cluster for our Deployment:
    ```
    k get deployment -n default default guestbook-sample   -oyaml
    ```

    With this result notice the updated `spec` for deployment's replicas is now 3, not 1:
    ```
    spec:
      replicas: 3
    ```

    Also, we probably have an updated state:
    ```
    status:
      availableReplicas: 3
      replicas: 3
      updatedReplicas: 3
    ```

## We did it!