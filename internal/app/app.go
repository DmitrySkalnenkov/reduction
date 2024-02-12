package app

import (
	"github.com/DmitrySkalnenkov/reduction/config"
	//"github.com/DmitrySkalnenkov/reduction/internal/controllers"
	router "github.com/DmitrySkalnenkov/reduction/internal/controllers/http"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	//"github.com/DmitrySkalnenkov/reduction/internal/repos/filerepo"
	//"github.com/DmitrySkalnenkov/reduction/internal/repos/memrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/repos/userrepo"
	"log"
	"net/http"
)

// Run creates objects via constructors.

func Run(cfg *config.ServerParameters) {
	//Global variables
	config.SetGlobalVariables(cfg)
	models.HostSocketAddrStr = cfg.HostSocketAddrStr
	models.BaseURLStr = cfg.BaseURLStr
	models.RepoFilePathStr = cfg.RepoFilePathStr
	//Repositories
	/*if models.RepoFilePathStr == "" {
		models.URLStorage = new(memrepo.MemRepo)
		models.URLStorage.InitRepo("")
	} else {
		models.URLStorage = new(filerepo.FileRepo)
		models.URLStorage.InitRepo(models.RepoFilePathStr)
	}*/
	models.UserKeyStorage = new(userrepo.UserRepo) //For user keys
	models.UserKeyStorage.InitRepo("")

	//Router
	r := router.NewRouter()
	//Use case
	s := &http.Server{
		Addr: models.HostSocketAddrStr,
	}
	s.Handler = r
	log.Fatal(s.ListenAndServe())
}
