package main

import (
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

func getPartitions() []FS {
	allPartitions, _ := disk.Partitions(true)
	partitions := []FS{}

	for _, p := range allPartitions {

		if locals[p.Fstype] {

			usage, _ := disk.Usage(p.Mountpoint)

			fs := FS{
				p.Device,
				p.Mountpoint,
				strconv.Itoa(int(usage.Free) / MB),
			}

			partitions = append(partitions, fs)
		}
	}

	return partitions
}
