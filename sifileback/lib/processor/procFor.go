package processor


import (
	"fmt"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/dockerApi"

	"github.com/xela07ax/toolsXela/tp"
	"path/filepath"
	"sync"
	"time"
)
type ToDocker struct {
	Data []byte
	Acton string // Приходит команда на распаковку и запуск
	Out chan ToDocker
}
const (
	fn = "Processor"
	limit = 15
)
type ProcessorCfg struct {
	SignOff chan struct{}
	PortsStart int // С какого порта начинаем выдачу под докер
	PortsEnd int // Это число минус верхнее, будет количество потоков
	Loger chan <-[4]string
	WorkDir string // С этой папки смотрим остальные пути и ее подключим докеру
}
type Box struct {
	BoxCfg ProcessorCfg
	TxInput chan ToDocker
	process *mutexRunner
	loger chan <-[4]string
	stopPrepare *mutexStopper
	stopX chan bool
	GoodBy chan struct{}
}

func NewProcessor(cfg ProcessorCfg) (Processor *Box, errTx string) {
	if cfg.Loger == nil {
		return nil, "логер не передан"
	}
	cfg.Loger  <- [4]string{fn,"nil","Модуль-балансировщик инициализация"}
	// Обработчик команд для поднятия докера, распаковка пакета и файлов настроек, завершение   докера и сохранение, п также удаление остатков
	cntMinions := cfg.PortsEnd - cfg.PortsStart

	if cntMinions > 15 || cntMinions < 1 {
		return nil, fmt.Sprintf("количество миньонов не может быть больше %d или меньше 1, проверьте диапаон портов",limit)
	}
	box :=  &Box {
		TxInput:make(chan ToDocker),
		loger:cfg.Loger,
		stopPrepare:new(mutexStopper),
		process:new(mutexRunner),
		stopX:make(chan bool,1000),
		GoodBy:make(chan struct{}), // todo:Про завершение надо еще подумать
	}
	err := tp.CheckMkdir(box.BoxCfg.WorkDir)
	if err != nil {
		box.loger <- [4]string{fn,"nil",fmt.Sprintf("CheckMkdir:%s",err),"1"}
	}
	for i := 0; i < cntMinions; i++ {
		// Запускаем миньонов
		go box.RunMinion(cfg.PortsStart-1+i)
	}
	if cfg.SignOff != nil {
		go box.SignalStoper(cfg.SignOff)
	}
	return box, ""
}
type mutexStopper struct {
	mu sync.Mutex
	x  bool
}
func (c *mutexStopper) stopSignal()  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x = true
	return
}
func (p *Box) Stop()  {
	p.stopPrepare.stopSignal()
	for i:=0; i <  p.GetCores(); i++ {
		<-p.stopX
	}
	// p.GoodBy <- true
	//fmt.Println("Модуль-балансировщик завершил все миньоны")
	return
}
func (p *Box) SignalStoper(off <-chan struct{})  {
	<- off
	p.Stop()
	p.loger <- [4]string{fn,"nil","Модуль-балансировщик завершил все миньоны"}
	p.GoodBy <- struct{}{}
	return
}
type mutexRunner struct {
	mu sync.Mutex
	x  int
}
func (c *mutexRunner) addRun()(i int)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x++
	i = c.x
	return
}
func (c *Box) GetCores()(cores int)  {
	c.process.mu.Lock()
	cores = c.process.x
	c.process.mu.Unlock()
	return
}
func (p *Box) RunMinion(port int)  {
	p.process.addRun()
	go p.minion(port)
	return
}
func (p *Box) minion(gophere int)  {
	p.loger <- [4]string{fn,fmt.Sprintf("%d",gophere),"Миньен инициализирован"}
	workMinion := filepath.Join(p.BoxCfg.WorkDir, fmt.Sprintf("%d",gophere))
	err := tp.CheckMkdir(workMinion)
	if err != nil {
		p.loger <- [4]string{fn,fmt.Sprintf("%d",gophere),fmt.Sprintf("CheckMkdir:%s",err),"1"}
	}
	/*
		Какой то код, который нужно выполнить до того как начнется сам процесс считывания данных со входящкго канала
	*/
	var counter int
	for  {
		select {
		case task := <-p.TxInput:
			// Оповестим, что получили команду на исполнение
			fmt.Println(task)
			p.loger <- [4]string{fn,fmt.Sprintf("%d",gophere),"получили команду на исполнение"}
			//ectractor(p.BoxCfg.KodConfigZip,workMinion,gophere,p.loger)
			//task.Out <- ToDocker{Acton: "Unarchive"}


			continue
			//fmt.Println(err)
			// Распакуем пришедший проект, развернем все

			//task.Data
			// И надо дождаться команды на остановку в task.ch
			p.loger <- [4]string{fn,fmt.Sprintf("%d",gophere),fmt.Sprintf("Пришла команда")}


			dockerApi.RunContainer(dockerApi.Kod{
				Image:      "xaljer/kodexplorer",
				DockerPort: 80,
				HostPort:   0,
				DockerDir:  "/var/www/html",
				HostDir:    "",
			})
			//newElement := fmt.Sprintf("%s - ок!", elem)

			if counter >= 50 {
				fmt.Printf("Мы достигли %d операций, все идет хорошо, продолжаю работу, я это миньен %d\n",counter,gophere)
				counter = 0
			}
			counter++
		default:
			if p.stopPrepare.x {
				p.loger <- [4]string{fn,fmt.Sprintf("%d",gophere),"Завершение"}
				// fmt.Printf("Мы достигли %d операций, и я завершаю работу, я это миньен %d\n",counter,gophere)
				p.stopX <- true
				return
			}
			time.Sleep(1*time.Second)
		}
	}
}