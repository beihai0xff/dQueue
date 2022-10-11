package log

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	logChanSize = 10240

	cacheTimeout = 2 * time.Second
)

type (
	// logFile add features upon the regular file:
	// 1. Reopen the file after receiving SIGHUP, for log rotate.
	// 2. Reduce execution time of callers by asynchronous log(return after only memory copy).
	// 3. Batch write logs by cache them with timeout.
	logFile struct {
		filename string
		file     *os.File

		logChan       chan []byte
		syncEventChan chan *syncEvent

		cacheCount    uint32
		maxCacheCount uint32
		cache         *bytes.Buffer
	}

	syncEvent struct {
		resultChan chan error
	}
)

// newLogFile can not open /dev/stderr, it will cause deadlock.
func newLogFile(filename string, maxCacheCount uint32) (*logFile, error) {
	lf := &logFile{
		filename:      filename,
		logChan:       make(chan []byte, logChanSize),
		syncEventChan: make(chan *syncEvent),
		maxCacheCount: maxCacheCount,
		cache:         bytes.NewBuffer(nil),
	}

	err := lf.openFile()
	if err != nil {
		return nil, err
	}

	go lf.run()

	return lf, nil
}

func (lf *logFile) openFile() error {
	file, err := os.OpenFile(lf.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o640)
	if err != nil {
		return err
	}

	lf.file = file
	return nil
}

func (lf *logFile) closeFile() {
	err := lf.file.Close()
	if err != nil {
		defaultLogger.Errorf("close %s failed: %v", lf.filename, err)
	}
}

func (lf *logFile) reopenFile() {
	lf.closeFile()
	err := lf.openFile()
	if err != nil {
		defaultLogger.Errorf("open %s failed: %v", lf.filename, err)
		return
	}
}

func (lf *logFile) run() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP)

	for {
		select {
		case <-signalChan:
			lf.reopenFile()
		case p := <-lf.logChan:
			lf.writeLog(p)
		case syncEvent := <-lf.syncEventChan:
			err := lf.flush()
			if err != nil {
				syncEvent.resultChan <- err
			} else {
				syncEvent.resultChan <- lf.file.Sync()
			}
		case <-time.After(cacheTimeout):
			lf.flush()
		}
	}
}

// Write writes log asynchronously, it always returns successful result.
func (lf *logFile) Write(p []byte) (int, error) {
	// NOTE: The memory of p may be corrupted after Write returned
	// So it's necessary to do copy.
	buff := make([]byte, len(p))
	copy(buff, p)
	lf.logChan <- buff
	return len(p), nil
}

// Sync flushes all cache to file with os-level flush.
func (lf *logFile) Sync() error {
	event := &syncEvent{
		resultChan: make(chan error, 1),
	}
	lf.syncEventChan <- event

	return <-event.resultChan
}

func (lf *logFile) writeLog(p []byte) {
	// No need to copy twice for non-cacheable log file.
	if lf.maxCacheCount == 0 {
		_, err := lf.file.Write(p)
		if err != nil {
			defaultLogger.Errorf("%v", err)
		}
		return
	}

	n, err := lf.cache.Write(p)
	if err != nil || len(p) != n {
		defaultLogger.Errorf("write %s to cache failed: %v", p, err)
	}
	lf.cacheCount++

	if lf.cacheCount < lf.maxCacheCount {
		return
	}

	err = lf.flush()
	if err != nil {
		defaultLogger.Errorf("%v", err)
	}
}

// flush flushes all cache to file without os-level flush.
func (lf *logFile) flush() error {
	if lf.cache.Len() == 0 {
		return nil
	}

	// NOTE: Discard all buffer regardless of it succeed or failed.
	defer func() {
		lf.cache.Reset()
		lf.cacheCount = 0
	}()

	n, err := lf.file.Write(lf.cache.Bytes())
	if err != nil || n != lf.cache.Len() {
		return fmt.Errorf("write buffer to %s failed: %d, %v", lf.filename, n, err)
	}

	return nil
}
