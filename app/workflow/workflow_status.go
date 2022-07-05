package workflow

type StatusType uint32

const (
	TaskReady StatusType = 1 << iota
	TaskPause
	TaskTerm
	TaskRun
	TaskFin

	JobReady StatusType = 1 << iota
	JobPause
	JobTerm
	JobRun
	JobFin

	TaskStatus = TaskPause | TaskTerm | TaskFin | TaskRun | TaskReady
	JobStatus  = JobPause | JobTerm | JobFin | JobRun | JobReady
)

type TaskNodeType uint32

const (
	TaskStart TaskNodeType = 1 << iota
	TaskEnd

	FlowNodeType = TaskStart | TaskEnd
)
