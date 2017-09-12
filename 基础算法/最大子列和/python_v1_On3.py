# 接收用户输入的数据, 转为列表
input_list = []
user_input_list = input("输入一列的整数以逗号隔开，如:11，2，3，5. :")
for i in user_input_list.split(','):
    tmp = int(i)
    input_list.append(tmp)

# 列表总长度
list_len = len(input_list)

# 不初始化成0,随意取列表中的某个元素作为最大值,防止输入的数都是小于0（都是负数）
MaxSum = input_list[0]
for left_index in range(list_len):  # 子列左端的位置
    for right_index in range(left_index, list_len):  # 子列右端的位置
        This_Sum = 0
        for index in range(left_index, right_index + 1):  # 计算 子列左端到右端所有元素之和
            This_Sum += input_list[index]
            if This_Sum > MaxSum:  # 子列的和比较比MaxSum大则更新最大值
                MaxSum = This_Sum

print(MaxSum)

'''


'''
