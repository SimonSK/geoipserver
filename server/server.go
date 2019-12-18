package server

import (
	"fmt"
	"net/http"

	"github.com/SimonSK/geoip2-webapi/info"
	"github.com/SimonSK/geoip2-webapi/mymaxminddb"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Log              *logrus.Logger
	DBBinaryFilepath string
	ListenPort       uint16
}

type Server struct {
	Config
	db *mymaxminddb.Reader
}

func (s *Server) info() string {
	return fmt.Sprintln(info.NameWithVersion, "on port", s.ListenPort) +
		fmt.Sprintln("Description:", info.Description) +
		fmt.Sprintln("") +
		fmt.Sprintln("Database info:") +
		fmt.Sprintln("\tType:", s.db.Metadata.DatabaseType) +
		fmt.Sprintln("\tFormat Version:", s.db.Metadata.FormatVersion()) +
		fmt.Sprintln("\tBuild Timestamp:", s.db.Metadata.BuildTimestamp()) +
		fmt.Sprintln("") +
		fmt.Sprintln("Usage:") +
		fmt.Sprintln("\tTo get all fields:", apiIPDetailsPath) +
		fmt.Sprintln("\tTo get \"location\" fields:", apiIPLocationPath) +
		fmt.Sprintln("\tTo get GPS coordinates:", apiIPGPSCoordsPath)
}

func (s *Server) Start() error {
	// Open database binary
	s.Log.Debugf("opening %s", s.DBBinaryFilepath)
	db, err := mymaxminddb.Open(s.DBBinaryFilepath)
	defer func() {
		if err := db.Close(); err != nil {
			s.Log.Error(err)
		}
	}()
	if err != nil {
		return err
	}
	s.db = db
	dbType := s.db.Metadata.DatabaseType
	dbFormatVersion := s.db.Metadata.FormatVersion()
	dbBuildTimestamp := s.db.Metadata.BuildTimestamp()
	s.Log.Debugf("[type=%s formatVersion=%s buildTimestamp=%s] database info", dbType, dbFormatVersion, dbBuildTimestamp)

	// Print server info
	fmt.Printf(s.info())

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(rootPath, s.handleDefault)
	router.HandleFunc(apiRootPath, s.handleDefault)
	router.HandleFunc(apiIPDetailsPath, s.handleIPDetails)
	router.HandleFunc(apiIPLocationPath, s.handleIPLocation)
	router.HandleFunc(apiIPGPSCoordsPath, s.handleIPGPSCoords)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.ListenPort), router); err != nil {
		return err
	}
	return nil
}
