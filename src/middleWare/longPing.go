package middleWare

import (

)

func LongPing(f func()bool){
	for {
		if f() {
			break
		}
	}
}
