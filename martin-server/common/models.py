from django.db import models
from django.utils.translation import ugettext_lazy as _
from rest_framework.serializers import ModelSerializer

from model_utils.fields import AutoCreatedField, AutoLastModifiedField


class IndexedTimeStampedModel(models.Model):
    created = AutoCreatedField(_('created'), db_index=True)
    modified = AutoLastModifiedField(_('modified'), db_index=True)

    class Meta:
        abstract = True

class Service(models.Model):
    name = models.CharField(max_length=32)
    description = models.TextField()

class ServiceSerializer(ModelSerializer):
    class Meta:
        model = Service
        fields = ('name', 'description')

class DeviceState(models.Model):
    value = models.CharField(max_length=32)

class DeviceStateSerializer(ModelSerializer):
    class Meta:
        fields = ('value')

class Device(IndexedTimeStampedModel):
    name = models.CharField(max_length=32, default='')
    uuid = models.UUIDField()
    addr = models.GenericIPAddressField()
    services = models.ManyToManyField(Service)
    state = models.ForeignKey(DeviceState, on_delete=models.DO_NOTHING)

class DeviceSerializer(ModelSerializer):
    services = ServiceSerializer(many=True, read_only=True)

    class Meta:
        model = Device
        fields = ('id', 'name', 'uuid', 'addr', 'services', 'created', 'modified')
