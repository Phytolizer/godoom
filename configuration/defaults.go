package configuration

import "godoom/errors"

func configVariableGeneric(name string, kind DefaultKind) Default {
	return Default{
		name,
		nil,
		kind,
		0,
		0,
		false,
	}
}

func configVariableKey(name string) Default {
	return configVariableGeneric(name, DefaultKey)
}

func configVariableInt(name string) Default {
	return configVariableGeneric(name, DefaultInt)
}

func configVariableIntHex(name string) Default {
	return configVariableGeneric(name, DefaultIntHex)
}

func configVariableFloat(name string) Default {
	return configVariableGeneric(name, DefaultFloat)
}

func configVariableString(name string) Default {
	return configVariableGeneric(name, DefaultString)
}

func configVariableBool(name string) Default {
	return configVariableGeneric(name, DefaultBool)
}

func searchCollection(collection *DefaultCollection, name string) *Default {
	for _, d := range collection.Defaults {
		if d.Name == name {
			return &d
		}
	}
	return nil
}

func GetDefaultForName(name string) *Default {
	result := searchCollection(&DoomDefaults, name)
	if result == nil {
		result = searchCollection(&ExtraDefaults, name)
	}
	if result == nil {
		errors.Error("Unknown configuration variable: '", name, "'")
	}
	return result
}
