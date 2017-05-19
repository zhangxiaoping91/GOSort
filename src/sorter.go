package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"os"
	"bufio"
	"io"
	"strconv"
	"context"
	"algorithms/bubblesort"
	"algorithms/qsort"
)

var infile *string = flag.String("i", "infile", "plesae choose infile")
var outfile *string = flag.String("o", "outfile", "please choose outfile")
var algorithm *string = flag.String("a", "algorithm", "please choose algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	var filename string
	//path := getCurrentPath()
	path := "E:\\workspace\\home\\go\\GO20170515\\"
	fmt.Println("currentPath->", path)
	filename = "unsort.txt"
	outfile := "sort.txt"
	values, err := readvalues(path + filename)
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	for index, value := range values {
		fmt.Println("index->", index + 1, "value->", value)
	}
	context.TODO()
	//冒泡排序
	values = bubblesort.BubbleSort(values)
	//快速排序
	values = qsort.QuickSort(values, 0, len(values) - 1)
	writevalues(values, path + outfile)
}


/**
获取当前项目路径
*/
func getCurrentPath() interface{} {
	str, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(str, "\\")
	path := string(str[0 : i + 1])
	return path
}
/**
读文件
 */
func readvalues(filename string) (values []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("failed to open the file ", filename)
		return nil, err
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("too long line,seems unexpected")
			return
		}
		str := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

/**
写文件
 */
func writevalues(values [] int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed create outfile", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
