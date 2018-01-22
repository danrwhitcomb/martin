import $ from 'jquery';
import DevicesView from '../components/views/DevicesView';

const DEVICE_API = '/api/device';

function basicApiSuccessHandler(data, textStatus, jqXHR) {
    return data.data;
}

function basicApiErrorHandler(jqXHR, textStatus, errorThrown) {
    return $.parseJSON(jqXHR.responseJSON).error;
}

function fetchDevices() {
    return $.ajax({url: DEVICE_API})
        .then(basicApiSuccessHandler, basicApiErrorHandler);
}

function discoverDevices() {
    return $.ajax({url: DEVICE_API + '/discover'})
        .then(basicApiSuccessHandler, basicApiErrorHandler);
}

function addDevice(deviceUuid) {
    return $.ajax({
        url: DEVICE_API,
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({'uuid': deviceUuid})
    }).then(basicApiSuccessHandler, basicApiErrorHandler);
}

export {fetchDevices, discoverDevices, addDevice};
