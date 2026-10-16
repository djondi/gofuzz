package fuzz

// ParsePacket processes untrusted packet bytes.
func ParsePacket(data []byte) byte {
	if len(data) == 0 {
		return 0
	}
	length := int(data[0])
	if length >= len(data) {
		return 0
	}
	return data[length]
}