package api

func RaiseInternalIfError(err error) {
	if err != nil {
		panic(err)
	}
}
