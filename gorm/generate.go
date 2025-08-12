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
	return filepath.Join(dirname, "../internal/infrastructure/database/gorm/generated/query")
}

func getTablesToGenerate() []string {
	return []string{
		"users",
		"orders",
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
	g.WithImportPkgPath("github.com/guregu/null/v6", "gorm.io/plugin/optimisticlock") // specify the 3rd party library import path

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
	for _, table := range getTablesToGenerate() {
		g.ApplyBasic(
			g.GenerateModel(
				table,
				gen.FieldType("id", "int"),
				gen.FieldType("user_id", "int"),
				gen.FieldType("version", "optimisticlock.Version"),
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
			return "null.Value[int64]"
		},
		"bigint": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[int64]"
		},
		"float4": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[float64]"
		},
		"float8": func(detailType gorm.ColumnType) (dataType string) {
			return "null.Value[float64]"
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
