local file = require("file")
--local env = require("env")

print("compute_file_md5:", file.ComputeFileMd5('test.txt'))
print("is_file_exist:", file.IsFileExist('test.txt'))
print("get_file_size", file.GetFileSize('test.txt'))
print("get_file_name_and_typ", file.GetFileNameAndType('test.txt'))