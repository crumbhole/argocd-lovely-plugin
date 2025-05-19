package main

import (
	"bytes"
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v3"
	"jaytaylor.com/html2text"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func main() {
	err := parameterMarkdown()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	err = pluginYaml()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func parameterMarkdown() error {
	f, err := os.OpenFile("doc/parameter.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	err = appendFile(f, "doc/.snippets/parameterHeader.md")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "|Name | Environment variable | Description | Default |\n")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "| ---- | -------------------- | ----------- | ------- |\n")
	if err != nil {
		return err
	}
	for featNum := features.FirstFeature; featNum <= features.LastFeature; featNum++ {
		feature := features.Features()[featNum]
		_, err = fmt.Fprintf(f, "| %s | %s | %s | %s |\n",
			feature.Title,
			feature.EnvName(),
			feature.Description,
			feature.DefaultVal)
		if err != nil {
			return err
		}
	}
	err = appendFile(f, "doc/.snippets/parameterFooter.md")
	return err
}

func appendFile(file *os.File, name string) error {
	// #nosec - G304 documentation generator, not in product
	content, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	return err
}

func pluginYaml() error {
	// return nil
	plugin := PluginConfig{
		APIVersion: `argoproj.io/v1alpha1`,
		Kind:       ConfigManagementPluginKind,
		Metadata: metav1.ObjectMeta{
			Name: `argocd-lovely-plugin`,
		},
		Spec: PluginConfigSpec{
			Version: `v1.0`,
			Generate: Command{
				Command: []string{`argocd-lovely-plugin`},
			},
		},
	}
	for featNum := features.FirstFeature; featNum <= features.LastFeature; featNum++ {
		feature := features.Features()[featNum]
		var param ParameterAnnouncement
		param.Name = feature.Name
		param.Title = feature.Title
		param.Required = false
		param.ItemType = `string`
		param.CollectionType = `string` // feature.CollectionType.String()
		tooltip, err := html2text.FromString(
			string(markdown.ToHTML([]byte(feature.Description), nil, nil)),
			html2text.Options{OmitLinks: true})
		if err != nil {
			return err
		}
		param.Tooltip = tooltip
		param.StringDefault = feature.DefaultVal
		plugin.Spec.Parameters.Static = append(plugin.Spec.Parameters.Static, &param)
	}
	var yamlText bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&yamlText)
	yamlEncoder.SetIndent(2)
	err := yamlEncoder.Encode(&plugin)
	if err != nil {
		return err
	}
	// #nosec - G306 - this needs to be readable by argocd
	err = os.WriteFile(`plugin.yaml`, yamlText.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
