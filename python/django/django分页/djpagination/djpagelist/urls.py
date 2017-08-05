from django.conf.urls import url 
from . import views

urlpatterns = [

	url(r'^fbvlist/$', views.fbvlist, name='fbvlist'),
	url(r'^cbvlist/$', views.UserListView.as_view(), name='cbvlist'),
	url(r'^cbvlistfix/$', views.UserListView_Fix.as_view(), name='cbvlist-fix'),

]