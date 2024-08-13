package control
import(
	"fmt"
	//"example.com/kpimon/f1apPackage"
	"example.com/kpimon/isolation"
)
type F1ap struct {
}
func (c *F1ap)F1apGetCellIds(message string)(cellList [] string,plmnList [] string){
                        fmt.Println("Ponga2 = ",message)
			//cellList,plmnList=f1apPackage.GetCellIdsPlmnIds(message)
			cellList,plmnList=isolation.IntermediateCaller(message)
			return
}
