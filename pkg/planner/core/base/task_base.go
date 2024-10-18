// Copyright 2024 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

// Note: appending the new adding method to the last, for the convenience of easy
// locating in other implementor from other package.

// Task 是 `PhysicalPlanInfo` 的新版本。它存储任务的代价信息。
// 一个任务可能是 CopTask、RootTask、MPPTaskMeta 或 ParallelTask。
type Task interface {
	// Count 返回当前任务的行数估计。
	Count() float64
	// Copy 返回当前任务的浅拷贝，与 p 指向相同的计划。
	Copy() Task
	// Plan 返回当前任务的物理计划。
	Plan() PhysicalPlan
	// Invalid 返回当前任务是否无效。
	Invalid() bool
	// ConvertToRootTask 将当前任务转换为根任务类型。
	// 这里我们将返回类型更改为接口以避免导入循环。
	// 基本接口定义不应依赖于具体的实现结构。
	ConvertToRootTask(ctx PlanContext) Task
	// MemoryUsage 返回当前任务的内存使用量估计。
	MemoryUsage() int64
	//Count(): 返回任务的估计行数，用于优化器评估执行计划的代价。
	//Copy(): 创建一个任务的浅拷贝，新的任务与原任务共享底层的物理计划，但可以拥有不同的代价信息或其他属性。
	//Plan(): 返回任务对应的物理执行计划 (PhysicalPlan)。
	//Invalid(): 判断任务是否有效。当优化器无法生成有效的执行计划时，会返回一个无效的任务。
	//ConvertToRootTask(): 将当前任务转换为根任务类型 (RootTask)。根任务通常是整个查询计划的顶层任务。
	//MemoryUsage(): 返回任务的估计内存使用量，用于优化器评估执行计划的资源消耗。
}

// InvalidTask is just a common invalid singleton instance initialized by core's empty RootTask.
var InvalidTask Task
