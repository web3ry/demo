# demo
Docker: webery/demo

	router.HandleFunc("/", homeLink)
	
	router.HandleFunc("/demo", createdemo).Methods("POST")
	
	router.HandleFunc("/demos", getAlldemos).Methods("GET")
	
	router.HandleFunc("/demos/{id}", getOnedemo).Methods("GET")
	
	router.HandleFunc("/demos/{id}", updatedemo).Methods("PATCH")
	
	router.HandleFunc("/demos/{id}", deletedemo).Methods("DELETE")
