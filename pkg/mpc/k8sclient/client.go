package k8sclient

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"nanto.io/application-auto-scaling-service/pkg/common/utils/logutil"
	apiextensionsclientset "nanto.io/application-auto-scaling-service/pkg/mpc/k8sclient/clientset/versioned"
)

var logger = logutil.GetLogger()

var clientSet *K8sClientSet

// K8sClientSet 包含 标准kube clientset 和 自定义资源的 clientset
type K8sClientSet struct {
	// kubeClientset is a standard kubernetes clientset
	kubeClientset *kubernetes.Clientset
	// crdClientset is a clientset for our own API group
	crdClientset apiextensionsclientset.Interface
}

// GetKubeClientSet 获取 标准kube clientset
func GetKubeClientSet() *kubernetes.Clientset {
	if clientSet == nil {
		logger.Panic("K8sClientSet invalid")
	}
	return clientSet.kubeClientset
}

// GetCrdClientSet 获取 自定义资源的 clientset
func GetCrdClientSet() apiextensionsclientset.Interface {
	if clientSet == nil {
		logger.Panic("K8sClientSet invalid")
	}
	return clientSet.crdClientset
}

// InitK8sClientSet 初始化 k8s client set
func InitK8sClientSet(kubeconfig string) error {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return errors.Wrap(err, "Error building kubeConfig")
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "Error building kubernetes clientset")
	}
	crdClient, err := apiextensionsclientset.NewForConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "Error building example clientset")
	}
	clientSet = &K8sClientSet{
		kubeClientset: kubeClient,
		crdClientset:  crdClient,
	}
	return nil
}
