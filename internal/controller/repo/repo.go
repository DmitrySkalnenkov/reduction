package repo

import (
	"github.com/DmitrySkalnenkov/reduction/internal/controller/repo/filerepo"
	"github.com/DmitrySkalnenkov/reduction/internal/controller/repo/memrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
)

func URLStorageInit(filePath string) (ur entity.Keeper) {
	if filePath != "" {
		ur := new(filerepo.FileRepo)
		ur.InitRepo(filePath)
		return ur
	} else {
		ur := new(memrepo.MemRepo)
		ur.InitRepo(filePath)
		return ur
	}
}
