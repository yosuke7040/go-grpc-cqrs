package products

type ProductAdapter interface {
	Convert(source *Product) any
	ReBuild(source any) (dest *Product, err error)
}
