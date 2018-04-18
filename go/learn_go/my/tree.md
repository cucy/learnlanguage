```
.
├─01_数据类型
├─02_语句
│  └─01_循环
├─03_包_package
│  ├─01_变量可见与不可见
│  └─02_stringutil
├─04_变量_variables
│  ├─00_变量
│  │  └─1_定义
│  ├─01_编译器自动识别推导_短声明
│  ├─02_变量_0值
│  └─03_scope作用域
│      ├─01_package-scope
│      ├─02_block-scope
│      │  └─01_this-does-not-compile
│      ├─02_visibility
│      │  ├─main
│      │  └─vis
│      ├─03_closure
│      │  ├─01
│      │  ├─02
│      │  ├─03
│      │  └─04
│      ├─03_order-matters
│      └─04_variable-shadowing
├─05_constants
│  ├─01_constant
│  └─02_iota
├─06_memory-address
│  └─01_showing-address
├─07_pointers
│  ├─01_referencing
│  ├─02_using-pointers
│  └─03_using-pointers
│      ├─no-pointer
│      └─pointer
├─08_for-loop
│  ├─01_init-condition-post初始化条件
│  └─02_rune-loop_UTF8
├─09_switch-statements
│  ├─01_switch
│  └─02_on_type
├─10_if_else-if_else
│  ├─01_base
│  └─02_just-fyi
│      ├─1_Benchmark
│      ├─2_benchMark
│      └─3_utf
├─11_function
│  ├─01_base
│  ├─02_func-expression函数表达式
│  │  └─1base
│  ├─03_closure
│  ├─04_callback
│  │  └─01_print_number
│  ├─05_recursion递归
│  ├─06_defer
│  └─07_passing-by-value通过值_引用_传递
│      ├─03_string
│      ├─04_string-pointer
│      ├─05_REFERENCE-TYPE
│      ├─06_REFERENCE-TYPE
│      ├─07_struct-pointer
│      ├─08_anon_self-executing立刻执行
│      ├─1_int
│      └─2_int-pointer
├─12_bool-expressions布尔表达式
│  └─01base
├─13_array
│  └─01
├─14_slice
│  ├─01_init-slice
│  ├─02_slicing-a-slice
│  ├─03_make-slice
│  ├─04_append
│  ├─05_delete
│  └─06_multi-dimensional多维slice
│      └─01_shorthand-slice
├─15_map
│  ├─01_base
│  └─02_hash_table
│      ├─02_get
│      ├─03_scanner
│      └─04_moby-dicks-words_莫比迪克斯的话
├─16_struct
│  ├─01_user_defined_type
│  ├─02_struct_fields_values_initialization_struct字段值初始化
│  ├─03_methods
│  ├─04_embedded-types嵌入类型
│  ├─05_promotion
│  │  └─01_overriding-fields
│  ├─06_struct-pointer
│  ├─07_marshal_unmarshal
│  │  ├─01_marshal
│  │  └─02_unmarshal
│  └─08_encode_decode
│      └─01_encode
├─17_interfaces
│  ├─01_base
│  │  ├─01_no-interface
│  │  ├─02_interface
│  │  ├─03_interface
│  │  ├─04_interface
│  │  └─05_io-copy
│  │      ├─01_no-error-checking
│  │      └─02_error-checking
│  ├─02_package-sort
│  │  ├─01_sort_name
│  │  ├─02_sort-names_type-StringSlice
│  │  ├─03_sort-Strings
│  │  ├─04_sort-names_type-StringSlice_reverse
│  │  ├─05_sort-int_type-IntSlice
│  │  ├─06_sort-int_type-IntSlice_reverse
│  │  ├─07_sort-Ints
│  │  └─08_standard-library-example
│  ├─03_empty-interface
│  │  ├─01_no-interface
│  │  ├─02_empty-interface
│  │  ├─03_param-accepts-any-type
│  │  └─04_slice-of-any-type
│  ├─04_method-sets
│  │  ├─01_value-receiver_value-type
│  │  ├─02_value-receiver_pointer-type
│  │  ├─03_pointer-receiver_pointer-type
│  │  └─04_pointer-receiver_value-type
│  └─05_conversion-vs-assertion 转换和断言
│      └─01_转换
│          ├─01_int-to-float
│          ├─02_fload_to_int
│          ├─03_rune_to_string
│          ├─04_rune_to_slice_of_bytes_to_string
│          ├─05_string_to_slice_of_bytes
│          ├─06_strconv
│          │  ├─01_Atoi
│          │  ├─02_Itoa
│          │  └─03_ParseInt
│          └─07_assertion断言
│              ├─01_no_interface_error_invalid_code
│              ├─02_interface_string
│              └─03_interface_int_print-type
├─18_go_routines
│  ├─01_no_go
│  ├─02_go_concurrency_go并发
│  ├─03_wait_group
│  ├─04_time_sleep
│  ├─05_gomaxprocs_parallelism
│  ├─06_race-condition竞争条件
│  ├─07_mutex互斥锁
│  ├─08_atomicity原子性
│  ├─09_channels
│  │  ├─00_unbuffered-channels-block
│  │  ├─01_range
│  │  ├─02_n-to-1
│  │  │  ├─01_race-condition竞争条件
│  │  │  ├─02_wait-group
│  │  │  ├─03_semaphore信号量
│  │  │  ├─04_semaphore_wrong-way错误的操作方式
│  │  │  └─05_n-times_to_1
│  │  ├─03_1-to-n
│  │  │  ├─01_1_to_2-times
│  │  │  └─02_1_to_n-times
│  │  ├─04_pass-return-channels 通过返回通道
│  │  ├─05_channel-direction方向
│  │  ├─06_refactor重构
│  │  ├─07_incrementor
│  │  └─08_closures
│  │      ├─01_no-closure-binding
│  │      ├─02_closure-binding关闭绑定
│  │      └─03_closure-binding
│  ├─10_deadlock-challenges死锁挑战
│  │  ├─01_deadlock-challenge
│  │  ├─02_deadlock-solution死锁解决方案
│  │  ├─03_deadlock-challenge
│  │  ├─04_deadlock-challenge
│  │  └─05_deadlock-solution
│  ├─11_factorial-challenge进阶挑战
│  │  ├─01_challenge-description
│  │  └─02_challenge-solution
│  ├─12_channels_pipeline
│  │  ├─01_sq-output
│  │  ├─02_sq-output
│  │  ├─03_challenge-description
│  │  └─04_challenge-solution
│  │      ├─01_original-solution
│  │      └─02_another-solution
│  ├─13_channels_fan-out_fan-in
│  │  ├─01_boring
│  │  ├─02_sq-output
│  │  ├─03_sq-output_variation
│  │  ├─04_challenge-description
│  │  ├─05_challenge-solution
│  │  ├─06_challenge-description
│  │  ├─07_challenge-solution
│  │  ├─08_challenge-description
│  │  ├─09_challenge-solution
│  │  │  ├─01_troubleshooting-step
│  │  │  └─02_solution
│  │  └─10_van-sickle_fan-out_fan-in
│  ├─14_incrementor-challenge
│  │  ├─01_description
│  │  └─02_solution
│  └─15_for-fun
│      └─01
├─19_error_handling
│  ├─01_golint
│  │  ├─01_before
│  │  └─02_after
│  ├─02_err-not-nil
│  │  ├─01_fmt-println
│  │  ├─02_log-println
│  │  ├─03_log-set-output
│  │  ├─04_log-fatalln
│  │  └─05_panic
│  └─03_custom-errors
│      ├─01_errors-new
│      ├─02_errors-new_var
│      ├─03_fmt-errorf
│      ├─04_fmt-errorf_var
│      └─05_custom-type
└─20_testing
```