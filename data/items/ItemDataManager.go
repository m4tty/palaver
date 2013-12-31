package data

type ItemDataManager interface {
	GetItems() (results []*Item, err error)
	GetItemById(id string) (result Item, err error)
	SaveItem(bundle *Item) (key string, err error)
	DeleteItem(id string) (err error)
}
