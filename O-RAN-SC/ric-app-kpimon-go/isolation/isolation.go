package isolation
import(
	"fmt"
	"example.com/kpimon/f1apPackage"
)
func IntermediateCaller(message string)(cellList [] string,plmnList [] string){
                        fmt.Println("Ponga3 = ",message)
			cellList,plmnList=f1apPackage.GetCellIdsPlmnIds(message)
			return
}
