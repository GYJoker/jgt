package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"sync"
)

// 定时任务工具类
type (
	YJCron struct {
		// Job 定时任务
		Job *YJJob
		// TimeCron 定时任务cron表达式
		TimeCron string
	}

	YJCronTime struct {
		// Job 定时任务
		Job *YJJob
		// TimeType 定时任务时间类型
		TimeType int
		// Space 定时任务执行间隔
		Space uint64
	}

	YJJob struct {
		// JobFunc 定时任务函数
		JobFunc   func()
		runningMu sync.Mutex
	}
)

var (
	_cron    *cron.Cron
	cronOnce sync.Once
)

const (
	JobTimeSecond = iota
	JobTimeMinute
	JobTimeHour
	JobTimeDay
)

func getCron() *cron.Cron {
	cronOnce.Do(func() {
		_cron = cron.New(cron.WithSeconds())
		_cron.Start()
	})
	return _cron
}

func (j *YJJob) Run() {
	j.runningMu.Lock()
	defer j.runningMu.Unlock()
	j.JobFunc()
}

func (c *YJCron) AddJob() (int, error) {
	entryID, err := getCron().AddJob(c.TimeCron, c.Job)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(entryID), nil
}

func (c *YJCron) RemoveJob(entryID int) {
	getCron().Remove(cron.EntryID(entryID))
}

func (c *YJCronTime) AddJob() (int, error) {
	curCron := getCron()
	var entryID cron.EntryID
	var err error
	switch c.TimeType {
	case JobTimeSecond:
		entryID, err = curCron.AddJob("@every "+strconv.FormatUint(c.Space, 10)+"s", c.Job)
		break
	case JobTimeMinute:
		entryID, err = curCron.AddJob("@every "+strconv.FormatUint(c.Space, 10)+"m", c.Job)
		break
	case JobTimeHour:
		entryID, err = curCron.AddJob("@every "+strconv.FormatUint(c.Space, 10)+"h", c.Job)
		break
	case JobTimeDay:
		entryID, err = curCron.AddJob("@every "+strconv.FormatUint(c.Space, 10)+"d", c.Job)
		break
	}
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(entryID), nil
}

func (c *YJCronTime) RemoveJob(entryID int) {
	getCron().Remove(cron.EntryID(entryID))
}
