package api

import (
	o7tapiauth "github.com/openshift/api/authorization/v1"
	o7tapiquota "github.com/openshift/api/quota/v1"
	o7tapiroute "github.com/openshift/api/route/v1"
	o7tapisecurity "github.com/openshift/api/security/v1"
	o7tapiuser "github.com/openshift/api/user/v1"
	"github.com/sirupsen/logrus"

	"k8s.io/api/apps/v1beta1"
	k8sapicore "k8s.io/api/core/v1"
	extv1b1 "k8s.io/api/extensions/v1beta1"
	k8sapistorage "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Resources represent api resources used in report
type Resources struct {
	QuotaList            *o7tapiquota.ClusterResourceQuotaList
	NodeList             *k8sapicore.NodeList
	PersistentVolumeList *k8sapicore.PersistentVolumeList
	StorageClassList     *k8sapistorage.StorageClassList
	NamespaceList        []NamespaceResources
	RBACResources        RBACResources
}

// RBACResources contains all resources related to RBAC report
type RBACResources struct {
	UsersList                      *o7tapiuser.UserList
	GroupList                      *o7tapiuser.GroupList
	ClusterRolesList               *o7tapiauth.ClusterRoleList
	ClusterRolesBindingsList       *o7tapiauth.ClusterRoleBindingList
	SecurityContextConstraintsList *o7tapisecurity.SecurityContextConstraintsList
}

// NamespaceResources holds all resources that belong to a namespace
type NamespaceResources struct {
	NamespaceName     string
	DaemonSetList     *extv1b1.DaemonSetList
	DeploymentList    *v1beta1.DeploymentList
	PodList           *k8sapicore.PodList
	ResourceQuotaList *k8sapicore.ResourceQuotaList
	RolesList         *o7tapiauth.RoleList
	RouteList         *o7tapiroute.RouteList
	PVCList           *k8sapicore.PersistentVolumeClaimList
}

var listOptions metav1.ListOptions

// ListNamespaces list all namespaces, wrapper around client-go
func ListNamespaces(client *kubernetes.Clientset, ch chan<- *k8sapicore.NamespaceList) {
	namespaces, err := client.CoreV1().Namespaces().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- namespaces
}

// ListPods list all pods in namespace, wrapper around client-go
func ListPods(client *kubernetes.Clientset, namespace string, ch chan<- *k8sapicore.PodList) {
	pods, err := client.CoreV1().Pods(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- pods
}

// ListPVs list all PVs, wrapper around client-go
func ListPVs(client *kubernetes.Clientset, ch chan<- *k8sapicore.PersistentVolumeList) {
	pvs, err := client.CoreV1().PersistentVolumes().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- pvs
}

// ListNodes list all nodes, wrapper around client-go
func ListNodes(client *kubernetes.Clientset, ch chan<- *k8sapicore.NodeList) {
	nodes, err := client.CoreV1().Nodes().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- nodes
}

// ListQuotas list all cluster quotas classes, wrapper around client-go
func ListQuotas(client *OpenshiftClient, ch chan<- *o7tapiquota.ClusterResourceQuotaList) {
	quotas, err := client.quotaClient.ClusterResourceQuotas().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- quotas
}

// ListResourceQuotas list all quotas classes, wrapper around client-go
func ListResourceQuotas(client *kubernetes.Clientset, namespace string, ch chan<- *k8sapicore.ResourceQuotaList) {
	quotas, err := client.CoreV1().ResourceQuotas(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- quotas
}

// ListRoutes list all routes classes, wrapper around client-go
func ListRoutes(client *OpenshiftClient, namespace string, ch chan<- *o7tapiroute.RouteList) {
	routes, err := client.routeClient.Routes(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- routes
}

// ListStorageClasses list all storage classes, wrapper around client-go
func ListStorageClasses(client *kubernetes.Clientset, ch chan<- *k8sapistorage.StorageClassList) {
	sc, err := client.StorageV1().StorageClasses().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- sc
}

// ListDeployments will list all deployments seeding in the selected namespace
func ListDeployments(client *kubernetes.Clientset, namespace string, ch chan<- *v1beta1.DeploymentList) {
	deployments, err := client.AppsV1beta1().Deployments(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- deployments
}

// ListDaemonSets will collect all DS from specific namespace
func ListDaemonSets(client *kubernetes.Clientset, namespace string, ch chan<- *extv1b1.DaemonSetList) {
	daemonSets, err := client.ExtensionsV1beta1().DaemonSets(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- daemonSets
}

// ListUsers list all users, wrapper around client-go
func ListUsers(client *OpenshiftClient, ch chan<- *o7tapiuser.UserList) {
	users, err := client.userClient.Users().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- users
}

// ListGroups list all users, wrapper around client-go
func ListGroups(client *OpenshiftClient, ch chan<- *o7tapiuser.GroupList) {
	groups, err := client.userClient.Groups().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- groups
}

// ListRoles list all storage classes, wrapper around client-go
func ListRoles(client *OpenshiftClient, namespace string, ch chan<- *o7tapiauth.RoleList) {
	roles, err := client.authClient.Roles(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- roles
}

// ListClusterRoles list all storage classes, wrapper around client-go
func ListClusterRoles(client *OpenshiftClient, ch chan<- *o7tapiauth.ClusterRoleList) {
	clusterRoles, err := client.authClient.ClusterRoles().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- clusterRoles
}

// ListClusterRolesBindings list all storage classes, wrapper around client-go
func ListClusterRolesBindings(client *OpenshiftClient, ch chan<- *o7tapiauth.ClusterRoleBindingList) {
	clusterRolesBindings, err := client.authClient.ClusterRoleBindings().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- clusterRolesBindings
}

// ListSCC list all security context constraints, wrapper around client-go
func ListSCC(client *OpenshiftClient, ch chan<- *o7tapisecurity.SecurityContextConstraintsList) {
	scc, err := client.securityClient.SecurityContextConstraints().List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- scc
}

// ListPVCs list all PVs, wrapper around client-go
func ListPVCs(client *kubernetes.Clientset, namespace string, ch chan<- *k8sapicore.PersistentVolumeClaimList) {
	pvcs, err := client.CoreV1().PersistentVolumeClaims(namespace).List(listOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	ch <- pvcs
}
