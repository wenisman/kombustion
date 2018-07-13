package plugins

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// ExtractResourcesFromPlugins and ensure there are no clashes for plugin resource names
func ExtractResourcesFromPlugins(
	loadedPlugins []*PluginLoaded,
	resources *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Resources != nil {
			for key, parserFunc := range *plugin.Resources {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*resources)[pluginKey]; ok { // Check for duplicates
					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load resource `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						resources, _ := loadResource(parserFunc(name, data))
						// TODO: print errs here as we know what plugin they came from
						return resources, nil
					}
					(*resources)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}
	return
}

// ExtractMappingsFromPlugins and ensure there are no clashes for plugin resource names
func ExtractMappingsFromPlugins(
	loadedPlugins []*PluginLoaded,
	mappings *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Mappings != nil {
			for key, parserFunc := range *plugin.Mappings {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*mappings)[pluginKey]; ok { // Check for duplicates
					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load mapping `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						mapping, _ := loadMapping(parserFunc(name, data))
						return mapping, nil
					}
					(*mappings)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}

	return
}

// ExtractOutputsFromPlugins and ensure there are no clashes for plugin resource names
func ExtractOutputsFromPlugins(
	loadedPlugins []*PluginLoaded,
	outputs *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Outputs != nil {
			for key, parserFunc := range *plugin.Outputs {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*outputs)[pluginKey]; ok { // Check for duplicates

					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load output `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						outputs, _ := loadOutput(parserFunc(name, data))
						return outputs, nil
					}
					(*outputs)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}

	return
}
