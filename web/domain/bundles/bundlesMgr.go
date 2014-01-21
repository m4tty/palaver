package bundlesDomain

import (
	"github.com/m4tty/palaver/data/bundles"
	"github.com/m4tty/palaver/web/resources"
	"time"
)

type BundlesMgr struct {
	bundleDataMgr.BundleDataManager
}

func NewBundlesMgr(bdm bundleDataMgr.BundleDataManager) *BundlesMgr {
	return &BundlesMgr{bdm}
}

func (dm BundlesMgr) GetBundleById(id string) (bundle *resources.BundleResource, err error) {
	dBundle, err := dm.BundleDataManager.GetBundleById(id)

	if err != nil {
		return nil, err
	}
	var bundleResource *resources.BundleResource = new(resources.BundleResource)

	mapDataToResource(&dBundle, bundleResource)

	return bundleResource, nil
}

func (dm BundlesMgr) GetBundlesByUserId(id string) (bundles []*resources.BundleResource, err error) {
	dBundles, err := dm.BundleDataManager.GetBundlesByUserId(id)
	if err != nil {
		return nil, err
	}

	bundles = make([]*resources.BundleResource, len(dBundles))
	for j, bundle := range dBundles {
		var bundleResource *resources.BundleResource = new(resources.BundleResource)
		mapDataToResource(bundle, bundleResource)
		bundles[j] = bundleResource
	}
	return bundles, nil
}

func (dm BundlesMgr) GetBundles() (bundles []*resources.BundleResource, err error) {
	dBundles, err := dm.BundleDataManager.GetBundles()
	if err != nil {
		return nil, err
	}

	bundles = make([]*resources.BundleResource, len(dBundles))
	for j, bundle := range dBundles {
		var bundleResource *resources.BundleResource = new(resources.BundleResource)
		mapDataToResource(bundle, bundleResource)
		bundles[j] = bundleResource
	}
	return bundles, nil
}

func (dm BundlesMgr) SaveBundle(bundle *resources.BundleResource) (key string, err error) {
	var dBundle *bundleDataMgr.Bundle = new(bundleDataMgr.Bundle)

	mapResourceToData(bundle, dBundle)

	key, saveErr := dm.BundleDataManager.SaveBundle(dBundle)
	// if saveErr != nil {
	// 	return key, saveErr
	// }
	return key, saveErr
}

func (dm BundlesMgr) DeleteBundle(id string) (err error) {
	deleteErr := dm.BundleDataManager.DeleteBundle(id)
	return deleteErr
}

// mapper...
func mapResourceToData(bundleResource *resources.BundleResource, bundleData *bundleDataMgr.Bundle) {
	bundleData.Id = bundleResource.Id
	bundleData.OwnerId = bundleResource.OwnerId
	bundleData.Name = bundleResource.Name
	bundleData.Description = bundleResource.Description
	bundleData.IsPublic = bundleResource.IsPublic
	bundleData.CreatedDate = bundleResource.CreatedDate
	bundleData.LastModified = time.Now().UTC()
	bundleData.Stars = bundleResource.Stars
	bundleData.Likes = bundleResource.Likes
	bundleData.Dislikes = bundleResource.Dislikes
	bundleData.LikedBy = bundleResource.LikedBy
	bundleData.DislikedBy = bundleResource.DislikedBy
	bundleData.Tags = bundleResource.Tags
}

func mapDataToResource(bundleData *bundleDataMgr.Bundle, bundleResource *resources.BundleResource) {
	bundleResource.Id = bundleData.Id
	bundleResource.OwnerId = bundleData.OwnerId
	bundleResource.Name = bundleData.Name
	bundleResource.Description = bundleData.Description
	bundleResource.IsPublic = bundleData.IsPublic
	bundleResource.CreatedDate = bundleData.CreatedDate
	bundleResource.LastModified = bundleData.LastModified
	bundleResource.Stars = bundleData.Stars
	bundleResource.Likes = bundleData.Likes
	bundleResource.Dislikes = bundleData.Dislikes
	bundleResource.LikedBy = bundleData.LikedBy
	bundleResource.DislikedBy = bundleData.DislikedBy
	bundleResource.Tags = bundleData.Tags
}
