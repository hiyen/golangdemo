/*
import xlsx file
@author:87733449@qq.com
@dep: webop
@date: 2015-03-10

*/
package importxlsx

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type position struct {
	X, Y int
	Name string
}

type Importxlsx struct {
	ExecelxFile      string //the path of xlsx file
	Sheet_row_start  string //start field
	Sheet_row_end    string //end field
	Sheet_column_end string // end column
	Sheet_num        int    //which sheet?
}

type TracingStaff struct {
	Date                      string //data's date
	Name                      string //Internal username
	Ch_Name                   string //Chinese Name
	Abeyance_Com_Info         int    //搁置公司信息
	SysAuto_Distri_WC_Com     int    //系统自动分配网才公司
	Remove_Dep_Info           int    //删除部门信息
	Check_Dep_Info            int    //检查部门信息
	Abeyance_Dep_Info         int    //搁置部门信息
	Sys_Auto_Distri_WC_Dep    int    //系统自动分配网才部门
	Remove_Job_Info           int    //删除职位信息
	Check_Job_Info            int    //检查职位信息
	Abeyance_Job_Info         int    //搁置职位信息
	Sys_Auto_Distri_WC_Job    int    //系统自动分配网才职位
	Finish_Task               int    //完成任务
	Get_Task                  int    //取任务
	Auto_Dirstri_Task         int    //系统自动分配任务
	Making_Web_Task_Start     int    //开始处理网页制作任务
	Making_TRJ_Task           int    //开始处理投录校任务
	Finish_Web_Task           int    //完成网页制作任务
	Making_Contact_Task_Start int    //开始处理contact任务
	Finish_TRJ_Task           int    //完成投录校任务
}

var (
	year string
	mon  string
	day  string
	date string
)

// the XLSX's file convert to slice's type
func (imxlsx *Importxlsx) FileToSlice() ([][]string, error) {
	sheet_myslice, err := xlsx.FileToSlice(imxlsx.ExecelxFile)
	if err != nil {
		return nil, err
	}
	return sheet_myslice[imxlsx.Sheet_num], err
}

// digging the data's zoom
func (imxlsx *Importxlsx) DigPosition(myslice [][]string) []position {

	pos := make([]position, 0)
	for s_id, row := range myslice {
		for c_id, c_row := range row {
			if c_row == imxlsx.Sheet_row_start {
				pos = append(pos, position{s_id, c_id, imxlsx.Sheet_row_start})
				// fmt.Println(pos)
			}

			if c_row == imxlsx.Sheet_row_end {
				// fmt.Println(pos)
				pos = append(pos, position{s_id, c_id, imxlsx.Sheet_row_end})
			}

			if c_row == imxlsx.Sheet_column_end {
				if c_id == 0 {
					pos = append(pos, position{s_id, c_id, imxlsx.Sheet_column_end})
				}
			}
		}
	}
	return pos
}

// slicing the TracingStaff to Array type
func SlicinStuff(pos []position, slice [][]string) (tracindb []TracingStaff) {
	// tracindb := make([]TracingStaff, 0)
	start := pos[0].X + 1
	col_end := pos[2].X
	for i := start; i < col_end; i++ {
		if slice[i][0] != "" {
			year = slice[i][0]
		}
		if slice[i][2] != "" {
			mon = slice[i][2]
		}
		if slice[i][3] != "" {
			day = slice[i][3]
		}
		date := year + "-" + mon + "-" + day
		tracing := TracingStaff{
			Date:                      date,
			Name:                      username(slice[i][5]),
			Ch_Name:                   ch_name(slice[i][5]),
			Abeyance_Com_Info:         atoi(slice[i][6]),
			SysAuto_Distri_WC_Com:     atoi(slice[i][7]),
			Remove_Dep_Info:           atoi(slice[i][8]),
			Check_Dep_Info:            atoi(slice[i][9]),
			Abeyance_Dep_Info:         atoi(slice[i][10]),
			Sys_Auto_Distri_WC_Dep:    atoi(slice[i][11]),
			Remove_Job_Info:           atoi(slice[i][12]),
			Check_Job_Info:            atoi(slice[i][13]),
			Abeyance_Job_Info:         atoi(slice[i][14]),
			Sys_Auto_Distri_WC_Job:    atoi(slice[i][15]),
			Finish_Task:               atoi(slice[i][16]),
			Get_Task:                  atoi(slice[i][17]),
			Auto_Dirstri_Task:         atoi(slice[i][18]),
			Making_Web_Task_Start:     atoi(slice[i][19]),
			Making_TRJ_Task:           atoi(slice[i][20]),
			Finish_Web_Task:           atoi(slice[i][21]),
			Making_Contact_Task_Start: atoi(slice[i][22]),
			Finish_TRJ_Task:           atoi(slice[i][23])}
		if tracing.Name != "" && checkdate(tracing.Date) {
			tracindb = append(tracindb, tracing)
		}

	}
	removesame(&tracindb)
	return tracindb
}

//String convert to Integer
func atoi(s string) int {
	con_i, err := strconv.Atoi(s)
	if err != nil {
		return con_i
	} else {
		// fmt.Println(err)
		return con_i
	}
}

//filting the username and chinese name.
func username(source string) string {
	pa := `[[:alpha:].[:alpha:]]{1,}`
	reg := regexp.MustCompile(pa)
	return strings.Join(reg.FindAllString(source, -1), "")
}

//checking the date formate
func checkdate(chk_date string) bool {
	num_pa := `\d$`
	reg := regexp.MustCompile(num_pa)
	return reg.MatchString(chk_date)
}

//chinese name
func ch_name(source string) string {
	ch_pa := `[\p{Han}]+`
	reg := regexp.MustCompile(ch_pa)
	return strings.Join(reg.FindAllString(source, -1), "")
}

// remove the same data
func removesame(tac *[]TracingStaff) {
	found := make(map[TracingStaff]bool)
	j := 0
	for s_id, slice := range *tac {
		if !found[slice] {
			found[slice] = true
			(*tac)[j] = (*tac)[s_id]
			j++
		}
	}
	*tac = (*tac)[:j]
	// return *tac
}
