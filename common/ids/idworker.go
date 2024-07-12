package ids

import (
	"crypto/md5"
	"errors"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

const (
	workerIDBits     = 5
	datacenterIDBits = 5
	sequenceBits     = 12

	workerIDShift      = sequenceBits
	datacenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits

	sequenceMask = int64(-1) ^ (int64(-1) << sequenceBits)

	twepoch          = int64(1710460800000) // 2024-03-15 00:00:00 UTC
	maxBackwardsTime = int64(10)
	maxWorkerID      = int64(-1) ^ (int64(-1) << workerIDBits)
	maxDatacenterID  = int64(-1) ^ (int64(-1) << datacenterIDBits)
)

type IDWorker struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	datacenterID  int64
	sequence      int64
}

func NewIDWorker() (*IDWorker, error) {
	workerID, datacenterID, err := generateIDs()
	if err != nil {
		return nil, err
	}
	return NewIDWorkerWith(workerID, datacenterID)
}
func NewIDWorkerWith(workerID, datacenterID int64) (*IDWorker, error) {
	return &IDWorker{
		lastTimestamp: 0,
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
	}, nil
}

func generateIDs() (int64, int64, error) {
	var id [16]byte
	hostname, err := os.Hostname()
	if err != nil {
		return 0, 0, err
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return 0, 0, err
	}

	var mac string
	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && !isLoopbackOrVirtual(i) {
			mac = i.HardwareAddr.String()
			break
		}
	}

	id = md5.Sum([]byte(hostname + mac))

	rand.Seed(time.Now().UnixNano())
	workerID := int64(id[0]) | int64(id[1])<<8
	workerID = workerID & maxWorkerID
	if workerID > maxWorkerID {
		workerID = rand.Int63n(maxWorkerID)
	}

	datacenterID := int64(id[2]) | int64(id[3])<<8
	datacenterID = datacenterID & maxDatacenterID
	if datacenterID > maxDatacenterID {
		datacenterID = rand.Int63n(maxDatacenterID)
	}

	return workerID, datacenterID, nil
}

func isLoopbackOrVirtual(i net.Interface) bool {
	return i.Flags&(net.FlagLoopback|net.FlagPointToPoint) != 0
}

func (w *IDWorker) NextID() (int64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	timestamp := time.Now().UnixNano() / 1e6

	if timestamp < w.lastTimestamp {
		if w.lastTimestamp-timestamp < maxBackwardsTime {
			time.Sleep(time.Duration(w.lastTimestamp-timestamp) * time.Millisecond)
			timestamp = time.Now().UnixNano() / 1e6
		} else {
			return 0, errors.New("clock moved backwards beyond tolerance")
		}
	}

	if w.lastTimestamp == timestamp {
		w.sequence = (w.sequence + 1) & sequenceMask
		if w.sequence == 0 {
			timestamp = w.tilNextMillis(w.lastTimestamp)
		}
	} else {
		w.sequence = 0
	}

	w.lastTimestamp = timestamp

	id := ((timestamp - twepoch) << timestampLeftShift) |
		(w.datacenterID << datacenterIDShift) |
		(w.workerID << workerIDShift) |
		w.sequence

	return id, nil
}

func (w *IDWorker) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / 1e6
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixNano() / 1e6
	}
	return timestamp
}
