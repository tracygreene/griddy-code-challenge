package main

import (
	"log"
	"net"
        "database/sql"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	
	ptypes "github.com/golang/protobuf/ptypes"	
	pb "grpc-web-service/data"
	"shared/database"
)

const (
	port = ":50051"
)


// server is used to implement data.DataServer.
type server struct {

}

// CreateData creates a new Data
func (s *server) CreateData(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	//connect to database
	req := in.Data
	DB, err := database.Connect()
	        
	sqlStatement := `
	INSERT INTO t1 (data)
	VALUES ($1)
	ON CONFLICT (data)
	DO NOTHING
	RETURNING key`
	        
	var key int32
	        
	//insert row into t1 table
	err = DB.QueryRow(sqlStatement, req).Scan(&key)
	if err != nil && err != sql.ErrNoRows {
	   panic(err)
	}
	
	//if a row was added to T1 then insert into T2
	if err != sql.ErrNoRows {
	
		//create t2 string made up of the data posted & the new key from the inserted t1 record
		t2Str := fmt.Sprintf("%s %d", req, key)

		t2Key := 0

		//insert into t2 
		sqlStatement = `
		INSERT INTO t2 (t1key, data)
		VALUES ($1, $2)
		RETURNING key`

		//insert into t2 
		err = DB.QueryRow(sqlStatement, key, t2Str).Scan(&t2Key)
		if err != nil {
		   panic(err)
		} 
		
		fmt.Println("New key is:", key)
	}
	        
	    
	//close database connection
    	DB.Close()
	
    
	return &pb.DataResponse{Data: req, Success: true, Key: key}, nil
}

// GetData returns all datas by given filter
func (s *server) GetData(filter *pb.DataFilter, stream pb.Data_GetDataServer) error {
	
	//connect to database
	DB, err := database.Connect()
	    	
	sqlStatement := `
	SELECT * 
	FROM t1
	INNER JOIN t2
	ON t1.key = t2.t1key
	ORDER BY t1.key ASC`
		    
	//get rows where t1.key = t2.t1key
	rows, err := DB.Query(sqlStatement)
	if err != nil  {
	   return err
	}   
	defer rows.Close()
	
	var records = make([]*pb.Record, 0)
	    	
	for rows.Next() {
		record := &pb.Record{}
			  var T1_Key_ int32
			  var T1_Data_ string
			  var T2_Key_  int32
			  var T2_T1Key_  int32
			  var T2_Data_ string
	  		  var T2_CreatedOn_ time.Time
	  
		err = rows.Scan(
			&T1_Key_, 
			&T1_Data_,
			&T2_Key_,
			&T2_T1Key_,
			&T2_Data_,
			&T2_CreatedOn_,
		)
		if err != nil && err != sql.ErrNoRows {
		   return err
		}
		
		record.T1_Key = T1_Key_
		record.T1_Data = T1_Data_
		record.T2_Key = T2_Key_
		record.T2_T1Key = T2_T1Key_
		record.T2_Data = T2_Data_
		
		ts, err := ptypes.TimestampProto(T2_CreatedOn_)
		if err != nil {
		   panic("ptypes: T2_CreatedOn_ out of Timestamp range")
		}
		record.T2_CreatedOn = ts
			
		records = append(records, record)
		        
	}
		        
    	DB.Close()
    	
    	for _, onedata := range records {
	    if err := stream.Send(onedata); err != nil {
	        return err
	    }
	}
	return nil
}

// DeleteData deletes a data given an ID
func (s *server) DeleteData(ctx context.Context, in *pb.DeleteDataRequest) (*pb.DeleteDataResponse, error) {
     req := in.Data
     
     DB, err := database.Connect()
	          
     sqlStatement := `
     DELETE FROM t1
     WHERE data = $1`
	          
    //delete record from t1 (the associated t2 record is auto-deleted) 
    res, err := DB.Exec(sqlStatement, req)
    if err != nil {
       	 panic(err)
    }

    count, err := res.RowsAffected()
    if err != nil {
       panic(err)
    }
    
    fmt.Printf("%d row(s) were affected", count)
	          	
    //close database connection
    defer DB.Close()
    
    return &pb.DeleteDataResponse{Data: req, Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server 
	s := grpc.NewServer()
	pb.RegisterDataServer(s, &server{})
	s.Serve(lis)
}