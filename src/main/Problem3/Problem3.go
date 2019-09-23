package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "log"
    "path/filepath"
)

func after(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

func filePathWalkDir(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}

func task1(source, target []string) {
    var counter, counter2 int
	for i:=0; i<len(source); i++{
    	for j:=0; j<len(target); j++{
    		if source[i] == target[j]{
    			counter++
    		}
    		if source[j] == target[i]{
    			counter2++
    		}
    	}
    	if counter == 0{
			fmt.Println(source[i]+"  NEW")
    	}
    	if counter2 == 0{
    		fmt.Println(target[i]+"  DELETED")
    	}	
    	counter = 0
    	counter2 = 0
    }
}

func getContent(temp string) string{
	// fmt.Println("\nTask 2\n")
	// var str string = file
	
    file, err := os.Open(temp)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var content string
    for scanner.Scan() {             // internally, it advances token based on sperator
        // fmt.Println(scanner.Text())  // token in unicode-char
    content = scanner.Text()
    }
    return content
}
var tempS,tempT []int;
func getSimilarSourcePosition(source, target []string) {
	// fmt.Println("\nTask 2\n")
    var counter, counter2  int
	for i:=0; i<len(source); i++{
    	for j:=0; j<len(target); j++{
    		if source[i] == target[j]{
    			counter++
    		}
    		if source[j] == target[i]{
    			counter2++
    		}
    	}
    	if counter == 0{
			fmt.Println(source[i]+"  NEW")
    	}else{
			tempS = append(tempS, i)
    	}
    	if counter2 == 0{
    		fmt.Println(target[i]+"  DELETED")
    	}else{
			tempT = append(tempT, i)
    	}
    	counter = 0
    	counter2 = 0
    }
}

func showModifiedFile(source, sourceContent, targetContent []string) {
    for i:=0; i<len(sourceContent); i++{
    	if sourceContent[i]!=targetContent[i]{
		fmt.Println(source[i]+"  MODIFIED")
		}
    }
}

func main() {
    var (
        root, temp string
        files, srcS, srcT []string
        err   error
        source, target []string
        sourceContent, targetContent []string
    )

    root = "src/main/Problem3"
    files, err = filePathWalkDir(root)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
    	// temp = string(file)
		temp = strings.ReplaceAll(string(file), `\`, "/")
    	if strings.Contains(temp, "Source"){
    		srcS = append(srcS, temp)
    		sourceContent = append(sourceContent, string(getContent(temp)))
            source =  append(source, after(temp, "Source"))
        }else if strings.Contains(temp, "Target"){
    		srcT = append(srcT, temp)
    		targetContent = append(targetContent, string(getContent(temp)))
            target =  append(target, after(temp, "Target"))
        }
    // fmt.Println(temp)
    }   
    
    fmt.Printf("Pilih tindakan :\n 1.Task 1\n 2.Task 2\n\nInputkan pilihan anda :" )
	var no string
	fmt.Scanln(&no)
	if (no == "1") {
		fmt.Println("Task 1\n")
    	task1(source, target);
	}else if (no == "2"){
		fmt.Println("Task 2\n")
		getSimilarSourcePosition(source, target);
	    sourceContent = nil;
	    targetContent = nil;
	    source = nil;
	    target = nil;
	    for i:=0; i<len(tempS); i++{
	    	sourceContent = append(sourceContent, string(getContent(srcS[tempS[i]])))
	    	source = append(source, after(srcS[tempS[i]], "Source"))
	    	// fmt.Println(source[i])    	
	    }
		for i:=0; i<len(tempT); i++{
	    	targetContent = append(targetContent,string(getContent(srcT[tempT[i]])))
	    	target = append(target, after(srcT[tempT[i]], "Target"))
	    	// fmt.Println(target[i])
	    }
	    // fmt.Println(len(source), len(target))
	    showModifiedFile(source, sourceContent, targetContent)
	}else{
		fmt.Println("Inputan salah, kembali keawal\n")
	}
}
