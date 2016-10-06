/* Embedded types in GoLang: This is the is-a relation ship unlike the 
   struct is the has-a relationship
*/

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is: ", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Person.Talk() //calling the talk function

	/* Since Person is a member of Android, that is Android is a person */
	a.Talk() //this is the is-a relationship through the transitive property

	// If a person and can and andriod is-a person, then android can talk :)
}

