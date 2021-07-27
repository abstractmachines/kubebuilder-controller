/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	// "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "github.com/abstractmachines/kubebuilder-tutorial/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.example.com,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.example.com,resources=guestbooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.example.com,resources=guestbooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guestbook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("guestbook", req.NamespacedName)

	// your logic here

	// 1. *** Let's just retrieve a resource. ***
	// Get the client object:
	var guestbook webappv1.Guestbook

	err := r.Get(ctx, req.NamespacedName, &guestbook)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// We now have our custom resource, its namespace, labels ...
	fmt.Print("successfully retrieved guestbook resource:")
	fmt.Println(guestbook)
	fmt.Println(req.Namespace)
	fmt.Println(guestbook.Labels)

	// 2. *** Let's create a bare-bones Deployment using only required fields ***
	deployment := appsv1.Deployment{}
	deployment.ObjectMeta = metav1.ObjectMeta{
		Name:      guestbook.Name,
		Namespace: guestbook.Namespace,
	}
	// Match labels / LabelSelector / match app to deployment
	deployment.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app":  "guestbook",
			"tier": "frontend",
		},
	}
	// Add those labels to deployment
	deployment.Spec.Template.ObjectMeta.Labels = map[string]string{
		"app":  "guestbook",
		"tier": "frontend",
	}
	// spec.template.spec.containers required fields
	deployment.Spec.Template.Spec.Containers = make([]corev1.Container, 1)
	deployment.Spec.Template.Spec.Containers[0].Name = "frontend"
	deployment.Spec.Template.Spec.Containers[0].Image = "gcr.io/google-samples/gb-frontend:v4"
	// create deployment
	err = r.Create(ctx, &deployment)

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}
