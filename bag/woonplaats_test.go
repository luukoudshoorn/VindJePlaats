package bag

import (
	"fmt"
	"testing"
)

func TestImportWoonplaats(t *testing.T) {
	//Woonplaatsen := ImportWoonplaats("D:/BAG/9999WPL08022023/example.xml")
	Woonplaatsen := ImportWoonplaats("D:/BAG/9999WPL08022023/9999WPL08022023-000001.xml")
	fmt.Println(len(Woonplaatsen))
}
