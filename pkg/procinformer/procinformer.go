package procinformer

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kinvolk/traceloop/pkg/podcgroup"
)

type ProcInfo struct {
	Utsns uint32

	PodUID         string
	ContainerIDSet map[string]struct{}
}

type ProcInformer struct {
	// key:   Utsns
	lookups map[uint32]ProcInfo
	mutex   *sync.Mutex

	stopChan         chan struct{}
	procInformerChan chan ProcInfo
}

func NewProcInformer(procInformerChan chan ProcInfo) (*ProcInformer, error) {
	c := &ProcInformer{
		lookups: make(map[uint32]ProcInfo),
		mutex:   &sync.Mutex{},

		stopChan:         make(chan struct{}),
		procInformerChan: procInformerChan,
	}

	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-c.stopChan:
				return
			case <-ticker.C:
				err := c.update()
				if err != nil {
					log.WithFields(log.Fields{
						"err": err,
					}).Error("Proc informer error")

					return
				}
			}
		}
	}()

	return c, nil
}

func (p *ProcInformer) update() error {
	p.mutex.Lock()
	if len(p.lookups) == 0 {
		p.mutex.Unlock()
		return nil
	}
	lookups := p.lookups
	p.lookups = make(map[uint32]ProcInfo)
	p.mutex.Unlock()

	procPath := "/proc"

	files, err := ioutil.ReadDir(procPath)
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		_, err := strconv.Atoi(fileInfo.Name())
		if err != nil {
			continue
		}

		var stat syscall.Stat_t
		if err := syscall.Stat(filepath.Join(procPath, fileInfo.Name(), "ns", "uts"), &stat); err != nil {
			// Process might have terminated
			continue
		}
		utsns := uint32(stat.Ino)

		// Defined in include/linux/proc_ns.h
		procUtsInitIno := uint32(0xEFFFFFFE)
		if utsns == procUtsInitIno {
			continue
		}

		procInfo, ok := lookups[utsns]
		if !ok {
			continue
		}

		cgroupProcFile := filepath.Join(procPath, fileInfo.Name(), "cgroup")
		podUID, containerID := podcgroup.ExtractIDFromCgroupProcFile(cgroupProcFile)
		if podUID == "" || containerID == "" {
			continue
		}

		procInfo.PodUID = podUID
		procInfo.ContainerIDSet[containerID] = struct{}{}

		lookups[utsns] = procInfo
	}

	for _, procInfo := range lookups {
		log.WithFields(log.Fields{
			"utsns":          procInfo.Utsns,
			"containerCount": len(procInfo.ContainerIDSet),
			"poduid":         procInfo.PodUID,
		}).Debug("Proc informer report")
		if procInfo.PodUID == "" {
			continue
		}
		for c := range procInfo.ContainerIDSet {
			log.WithFields(log.Fields{
				"containerID": c,
			}).Debug("Proc informer report detail")
		}
		p.procInformerChan <- procInfo
	}

	return nil
}

func (p *ProcInformer) LookupContainerID(utsns uint32) {
	log.WithFields(log.Fields{
		"utsns": utsns,
	}).Debug("Proc informer lookup for utsns")

	p.mutex.Lock()
	p.lookups[utsns] = ProcInfo{
		Utsns:          utsns,
		PodUID:         "",
		ContainerIDSet: make(map[string]struct{}),
	}
	p.mutex.Unlock()
}
