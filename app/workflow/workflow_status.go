package workflow

type StatusType uint32

const (
	TASK_READY StatusType = 1 << iota
	TASK_PAUSE
	TASK_TERM
	TASK_RUN
	TASK_FIN

	JOB_READY StatusType = 1 << iota
	JOB_PAUSE
	JOB_TERM
	JOB_RUN
	JOB_FIN

	TaskStatus = TASK_PAUSE | TASK_TERM | TASK_FIN | TASK_RUN | TASK_READY
	JobStatus  = JOB_PAUSE | JOB_TERM | JOB_FIN | JOB_RUN | JOB_READY
)

type TaskNodeType uint32

const (
	TASK_START TaskNodeType = 1 << iota
	TASK_END

	FlowNodeType = TASK_START | TASK_END
)
