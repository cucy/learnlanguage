from django.shortcuts import render

# Create your views here.
from django.contrib.auth.models import User
from django.core.paginator import Paginator, PageNotAnInteger, EmptyPage
from django.views.generic.list import ListView


def fbvlist(request):
    user_list = User.objects.all()
    page = request.GET.get('page', 1)  # 当前已经点击到的页面
    paginator = Paginator(user_list, 10)  # 每一页页数显示10条数据

    try:
        users = paginator.page(page)  # 拿出当前页的数据
    except PageNotAnInteger:
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


class UserListView(ListView):
    models = User
    template_name = 'cbv_list.html'         # 默认为: <app_label>/<model_name>_list.html
    context_object_name = 'users'           # 默认:  object_list
    paginate_by = 10                        # 每一页的最大查询返回列表
    queryset = User.objects.all()           # 默认: Model.objects.all()


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
