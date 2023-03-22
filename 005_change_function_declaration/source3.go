package main

type Address struct {
	State string
}

type Customer struct {
	Address Address
}

func InNewEngland(aCustomer Customer) bool {
	for _, s := range []string{"MA", "CT", "ME", "VT", "NH", "RI"} {
		if s == aCustomer.Address.State {
			return true
		}
	}
	return false
}

func findNewEnglanders(someCustomers []Customer) (result []Customer) {
	for _, c := range someCustomers {
		if InNewEngland(c) {
			result = append(result, c)
		}
	}
	return
}
