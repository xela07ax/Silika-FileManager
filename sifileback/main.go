package main

import (
	"./lib/daemon"
	"fmt"
	"github.com/xela07ax/toolsXela/chLogger"
	"log"
	"os"
)

func main()  {
	// Запускаем логер
	logEr := chLogger.NewChLoger(&chLogger.Config{
		IntervalMs:     300,
		ConsolFilterFn: map[string]int{"Front Http Server":  0},
		ConsolFilterUn: map[string]int{"Pooling": 1},
		Mode:           0,
		Dir:            "x-loger",
	})
	logEr.RunMinion()
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Вас приветствует \"Silika-FileManager Контроллер\" v1.1 (11112020) \n")}

	// Запустим сервер
	destroyed, err := daemon.Run(&daemon.Config{
		ServerPort: 3447,
		Close: make(chan os.Signal),
		Loger: logEr.ChInLog,
	})
	if err != nil {
		log.Fatalf("Error in main(): %v\n", err)
	}
	<- destroyed
	fmt.Println("Good By America")
}
