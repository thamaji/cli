package cli

import (
	"errors"
	"strconv"
)

type Option interface {
	SetDefaultValue(map[string]interface{})
	Keywords() []string
	Parse([]*Node, map[string]interface{}) (int, error)
	Help() [2]string
}

type BoolOption struct {
	Name        string
	Short       string
	Description string
	Usage       string
	ArgUsage    string
}

func (option *BoolOption) SetDefaultValue(options map[string]interface{}) {
}

func (option *BoolOption) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *BoolOption) Parse(args []*Node, options map[string]interface{}) (int, error) {
	options[option.Name] = true
	return 0, nil
}

func (option *BoolOption) Help() [2]string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}
			usage += "--" + option.Name
		}
	}

	description := option.Description

	return [2]string{usage, description}
}

type StringOption struct {
	Name         string
	Short        string
	DefaultValue string
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *StringOption) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == "" {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *StringOption) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *StringOption) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {
		return 0, errors.New("missing required value: " + option.usage())
	}

	v := args[0].Value.(string) // .Type == T_Value の .Value は常に string
	options[option.Name] = v
	return 1, nil
}

func (option *StringOption) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "string"
			}
		}
	}

	return usage
}

func (option *StringOption) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != "" {
		description += " (default: " + option.DefaultValue + ")"
	}

	return [2]string{usage, description}
}

type IntOption struct {
	Name         string
	Short        string
	DefaultValue int
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *IntOption) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == 0 {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *IntOption) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *IntOption) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {
		return 0, errors.New("missing required value: " + option.usage())
	}

	v, err := strconv.ParseInt(args[0].Value.(string), 10, 64) // .Type == T_Value の .Value は常に string
	if err != nil {
		return 0, err
	}

	options[option.Name] = int(v)
	return 1, nil
}

func (option *IntOption) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "number"
			}
		}
	}

	return usage
}

func (option *IntOption) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != 0 {
		description += " (default: " + strconv.FormatInt(int64(option.DefaultValue), 10) + ")"
	}

	return [2]string{usage, description}
}

type Int32Option struct {
	Name         string
	Short        string
	DefaultValue int32
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *Int32Option) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == 0 {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *Int32Option) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *Int32Option) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {
		return 0, errors.New("missing required value: " + option.usage())
	}

	v, err := strconv.ParseInt(args[0].Value.(string), 10, 32) // .Type == T_Value の .Value は常に string
	if err != nil {
		return 0, err
	}

	options[option.Name] = int32(v)
	return 1, nil
}

func (option *Int32Option) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "number"
			}
		}
	}

	return usage
}

func (option *Int32Option) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != 0 {
		description += " (default: " + strconv.FormatInt(int64(option.DefaultValue), 10) + ")"
	}

	return [2]string{usage, description}
}

type Int64Option struct {
	Name         string
	Short        string
	DefaultValue int64
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *Int64Option) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == 0 {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *Int64Option) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *Int64Option) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {

		return 0, errors.New("missing required value: " + option.usage())
	}

	v, err := strconv.ParseInt(args[0].Value.(string), 10, 64) // .Type == T_Value の .Value は常に string
	if err != nil {
		return 0, err
	}

	options[option.Name] = v
	return 1, nil
}

func (option *Int64Option) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "number"
			}
		}
	}

	return usage
}

func (option *Int64Option) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != 0 {
		description += " (default: " + strconv.FormatInt(option.DefaultValue, 10) + ")"
	}

	return [2]string{usage, description}
}

type Float32Option struct {
	Name         string
	Short        string
	DefaultValue float32
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *Float32Option) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == 0 {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *Float32Option) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *Float32Option) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {
		return 0, errors.New("missing required value: " + option.usage())
	}

	v, err := strconv.ParseFloat(args[0].Value.(string), 32) // .Type == T_Value の .Value は常に string
	if err != nil {
		return 0, err
	}

	options[option.Name] = float32(v)
	return 1, nil
}

func (option *Float32Option) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "number"
			}
		}
	}

	return usage
}

func (option *Float32Option) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != 0 {
		description += " (default: " + strconv.FormatFloat(float64(option.DefaultValue), 'f', -1, 32) + ")"
	}

	return [2]string{usage, description}
}

type Float64Option struct {
	Name         string
	Short        string
	DefaultValue float64
	Description  string
	Usage        string
	ArgUsage     string
}

func (option *Float64Option) SetDefaultValue(options map[string]interface{}) {
	if option.DefaultValue == 0 {
		return
	}
	options[option.Name] = option.DefaultValue
}

func (option *Float64Option) Keywords() []string {
	keywords := []string{}

	if option.Short != "" {
		keywords = append(keywords, "-"+option.Short)
	}

	if option.Name != "" {
		keywords = append(keywords, "--"+option.Name)
	}

	return keywords
}

func (option *Float64Option) Parse(args []*Node, options map[string]interface{}) (int, error) {
	if len(args) < 1 || args[0].Type != T_Value {
		return 0, errors.New("missing required value: " + option.usage())
	}

	v, err := strconv.ParseFloat(args[0].Value.(string), 64) // .Type == T_Value の .Value は常に string
	if err != nil {
		return 0, err
	}

	options[option.Name] = v
	return 1, nil
}

func (option *Float64Option) usage() string {
	usage := option.Usage

	if usage == "" {
		if option.Short != "" {
			usage = "-" + option.Short
		}

		if option.Name != "" {
			if usage != "" {
				usage += ","
			}

			usage += "--" + option.Name + "="

			if option.ArgUsage != "" {
				usage += option.ArgUsage
			} else {
				usage += "number"
			}
		}
	}

	return usage
}

func (option *Float64Option) Help() [2]string {
	usage := option.usage()

	description := option.Description
	if option.DefaultValue != 0 {
		description += " (default: " + strconv.FormatFloat(option.DefaultValue, 'f', -1, 64) + ")"
	}

	return [2]string{usage, description}
}
