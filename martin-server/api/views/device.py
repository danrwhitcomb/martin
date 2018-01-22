''' Device handling endpoints '''
from django.core.exceptions import ObjectDoesNotExist
from rest_framework.response import Response
from rest_framework.status import HTTP_400_BAD_REQUEST, HTTP_404_NOT_FOUND
from schema import Schema

from martin_lib.rest import ApiResponse, validate_body

from common.zeroconf import zeroconf_manager
from common.models import Device, DeviceSerializer, DeviceState, Service

request_registration_schema = Schema({
    'uuid': str
})

def search_network(request):
    added_uuids = [str(device.uuid) for device in list(Device.objects.all())]
    print(added_uuids)
    discovered_devices = [device.todict() for device in zeroconf_manager.services.values()
                          if device.uuid not in added_uuids]

    return ApiResponse(data=discovered_devices)

def get_devices(request):
    data = [DeviceSerializer(device).data for device in list(Device.objects.all())]
    return ApiResponse(data=data)

def get_device(request, uuid=None):
    if uuid is None:
        return Response(status=HTTP_400_BAD_REQUEST)

    try:
        device = Device.objects.get(uuid=uuid)
        return ApiResponse(data=DeviceSerializer(device).data)
    except ObjectDoesNotExist:
        return ApiResponse(error='Device with UUID {} could not be found'.format(uuid),
                           status=HTTP_404_NOT_FOUND)

@validate_body(request_registration_schema)
def request_device_registration(request):
    body = request.data
    devices_with_uuid = [device for device in zeroconf_manager.services.values() \
                         if device.uuid == body['uuid']]

    if len(devices_with_uuid) == 0:
        return ApiResponse(error='Device with UUID {} could not be found'.format(body['uuid']),
                           status=HTTP_404_NOT_FOUND)

    discovered_device = devices_with_uuid[0]

    services = Service.objects.filter(name__in=discovered_device.services)
    added_device = Device.objects.create(name='test!',
                                         uuid=body['uuid'],
                                         state=DeviceState.objects.get(value='pending'),
                                         addr=discovered_device.address)
    added_device.services.add(*list(services))

    return ApiResponse(data=DeviceSerializer(added_device).data)
