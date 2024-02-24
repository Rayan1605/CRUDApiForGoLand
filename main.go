package main

// like a box where you can store different kinds of things together,
// like a person's name and age. It helps you keep related information organized in one place
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // The * says it might be null
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
