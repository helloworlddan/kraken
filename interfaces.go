package kraken

// Inspectable graph item.
type Inspectable interface {
	Inspect()
}

// YamlSerializable graph item.
type YamlSerializable interface {
	ToYaml()
}

// JSONSerializable graph item.
type JSONSerializable interface {
	ToJSON()
}

// Serializable graph item.
type Serializable interface {
	Serializable()
}

// Sizable graph item computing rough memory size.
type Sizable interface {
	Size()
}
