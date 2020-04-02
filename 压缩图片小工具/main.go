package main

import (
	"compr"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	// 获取命令行参数

	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		cmd := fmt.Sprintln(`
			参数名              含义
			 -d             指定文件夹 
			 -f             指定单个文件
			 -r             使用降低图片分辨率的方式缩小图片
			 -p             打印详细信息   eg: ./compress -d 文件夹名 -p
			 -s             分辨率宽度     eg: ./compress -r 文件夹名 -s 960 720
			`)
		fmt.Println(cmd)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println(len(os.Args))
		fmt.Println("param err")
		return
	}

	param := os.Args[1]
	Name := os.Args[2]

	var gz compr.Gzip
	var bz compr.Bzip2
	var lzma compr.Lzma
	var lzw compr.Lzw
	var zlib compr.Zlib
	var zstd compr.Zstd
	compress := []compr.Compr{&gz, &bz, &lzma, &lzw, &zlib, &zstd}

	switch param {
	case "-d":
		var isPrintData string
		if len(os.Args) == 4 {
			isPrintData = os.Args[3]
		} else {
			isPrintData = "no"
		}

		var err error
		for _, v := range compress {
			err = readFolder(Name, isPrintData, v)
		}
		if err != nil {
			fmt.Println("readFolder err:", err)
			return
		}

	case "-f":
		var err error
		for _, v := range compress {
			err = readFile(Name, v)
		}
		if err != nil {
			fmt.Println("readFolder err:", err)
			return
		}
		return
	case "-r":
		var width, high int
		isPrintData := "no"
		if len(os.Args) >= 6 && os.Args[3] == "-s" {
			width, _ = strconv.Atoi(os.Args[4])
			high, _ = strconv.Atoi(os.Args[5])
		}
		if len(os.Args) == 7 && os.Args[6] == "-p" {
			isPrintData = os.Args[6]
		}
		err := reducedResolution(Name, width, high, isPrintData)
		if err != nil {
			fmt.Println("reducedResolution err:", err)
			return
		}
	}
}

//通过降低图片分辨率压缩图片
func reducedResolution(folderName string, width, high int, isPrintData string) error {
	err := CreateDirIfNotExists("images2", 0777)
	if err != nil {
		fmt.Printf("CreateDirIfNotExists err:[%s]\n", err.Error())
		return errors.New("CreateDirIfNotExists err")
	}

	fileTxt, err := os.Create("记录.txt")
	if err != nil {
		fmt.Printf("Create err:%s\n", err.Error())
		return errors.New("Create err")
	}
	defer fileTxt.Close()

	//1 读取文件夹
	fileSlice, err := ioutil.ReadDir(folderName)
	if err != nil {
		fmt.Printf("ReadDir err:[%s]\n", err.Error())
		return errors.New("reducedResolution err")
	}

	fmt.Printf("********* [文件数量:%d] *********\n", len(fileSlice))

	var compressRate []float64
	var runTimeSlice []float64
	var beforeFileSize []int
	var afterFileSize []int

	for i, v := range fileSlice {
		//调用算法前数据
		startTime := time.Now()

		file, err := os.Open(fmt.Sprintf("%s/%s", folderName, v.Name()))
		if err != nil {
			fmt.Printf("Open err:%s\n", err.Error())
			return errors.New("Open err")
		}

		//fmt.Println(file.Name())

		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Printf("jpeg Decode err:%s\n", err.Error())
			continue
			//return errors.New("jpeg Decode err")
		}
		file.Close()

		m := resize.Resize(uint(width), uint(high), img, resize.Lanczos3)

		fileName := "./images2/" + v.Name() + ".jpg"
		out, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Create err: %s\n", err.Error())
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)

		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("ioutil ReadFile err:%s\n", err.Error())
			continue
			//return errors.New("ioutil ReadFile err")
		}

		beforeSize := v.Size()
		beforeFileSize = append(beforeFileSize, int(beforeSize))

		afterSize := len(data)
		afterFileSize = append(afterFileSize, afterSize)

		rate := float32(afterSize) / float32(beforeSize)
		compressRate = append(compressRate, float64(rate))

		spendTime := time.Now().Sub(startTime).Seconds()
		runTimeSlice = append(runTimeSlice, spendTime)

		if isPrintData == "-p" {
			fmt.Printf("[文件索引:%d || 文件压缩前尺寸:%d || 文件压缩后尺寸:%d || 文件压缩比:%f || 压缩花费时间:%f]\n",
				i, beforeSize, afterSize, rate, spendTime)

			temp := fmt.Sprintf("[文件索引:%d || 文件压缩前尺寸:%d || 文件压缩后尺寸:%d || 文件压缩比:%f || 压缩花费时间:%f]\n",
				i, beforeSize, afterSize, rate, spendTime)
			_, err := fileT.WriteString(temp)
			if err != nil {
				fmt.Printf("WriteString err:%s\n", err)
				return errors.New("WriteString err")
			}
		}
	}

	//计算各种平均值 最大值 最小值 average
	sort.Ints(beforeFileSize)
	sort.Ints(afterFileSize)
	sort.Float64s(compressRate)
	sort.Float64s(runTimeSlice)

	//压缩前文件平均大小
	averageBeforeFileSize := averageInt(beforeFileSize)
	//压缩后文件平均大小
	averageAfterFileSize := averageInt(afterFileSize)
	//平均压缩率
	averageCompressRate := averageFloat64(compressRate)
	//平均运行时间
	averageRunTimeSlice := averageFloat64(runTimeSlice)

	fmt.Println("=======================================================================")
	temp := "降低分辨率"
	fmt.Printf("[压缩算法:%s]\n", temp)

	fmt.Printf("[压缩前最小文件尺寸:%d || 压缩前最大文件尺寸:%d || 压缩前平均文件尺寸:%f]\n", beforeFileSize[0], beforeFileSize[len(beforeFileSize)-1], averageBeforeFileSize)

	fmt.Printf("[压缩后最小文件尺寸:%d || 压缩后最大文件尺寸:%d || 压缩后平均文件尺寸:%f]\n", afterFileSize[0], afterFileSize[len(afterFileSize)-1], averageAfterFileSize)

	fmt.Printf("[最小压缩率:%f || 最大压缩率:%f || 平均压缩率:%f]\n", compressRate[0], compressRate[len(compressRate)-1], averageCompressRate)

	fmt.Printf("[最小运行时间:%f || 最大运行时间:%f || 平均运行时间:%f]\n", runTimeSlice[0], runTimeSlice[len(runTimeSlice)-1], averageRunTimeSlice)

	fmt.Println("=======================================================================")

	return nil
}

//读取单个图片
func readFile(fileName string, gz compr.Compr) error {
	startTime := time.Now()
	data, _ := ioutil.ReadFile(fileName)
	res, err := gz.Compress(data)
	if err != nil {
		fmt.Printf("Compression algorithm err:[%s]\n", err.Error())
		return errors.New("readFile err")
	}
	spendTime := time.Now().Sub(startTime).Seconds()
	beforeSize := len(data)
	afterSize := len(res)
	rate := float32(afterSize) / float32(beforeSize)

	fmt.Println("***********************************************************************************")
	fmt.Printf("[压缩算法: %s || 文件压缩前尺寸:%d || 文件压缩后尺寸:%d || 文件压缩比:%f || 压缩花费时间:%f]\n", gz.Name(), beforeSize, afterSize, rate, spendTime)
	fmt.Println("***********************************************************************************")

	return nil
}

/*
	1 读取文件夹
	2 循环文件夹内的文件,读取文件各项数据
	3 调用压缩算法进行文件处理
	4 读取处理后的文件各项数据
	5 计算数据各项指标  压缩比  平均值  文件大小
	6 合理的打印出来
*/
func readFolder(folderName string, isPrintData string, gz compr.Compr) error {
	//1 读取文件夹
	fileSlice, err := ioutil.ReadDir(folderName)
	if err != nil {
		fmt.Printf("ReadDir err:[%s]\n", err.Error())
		return errors.New("readFolder err")
	}

	fmt.Printf("********* [文件数量:%d] *********\n", len(fileSlice))

	var compressRate []float64
	var runTimeSlice []float64
	var beforeFileSize []int
	var afterFileSize []int
	//2 循环文件夹内的文件,读取文件各项数据
	for i, v := range fileSlice {
		//调用算法前数据
		startTime := time.Now()

		//3 调用压缩算法进行文件处理
		fileName := folderName + "/" + v.Name()
		data, _ := ioutil.ReadFile(fileName)
		res, err := gz.Compress(data)
		if err != nil {
			fmt.Printf("Compression algorithm err:[%s]\n", err.Error())
			return errors.New("readFolder err")
		}

		//5 计算数据各项指标  压缩比  平均值  文件大小
		beforeSize := len(data)
		beforeFileSize = append(beforeFileSize, beforeSize)

		afterSize := len(res)
		afterFileSize = append(afterFileSize, afterSize)

		rate := float32(afterSize) / float32(beforeSize)
		compressRate = append(compressRate, float64(rate))

		spendTime := time.Now().Sub(startTime).Seconds()
		runTimeSlice = append(runTimeSlice, spendTime)

		//6 合理的打印出来
		if isPrintData == "-p" {
			fmt.Printf("[文件索引:%d || 文件压缩前尺寸:%d || 文件压缩后尺寸:%d || 文件压缩比:%f || 压缩花费时间:%f]\n",
				i, beforeSize, afterSize, rate, spendTime)
		}
	}

	if len(beforeFileSize) <= 0 || len(afterFileSize) <= 0 || len(compressRate) <= 0 || len(runTimeSlice) <= 0 {
		fmt.Println("文件夹为空")
		return errors.New("readFolder err")
	}

	//计算各种平均值 最大值 最小值 average
	sort.Ints(beforeFileSize)
	sort.Ints(afterFileSize)
	sort.Float64s(compressRate)
	sort.Float64s(runTimeSlice)

	//压缩前文件平均大小
	averageBeforeFileSize := averageInt(beforeFileSize)
	//压缩后文件平均大小
	averageAfterFileSize := averageInt(afterFileSize)
	//平均压缩率
	averageCompressRate := averageFloat64(compressRate)
	//平均运行时间
	averageRunTimeSlice := averageFloat64(runTimeSlice)

	fmt.Println("=======================================================================")
	fmt.Printf("[压缩算法: %s]\n", gz.Name())

	fmt.Printf("[压缩前最小文件尺寸:%d || 压缩前最大文件尺寸:%d || 压缩前平均文件尺寸:%f]\n", beforeFileSize[0], beforeFileSize[len(beforeFileSize)-1], averageBeforeFileSize)

	fmt.Printf("[压缩后最小文件尺寸:%d || 压缩后最大文件尺寸:%d || 压缩后平均文件尺寸:%f]\n", afterFileSize[0], afterFileSize[len(afterFileSize)-1], averageAfterFileSize)

	fmt.Printf("[最小压缩率:%f || 最大压缩率:%f || 平均压缩率:%f]\n", compressRate[0], compressRate[len(compressRate)-1], averageCompressRate)

	fmt.Printf("[最小运行时间:%f || 最大运行时间:%f || 平均运行时间:%f]\n", runTimeSlice[0], runTimeSlice[len(runTimeSlice)-1], averageRunTimeSlice)

	fmt.Println("=======================================================================")

	return nil
}

func averageInt(intSlice []int) float64 {
	sum := 0
	for _, i := range intSlice {
		sum += i
	}

	return float64(sum) / float64(len(intSlice))
}

func averageFloat64(float64Slice []float64) float64 {
	sum := 0.0
	for _, i := range float64Slice {
		sum += i
	}

	return float64(sum) / float64(len(float64Slice))
}

func CreateDirIfNotExists(path string, perm os.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, perm)
	}

	return err
}
