package workflow

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"math"
	"time"
)

func Power(n float64) float64 {
	return math.Pow(2, n)
}

//负责创建和执行作业
//进行站点服务注册
//路径：/flowsite/node1 /flowsite/node2
type FlowExecutorSite struct {
}

type TaskDescInfo struct {
	PumpFreq int8
	IsCron   bool
	Name     string
	AndPrev  []string
	OrPrev   []string
	Next     string
	NodeType TaskNodeType
}

type FlowTemplateManager struct {
}

type FlowTemplate struct {
	Key           string
	Name          string
	taskDescInfos g.List
}

func (ft *FlowTemplate) ParseTaskInfo() {

}

//one routine each job
//one job have muti tasks
//作业注册 /flowsite/node1/jobs/run  /flowsite/node1/jobs/fin   /flowsite/node1/jobs/term
type Job struct {
	tasks g.Map
}

type JobContext struct {
	context.Context
	Tasks g.Map
}

func (jobCtx *JobContext) CheckAndPrev(prevAndTasks []string) bool {
	return true
}
func (jobCtx *JobContext) CheckOrPrev(prevOrTasks []string) bool {
	return true
}
func (jobCtx *JobContext) Next(nextTask string) {

}

//one routine each task

func NewTask(ctx JobContext, taskDescInfo TaskDescInfo) *Task {
	task := &Task{PumpFreq: taskDescInfo.PumpFreq,
		IsCron:  taskDescInfo.IsCron,
		Name:    taskDescInfo.Name,
		AndPrev: taskDescInfo.AndPrev,
		OrPrev:  taskDescInfo.OrPrev,
		Next:    taskDescInfo.Next,
	}
	return task
}

type Task struct {
	PumpFreq int8
	IsCron   bool
	Name     string
	AndPrev  []string
	OrPrev   []string
	Next     string
	NodeType TaskNodeType
	Status   StatusType
}

func (t *Task) ifCanDoCheck(ctx JobContext) bool {
	return true
}

func (t *Task) DoExec(ctx JobContext) {
	for {
		select {
		case <-ctx.Done():
			t.Terminate(ctx)
			return
		default:
			time.Sleep(time.Second * time.Duration(t.PumpFreq))
			if t.ifCanDoCheck(ctx) == true {
				//完成核心任务逻辑
				t.DoTask(ctx)
				//处理任务完成后的事务
				t.End(ctx)
				if t.IsCron == false {
					return
				}
			}
		}
	}
}
func (t *Task) DoTask(ctx JobContext) {
	fmt.Println("task is running....")
}
func (t *Task) Terminate(ctx JobContext) {
	fmt.Println("task is Terminate....", ctx.Err())
}
func (t *Task) End(ctx JobContext) {
	ctx.Next(t.Next)
	fmt.Println("task is ending....")
}
