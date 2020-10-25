package file

import (
	file "github.com/hanjingo/golib/file"
	glua "github.com/yuin/gopher-lua"
)

func ComputeFileMd5(state *glua.LState) int {
	filePath := state.CheckString(1)
	data, err := file.ComputeMD5(filePath)
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LString(string(data)))
	return 1
}

func IsFileExist(state *glua.LState) int {
	filePathName := state.CheckString(1)
	ok := file.IsExist(filePathName)
	state.Push(glua.LBool(ok))
	return 1
}

func GetFileSize(state *glua.LState) int {
	filePathName := state.CheckString(1)
	sz, err := file.GetSize(filePathName)
	if err != nil {
		state.Push(glua.LNumber(0))
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LNumber(sz.KB()))
	return 1
}

func GetFileNameAndType(state *glua.LState) int {
	filePathName := state.CheckString(1)
	name, typ := file.GetNameAndType(filePathName)
	state.Push(glua.LString(name))
	state.Push(glua.LString(typ))
	return 2
}
