package domain

import (
	"github.com/m4tty/palaver/data/bundles"
	"github.com/m4tty/palaver/web/resources"
)

type BundlesMgr struct {
	BundleDataManager *data.BundleDataManager
}

func (dm BundlesMgr) GetBundleById(id string) (bundle *resources.Bundle, err error) {
	return
}

func (dm BundlesMgr) GetBundles() (bundles *[]resources.Bundle, err error) {
	return
}

func (dm BundlesMgr) SaveBundle(bundle *resources.Bundle) (key string, err error) {
	return
}

func (dm BundlesMgr) DeleteBundle(id string) (err error) {
	return
}
