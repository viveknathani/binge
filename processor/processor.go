package processor

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

// An Event is a new video which is uploaded and has to be broken down into
// chunks according to the MPEG-DASH format.
type Event struct {
	Path    string
	VideoId string
}

// Processor is the data structure that will accept incoming events
// of type Event. It will put these events into the queue and the pendingJobs list.
// maxWorkers represents the maximum number of goroutines that will fire up.
// These goroutines will wait for events to be added onto the queue and process
// them.
type Processor struct {
	pendingJobs sync.Map
	maxWorkers  int
	jobQueue    chan Event
}

func makeDirectoryIfNotExists(path string) {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

// ProcessVideo will take the path to a video and produce its chunks
// and the .mpd file. The file will have the name $VideoId.mpd.
func (e *Event) ProcessVideo() {

	makeDirectoryIfNotExists("./content/" + e.VideoId)
	cmd := exec.Command("ffmpeg", "-re", "-i", e.Path, "-map", "0", "-map", "0",
		"-c:a", "libfdk_aac", "-c:v", "libx264", "-b:v:0", "800k", "-b:v:1", "300k",
		"-s:v:1", "320x170", "-profile:v:1", "baseline", "-profile:v:0", "main",
		"-bf", "1", "-keyint_min", "120", "-g", "120", "-sc_threshold", "0",
		"-b_strategy", "0", "-ar:a:1", "22050", "-use_timeline", "1", "-use_template", "1",
		"-window_size", "5", "-adaptation_sets", "id=0,streams=v id=1,streams=a",
		"-seg_duration", "10", "-f", "dash", "./content/"+e.VideoId+"/"+e.VideoId+".mpd")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

// Push will add the event to the processor's queue.
func (d *Processor) Push(e Event) {
	d.pendingJobs.Store(e.VideoId, true)
	d.jobQueue <- e
}

// waitForEvents is a function that infinitely waits for an event to be
// added onto the queue and process it once found.
func (d *Processor) waitForEvents() {
	for e := range d.jobQueue {
		e.ProcessVideo()
		d.pendingJobs.Delete(e.VideoId)
	}
}

// Run is the function that fires up the workers for the processor.
func (d *Processor) Run() {

	wg := sync.WaitGroup{}
	for i := 0; i < d.maxWorkers; i++ {

		wg.Add(1)
		go d.waitForEvents()
	}
	wg.Wait()
}

// GetAllPendingJobs will return all the ids of all the videos that are being
// processed or will be processed soon.
func (d *Processor) GetAllPendingJobs() []string {

	result := make([]string, 0)

	d.pendingJobs.Range(func(key, value any) bool {
		result = append(result, key.(string))
		return true
	})
	return result
}

// New returns a new initialised Processor.
func New(maxWorkers int, maxJobs int) *Processor {

	return &Processor{
		maxWorkers:  maxWorkers,
		pendingJobs: *new(sync.Map),
		jobQueue:    make(chan Event, maxJobs),
	}
}
