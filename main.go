package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/ca-gip/kubi-members/internal/controller"
	"github.com/ca-gip/kubi-members/internal/ldap"
	membersclientset "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned"
	projectclientset "github.com/ca-gip/kubi/pkg/generated/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.StringVar(&kubeconfig, "kubeconfig", defaultKubeconfig(), "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

	klog.InitFlags(nil)

	flag.Parse()

	// Load kube config
	cfg, err := rest.InClusterConfig()
	if err != nil {
		cfg, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
		if err != nil {
			klog.Fatalf("Error building kubeconfig: %s", err.Error())
		}
	}

	// Generate clientsets
	configMapClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes configMapClient: %s", err.Error())
	}

	projectClient, err := projectclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes projectClient: %s", err.Error())
	}

	membersClient, err := membersclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes membersClient: %s", err.Error())
	}

	klog.Info("Creating LDAP client")

	ldapClient := ldap.NewLdap()

	controller := controller.NewController(configMapClient, projectClient, membersClient, ldapClient)

	if err := controller.Run(); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func defaultKubeconfig() string {
	fname := os.Getenv("KUBECONFIG")
	if fname != "" {
		return fname
	}
	home, err := os.UserHomeDir()
	if err != nil {
		klog.Warningf("failed to get home directory: %v", err)
		return ""
	}
	return filepath.Join(home, ".kube", "config")
}
