package main

import (
	log "github.com/sirupsen/logrus"
	"fmt"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetLevel(log.WarnLevel) //ubah jenis log untuk mendapatkan output yang dibutuhkan
}


func main() {
	// log.Info("Starting")
	log.WithFields(log.Fields{
		"name": "John Doe",
		"age": 25,
	}).Info("biodata")

	log.WithFields(log.Fields{
		"name": "Amelia",
		"age": 25,
	}).Error("biodata")

	log.WithFields(log.Fields{
		"name": "Sarah",
		"age": 25,
	}).Warn("biodata")

	fmt.Println("Hello world")

	// Global Logger (jika ingin menggunakan ini, harus matikan/hilangkan dulu log set level pada func init diatas)
	ContextLogger := log.WithFields(log.Fields{
		"message": "harus di bawa terus",
	})

	ContextLogger.Debug("result 1")
	ContextLogger.Info("result 2")
	ContextLogger.WithFields(log.Fields{
		"data":"response",
	}).Info("result 2")
}