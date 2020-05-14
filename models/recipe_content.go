package models

import "github.com/astaxie/beego/orm"

type RecipeContent struct {
	Id          int
	Name        string `json:"name" orm:"column(german_name)"`
	ImgUrl      string
	ImgUrlLarge string
	VideoId     string
	Preparation string `json:"preparation" orm:"column(german_preparation)"`
	Ingredients string `json:"ingredients" orm:"column(german_ingredients)"`
	ListId      int
	VideoPath   string
}

func (recipeContent *RecipeContent) TableName() string {
	return "recipe_content"
}

func GetRecipeContentById(id int) RecipeContent {
	o := orm.NewOrm()
	var recipeContent RecipeContent
	_ = o.QueryTable("recipe_content").Filter("list_id", id).One(&recipeContent)
	return recipeContent
}
