package main

type WE struct {
	i int
}

var testFilePath = "/home/kamil/Downloads/test.hdbd"

func main() {
	//dbsystem.Config.DBHomePath()
	//dbsystem.DBHomePath()
	array := make([]WE, 2)
	value := &array[1]
	value.i = 10
	println(array[1].i)

	//file, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(1)
	//}
	//
	//data1 := "Data 1 for this file"
	//fmt.Printf("Writing following string with len = %d,\n%s\n", len(data1), data1)
	//_, err = file.Write([]byte(data1))
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(2)
	//}
	//
	//data2 := "Data 2 for this file2"
	//fmt.Printf("Writing following string with len = %d,\n%s\n", len(data2), data2)
	//_, err = file.WriteAt([]byte(data2), 0)
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(3)
	//}
	//
	//fileInf, err := file.Stat()
	//if err != nil {
	//	println(err)
	//	os.Exit(4)
	//}
	//
	//fLen := fileInf.Size()
	//fmt.Printf("File len = %d\n", fLen)
	//contentBuffer := make([]byte, fLen)
	//_, err = file.ReadAt(contentBuffer, 0)
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(5)
	//}
	//fmt.Printf("File content:\n%s\n", string(contentBuffer))
	//
	//err = file.Truncate(8)
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(6)
	//}
	//
	//fileInf, _ = file.Stat()
	//fLen = fileInf.Size()
	//fmt.Printf("File len = %d\n", fLen)
}
