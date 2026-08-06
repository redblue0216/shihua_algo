package main

import "shihua_algo/observer"

func main() {
	//event := observer.Event{"event"}
	//observer_instance := observer.NewObserver("Barry")
	//observer_instance.Update(event)

	notifier := observer.Subject{}
	observers := []observer.Observer{
		observer.NewObserver("Barry"),
		observer.NewObserver("Jack"),
		observer.NewObserver("Shirdon"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}
	notifier.Unregister(observers[1])
	notifier.NotifyObservers(observer.Event{"Received an email!"})
}

//$ go run main.go
//ConcreteObserver 'Barry' received event 'Received an email!'
//ConcreteObserver 'Shirdon' received event 'Received an email!'