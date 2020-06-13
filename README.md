# sample-scheduler-extender
extender工作逻辑:
在过滤函数中,循环遍历每个节点，然后检查随机数是否为奇数,是则判断节点满足候选标准
在打分函数中,通过计算每个节点的Cpu和Memory剩余量,通过一个加权函数获得最终的分数，在这里我用的是score = ((cpu*3) + (memory*7)) / 10
得分高的节点作为最终的调度节点
