package main

type Customer struct {
	Name     string
	Location string
}

func reportLines(aCustomer Customer) map[string]string {
	lines := make(map[string]string)
	gatherCustomerData(lines, aCustomer)
	return lines
}

func gatherCustomerData(out map[string]string, aCustomer Customer) {
	out["name"] = aCustomer.Name
	out["location"] = aCustomer.Location
}
