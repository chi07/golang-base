package debuger

import "log"

func Dump(data interface{})  {
	log.Println(data)
}

func CheckError(err error)  {
	if err != nil {
		log.Println(err)
	}
}

