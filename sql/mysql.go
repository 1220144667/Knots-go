/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2025-04-20 11:27:05
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2025-04-20 16:47:04
 * @FilePath: /knots-go/sql/mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package sql

import (
	"database/sql"
	"errors"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int":              "int",
	"bigint":           "int64",
	"smallint":         "int16",
	"tinyint":          "int8",
	"numeric":          "float64",
	"decimal":          "float64",
	"float":            "float32",
	"real":             "float32",
	"double":           "float64",
	"double precision": "float64",
	"char":             "string",
	"varchar":          "string",
	"text":             "string",
	"date":             "time.Time",
	"time":             "time.Time",
	"timestamp":        "time.Time",
	"timestamptz":      "time.Time",
	"boolean":          "bool",
	"binary":           "[]byte",
	"varbinary":        "[]byte",
	"blob":             "[]byte",
	"clob":             "[]byte",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (model *DBModel) Open() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		model.DBInfo.UserName,
		model.DBInfo.Password,
		model.DBInfo.Host,
		model.DBInfo.Charset)
	model.DBEngine, err = sql.Open(model.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (model *DBModel) GetColumns(dnName, tableName string) ([]*TableColumn, error) {
	query := "SELECT " + "COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " + "FROM COLUMNS WHERE TABLE_NAME = ? AND TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	rows, err := model.DBEngine.Query(query, dnName, tableName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
