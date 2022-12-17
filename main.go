package main

import(
	"database/sql"
	"log"
	"github.com/techschool/simplebank/util"
	
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"

	_"github.com/lib/pq"
)  

func main(){
	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load the config:",err)
	}
	conn,err := sql.Open(config.DBDriver,config.DBSource)
	if err != nil{
		log.Fatal("cannot connect to the db:",err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if(err != nil){
		log.Fatal("Can not create server")
	}


	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("Cannot start server: ",err)
	}
}