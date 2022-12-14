package main

/*
首先画出全排列的树形结构，以123为例，一开始排列为空列表，第一个位置有三种可能，分别是
1、2、3画出三个分支；由于第一个位置已经被占用，那么第二个位置可选择的就只有两个，所以又可以
展开两个分支，如：1下的1,2和1,3；选出两个数字之后，最后就只剩下一个数字了，所以最后一个位置上
的数就是唯一确定的了。之后这个树的所有叶子结点就是全排列的结果。
*/
/*
回溯过程：先选择1，之后按顺序选择2，最后没有可选数字就得到了1,2,3；为了得到
所有的排列，这时候就要进行回溯。最后一步选择的是3那么回退的时候就要撤回3，回到1,2结点
由于1,2阶段3已经被选择过了，所以继续撤销2，回退到1结点，这个阶段本可以选择2或者3
但是2已经选择过了，所以下一步就要选择3，得到1,3结点，之后再进行刚才的选择回退操作
*/
/*
这个树除了叶子结点以外，其他结点做的事情都是一样的，也就是在已经选了数 的前提下
需要在剩下还没有选择的数里，按照顺序选择一棵树，所以这就是一个递归。那么
递归终止的条件就是数字的个数已经选完了。所以我们需要一个变量来记录已经选了
多少个数字，其实这个变量等价递归到了第几层depth，当遍历的层数和输入数组
的个数相等的时候，所有的元素就都被考虑完了，就可以退出递归。
将已经选择的数放进一个列表里temp，这个其实就是树的路径，因为要不断地添加删除
所以这个应该是个栈。在设置一个布尔数组used表示当前已经考虑的数字是否在之前
已经选择过，也就是判断是否在path变量里，初始化都为FALSE，表示都未被选择。

*/
func main() {
}
func permute(nums []int) [][]int  {
	//保存输入数组的长度
	nlen := len(nums)
	//初始化，用来存放结果
	var result [][]int
	//如果传入长度为0，那就直接返回空数组（要对空列表进行初始化）
	if nlen==0{
		return result
	}
	//创建中间变量，存放临时结果
	var temp []int
	//创建bool值，判断该位置数字是否用过
	used := make([]bool, nlen)
	//回溯函数
	BackTrack(used, temp, nums, &result,nlen,0)
	return result
}
func BackTrack(used []bool, temp []int, nums []int, result *[][]int,nlen int,depth int) {
	//判断回溯函数结束条件
	//当临时temp长度和所给的数字长度相等时(也就是递归到了第几层)，将该temp加入结果
	if depth == nlen {
		//由于go语言的特性如果不特别说明创建的切片本质上都是指向同一个内存空间
		//如果想要循环赋值的切片与原来切片不相关，需要另外开辟空间，这里用到copy函数，开辟独立空间
		current := make([]int, depth)
		copy(current, temp)
		*result = append(*result, current)
	}
	//遍历数组中的数字，进行排列组合
	for i := 0; i < nlen; i++ {
		//减枝，当该位置数字使用过时则跳过
		if used[i] {
			continue
		}
		//没有使用过就添加数字
		temp = append(temp, nums[i])
		//将该位置数字设置为访问过的状态
		used[i] = true
		//递归继续搜索该支线
		BackTrack(used, temp, nums, result,nlen,depth+1)
		//回溯，恢复到之前的状态
		temp = temp[:len(temp)-1]
		used[i] = false
	}
}

