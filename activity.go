package Oracle

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"database/sql"
    _ "github.com/mattn/go-oci8"
)

var log = logger.GetLogger("activity-Oracle")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	
	
	dbUrl := context.GetInput("dbUrl").(string)
	log.Debugf("dbUrl");

	
	db, err := sql.Open("oci8",dbUrl)	
	log.Debugf("Connection Successful");
    if err != nil {
         log.Debugf("Connection Refused");
        return
    }
    defer db.Close()
	
	if err = db.Ping(); err != nil {
        log.Debugf("Error connecting to the database: %s\n", err);
        return
    }
    
	context.SetOutput("output","")
	return true, nil
}
