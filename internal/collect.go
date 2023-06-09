package internal

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sky0621/kreu-crud-for-sqlc/internal/parser"
)

func CollectSQLParseResult(root string, targetSQLNames []string) ([]*parser.SQLParseResult, error) {
	var sqlParseResults []*parser.SQLParseResult
	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		log.Printf("root: %s, path: %s, error: %v\n", root, path, err)
		if d.IsDir() {
			log.Println("is directory, so skip")
			return nil
		}

		fl, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer func() {
			if fl != nil {
				if err := fl.Close(); err != nil {
					fmt.Println(err)
				}
			}
		}()

		stat, err := fl.Stat()
		if err != nil {
			panic(err)
		}

		sqlFileName := stat.Name()
		log.Printf("sqlFileName:%s\n", sqlFileName)
		if !strings.HasSuffix(sqlFileName, ".sql") {
			log.Println("it is not target suffix, so skip")
			return nil
		}

		sqlName := ""
		sql := strings.Builder{}

		sc := bufio.NewScanner(fl)
		for sc.Scan() {
			line := sc.Text()

			if isBlankLine(line) {
				continue
			}

			if isSQLNameLine(line) {
				sqlName = getSQLName(line)
				continue
			}

			sql.WriteString(line + " ")

			if isEndSQL(line) {
				if contains(sqlName, targetSQLNames) {
					res, err := NewSQLParser().Parse(sqlName, sqlFileName, sql.String())
					if err != nil {
						panic(err)
					}
					sqlParseResults = append(sqlParseResults, res)
				}

				sqlName = ""
				sql.Reset()
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return sqlParseResults, nil
}

func contains(str string, strs []string) bool {
	if len(strs) == 0 {
		return true
	}
	for _, s := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func isBlankLine(line string) bool {
	return len(strings.Trim(line, " ")) == 0
}

func isSQLNameLine(line string) bool {
	return strings.HasPrefix(strings.Trim(line, " "), "-- name")
}

func isEndSQL(line string) bool {
	return strings.HasSuffix(strings.Trim(line, " "), ";")
}

func getSQLName(line string) string {
	// 形式　-- name: CreateGuestToken :one
	tLine := strings.Trim(line, " ")
	tpLine := strings.TrimPrefix(tLine, "--")
	tokens := strings.Split(tpLine, ":")
	if len(tokens) != 3 {
		return ""
	}
	return strings.Trim(tokens[1], " ")
}

func CollectTableNames(sqlParseResults []*parser.SQLParseResult) []string {
	m := map[string]struct{}{}
	for _, x := range sqlParseResults {
		for _, y := range x.TableNameWithCRUDSlice {
			m[y.TableName.ToString()] = struct{}{}
		}
	}
	var r []string
	for k, _ := range m {
		r = append(r, k)
	}
	sort.Strings(r)
	return r
}
