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
- Note that `replicas` in both `state` and `spec` for our YAML def here _all say one, not 3._
- Further, our output from running `make run` for our controller won't print out "hey, 3 spec replicas!" It'll be some HEX memory location mess.
> **This is because we didn't update our deployment.**

## What's a Deployment again?


7. Note that once our CRD `guestbook` is up again and so 