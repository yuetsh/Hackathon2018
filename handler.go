package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

type Asset struct {
	Name         string   `json:"name"`
	Cover        string   `json:"cover"`
	Gif          string   `json:"gif"`
	Placeholders []string `json:"placeholders"`
}

func (h *Handler) listMemes(r *http.Request) interface{} {
	return [3]Asset{
		{
			Name:  "zhenxiang",
			Cover: "https://i.loli.net/2018/10/11/5bbeba99d8ce4.png",
			Gif:   "https://i.loli.net/2018/10/11/5bbeba9a6fae7.gif",
			Placeholders: []string{
				"我王境泽就是饿死",
				"死外面 从这里跳下去",
				"也不会吃你们一点东西",
				"真香",
			},
		},
		{
			Name:  "sorry",
			Cover: "https://i.loli.net/2018/10/11/5bbeba99cf285.png",
			Gif:   "https://i.loli.net/2018/10/11/5bbeba9ab1b99.gif",
			Placeholders: []string{
				"好啊",
				"别说我是一等良民",
				"就算你们真的想诬告我",
				"我有的是钱请律师打官司",
				"我想我根本不用坐牢",
				"你别以为有钱了不起啊",
				"Sorry 有钱真的可以为所欲为",
				"不过 我想你不会明白这种感觉",
				"不会 不会",
			},
		},
		{
			Name:  "dagong",
			Cover: "https://i.loli.net/2018/10/11/5bbf10ca7100c.png",
			Gif:   "https://i.loli.net/2018/10/11/5bbf10937b454.gif",
			Placeholders: []string{
				"没有钱啊肯定要做啊",
				"不做的话没有钱用",
				"那你不会打工啊",
				"有手有脚的",
				"打工是不可能打工的",
				"这辈子不可能打工的",
			},
		},
	}
}

func (h *Handler) createMeme(r *http.Request) interface{} {
	m := new(Meme)
	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		return err
	}
	if err := m.New(); err != nil {
		return err
	}
	data, err := UploadGif(m.paths.output.gif)
	if err != nil {
		return err
	}
	return data
}
