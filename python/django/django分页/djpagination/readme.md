## 建立多个数据

```python
from django.contrib.auth.models import User
for i in range(200):
	User.objects.create_user('admin{}'.format(i), "admin{}@zrd.com".format(i), '123456')
```


## django 分页
### Paginator类

`主要作用是页码相关`

```python
$ python manage.py shell

In [1]: from django.contrib.auth.models import User
In [2]: from django.core.paginator import Paginator
In [3]:
In [3]: # queryset
In [4]: user_list = User.objects.all()
In [4]: # class Paginator(object):
In [4]: #     def __init__(self, object_list, per_page, orphans=0, allow_empty_first_page=True):
In [5]: # 分页对象  实例传入参数(一般是数据集合, 每页显示多少条数据)
In [6]: paginator = Paginator(user_list, 10)
```
- `必传入参数`
  - `object_list`  一个list，tuple，django的QuerySet
  - `per_page` 每一页显示最大的数据对象个数
- `可选参数`
   - ` orphans` 最后页显示多少条 如果`per_page = 10 orphans=2` 会显示13条
   - `allow_empty_first_pag` 首页是否能为空

- 分页对象所有方法预览
```python
In [7]: paginator.
                   allow_empty_first_page object_list            page_range
                   count                  orphans                per_page
                   num_pages              page()                 validate_number()
```

| 输入                                       |       输出       |       类型       |    含义     |
| :--------------------------------------- | :------------: | :------------: | :-------: |
| In [7]: paginator.allow_empty_first_page |      True      |      bool      |           |
| In [9]: paginator.count                  |      219       |      Int       | 总共有219条数据 |
| In [11]: In [10]: paginator.num_pages    |       22       |      Int       | 数据可以分为22页 |
| In [11]: paginator.object_list           | query_set_list | query_set_list |  返回查询列表   |
| In [12]: paginator.orphans               |       0        |      Int       | 最后一页数据总和  |
| In [18]: paginator.page(4)               | <Page 4 of 22> |                |           |
| In [23]: paginator.page_range            |  range(1, 23)  |                |           |
| In [24]: paginator.per_page              |       10       |      Int       | 每页最大的查询集合 |
| In [30]: paginator.validate_number(22)   |       22       |                |  检查所给的页数  |

### Page类

​	`一般是是由Paginator生成`

```python
from django.contrib.auth.models import User
from django.core.paginator import Paginator
user_list = User.objects.all()
paginator = Paginator(user_list, 10)

users = paginator.page(2) # 取第二页的数据
```

- 属性方法集
```python
In [60]: users.
                count()                has_other_pages()      next_page_number()     paginator
                end_index()            has_previous()         number                 previous_page_number()
                has_next()             index()                object_list            start_index()
```

```python
In [109]: # 起始条目
In [110]: users.start_index()
Out[110]: 11
    
In [73]: # 最后查询集条目
In [74]: users.end_index()
Out[74]: 20

In [75]: # 是否有下一页
In [76]: users.has_next()
Out[76]: True

In [77]: # 是否有 上一页 或者 下一页
In [78]: users.has_other_pages()
Out[78]: True

In [79]: # 是否有上一页
In [80]: users.has_previous()
Out[80]: True
    
In [84]: # 下一页页码
In [85]: users.next_page_number()
Out[85]: 3

In [87]: # 当前对象所在的页码
In [88]: users.number
Out[88]: 2

In [93]: # 返回当前页的对象列表 可迭代每个元素的属性方法
In [94]: users.object_list
Out[94]:
[<User: w1>,
 <User: w11>,
 <User: w12>,
 <User: s1>,
 <User: s2>,
 <User: s5>,
 <User: admin0>,
 <User: admin1>,
 <User: admin2>,
 <User: admin3>]

In [104]: # 上一页页码
In [105]: users.previous_page_number()
Out[105]: 1
```

**属性**

*Page.object_list*

当前页对象列表

*Page.number*

当前页的索引

*Page.paginator*

和page相关的分页类





##  fbv例子


- 	`views.py`
```python
from django.shortcuts import render

# Create your views here.
from django.contrib.auth.models import User
from django.core.paginator import Paginator, PageNotAnInteger, EmptyPage


def fbvlist(request):
	user_list = User.objects.all()
	page = request.GET.get('page', 1) # 当前已经点击到的页面
	paginator = Paginator(user_list, 10) # 每一页页数显示10条数据

	try:
		users = paginator.page(page) # 拿出当前页的数据
	except PageNotAnInteger :
		users = paginator.page(1)  # 异常则返回第一页的数据
	except EmptyPage:
		users = paginator.page(paginator.num_pages)

	# 【1.2.3.4.5.当前页.7.8.9.10】
	# before_index = 6 
	# after_index = 5
	# start_index = users.number - before_index
	# if start_index < 0:
	# 	start_index = 0
	# users.paginator.page_rang[start_index: users.number + after_index]
	return render(request, 'fbv_list.html', {'users': users})
```
- 模板`fbv_list.html`
```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>用户列表</title>
    <link href="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/css/bootstrap.min.css"  type="text/css" rel="stylesheet">
</head>

<body>
<h1 class="text-center">用户列表</h1>
<table class="table table-bordered">
    <thead>
    <tr>
        <th>用户名</th>
        <th>密码</th>
    </tr>
    </thead>
    <tbody>
    {% for user in users %}
    <tr>
        <td>{{ user.username }}</td>

        <td>{{ user.password }}</td>
    </tr>
    {% endfor %}
    </tbody>
</table>

{% if users.has_other_pages %}
<ul class="pagination">
    {% if users.has_previous %}
        <li><a href="?page={{ users.previous_page_number }}">前一页</a></li>
    {% else %}
        <li class="disabled"><span>&laquo;</span></li>
    {% endif %}

    {% for i in users.paginator.page_range %}
        {% if users.number == i %}
            <li class="active"><span>{{ i }} <span class="sr-only">(current)</span></span></li>
        {% else %}
            <li><a href="?page={{ i }}">{{ i }}</a></li>
        {% endif %}
    {% endfor %}

    {% if users.has_next %}
        <li><a href="?page={{ users.next_page_number }}">后一页</a></li>
    {% else %}
        <li class="disabled"><span>&raquo;</span></li>
    {% endif %}
</ul>
{% endif %}

<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/jquery.min.js" type="text/javascript"></script>
<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
</body>
</html>
```
- `urls.py`
```python
from django.conf.urls import url 
from . import views

urlpatterns = [

	url(r'^fbvlist/$', views.fbvlist, name='fbvlist'),

]
```


## CBV例子
- `views.py`
```python
from django.contrib.auth.models import User
from django.views.generic.list import ListView


class UserListView(ListView):
	models = User 
	template_name = 'cbv_list.html' # 默认为: <app_label>/<model_name>_list.html
	context_object_name = 'users'	# 默认:  object_list
	paginate_by = 10   				# 每一页的最大查询返回列表
	queryset = User.objects.all()  	# 默认: Model.objects.all()
```
- `urls.py`
```python
from django.conf.urls import url 
from . import views

urlpatterns = [
	url(r'^cbvlist/$', views.UserListView.as_view(), name='cbvlist'),

]
```
- 模板`cbv_list.html`
```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>用户列表</title>
    <link href="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/css/bootstrap.min.css"  type="text/css" rel="stylesheet">
</head>

<body>
<h1 class="text-center">用户列表</h1>

<table class="table table-bordered">
  <thead>
    <tr>
      <th>用户名</th>
      <th>邮箱</th>
      <th>密码</th>
    </tr>
  </thead>
  <tbody>
    {% for user in users %}
      <tr>
        <td>{{ user.username }}</td>
        <td>{{ user.email }}</td>
        <td>{{ user.password }}</td>
      </tr>
    {% endfor %}
  </tbody>
</table>

{% if is_paginated %}
  <ul class="pagination">

    {% if page_obj.has_previous %}
      <li><a href="?page={{ page_obj.previous_page_number }}">前一页</a></li>
    {% else %}
      <li class="disabled"><span>&laquo;</span></li>
    {% endif %}

    {% for i in paginator.page_range %}
      {% if page_obj.number == i %}
        <li class="active"><span>{{ i }} <span class="sr-only">(current)</span></span></li>
      {% else %}
        <li><a href="?page={{ i }}">{{ i }}</a></li>
      {% endif %}
    {% endfor %}

    {% if page_obj.has_next %}
      <li><a href="?page={{ page_obj.next_page_number }}">后一页</a></li>
    {% else %}
      <li class="disabled"><span>&raquo;</span></li>
    {% endif %}
      
  </ul>
{% endif %}


<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/jquery.min.js" type="text/javascript"></script>
<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
</body>
</html>
```
## CBV修改版

- `views.py`
```python
class UserListView_Fix(ListView):
    models = User
    template_name = 'cbv_list_fix.html'  # 默认为: <app_label>/<model_name>_list.html
    context_object_name = 'users'           # 默认:  object_list
    paginate_by = 10                        # 每一页的最大查询返回列表
    queryset = User.objects.all()           # 默认: Model.objects.all()

    def get_page_range(self, page_obj):
        before_index = 6  # 前展示页数
        after_index = 5  # 后展示页数
        start_index = page_obj.number - before_index
        end_index = page_obj.number + after_index

        if start_index < 0:
            start_index = 0

        ''' 范围等于 [ 当前页-前展示页数 : 当前页 + 后展示页数 ] '''
        '''  << 1,2,3,4,5,当前页,7,8,9,10 >>'''
        return page_obj.paginator.page_range[start_index: end_index]

    def get_context_data(self, **kwargs):
        context = super(UserListView_Fix, self).get_context_data(**kwargs)
        context['page_range'] = self.get_page_range(context['page_obj'])
        return context
```
- `urls.py`
```python
from django.conf.urls import url 
from . import views

urlpatterns = [
	url(r'^cbvlistfix/$', views.UserListView_Fix.as_view(), name='cbvlist-fix'),

]
```
- 模板 `cbv_list_fix.html`
```php+HTML
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>用户列表</title>
    <link href="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/css/bootstrap.min.css"  type="text/css" rel="stylesheet">
</head>

<body>
<h1 class="text-center">用户列表</h1>

<table class="table table-bordered">
  <thead>
    <tr>
      <th>用户名</th>
      <th>邮箱</th>
      <th>密码</th>
    </tr>
  </thead>
  <tbody>
    {% for user in users %}
      <tr>
        <td>{{ user.username }}</td>
        <td>{{ user.email }}</td>
        <td>{{ user.password }}</td>
      </tr>
    {% endfor %}
  </tbody>
</table>

{% if is_paginated %}
  <ul class="pagination">
  {% if page_obj.number != 1 %}
    <li><a href="?page=1">首页</a></li>
  {% endif %}

    {% if page_obj.has_previous %}
      <li><a href="?page={{ page_obj.previous_page_number }}">前一页</a></li>
    {% else %}
      <li class="disabled"><span>前一页</span></li>
    {% endif %}

    {% for i in page_range %}
      {% if page_obj.number == i %}
        <li class="active"><span>{{ i }} <span class="sr-only">(current)</span></span></li>
      {% else %}
        <li><a href="?page={{ i }}">{{ i }}</a></li>
      {% endif %}
    {% endfor %}

    {% if page_obj.has_next %}
      <li><a href="?page={{ page_obj.next_page_number }}">后一页</a></li>
    {% else %}
      <li class="disabled"><span>后一页</span></li>
    {% endif %}

  {% if page_obj.number != page_obj.paginator.num_pages %}
    <li><a href="?page={{ page_obj.paginator.num_pages }}">末页</a></li>
  {% endif %}
  </ul>
{% endif %}


<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/jquery.min.js" type="text/javascript"></script>
<script src="http://ou04217kp.bkt.clouddn.com/assets/global/plugins/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
</body>
</html>
```