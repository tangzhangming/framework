package storage

import (
	"context"
)

var disks map[string]driverface


//动态配置磁盘
func Build(name string, driver string) {
	if disks == nil {
		disks = make(map[string]driverface)
	}

	var disk driverface

	if driver == "s3" {
		disk = &s3_storage{}

	} else if driver == "file" {
		disk = &file_storage{}
	}

	disks[name] = disk
}

func WithContext(ctx context.Context) {
	return
}

func Disk(name string) driverface {
	disk, ok := disks[name]

	if ok {
		return disk
	}

	return nil
}
