package dao

import "github.com/astaxie/beego/orm"
import "portal/model"

func Count(sql string, args []interface{}, dao orm.Ormer) int {
	temp := "select count(1) from (" + sql + ") temp";
	var count int
	dao.Raw(temp, args).QueryRow(&count)
	return count
}

func Paging(sql, orderSql string, args []interface{}, term *model.Term, dao orm.Ormer) orm.RawSeter {
	term.Total = Count(sql, args, dao)
	selectSql := sql + orderSql
	if term.Start() >= 0 {
		selectSql += " limit ?, ?"
		args = append(args, term.Start(), term.Size)
	} else {
		selectSql += " limit ?"
		args = append(args, term.Size)
	}
	return dao.Raw(selectSql, args)
}
