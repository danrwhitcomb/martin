''' Zeroconf service monitor'''
import ipaddress

from zeroconf import ServiceBrowser, Zeroconf

class DiscoveredDevice:
    def __init__(self, name, address, service_info):
        self.name = name
        self.address = address
        self._service_info = service_info

        # Device properties
        properties = {to_utf8(key): to_utf8(value) for key, value in service_info.properties.items()}
        self.uuid = properties['uuid']
        self.services = [service.strip() for service in properties['services'].split(',')]

    def todict(self):
        return {
            'name': self.name,
            'address': self.address,
            'uuid': self.uuid,
            'services': self.services
        }

class ZeroconfDeviceManager:

    def __init__(self):
        self.services = {}
        self.zeroconf = Zeroconf()
        self.browser = ServiceBrowser(self.zeroconf, "_martin._tcp.local.", self)

    def __del__(self):
        self.zeroconf.close()

    def remove_service(self, zeroconf, type, name):
        if name in self.services:
            del self.services[name]

    def add_service(self, zeroconf, type, name):
        info = zeroconf.get_service_info(type, name)
        self.services[name] = DiscoveredDevice(name, bytes_to_ip(info.address), info)
        print("Service %s added, service info: %s" % (name, info))


def to_utf8(bytes):
    return bytes.decode('utf-8')

def bytes_to_ip(bytes):
    return ipaddress.ip_address(bytes).exploded

zeroconf_manager = ZeroconfDeviceManager()
