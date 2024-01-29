package config

import (
	"flag"
	"os"
)

const (
	DefaultHostTCPPort    string = ":8080"
	DefaultHostIPAddr     string = "localhost"
	DefaultSocketAddr            = DefaultHostIPAddr + DefaultHostTCPPort
	DefaultHostURL               = "http://" + DefaultHostIPAddr + DefaultHostTCPPort
	DefaultRepoFilePath   string = "/tmp/temp_repository_file.json"
	DefaultShortURLLength        = 15
)

type ServerParameters struct {
	HostSocketAddrStr string
	BaseURLStr        string
	RepoFilePathStr   string
	ShortURLLength    int
	envSocketAddr     string
	isEnvSocketAddr   bool
	envBaseURL        string
	isEnvBaseURL      bool
	envRepoFilePath   string
	isRepoFilePath    bool
	flagSocketAddr    string
	flagBaseURL       string
	flagRepoFilePath  string
}

// (i5) Добавьте возможность конфигурировать сервис с помощью переменных окружения:

// GetEnvs get environment variables SERVER_ADDRESS and BASE_URL check values and set them to global variables HostSocketAddrStr and BaseURLStr
func (ps *ServerParameters) GetEnvs() {
	ps.envSocketAddr, ps.isEnvSocketAddr = os.LookupEnv("SERVER_ADDRESS")     //(i5) адрес запуска HTTP-сервера с помощью переменной SERVER_ADDRESS (видимо адрес и порт)
	ps.envBaseURL, ps.isEnvBaseURL = os.LookupEnv("BASE_URL")                 //(i5) базовый адрес результирующего сокращённого URL с помощью переменной BASE_URL.
	ps.envRepoFilePath, ps.isRepoFilePath = os.LookupEnv("FILE_STORAGE_PATH") // (i6) Путь до файла должен передаваться в переменной окружения FILE_STORAGE_PATH
}

// (i7) Поддержите конфигурирование сервиса с помощью флагов командной строки наравне с уже имеющимися переменными окружения:

// GetFlags gets flags value -a, -b, -f
func (ps *ServerParameters) GetFlags() {
	flag.StringVar(&(ps.flagSocketAddr), "a", DefaultSocketAddr, "Value for -a sets socket address for the server. It should be in 'ip:port' format, for example: 127.0.0.1:8080")                         // (i7)	- флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
	flag.StringVar(&(ps.flagBaseURL), "b", DefaultHostURL, "Value for -b sets base URL for shortened URL. It should be in 'http://ip:port/' format, for example: http://127.0.0.1:8080/")                  // (i7) - флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
	flag.StringVar(&(ps.flagRepoFilePath), "f", DefaultRepoFilePath, "Value for -f sets path to the file for data storing. It should be in '/absolute/path/' format, for example: http://127.0.0.1:8080/") // (i7) - флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH)
	flag.Parse()
}

// (i7) Во всех случаях должны быть:
// (i7) значения по умолчанию,
// (i7) приоритет значений, полученных через ENV, перед значениями, задаваемыми посредством флагов.

// CheckParamPriority checks priorities of env and flag parameters. Sets values of global variables HostSocketAddrStr, BaseURLStr, RepoFilePathStr   Sets default value if no parameters exist.
func (ps *ServerParameters) CheckParamPriority() {
	//sets HostSocketAddrStr global parameter
	if ps.isEnvSocketAddr && ps.envSocketAddr != "" {
		ps.HostSocketAddrStr = ps.envSocketAddr
	} else if ps.flagSocketAddr != "" {
		ps.HostSocketAddrStr = ps.flagSocketAddr
	} else {
		ps.HostSocketAddrStr = DefaultSocketAddr
	}
	//sets BaseURLStr global parameter
	if ps.isEnvBaseURL && ps.envBaseURL != "" {
		ps.BaseURLStr = ps.envBaseURL
	} else if ps.flagBaseURL != "" {
		ps.BaseURLStr = ps.flagBaseURL
	} else {
		ps.BaseURLStr = DefaultHostURL
	}
	//sets RepoFilePathStr global parameter
	if ps.isRepoFilePath && ps.envRepoFilePath != "" {
		ps.RepoFilePathStr = ps.envRepoFilePath
	} else if ps.flagRepoFilePath != "" {
		ps.RepoFilePathStr = ps.flagRepoFilePath
	} else {
		ps.RepoFilePathStr = DefaultRepoFilePath
	}
	ps.ShortURLLength = DefaultShortURLLength
}

// NewConfig returns app config.
func NewConfig() (*ServerParameters, error) {
	cfg := &ServerParameters{}

	//var sp app.ServerParameters
	cfg.GetEnvs()
	cfg.GetFlags()
	cfg.CheckParamPriority()

	return cfg, nil
}
