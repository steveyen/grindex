package grouter

import (
	"time"

	"github.com/dustin/gomemcached"
)

func WorkLoad(sourceSpec string, params Params, targetChans []chan []Request,
	statsChan chan Stats) {
	for i := 1; i < params.TargetConcurrency; i++ {
		go run(i, sourceSpec, targetChans[i % len(targetChans)], statsChan)
	}
	run(0, sourceSpec, targetChans[0], statsChan)
}

func run(clientNum int, sourceSpec string, targetChan chan []Request,
	statsChan chan Stats) {
	report_every := 1000
	ops_per_round := 100
	tot_workload_ops_nsecs := int64(0) // In nanoseconds.
	tot_workload_ops := 0
	res := make(chan *gomemcached.MCResponse)
	for {
		reqs := make([]Request, ops_per_round)
		for i := 0; i < ops_per_round; i++ {
			reqs[i] = Request{
				Bucket: "default",
				Req: &gomemcached.MCRequest{
					Opcode: gomemcached.GET,
					Key:    []byte("hello"),
				},
				Res:       res,
				ClientNum: uint32(clientNum),
			}
		}
		reqs_start := time.Now()
		targetChan <- reqs
		for i := 0; i < ops_per_round; i++ {
			<-res
		}
		reqs_end := time.Now()

		tot_workload_ops_nsecs += reqs_end.Sub(reqs_start).Nanoseconds()
		tot_workload_ops += ops_per_round
		if tot_workload_ops%report_every == 0 {
			statsChan <- Stats{
				Keys: []string{
					"tot_workload_ops",
					"tot_workload_ops_usecs",
				},
				Vals: []int64{
					int64(tot_workload_ops),
					int64(tot_workload_ops_nsecs / 1000),
				},
			}
			tot_workload_ops_nsecs = int64(0)
			tot_workload_ops = 0
		}
	}
}
