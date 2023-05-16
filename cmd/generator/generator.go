package main

import (
	"fmt"
	"os"
	// "github.com/argoproj/argo-cd/v2/cmpserver/plugin"
	// "github.com/argoproj/argo-cd/v2/reposerver/apiclient"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "gopkg.in/yaml.v3"
)


func main() {
	err := configMarkdown()
	if err != nil{
		fmt.Printf("%s\n",err)
		os.Exit(1)
	}
	err = pluginYaml()
	if err != nil{
		fmt.Printf("%s\n",err)
		os.Exit(1)
	}
// 		var param apiclient.ParameterAnnouncement;
// 		param.Name = feature.Name
// 		param.Title = feature.Title
// 		param.Required = false
// 		param.CollectionType = `string`
// 		param._String = feature.Default
// 		plugin.Spec.Parameters.Static = append(&param, yaml.Spec.Parameters.Static)
// 	}
}

func configMarkdown() error {
	f, err := os.OpenFile("config.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil{
		return err
	}
	defer f.Close()
	err = appendFile(f, ".docs/configHeader.md")
	if err != nil{
		return err
	}
	fmt.Fprintf(f,"|Name | Environment variable | Description | Default |\n")
	fmt.Fprintf(f,"| ---- | -------------------- | ----------- | ------- |\n")
	for _, feature := range features.Features {
		fmt.Fprintf(f,"| %s | %s | %s | %s |\n",
			feature.Title,
			feature.EnvName(),
			feature.Description,
			feature.DefaultVal)
	}
	err = appendFile(f, ".docs/configFooter.md")
	return err
}

func appendFile(file *os.File, name string) error {
 	content, err := os.ReadFile(name)
	if err != nil{
		return err
	}
	_, err = file.Write(content)
	return err
}

func pluginYaml() error {
	return nil
// 	plugin := plugin.PluginConfig{
// 		ApiVersion: `argoproj.io/v1alpha1`,
// 		Kind: plugin.ConfigManagementPluginKind,
// 		Metadata: metav1.ObjectMeta{
// 			name: `argocd-lovely-plugin`,
// 		},
// 		Spec: plugin.PluginConfigSpec{
// 			Version: `v1.0`,
// 			generate: plugin.Command{
// 				Command: [`argocd-lovely-plugin`],
// 			},
// 		},
// 	}
// 	parameters := plugin.Parameters
// 	for id, feature := range features.Features {
// 		var param apiclient.ParameterAnnouncement;
// 		param.Name = feature.Name
// 		param.Title = feature.Title
// 		param.Required = false
// 		param.CollectionType = `string`
// 		param._String = feature.Default
// 		plugin.Spec.Parameters.Static = append(&param, yaml.Spec.Parameters.Static)
// 	}
// 	pluginYaml, err := yaml.Marshal(&plugin)
// 	if err != nil {
// 		return err
// 	}
// 	printf("%s\n", pluginYaml)
}
