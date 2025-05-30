package cache

import "admin/pb"

type AdminCache interface {
	GetQueue() (*pb.QueueList, bool)
	SetQueue(*pb.QueueList) error

	GetSchedule(studentID string) (*pb.ScheduleResponse, bool)
	SetSchedule(studentID string, schedule *pb.ScheduleResponse) error
}
