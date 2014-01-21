package bundleDataMgr

type BundleDataManager interface {
	GetBundles() (results []*Bundle, err error)
	GetBundleById(id string) (result Bundle, err error)
	GetBundlesByUserId(id string) (results []*Bundle, err error)
	SaveBundle(bundle *Bundle) (key string, err error)
	DeleteBundle(id string) (err error)
}
