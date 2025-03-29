package file

import "os"

// IsPathExist checks if a file or directory exists at the specified path.
// It returns true if the path exists, false if it doesn't exist, and an error
// if there was a problem checking the path (other than non-existence).
//
// 判断文件是否存在
// 不报错的情况下，false表示不存在，true表示存在
func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	// 报错的情况分两种，一种是文件不存在，另一种是其他报错
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// IsDir checks if the specified path is a directory.
// It returns true if the path is a directory, false otherwise, and an error
// if there was a problem checking the path (including if the path doesn't exist).
//
// Note: If you need to check both existence and directory status, using os.Stat directly
// would be more efficient than calling this function.
//
// 检查路径是否是目录
// 如果既要检查路径是否存在，又要判断是否是目录，不用这个函数，直接os.Stat会更快
func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
		return false, err
	}

	return fileInfo.IsDir(), nil
}
