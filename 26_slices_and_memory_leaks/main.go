package main

func main() {
	consumeMessages()
}

func consumeMessages() {
	for {
		msg := receiveMessage() // Do something with msg
		storeMessageType(getMessageType(msg))
	}
}

func getMessageType(msg []byte) []byte {
	return msg[:5] // len: 5, cap: 1 000 000
}

// solution
func getMessageType2(msg []byte) []byte {
	dst := make([]byte, 5)
	copy(dst, msg)
	return dst // len: 5, cap: 5
}

func receiveMessage() []byte {
	return []byte{}
}

func storeMessageType(_ []byte) {
}
