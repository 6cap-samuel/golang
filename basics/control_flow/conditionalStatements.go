package main

func main() {
	if "1" == "1" {
		println("same")
	}

	if "1" != "2" {
		println("not same")
	}

	if "1" == "2" {
		println("same")
	} else {
		println("not same")
	}

	if "1" == "2" {
		println("same")
	} else if "1" == "1" {
		println("same 2")
	} else {
		println("not same")
	}

}
