package kraken

// Inspectable graph item.
type Inspectable interface {
	Inspect()
}

// YamlSerializable graph item.
type YamlSerializable interface {
	ToYaml()
}

// Sizable graph item computing rough memory size.
type Sizable interface {
	Size()
}
