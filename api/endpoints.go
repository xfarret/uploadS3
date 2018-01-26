package api

import (
	"net/http"
	"encoding/json"
	"uploadS3/server/beanstalk"
	//"bytes"
	//"encoding/gob"
	"uploadS3/models"
)

func AddSoftware(w http.ResponseWriter, r *http.Request) {
	/* get software from body */
	decoder := json.NewDecoder(r.Body)
	var soft *models.Software
	err := decoder.Decode(&soft)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	/* convert soft to byte array and write it on beanstalk queue */
	bSoft, _ := json.Marshal(soft)
	id, err := beanstalk.Put(bSoft)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	sendJsonResponse(w, id)
}

func sendJsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)


	/* convert interface to byte array and send response */
	//var buf bytes.Buffer
	//enc := gob.NewEncoder(&buf)
	//
	//err := enc.Encode(v)
	//if err != nil {
	//	w.Write([]byte("Error converting data"))
	//	return
	//}
	//
	//w.Write(buf.Bytes())
}
