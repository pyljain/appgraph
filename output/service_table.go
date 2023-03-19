package output

import (
	"appgraph/db"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func PrintServiceTable(services []db.Service) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "NAME", "TYPE", "CSI", "STATUS", "REPOSITORY")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, svc := range services {
		tbl.AddRow(svc.Id, svc.Name, svc.Type, svc.Csi, svc.Status, svc.Repository)
	}

	tbl.WithWriter(color.Output)
	tbl.Print()
}
