package ldm

import (
	// "io"
	// "io/ioutil"
	// "os"
	// "path"
	// "sort"
	// "strconv"
	// "strings"
	"sync"

	pb "github.com/xcheng85/go-ldm/server/api/v1"
)

// // dependencies injection for LDMServer
// type LDMManager interface {
// 	Write(*pb.Tile) (uint64, error)
// 	Read(offset uint64) (*pb.Tile, error)
// }


type LDMManager struct {
	mu sync.RWMutex
	// Dir    string
	// activeSegment *segment
	// segments      []*segment
}

func NewLDMManager() (*LDMManager, error) {
	l := &LDMManager{
	}
	return l, nil
}

func (l *LDMManager) Write(Tile *pb.Tile) (uint64, error) {
	// wr lock
	l.mu.Lock()
	defer l.mu.Unlock()
	return 0, nil
	// off, err := l.activeSegment.Append(record)
	// if err != nil {
	// 	return 0, err
	// }
	// if l.activeSegment.IsMaxed() {
	// 	err = l.newSegment(off + 1)
	// }
	// return off, err
}

func (l *LDMManager) Read(off uint64) (*pb.Tile, error) {
	// read lock
	l.mu.RLock()
	defer l.mu.RUnlock()
	return nil, nil
	// find the segment
	// var s *segment
	// for _, segment := range l.segments {
	// 	if segment.baseOffset <= off && off < segment.nextOffset {
	// 		s = segment
	// 		break
	// 	}
	// }
	// // START: after
	// if s == nil || s.nextOffset <= off {
	// 	return nil, api.ErrOffsetOutOfRange{Offset: off}
	// }
	// // END: after
	// return s.Read(off)
}
