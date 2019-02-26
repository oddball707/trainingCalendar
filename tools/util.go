package util

func checkError(message string, err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf(err.Error())
	}
}
