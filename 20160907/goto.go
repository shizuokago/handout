package main

func main() {
L:
	for {
		select {
		case t1 := <-time.After(time.Second):
			fmt.Println("hello", t1)
			if t1.Second()%4 == 0 {
				break L
			}
		}
	}
	fmt.Println("ending")
}
