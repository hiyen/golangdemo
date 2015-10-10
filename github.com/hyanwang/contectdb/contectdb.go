package contectdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyanwang/importxlsx"
)

type Contectdb struct {
	Server_add   string
	Username     string
	Passwd       string
	DatabaseName string
	DB           *sql.DB
}

func (db_conf *Contectdb) Contect() error {
	con := db_conf.Username + ":" + db_conf.Passwd + "@tcp(" + db_conf.Server_add + ":3306)/" + db_conf.DatabaseName
	fmt.Println(con)
	db, err := sql.Open("mysql", con)
	if err == nil {
		db_conf.DB = db
	}
	// defer db.Close()

	return err
}

func (db_conf *Contectdb) Close() {
	db_conf.DB.Close()
}

func (db_conf *Contectdb) Test() {
	fmt.Println(db_conf.DB)
	d, _ := db_conf.DB.Query("select * from sync_tasks")

	fmt.Println(d.Columns())
}

func (db_conf *Contectdb) InsertRow(tracindb []importxlsx.TracingStaff) error {
	for _, v := range tracindb {
		if db_conf.ConR(v.Date, v.Name) {
			fmt.Println(v)
			stmi, err := db_conf.DB.Prepare("insert into track_worker (`trac_date`,`trac_chinese_name`,`trac_username`,`trac_abenyance_com_info`,`trac_sysauto_dist_wc_com`,`trac_rm_dep_info`,`trac_chk_dep_info`,`trac_abey_dep_info`,`trac_sysauto_dist_wc_dep`,`trac_rm_job_info`,`trac_chk_job_info`,`trac_abey_job_info`,`trac_sysauto_dist_wc_job`,`trac_finish_task`,`trac_get_task`,`trac_auto_dirst_task`,`trac_makin_w_task_start`,`trac_makin_trj_task`,`trac_finish_web_task`,`trac_makin_cont_task_start`,`trac_finish_trj_task`) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
			if err != nil {
				return err
			}
			_, err = stmi.Exec(v.Date, v.Ch_Name, v.Name, v.Abeyance_Com_Info, v.SysAuto_Distri_WC_Com, v.Remove_Dep_Info, v.Check_Dep_Info, v.Abeyance_Dep_Info, v.Sys_Auto_Distri_WC_Dep, v.Remove_Job_Info, v.Check_Job_Info, v.Abeyance_Job_Info, v.Sys_Auto_Distri_WC_Job, v.Finish_Task, v.Get_Task, v.Auto_Dirstri_Task, v.Making_Web_Task_Start, v.Making_TRJ_Task, v.Finish_Web_Task, v.Making_Contact_Task_Start, v.Finish_TRJ_Task)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (db_conf *Contectdb) ConR(con_date string, con_username string) bool {
	row, err := db_conf.DB.Prepare("select `trac_date`,`trac_username` from `track_worker` where `trac_date`=? and `trac_username`=?")

	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	var find_username string
	var find_date string
	row.QueryRow(con_date, con_username).Scan(&find_date, &find_username)
	if len(find_date) != 0 && len(find_username) != 0 {
		fmt.Printf("***数据库已经有这条记录: %+s,%+s \n", find_date, find_username)
		return false
	} else {
		return true
	}

}
