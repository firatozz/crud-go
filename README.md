# crud-go
Restful API with Golang


<b>POST</b> /{collection}: Creates a resource of this collection and stores the data you POSTed, then returns the ID
<br>
<b>GET</b> /{collection}: Returns all resources of this collection
<br>
<b>GET</b> /{collection}/{id}: Returns the resource of this collection with that id
<br>
<b>PUT</b> /{collection}/{id}: Updates the resource of this collection with that id
<br>
<b>DELETE</b> /{collection}: Deletes all resources of this collection
<br>
<b>DELETE</b> /{collection}/{id}: Deletes the resource of this collection with that id
<br>
It also adds OPTIONS routes for easy discovery of your API:
<br>

<b>OPTIONS</b> /{collection}: Returns Allow: POST, GET, DELETE in an HTTP header
<br>
<b>OPTION</b>S /{collection}/{id]: Returns Allow: PUT, GET, DELETE in an HTTP header
<br>
Last but not least, pass the *mux.Router to your http server's ListenAndServe() as usual:
<br>
<br>
http.ListenAndServe(":8080", router)
<br>
Since the API is mounted on top of your router, you can also define additional custom handlers, like so:
<br>
router.HandleFunc("/", index)<br>
router.HandleFunc("/search", search)<br>
