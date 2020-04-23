package book

type Book struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

func NewBook(name string, price float64) Book {
	return Book{
		Name:  name,
		Price: price,
	}
}

func (b Book) getName() string {
	return b.Name
}

func (b Book) getPrice() float64 {
	return b.Price
}
