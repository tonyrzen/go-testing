package channels

type Publisher interface {
	Publish(*Thing)
}

type Object struct {
	Channel chan *Thing
}

func (o *Object) Publish(t *Thing) {
	o.Channel <- t
}

// Thing is the type of object that will be published in the channel
type Thing struct {
	name string
}

func NewThing(name string) *Thing {
	return &Thing{
		name: name,
	}
}

type someImportantObject struct {
	publisher Publisher
}

func SomeImportantLogic(sio *someImportantObject) error {
	// after some business logic, I want to publish something in my Publisher
	// create a new thing
	thing := NewThing("thing")
	sio.publisher.Publish(thing)

	return nil
}
