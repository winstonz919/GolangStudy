/*
 * @Author: WinstonRD
 * @Date: 2025-07-10 00:25:14
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-10 01:25:07
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/9-模板方法模式/main.go
 * @Description: 模板方法模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type DownLoader interface {
	Download()
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func NewTemplate(uri string, impl implement) *template {
	return &template{uri: uri, implement: impl}
}

func (t *template) Download() {
	fmt.Println("ready to download")
	t.implement.download()
	t.implement.save()
	fmt.Println("download done")
}

type HttpDownLoader struct {
	*template
}

func NewHttpDownLoader(uri string) DownLoader {
	httpDownLoader := &HttpDownLoader{}
	template := NewTemplate(uri, httpDownLoader)
	httpDownLoader.template = template
	return httpDownLoader
}

func (hdl *HttpDownLoader) download() {
	fmt.Println("download", hdl.uri, "via http")
}

func (hdl *HttpDownLoader) save() {
	fmt.Println("http save")
}

type FtpDownLoader struct {
	*template
}

func NewFtpDownLoader(uri string) DownLoader {
	ftpDownLoader := &FtpDownLoader{}
	template := NewTemplate(uri, ftpDownLoader)
	ftpDownLoader.template = template
	return ftpDownLoader
}

func (fdl *FtpDownLoader) download() {
	fmt.Println("download", fdl.uri, "via ftp")
}

func (fdl *FtpDownLoader) save() {
	fmt.Println("ftp save")
}

func main() {
	var downloader DownLoader = NewHttpDownLoader("www.google.com")
	downloader.Download()

	/* ------------------------ */
	fmt.Println("------------------------")

	downloader = NewFtpDownLoader("www.baidu.com")
	downloader.Download()
}
