package main

import (
	"fmt"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/communicator"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/daemon"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/processor"
	"github.com/xela07ax/toolsXela/chLogger"
	"github.com/xela07ax/toolsXela/hub"
	"github.com/xela07ax/toolsXela/tp"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main()  {
	//dat, _ := tp.OpenReadFile("zip.zip")
	//nt := Notyfy{
	//	Name: "RunKod",
	//	Text: "Запуск проекта с вложением",
	//	Data: dat,
	//}
	//bt, _ := json.Marshal(nt)
	//fmt.Printf("%s",bt)
	//return
	// 1. Настройка хаба для клиентов
	hubib := hub.NewHub(false)
	go hubib.Run()

	// Для демона в роут вэбсокета надо отдать ссылку на Хаб
	dir, err := tp.BinDir()
	tp.Fck(err)

	// Запускаем логер
	fullLogPath := filepath.Join(dir,"chloger")
	err = tp.CheckMkdir(fullLogPath)
	tp.Fck(err)
	logEr := chLogger.NewChLoger(&chLogger.Config{
		IntervalMs:     300,
		ConsolFilterFn: map[string]int{"Front Http Server":  0},
		ConsolFilterUn: map[string]int{"Pooling": 1},
		Mode:           0,
		Dir:            fullLogPath,
		Broadcast: hubib.Input, // Трансляция в вэбсокет
	})
	logEr.RunMinion()
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Вас приветствует \"Silika-FileManager Контроллер\" v1.1 (11112020) \n")}
	// Реализуем небольшой обработчик входящих Вэбсокет клиентов

	// Запустим перехватчик системного сигнала о выходе, размножит сигнал всем нуждающимся
	closer := NewCloser(logEr.ChInLog)

	// Запустим обработчик докера
	tempPath := filepath.Join(dir,"dockerFile")
	err = tp.CheckMkdir(fullLogPath)
	if err != nil {
		logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("CheckMkdir:%s",err)}
		tp.ExitWithSecTimeout(1)
	}
	kodConf := filepath.Join(dir,"res", "kodConfig.zip")
	// KodConfigZip string // zip файл, который распаковывается в рабочую папку докера
	proc, ertx := processor.NewProcessor(processor.ProcessorCfg{
		PortsStart:   6100,
		PortsEnd:     6103,
		SignOff: 	closer.GetDestroyChan(),//SignOff <- Когда придет сюда сообщение, начнется остановка
		Loger:        logEr.ChInLog,
		WorkDir:      tempPath,
	})
	if ertx != "" {
		logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("NewProcessor:%s",ertx)}
		tp.ExitWithSecTimeout(1)
	}
	fmt.Println(proc)
	// Запустим сервер
	server := &daemon.Config {
		ServerPort: 3447,
		Destroy: nil,
		Loger: logEr.ChInLog,
		Hub: hubib,
	}


	go daemon.Run(server)

	commr := communicator.Coming{
		Bag: &communicator.Box{ZipKodConfigDefault: kodConf},
		Loger: logEr.ChInLog,
		InputClients: hubib.WebSocketOutput,
		OutputTxDock: proc.TxInput,
	}
	go commr.ProcessRun()
	//fmt.Println("get destroi")
	//ch := closer.GetDestroyChan()
	//fmt.Println("get destroi")
	//<- ch
	<- proc.GoodBy
	fmt.Println("Off")
	//<- server.Destroy // Не будем отключаться пока сервер не завершится
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Good by")}
	fmt.Println("Good By America")
	time.Sleep(500*time.Millisecond) // Подождем, не все успевает завершиться
}

type Closer struct {
	watchers []chan struct{}
}
func NewCloser(loger chan<- [4]string) *Closer {
	c := &Closer{watchers: make([]chan struct{}, 0, 5)}
	go c.waitForSignal(loger)
	return c
}

func (c *Closer)SendMessages()  {
	for i, v := range c.watchers {
		fmt.Printf("Безопасное завершение программы | SendMessages:%d|ln:%d\n",i+1,len(c.watchers))
		v <- struct {}{}
	}
}
func (c *Closer)GetDestroyChan() chan struct{}  {
	destroy := make(chan struct{})
	c.watchers = append(c.watchers, destroy)
	return destroy
}
func (c *Closer)waitForSignal(loger chan<- [4]string) {
	close := make(chan os.Signal)
	subNmae := "Daemon_Closer"
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)
	s := <-close
	fmt.Printf("%s | Безопасное завершение программы\n",subNmae)
	// Диплога могло вообще не создасться
	// !!! Мы не проверяем записаны ли все данные, просто закрываем это дело
	// loger <- [4]string{"Front Http Server","nil",fmt.Sprintf("Got signal: %v, exiting.", s)}
	// Пока программа при обрыве завершается через этот блок
	loger <- [4]string{"Front Http Server", "nil", fmt.Sprintf("%s | COM:Безопасное завершение программы | DT:%s", subNmae, s)}
	loger <- [4]string{"Front Http Server", "nil", fmt.Sprintf("%s | COM:Завершаем сессии", s)}
	// fmt.Printf("%d | Закрываем базу с пользователями\n",operInc())

	//ReactBro.Close()
	c.SendMessages()
}