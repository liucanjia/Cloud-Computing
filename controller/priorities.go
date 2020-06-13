package controller

import (
	"log"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const (
	luckyPrioMsg = "pod %v/%v is lucky to get score %v\n"
)

func prioritize(args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
	pod := args.Pod
	nodes := args.Nodes.Items

	hostPriorityList := make(schedulerapi.HostPriorityList, len(nodes))
	for i, node := range nodes {
		cpu := node.Status.Allocatable.Cpu().MilliValue()
		memory := node.Status.Allocatable.Memory().MilliValue()
		score := int((cpu*3 + memory*7)/10)
		log.Printf(luckyPrioMsg, pod.Name, pod.Namespace, score)
		hostPriorityList[i] = schedulerapi.HostPriority{
			Host:  node.Name,
			Score: score,
		}
	}

	return &hostPriorityList
}