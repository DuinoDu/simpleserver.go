// Package main provides ...
package main

import (
    "time"
    "math/rand"
    "fmt"
    "os"
    "os/exec"
    "bytes"
    "strings"
    "net/http"
    "strconv"
    "io"
    "net"
    "reflect"
    "sort"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func NowString() string {
    return time.Now().Format("20060102_150405")
}

func RandFileName() string {
    name := fmt.Sprintf("%s_%s", NowString(), RandString(6))
    return name
}


// os related
func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()
}

func Exists(path string) bool {
    _, err := os.Stat(path)    //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func Tar(src string, dst string) bool {
    cmd := fmt.Sprintf("tar -zcvf %s %s", dst, src)
    _, err := Runshell(cmd)
    if err != nil {
        fmt.Println(fmt.Sprint(err))
        return false
    }
    return true
}


func GetPWD() string {
    pwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return pwd
}

func Runshell(cmdstr string) (string, error) {
    cmds := strings.Split(cmdstr, " ")
    cmd := exec.Command(cmds[0], cmds[1:]...)
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println(cmdstr)
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return stderr.String(), err
    }
    return stdout.String(), err
}

func WriteStringToFile(text string, filename string) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }
}

func WriteFileToResponse(Filename string, writer http.ResponseWriter) {
    //Check if file exists and open
    Openfile, err := os.Open(Filename)
    defer Openfile.Close() //Close after function return
    if err != nil {
        http.Error(writer, "File not found.", 404)
        return
    }
    //Get the Content-Type of the file
    //Create a buffer to store the header of the file in
    FileHeader := make([]byte, 512)
    //Copy the headers into the FileHeader buffer
    Openfile.Read(FileHeader)
    //Get content type of file
    FileContentType := http.DetectContentType(FileHeader)

    //Get the file size
    FileStat, _ := Openfile.Stat()                     //Get info from file
    FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

    //Send the headers
    writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
    writer.Header().Set("Content-Type", FileContentType)
    writer.Header().Set("Content-Length", FileSize)

    //Send the file
    //We read 512 bytes from the file already, so we reset the offset back to 0
    Openfile.Seek(0, 0)
    io.Copy(writer, Openfile) //'Copy' the file to the client
    return
}

func GetIp() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, address := range addrs {
        // 检查ip地址判断是否回环地址
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}


// python-like operate
func StringInSlice(item string, list []string) bool {
    for _, l := range list {
        if l == item {
            return true
        }
    }
    return false
}

/* 
usage:
key_slice = MapKeys2StringSlice(reflect.ValueOf(custom_map).MapKeys())
*/
func MapKeys2StringSlice(keys []reflect.Value) []string {
    strings := make([]string, len(keys))
    for i := 0; i < len(keys); i++ {
        strings[i] = keys[i].String()
    }
    sort.Strings(strings)
    return strings
}

//func main() {
//    //item := "aa,"
//    //list := []string{"aa", "bb"}
//    //fmt.Println(StringInSlice(item, list))
//
//    a := map[string]int{"aa":1, "bb":2}
//    strings := MapKeys2StringSlice(reflect.ValueOf(a).MapKeys())
//    fmt.Println(strings)
//}
