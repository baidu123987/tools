package compress

import (
	"context"
	"path"
	"sync"
)

type CompressJob struct {
	Ctx         context.Context
	File        string
	VersionPath string
}

func Worker(wg *sync.WaitGroup, jobs <-chan CompressJob) {
	for j := range jobs {
		switch path.Ext(j.File) {
		case ".jpg", ".jpeg":
			Jpeg(j.Ctx, path.Join(j.VersionPath, j.File))
		case ".png":
			Png(j.Ctx, path.Join(j.VersionPath, j.File))
		case ".js":
			Js(j.Ctx, path.Join(j.VersionPath, j.File))
		case ".css":
			CSS(j.Ctx, path.Join(j.VersionPath, j.File))
		}
		wg.Done()
	}
}