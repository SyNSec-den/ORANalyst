package f1apPackage

/*
#include <f1ap/wrapper.h>
#cgo LDFLAGS: -lf1apwrapper -lm
#cgo CFLAGS: -I/usr/local/include/f1ap
*/
import "C"
import(
	"fmt"
	"unsafe"
)
func GetCellIdsPlmnIds(message string)(cellList [] string,plmnList [] string){
                        cString2 := C.CString(message)
                        fmt.Println("Ponga3 = ",message)
                        defer C.free(unsafe.Pointer(cString2)) // Free the allocated C string when done
                        result2:=C.decodeF1apGetCellIds(cString2)
                        cellList=make([]string, result2.size)
                        for _, v := range unsafe.Slice(result2.cellids, result2.size) {
                                cellList = append(cellList, C.GoString(v))
                        }
                        plmnList=make([]string, result2.size)
                        for _, v := range unsafe.Slice(result2.plmn, result2.size) {
                                plmnList = append(plmnList, C.GoString(v))
			}
			return
}
