package singleton

import (
	"sync"
)

type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			data: "Exemplo de dados para o Singleton",
		}
	})
	return instance
}
