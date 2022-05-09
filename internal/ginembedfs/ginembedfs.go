package ginembedfs

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

type embedFileSystem struct {
	http.FileSystem
}

// Exists returns a bool if the path exists inside the filesystem.
//
// Warning: Signature must not be changed, it's necessary for implementing the gin static's "ServeFileSystem" interface.
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func embedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

// NewHandler returns a new gin.HandlerFunc capable of serving a dir inside embed.FS
//
// This solves triggering Gin's wildcard bug:
//
// "conflicts with existing wildcard '/*filepath' in existing prefix '/*filepath'"
//
// More info:
//
// https://stackoverflow.com/a/68613803 https://github.com/gin-gonic/gin/issues/360
// https://github.com/gin-contrib/static/issues/19
func NewHandler(route string, fsPath string, fs embed.FS) gin.HandlerFunc {
	return static.Serve(route, embedFolder(fs, fsPath))
}
