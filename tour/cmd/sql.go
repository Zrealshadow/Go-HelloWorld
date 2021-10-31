package cmd

import (
	"log"

	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sqlTable2struct",
	Long:  "sqlTable2struct",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql translation",
	Long:  "sql translation",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel connect fail: %v", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)

		if err != nil {
			log.Fatalf("dbModel.GetColumns err :%v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("Generate Err :%v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "root", "Input database username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "Input database password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "Input database host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "Input database coding")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "Input database driver type")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "Input database name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "Input database Tablename")
}
