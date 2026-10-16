package fuzz

// ParsePacket processes untrusted packet bytes.
func ParsePacket(data []byte) byte {
	length := int(data[0])
	return data[length]
}