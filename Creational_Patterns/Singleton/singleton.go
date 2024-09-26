package main

import (
	"fmt"
	"sync"
)

/* Singleton Patterns */
/* ใช้เพื่อให้แน่ใจว่ามีเพียงอ็อบเจ็กต์เดียวของคลาสในโปรแกรมทั้งหมด */

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
	s1       = GetInstance()
	s2       = GetInstance()
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	fmt.Println(s1 == s2)
}
