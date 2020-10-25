package component

import (
	v1 "github.com/hanjingo/golib/component/file_manager/v1"
	glua "github.com/yuin/gopher-lua"
)

func getFm1() *v1.FileManager {
	return nil
}

func Fm1OpenFile(state *glua.LState) int {
	fm1 := getFm1()
	return 0
}

func Fm1Write(state *glua.LState) int {
	fm1 := getFm1()

	return 0
}

func Fm1CloseFile(state *glua.LState) int {
	fm1 := getFm1()

	return 0
}

func Fm1CleanFile(state *glua.LState) int {
	fm1 := getFm1()

	return 0
}
