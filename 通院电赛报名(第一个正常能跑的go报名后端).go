package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func check(formData map[string]interface{}) string { //新生查重
	print("正在查重")
	db, _ := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/aa?charset=utf8")
	rows, _ := db.Query("SELECT ID FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["md5Code"] {
			return "标识码重复"
		}
	}

	rows, _ = db.Query("SELECT studentNum FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["studentNum"] {
			return "学号重复"
		}
	}

	rows, _ = db.Query("SELECT tele FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tele"] {
			return "手机号重复"
		}
	}

	rows, _ = db.Query("SELECT cNum FROM others") //新生跨老生表查询
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["studentNum"] {
			return "学号重复"
		}
	}

	rows, _ = db.Query("SELECT cTele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tele"] {
			return "电话重复"
		}
	}

	rows, _ = db.Query("SELECT tm1Num FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["studentNum"] {
			return "学号重复"
		}
	}

	rows, _ = db.Query("SELECT tm1Tele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tele"] {
			return "电话重复"
		}
	}

	rows, _ = db.Query("SELECT tm2Num FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["studentNum"] {
			return "学号重复"
		}
	}

	rows, _ = db.Query("SELECT tm2Tele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tele"] {
			return "电话重复"
		}
	}
	rows, _ = db.Query("SELECT studentNum FROM freshman") //老生跨新生表查询
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["cNum"] {
			return "队长学号重复"
		}
	}

	rows, _ = db.Query("SELECT tele FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["cTele"] {
			return "队长手机号重复"
		}
	}

	rows, _ = db.Query("SELECT studentNum FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm1Num"] {
			return "队员一学号重复"
		}
	}

	rows, _ = db.Query("SELECT tele FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm1Tele"] {
			return "队员一手机号重复"
		}
	}

	rows, _ = db.Query("SELECT ID FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["md5Code"] {
			return "队伍标识重复"
		}
	}

	rows, _ = db.Query("SELECT teamName FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["teamName"] {
			return "队伍名称重复"
		}
	}

	rows, _ = db.Query("SELECT cNum FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["cNum"] {
			return "队长学号重复"
		}
	}

	rows, _ = db.Query("SELECT cTele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["cTele"] {
			return "队长电话重复"
		}
	}

	rows, _ = db.Query("SELECT tm1Num FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm1Num"] {
			return "队员一学号重复"
		}
	}

	rows, _ = db.Query("SELECT tm1Tele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm1Tele"] {
			return "队员一电话重复"
		}
	}
	return "查重通过"
}

func checkTm2(formData map[string]interface{}) string {
	db, _ := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/aa?charset=utf8")
	rows, _ := db.Query("SELECT studentNum FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm2Num"] {
			return "队员二学号重复"
		}

	}

	rows, _ = db.Query("SELECT tele FROM freshman")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm2Tele"] {
			return "队员二电话重复"
		}

	}

	rows, _ = db.Query("SELECT tm2Num FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm2Num"] {
			return "队员二学号重复"
		}
	}

	rows, _ = db.Query("SELECT tm2Tele FROM others")
	for rows.Next() {
		var fg string
		rows.Scan(&fg)
		if fg == formData["tm2Tele"] {
			return "队员二电话重复"
		}
	}

	return "查重通过"

}

func freshman(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/aa?charset=utf8")
	r.ParseForm()

	err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseForm: ", err)
	}

	formData := make(map[string]string)
	json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		fmt.Print(" key:", key, " = value :", value, " ")
	}
	fmt.Println("")
	fmt.Println("数据获取成功")
	flag := "查重通过"
	w.Write([]byte(flag))
	//获取表单数据
	if flag == "查重通过" {
		studentNum := formData["studentNum"]
		studentName := formData["studentName"]
		studentClass := formData["studentClass"]
		majorDire := formData["majorDire"]
		tele := formData["tele"]
		prefer := formData["prefer"]
		ID := formData["md5Code"]
		subject := formData["subject"]
		stmt, _ := db.Prepare("INSERT freshman SET ID=?,studentNum=?,studentClass=?,studentName=?,prefer=?,tele=?,majorDire=?,subject=?")
		stmt.Exec(ID, studentNum, studentClass, studentName, prefer, tele, majorDire, subject)
		fmt.Print(" ID: ", ID)
		fmt.Print(" name: ", studentName)
		fmt.Print(" number: ", studentNum)
		fmt.Print(" class: ", studentClass)
		fmt.Print(" prefer: ", prefer)
		fmt.Print(" tele: ", tele)
		fmt.Println("majorDire", majorDire)
	}

}

func others(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/aa?charset=utf8")
	err := r.ParseForm() // 解析 url 传递的参数，对于 POST解析响应包的主体（request body）
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseForm: ", err)
	}
	// 初始化请求变量结构
	formData := make(map[string]interface{})
	// 调用json包的解析，解析请求body
	json.NewDecoder(r.Body).Decode(&formData)
	flag := check(formData)
	//录入队长数据
	w.Write([]byte(flag))
	if flag == "查重通过" {
		ID := formData["md5Code"]
		teamName := formData["teamName"]
		studentName := formData["cName"]
		studentNum := formData["cNum"]
		studentClass := formData["cClass"]
		prefer := formData["cPrefer"]
		tele := formData["cTele"]
		majorDire := formData["cMajorDire"]
		subject := formData["subject"]
		stmt, _ := db.Prepare("INSERT others SET ID=?,teamName=?, cName=?, cNum=?,cClass=?,cPrefer=?,cTele=?,cMajorDire=?,subject=?")
		stmt.Exec(ID, teamName, studentName, studentNum, studentClass, prefer, tele, majorDire, subject)
		fmt.Print("succeed TID: ", ID)
		fmt.Print(" name: ", studentName)
		fmt.Print(" number: ", studentNum)
		fmt.Print(" class: ", studentClass)
		fmt.Print(" prefer: ", prefer)
		fmt.Print(" tele: ", tele)
		fmt.Print(" majorDire", majorDire)
		//录入队员1数据
		studentName = formData["tm1Name"]
		studentNum = formData["tm1Num"]
		studentClass = formData["tm1Class"]
		prefer = formData["tm1Prefer"]
		tele = formData["tm1Tele"]
		majorDire = formData["tm1MajorDire"]
		stmt, _ = db.Prepare("UPDATE others SET tm1Name=?,tm1Num=?,tm1Class=?,tm1Prefer=?,tm1Tele=?,tm1MajorDire=? WHERE ID=?")
		stmt.Exec(studentName, studentNum, studentClass, prefer, tele, majorDire, ID)
		fmt.Print(" name: ", studentName)
		fmt.Print(" number: ", studentNum)
		fmt.Print(" class: ", studentClass)
		fmt.Print(" prefer: ", prefer)
		fmt.Print(" tele: ", tele)
		fmt.Print(" majorDire", majorDire)
		//录入队员2数据
		studentName = formData["tm2Name"]
		if studentName != nil {
			checkTm2(formData)
			studentNum = formData["tm2Num"]
			studentClass = formData["tm2Class"]
			prefer = formData["tm2Prefer"]
			tele = formData["tm2Tele"]
			majorDire = formData["tm1MajorDire"]
			stmt, _ = db.Prepare("UPDATE others SET tm2Name=?,tm2Num=?,tm2Class=?,tm2Prefer=?,tm2Tele=?,tm2MajorDire=? WHERE ID=?")
			stmt.Exec(studentName, studentNum, studentClass, prefer, tele, majorDire, ID)
			fmt.Print(" name: ", studentName)
			fmt.Print(" number: ", studentNum)
			fmt.Print(" class: ", studentClass)
			fmt.Print(" prefer: ", prefer)
			fmt.Print(" tele: ", tele)
			fmt.Println(" majorDire", majorDire)
		}
	}
}

func push(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
		log.Fatal("ParseForm: ", err)
	}

}

func main() {
	http.HandleFunc("/freshman", freshman)
	http.HandleFunc("/others", others)
	print("准备开始运行服务")
	err := http.ListenAndServe(":8080", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Println("http server listen err :", err)
	}
	print("服务意外退出")
}
