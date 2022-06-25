package erratum

func Use(opener ResourceOpener, input string) error {
	resource, err := open(opener)
	if err != nil {
		return err
	}
	return frob(resource)
}

func open(opener ResourceOpener) (Resource, error) {
	resource, err := opener()
	if _, ok := err.(TransientError); ok {
		return open(opener)
	}
	return resource, err
}

func frob(resource Resource) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if ferr, ok := r.(FrobError); ok {
				resource.Defrob(ferr.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()
	resource.Frob("hello")
	return
}
