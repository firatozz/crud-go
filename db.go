package crud

type DbStatusResponse interface {
	error
	StatusCode()
}


type Db intarface { //Db describe the methods

//first argument is for resource
//second argument is for resource ID

//create a resource and returns ID
Create(string, intarface{}) (string, DbStatusResponse)

//retrieves a resource
Get(string, string) (intarface{}) (string, DbStatusResponse)

//all resources
GetAll(string) ([]interface{}, DbStatusResponse)

//update resource
Update(string, string, interface{}, DbStatusResponse)

//delete resource
Delete(string, string) DbStatusResponse

//delete all
DeleteAll(string) DbStatusResponse
}
