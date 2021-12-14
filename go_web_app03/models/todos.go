package models

import "go_web_app03/dao"


type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//把todo这个model的增删改查操作都放在这里

func CreateATodoInModel(todo *Todo)(err error){
	err = dao.DB.Create(&todo).Error
	if err!=nil{
		return err
	}
	return
}

func GetTodoListInModel()(todos []*Todo,err error){
	err = dao.DB.Find(&todos).Error
	if err!=nil{
		return nil,err
	}
	return
}

func GetATodoInModel(anum string)(todo *Todo,err error){
	todo=new(Todo)
	err = dao.DB.First(&todo,anum).Error
	if err!=nil{
		return nil,err
	}
	return
}

func UpdateATodoInModel(todo *Todo)(err error){
	err=dao.DB.Save(&todo).Error
	if err!=nil{
		return err
	}
	return
}

func DeleteATodoInModel(anum string)(err error){
	err = dao.DB.Where("id=?", anum).Delete(Todo{}).Error
	if err!=nil{
		return err
	}
	return
}