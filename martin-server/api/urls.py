''' API urls '''
from django.conf.urls import url  # noqa

from martin_lib.rest import route

from api.views.device import search_network,get_devices, \
                             get_device, request_device_registration

urlpatterns = [
    url(r'^device$', route(GET=get_devices, POST=request_device_registration)),
    url(r'^device/discover$', route(GET=search_network)),
    url(r'^device/<uuid:uuid>', route(GET=get_device)),
]
