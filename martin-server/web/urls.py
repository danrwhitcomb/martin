''' Web URL conf '''
from django.conf.urls import url
from django.views.generic import TemplateView

from web import views

urlpatterns = [
    url(r'^login$', views.handle_login),
    url(r'^logout$', views.handle_logout),
    url(r'', views.home)
]
