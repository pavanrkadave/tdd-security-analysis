package valueobjects

type Asset struct {
	name string
}

func NewAsset(name string) *Asset {
	return &Asset{
		name: name,
	}
}

func (asset *Asset) Name() string {
	return asset.name
}
