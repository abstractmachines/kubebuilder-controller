# Kubebuilder

## Runbook

> 1. Kubebuilder init

  ```
  kubebuilder init --domain example.com --repo github.com/abstractmachines/kubebuilder-tutorial
  ```

  Or
  ```
  kubebuilder init --domain example.com --repo github.com/abstractmachines/kubebuilder-tutorial --skip-go-version-check
  ```
> Result:

That'll build kustomize layers and scaffolding.


> 2. Create API (also scaffold controllers and resource).

```
kubebuilder create api --group webapp --version v1 --kind Guestbook --resource --controller
```

> Result: codegen api/v1/guestbook_types.go controllers/guestbook_controller.go

