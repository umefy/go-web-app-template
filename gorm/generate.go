package main

import (
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func getGeneratePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	dirname := filepath.Dir(filePath)
	return filepath.Join(dirname, "./generated/query")
}

func getTablesForModelGen() []string {
	return []string{
		"users",
	}
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           getGeneratePath(),
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
	})

	g.WithDataTypeMap(getDataTypeMap())
	g.WithImportPkgPath("github.com/guregu/null/v6") // specify the 3rd party library import path

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}
	g.UseDB(db)

	err = os.RemoveAll(getGeneratePath())
	if err != nil {
		panic(err)
	}

	generateModels(g)
	g.Execute()
}

func generateModels(g *gen.Generator) {
	for _, tableName := range getTablesForModelGen() {
		// override id as int as it's never going to be 0
		g.ApplyBasic(g.GenerateModel(tableName,
			gen.FieldType("id", "int"),
		),
		)
	}
}

func getDataTypeMap() map[string]func(detailType gorm.ColumnType) (dataType string) {
	// nolint:goconst
	return map[string]func(detailType gorm.ColumnType) (dataType string){
		"json":  func(detailType gorm.ColumnType) (dataType string) { return "datatypes.JSON" },
		"jsonb": func(detailType gorm.ColumnType) (dataType string) { return "datatypes.JSON" },
		"int2": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[int]"
		},
		"int4": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[int]"
		},
		"int8": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[int]"
		},
		"varchar": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[string]"
		},
		"text": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[string]"
		},
		"tinyint": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[bool]"
		},
	}
}
