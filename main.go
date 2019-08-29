package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io"
	"os"

	//	"database/sql"
	_ "github.com/alexbrainman/odbc"
	"log"
	"net"

	//	"os"
//	"encoding/json"
	"fmt"
	//"github.com/rs/cors"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"text/template"
)

//var db *sql.DB
var r = httprouter.New()
var err error
var tpl *template.Template

//func init(){
//	tpl = template.Must(template.ParseGlob("html/*.html"))
//}

type Configuration struct {
	HttpPort string
	ConnectionString string
	Appname string
	Runmode string
	sqluser string
	sqlpass string
	sqldb string
	SessionName string
}

func main() {
	//configuration := Configuration{}
	//filename := "app.json"
	//log.Println(filename)
	//pwd, _ := os.Getwd()
	//fto := pwd+"\\"+filename
	//log.Println(fto)
	//file, err := os.Open(fto)
	//if err != nil {
	//	log.Println("File Open error")
	//	os.Exit(500) //return err
	//}
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&configuration)
	//if err != nil {
	//	log.Println("json erorr")
	//	os.Exit(500)//return err
	//}

	//// ********************************************************
	//// Create the database handle, confirm driver is present
	//// *********************************************************
	//connectString := configuration.sqluser + ":" + configuration.sqlpass + configuration.sqldb
	//log.Println(connectString)
	//db, err = sql.Open("odbc", "DSN=DYLT_IMP"  )
	//if err != nil {
	//	log.Fatalf("Error on initializing database connection: %s", err.Error())
	//}
	//fmt.Println("db opened at root:****@/test")
	//db.SetMaxIdleConns(100)
	//defer db.Close()
	//// make sure connection is available
	//err = db.Ping()
	//if err != nil {
	//	log.Fatalf("Error on opening database connection: %s", err.Error())
	//}else {fmt.Println("verified db is open")}

	r.GET("/callSFTP", CallSFTPservice)
	//r.GET("/irefPeriod", irefPeriod)
	//r.GET("/bpHitRatio", bpHitRatio)
	//r.GET("/iref", iref)
	//r.GET("/irefData", irefData)
	// open for business
	fmt.Println("Router is open for business on port " + "8081")
	port := ":8083" //":"+ configuration.HttpPort
	log.Println(port)
	//handler := cors.Default().Handler(r)
	http.ListenAndServe(port, r)
}

func CallSFTPservice(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	//fmt.Println("hi")
	fmt.Fprint(w, "Hi!\n")

	x := handleProtoClient()

}
func handleProtoClient(conn net.Conn, c chan *PbTest.TestMessage) {
	fmt.Println(“Connected!”)
	defer conn.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, “Fatal error: %s”, err.Error())
		os.Exit(1)
	}
	pdata := new(PbTest.TestMessage)
	err = proto.Unmarshal(buf.Bytes(), pdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, “Fatal error: %s”, err.Error())
		os.Exit(1)
	}
	c <- pdata
}

getSystemId
setSystemId(String systemId)
RemoteSystem getSystem()
setSystem(RemoteSystem system)
getCurrentUser()
setCurrentUser(String currentUser)
getPath()
setPath(String path)
func getRanges() Range {

}
getPublicLink(RemoteSystem system, String absolutePath)
checkGroups(String[] allowedGroups, String[] deniedGroups, List<Role> userRoles)
hasRole(List<Role> roles, String roleName)

/**
 * On public storage systems, the first token of any relative path will
 * indicate an implied ownership permission for the resource.
 * @return
 */
private String resolveResourceOwnerFromPathIfPublicSystem()

/**
 * Parses the request URL and returns the relative or absolute path reference
 * by this request. The leading slash is stripped on relative paths.
 *
 * @return
 */
private String resolveRequestedSystemPathFromUrl()

/**
     * Fetches a {@link RemoteDataClient} for the requested system and authenticates
     * for use in the remainder of this session.
 */
initRemoteDataClientForSystem(RemoteSystem remoteSystem, String internalUsername)

resolveRemoteSystemFromURL(String systemId, String username)

/**
	 * This method represents the HTTP GET action. Using the file id from the URL, the
	 * input file is streamed to the user from the local cache. If the file id is invalid
	 * for any reason, a HTTP {@link org.restlet.data.Status#CLIENT_ERROR_BAD_REQUEST 400} code is sent.
 */
get()

/**
 * This method represents the HTTP POST action. Posting a file upload form
 * to this service will cache a file on behalf of the authenticated
 * user and submit it to the i/o processing queue. While this method does
 * not return a value internally, a {@link org.json.JSONObject JSONObject} representation
 * of the successfully uploaded file is written to the output stream. If the
 * upload fails due to transport issue, "no file uploaded" is returned in the body
 * of the response. If any other error occurs a HTTP {@link org.restlet.data.Status#CLIENT_ERROR_BAD_REQUEST 400} code will be
 * sent.
 * <p>
 * For the file upload form, only multipart form data is accepted. The form
 * should contain the following field:
 * <p>
 * <ul><li>format: the input file type. If not type is given, it's treated as raw data</li></ul>
 * <ul><li>url: the url of the file.</li></ul>
 * <ul><li>fileToUpload: the path on the user's system to the file (if uploading).</li></ul>
 * </p>
 */

add()

/**
 * Deletes a file or folder with provenance and alerts
 *
 */
delete()

/**
 * Accepts mkdir, copy, move, touch, index, and rename functionality.
 */
Put()

/**
 * @param jsonInput
 * @param absolutePath
 * @param logicalFile
 * @param pm
 * @return the new file item representation
 * @throws ResourceException
 * @throws PermissionException
 * @throws FileNotFoundException
 * @throws IOException
 * @throws RemoteDataException
 * @throws HibernateException
 * @throws JSONException
 * @throws RemoteDataSyntaxException
 */
doRenameOperation(JsonNode jsonInput, String absolutePath,LogicalFile logicalFile, PermissionManager pm)

/**
 * @param jsonInput
 * @param logicalFile
 * @param pm
 * @return
 * @throws PermissionException
 * @throws FileNotFoundException
 * @throws ResourceException
 * @throws IOException
 * @throws RemoteDataException
 * @throws HibernateException
 * @throws JSONException
 * @throws RemoteDataSyntaxException
 */
doMoveOperation(JsonNode jsonInput,LogicalFile logicalFile, PermissionManager pm)


/**
 * @param jsonInput
 * @param logicalFile
 * @param pm
 * @return
 * @throws PermissionException
 * @throws FileNotFoundException
 * @throws ResourceException
 * @throws IOException
 * @throws RemoteDataException
 * @throws HibernateException
 */
doMkdirOperation(JsonNode jsonInput, LogicalFile logicalFile, PermissionManager pm)

/**
 * @param jsonInput
 * @param logicalFile
 * @param pm
 * @return
 * @throws PermissionException
 * @throws FileNotFoundException
 * @throws ResourceException
 * @throws IOException
 * @throws RemoteDataException
 * @throws JSONException
 * @throws RemoteDataSyntaxException
 * @throws HibernateException
 */
doCopyOperation(JsonNode jsonInput, String absolutePath, LogicalFile logicalFile, PermissionManager pm)

/**
 * Lists the permissions of the file or directory given in the path.
 */
getPermissons()

/*
 * Sets the share permissions operations on a particular file or folder. This will overwrite the existing
 * permissions for that file or directory.
 */
setPrermissions()

/**
 * Processes the delete operation on a particular file or folder.
 * This will clear all permissions for that file or directory, but
 * the file or directory itself will remain unchanged. Only admin
 * and above can grant or change permissions.
 */
removePermissions()

/**
 * This method represents the HTTP GET action. Using the file id from the URL, the
 * input file is streamed to the user from the local cache. If the file id is invalid
 * for any reason, a HTTP {@link org.restlet.data.Status#CLIENT_ERROR_BAD_REQUEST 400} code is sent.
 */
retreiveFile()







