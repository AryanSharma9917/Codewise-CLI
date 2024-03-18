package kubernetes

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	k80bj string
)

// kubernetesManifestCmd is the command for generating manifest files for kubernetes objects.
var kubernetesManifestCmd = &cobra.Command{
	Use:   "manifest [flags]",
	Short: "Generates manifest file for different objects.",
	Run: func(cmd *cobra.Command, args []string) {

		k80bj = strings.ToLower(k80bj)
		switch {
		case k80bj == "deployment":
			createManifestFile("deployment.yaml", deployment)
		case k80bj == "pod":
			createManifestFile("pod.yaml", pod)
		case k80bj == "service":
			createManifestFile("service.yaml", service)
		case k80bj == "ingress":
			createManifestFile("ingress.yaml", ingress)
		case k80bj == "secret":
			createManifestFile("secret.yaml", secret)
		case k80bj == "configmap":
			createManifestFile("configmap.yaml", configmap)
		case k80bj == "persistentvolume" || k80bj == "pv":
			createManifestFile("persistentvolume.yaml", pv)
		case k80bj == "persistentvolumeclaim" || k80bj == "pvc":
			createManifestFile("persistentvolumeclaim.yaml", pvc)
		default:
			log.Print("Currently we don't support manifest generation for " + k80bj + ".")
		}
	}}

func createManifestFile(filename string, obj string) {
	file, err := os.Create(filename)
	checkNilErr(err)

	defer file.Close()

	_, err = file.WriteString(obj)
	checkNilErr(err)

	log.Print(filename + " created successfully.")
}

func init() {
	kubernetesManifestCmd.Flags().StringVarP(&k80bj, "obj", "o", "", "Kubernetes object to generate manifest for.")
	err := kubernetesManifestCmd.MarkFlagRequired("obj")
	checkNilErr(err)

}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
