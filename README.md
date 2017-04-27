# crud-go
Restful API with Golang


POST /{collection}: Creates a resource of this collection and stores the data you POSTed, then returns the ID
GET /{collection}: Returns all resources of this collection
GET /{collection}/{id}: Returns the resource of this collection with that id
PUT /{collection}/{id}: Updates the resource of this collection with that id
DELETE /{collection}: Deletes all resources of this collection
DELETE /{collection}/{id}: Deletes the resource of this collection with that id
It also adds OPTIONS routes for easy discovery of your API:

OPTIONS /{collection}: Returns Allow: POST, GET, DELETE in an HTTP header
OPTIONS /{collection}/{id]: Returns Allow: PUT, GET, DELETE in an HTTP header
Last but not least, pass the *mux.Router to your http server's ListenAndServe() as usual:

http.ListenAndServe(":8080", router)
Since the API is mounted on top of your router, you can also define additional custom handlers, like so:

router.HandleFunc("/", index)
router.HandleFunc("/search", search)
