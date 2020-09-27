package diskicon

var recentBytesRead = uint64(0)

func BeginRead(nbytes uint64) {
	recentBytesRead += nbytes
}
