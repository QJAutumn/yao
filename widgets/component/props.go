package component

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/yaoapp/gou"
	"github.com/yaoapp/kun/log"
)

// CloudProps parse CloudProps
func (p PropsDSL) CloudProps(xpath string) (map[string]CloudPropsDSL, error) {
	return p.parseCloudProps(xpath, p)
}

// ExecQuery execute query
func (cProp CloudPropsDSL) ExecQuery(process *gou.Process, query map[string]interface{}) (interface{}, error) {

	if query == nil {
		query = map[string]interface{}{}
	}

	// Process
	name := cProp.Process
	if name == "" {
		log.Error("[component] %s.$%s process is required", cProp.Xpath, cProp.Name)
		return nil, fmt.Errorf("[component] %s.$%s process is required", cProp.Xpath, cProp.Name)
	}

	// Create process
	p, err := gou.ProcessOf(name, query)
	if err != nil {
		log.Error("[component] %s.$%s %s", cProp.Xpath, cProp.Name, err.Error())
		return nil, fmt.Errorf("[component] %s.$%s %s", cProp.Xpath, cProp.Name, err.Error())
	}

	// Excute process
	res, err := p.WithGlobal(process.Global).WithSID(process.Sid).Exec()
	if err != nil {
		log.Error("[component] %s.$%s %s", cProp.Xpath, cProp.Name, err.Error())
		return nil, fmt.Errorf("[component] %s.$%s %s", cProp.Xpath, cProp.Name, err.Error())
	}

	return res, nil
}

// Replace xpath
func (cProp CloudPropsDSL) Replace(data interface{}, replace func(cProp CloudPropsDSL) interface{}) error {
	return cProp.replaceAny(data, "", replace)
}

func (cProp CloudPropsDSL) replaceAny(data interface{}, root string, replace func(cProp CloudPropsDSL) interface{}) error {
	switch data.(type) {
	case map[string]interface{}:
		return cProp.replaceMap(data.(map[string]interface{}), root, replace)
	}
	return nil
}

func (cProp CloudPropsDSL) replaceMap(data map[string]interface{}, root string, replace func(cProp CloudPropsDSL) interface{}) error {
	xpath := fmt.Sprintf(".%s.$%s", cProp.Xpath, cProp.Name)
	for key := range data {
		path := fmt.Sprintf("%s.%s", root, key)
		if !strings.HasPrefix(xpath, path) {
			continue
		}

		// Replace field
		if path == xpath {
			data[cProp.Name] = replace(cProp)
			delete(data, fmt.Sprintf("$%s", cProp.Name))
			continue
		}

		err := cProp.replaceAny(data[key], path, replace)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p PropsDSL) parseCloudProps(xpath string, props map[string]interface{}) (map[string]CloudPropsDSL, error) {

	res := map[string]CloudPropsDSL{}

	for name, prop := range props {

		fullname := fmt.Sprintf("%s.%s", xpath, name)
		if sub, ok := prop.(map[string]interface{}); ok {
			cProps, err := p.parseCloudProps(fullname, sub)
			if err != nil {
				return nil, err
			}
			for k, v := range cProps {
				res[k] = v
			}
		}

		if !strings.HasPrefix(name, "$") {
			continue
		}

		cProp := &CloudPropsDSL{
			Name:  strings.TrimPrefix(name, "$"),
			Xpath: xpath,
		}

		err := cProp.Parse(prop)
		if err != nil {
			return nil, fmt.Errorf("%s %s", fullname, err.Error())
		}

		cProp.Xpath = xpath
		res[fullname] = *cProp
	}

	return res, nil
}

// Parse parse cloud props
func (cProp *CloudPropsDSL) Parse(v interface{}) error {

	bytes, err := jsoniter.Marshal(v)
	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(bytes, cProp)
	if err != nil {
		return err
	}
	return nil
}
