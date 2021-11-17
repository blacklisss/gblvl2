package main_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

var (
	lock  sync.Mutex
	lock2 sync.RWMutex
)

var arr = make(map[int]struct{})

func BenchmarkTest1(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 1000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 1000 {
							lock.Lock()
							_, _ = arr[j]
							lock.Unlock()
						} else {
							lock.Lock()
							arr[j] = struct{}{}
							lock.Unlock()
						}

						lock.Lock()
						if len(ch) != 1000 {
							ch <- struct{}{}
						}
						lock.Unlock()
					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func BenchmarkTest2(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 1000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 1000 {
							lock2.RLock()
							_, _ = arr[j]
							lock2.RUnlock()
						} else {
							lock2.Lock()
							arr[j] = struct{}{}
							lock2.Unlock()
						}

						lock2.Lock()
						if len(ch) != 1000 {
							ch <- struct{}{}
						}
						lock2.Unlock()

					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func BenchmarkTest3(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 5000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 5000 {
							lock.Lock()
							_, _ = arr[j]
							lock.Unlock()
						} else {
							lock.Lock()
							arr[j] = struct{}{}
							lock.Unlock()
						}

						lock.Lock()
						if len(ch) != 5000 {
							ch <- struct{}{}
						}
						lock.Unlock()

					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func BenchmarkTest4(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 5000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 5000 {
							lock2.RLock()
							_, _ = arr[j]
							lock2.RUnlock()
						} else {
							lock2.Lock()
							arr[j] = struct{}{}
							lock2.Unlock()
						}

						lock2.Lock()
						if len(ch) != 5000 {
							ch <- struct{}{}
						}
						lock2.Unlock()

					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func BenchmarkTest5(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 9000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 9000 {
							lock.Lock()
							_, _ = arr[j]
							lock.Unlock()
						} else {
							lock.Lock()
							arr[j] = struct{}{}
							lock.Unlock()
						}

						lock.Lock()
						if len(ch) != 9000 {
							ch <- struct{}{}
						}
						lock.Unlock()

					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func BenchmarkTest6(b *testing.B) {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 9000)

	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for j := 0; j < 10000; j++ {
					wg.Add(1)

					go func(j int) {
						defer wg.Done()

						if len(ch) != 9000 {
							lock2.RLock()
							_, _ = arr[j]
							lock2.RUnlock()
						} else {
							lock2.Lock()
							arr[j] = struct{}{}
							lock2.Unlock()
						}

						lock2.Lock()
						if len(ch) != 9000 {
							ch <- struct{}{}
						}
						lock2.Unlock()

					}(j)
				}
				wg.Wait()
			}
		})
	})
}

func init() {
	for i := 0; i < 10000; i++ {
		arr[i] = struct{}{}
	}
}
