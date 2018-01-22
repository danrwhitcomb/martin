import React from 'react';
import $ from 'jquery';

import {addDevice, fetchDevices, discoverDevices} from 'services/device';

export default class DevicesView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            devices: [],
            discovered: [],
            devicesFetchError: null,
            discoveryError: null,
        };
    }

    componentWillMount() {
        this.getDiscoveredDevices();
        this.getOwnedDevices();
    }

    componentWillUnmount() {}

    getDiscoveredDevices() {
        discoverDevices().then((devices) => {
            this.setState({
                discovered: devices,
                discoveryError: null
            });
        }, (errorString) => {
            this.setState({
                discoveryError: errorString,
                discovered: []
            });
        });
    }

    getOwnedDevices() {
        fetchDevices().then((devices) => {
            this.setState({
                devices: devices,
                devicesFetchError: null
            });
        }, (errorString) => {
            this.setState({
                devices: [],
                devicesFetchError: null
            });
        });
    }

    addDeviceHandler(deviceUuid) {
        addDevice(deviceUuid)
        .then(null, (error) => {
            alert('An error occurred: ' + error);
        })
    }

    renderAlertBox(errorMsg) {
        return <div className='alert-box alert'>
            An error occurred: {this.state.discoveryError}
        </div>
    }

    render() {

        var discoveredDevices = null;
        if (this.state.discoveryError) {
            discoveredDevices = this.renderAlertBox(this.state.discoveryError);
        } else {
            discoveredDevices = this.state.discovered.map((device) => {
                return <div className='device' key={device.uuid}>
                    <div className='device-addr' >
                        <p>Address</p>
                        <div>{device.address}</div>
                    </div>
                    <div className='device-uuid'>
                        <p>UUID</p>
                        <div>{device.uuid}</div>
                    </div>
                    <div className='device-services'>
                        <p>Services</p>
                        <ul>{device.services.map((service) => <li key={service}>- {service}</li>)}</ul>
                    </div>
                    <div className='device-add'>
                        <button className='button' onClick={() => this.addDeviceHandler(device.uuid)}>Add</button>
                    </div>
                </div>
            });
        }

        var ownedDevices = null;
        if (this.state.devicesFetchError) {
            ownedDevices = this.renderAlertBox(this.state.devicesFetchError);
        } else {
            ownedDevices = this.state.devices.map((device) => {
                return <div className='device' key={device.uuid}>
                    <div className='device-name'>
                        <p>Name</p>
                        <div>{device.name}</div>
                    </div>
                    <div className='device-addr'>
                        <p>Address</p>
                        <div>{device.addr}</div></div>
                    <div className='device-services'>
                        <p>Services</p>
                        <ul>{device.services.map((service) => <li key={service.name}>{service.name}</li>)}</ul>
                    </div>
                </div>
            });
        }

        var discoveredWrapper = null;
        if (this.state.discovered.length > 0) {
            var discoveredWrapper = <div className='devices-discovered'>
                                        <h4>Discovered Devices</h4>
                                        {discoveredDevices}
                                    </div>
        }

        return <div>
            {discoveredWrapper}
            <div className='devices-owned'>{ownedDevices}</div>
        </div>;
    }
}
