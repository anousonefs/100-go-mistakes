package main

func main() {
	handleOperations("")
}

func handleOperations(id string) {
	operations := getOperations2(id)
	if operations != nil && len(operations) != 0 {
		println("handle operations")
	}
}

func getOperations(id string) []float32 {
	operations := make([]float32, 0)
	if id == "" {
		return operations
	}
	// Add elements to operations
	return operations
}

func getOperations2(id string) []float32 {
	operations := make([]float32, 0)
	if id == "" {
		return nil
	}
	// Add elements to operations
	operations = append(operations, 4.0)
	return operations
}
