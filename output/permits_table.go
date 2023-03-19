package output

import (
	"appgraph/db"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func PrintPermitsTable(permits []db.Permit) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "TYPE", "CSI", "STATUS", "SERVICES")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pmt := range permits {
		tbl.AddRow(pmt.Id, pmt.Type, pmt.Csi, pmt.Status, pmt.Services)
	}

	tbl.WithWriter(color.Output)
	tbl.Print()
}
