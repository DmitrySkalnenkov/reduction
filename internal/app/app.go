package app

import (
	"github.com/DmitrySkalnenkov/reduction/config"
	"github.com/DmitrySkalnenkov/reduction/internal/controller/filerepo"
	router "github.com/DmitrySkalnenkov/reduction/internal/controller/http"
	"github.com/DmitrySkalnenkov/reduction/internal/controller/memrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/controller/userrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
	"log"
	"net/http"
)

// Run creates objects via constructors.

func Run(cfg *config.ServerParameters) {
	//Global variables
	config.SetGlobalVariables(cfg)
	entity.HostSocketAddrStr = cfg.HostSocketAddrStr
	entity.BaseURLStr = cfg.BaseURLStr
	entity.RepoFilePathStr = cfg.RepoFilePathStr
	//Repositories
	if entity.RepoFilePathStr == "" {
		entity.URLStorage = new(memrepo.MemRepo)
		entity.URLStorage.InitRepo("")
	} else {
		entity.URLStorage = new(filerepo.FileRepo)
		entity.URLStorage.InitRepo(entity.RepoFilePathStr)
	}
	entity.UserKeyStorage = new(userrepo.UserRepo) //For user keys
	entity.UserKeyStorage.InitRepo("")

	//Router
	r := router.NewRouter()
	//Use case
	s := &http.Server{
		Addr: entity.HostSocketAddrStr,
	}
	s.Handler = r
	log.Fatal(s.ListenAndServe())
}
