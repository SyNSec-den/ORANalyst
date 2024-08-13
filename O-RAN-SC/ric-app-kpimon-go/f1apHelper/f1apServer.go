package main

/*
#include <f1ap/wrapper.h>
#cgo LDFLAGS: -lf1apwrapper -lm
#cgo CFLAGS: -I/usr/local/include/f1ap
*/
import "C"
import (
   "encoding/json"
   "net/http"
   "encoding/base64"
   "log"
   "unsafe"
   "fmt"
)
type RanFunctions struct {
        RanFunctionId         int
        RanFunctionDefinition string
        RanFunctionRevision   int
        RanFunctionOid        string
}

type GlobalNbId struct {
        PlmnId string
        NbId   string
}
type E2nodeComponentInterfaceTypeE1 struct {
}
type E2nodeComponentInterfaceTypeXn struct {
}
type E2nodeComponentInterfaceTypeF1 struct {
}

type NodeConfigs struct {
        E2nodeComponentInterfaceTypeE1 E2nodeComponentInterfaceTypeE1 `json:e2nodeComponentInterfaceTypeE1",omitempty"`
        E2nodeComponentInterfaceTypeXn E2nodeComponentInterfaceTypeXn `json:e2nodeComponentInterfaceTypeXn",omitempty"`
        E2nodeComponentInterfaceTypeF1 E2nodeComponentInterfaceTypeF1 `json:e2nodeComponentInterfaceTypeF1",omitempty"`
        E2nodeComponentInterfaceType   string
        E2nodeComponentRequestPart     string
        E2nodeComponentResponsePart    string `json:e2nodeComponentResponsePart",omitempty"`
}
type Gnb struct {
        RanFunctions []RanFunctions
        GnbType      string
        NodeConfigs  []NodeConfigs
}

/*
        type E2nodeComponentInterfaceType struct {
                E2nodeComponentInterfaceType string
        }

        type E2nodeComponentRequestPart struct {
                E2nodeComponentRequestPart string
        }
*/

type E2mgrResponse struct {
        RanName                      string
        ConnectionStatus             string
        GlobalNbId                   GlobalNbId
        NodeType                     string
        Gnb                          Gnb
        AssociatedE2tInstanceAddress string `json:associatedE2tInstanceAddress",omitempty"`
        SetupFromNetwork             bool
        StatusUpdateTimeStamp        string
}

func main() {
   http.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
	
       gnbid := r.URL.Query().Get("gnbid")
 	link:="http://service-ricplt-e2mgr-http.ricplt.svc.cluster.local:3800/v1/nodeb/"
        link=link+gnbid
         tmpr,err := http.Get(link)
                        if err != nil {
                                log.Fatalln(err)
                                panic("Unable to get cellid")
                        }
                        defer tmpr.Body.Close()
                        var resp E2mgrResponse

                        err=json.NewDecoder(tmpr.Body).Decode(&resp)
                        if err != nil {
                                log.Fatalln(err)
                                panic("Unable to get cellid")

                        }
			counter := 0
                        for i := 0; i < len(resp.Gnb.NodeConfigs); i++ {
                                if resp.Gnb.NodeConfigs[i].E2nodeComponentInterfaceType == "f1" {
                                        counter = i
                                        break
                                }
                        }

                        tm := resp.Gnb.NodeConfigs[counter].E2nodeComponentRequestPart
                        base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(tm)))
                        nl, _ := base64.StdEncoding.Decode(base64Text, []byte(tm))
                        message := string(base64Text[:nl])

       cString2 := C.CString(message)
       fmt.Println("server = ",message)
       defer C.free(unsafe.Pointer(cString2)) // Free the allocated C string when done
       result2:=C.decodeF1apGetCellIds(cString2)
       cellList:=make([]string, result2.size)
       for _, v := range unsafe.Slice(result2.cellids, result2.size) {
                cellList = append(cellList, C.GoString(v))
	}
	plmnList:=make([]string, result2.size)
       for _, v := range unsafe.Slice(result2.plmn, result2.size) {
                plmnList = append(plmnList, C.GoString(v))
       }

       // Simulate data
       data := map[string]interface{}{
           "CellIds": cellList,
           "PlmnIds": plmnList,
       }
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(data)
   })
   http.ListenAndServe(":8090", nil)
}
