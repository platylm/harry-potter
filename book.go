package book

type Book struct {
	ID    string
	Name  string
	Price int
	Stock int
}

func NewBook(name string, price int) Book {
	return Book{
		Name:  name,
		Price: price,
	}
}

func (b Book) getName() string {
	return b.Name
}

func (b Book) getPrice() int {
	return b.Price
}
